// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fetch provides a way to fetch modules from a proxy.
package fetch

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"path"
	"runtime/debug"
	"strings"
	"sync"

	"golang.org/x/mod/module"
	"golang.org/x/pkgsite/internal"
	"golang.org/x/pkgsite/internal/derrors"
	"golang.org/x/pkgsite/internal/godoc"
	"golang.org/x/pkgsite/internal/licenses"
	"golang.org/x/pkgsite/internal/log"
	"golang.org/x/pkgsite/internal/source"
	"golang.org/x/pkgsite/internal/trace"
	"golang.org/x/sync/errgroup"
)

// A goPackage is a group of one or more Go source files with the same
// package header. Packages are part of a module.
type goPackage struct {
	path              string
	name              string
	imports           []string
	isRedistributable bool
	licenseMeta       []*licenses.Metadata // metadata of applicable licenses
	// v1path is the package path of a package with major version 1 in a given
	// series.
	v1path string
	docs   []*internal.Documentation // doc for different build contexts
	err    error                     // non-fatal error when loading the package (e.g. documentation is too large)
}

// rel returns the relative path from the modulePath to the pkgPath
// returning "." if they're the same.
func rel(pkgPath, modulePath string) string {
	suff := internal.Suffix(pkgPath, modulePath)
	if suff == "" {
		return "."
	}
	return suff
}

// extractPackage returns a package from a filesystem arranged like a module zip.
// It matches against the given licenses to determine the subset of licenses
// that applies to each package.
// It returns a packageVersionState representing the status of doing the work
// of computing the package after the UnitMeta was computed. The packageVersionState
// of a package that failed to have a UnitMeta produced was produced by extractPackageMetas.
func extractPackage(ctx context.Context, modulePath, pkgPath string, contentDir fs.FS, d *licenses.Detector, sourceInfo *source.Info, modInfo *godoc.ModuleInfo) (*goPackage, *internal.PackageVersionState, error) {
	innerPath := rel(pkgPath, modulePath)
	f, err := contentDir.Open(innerPath)
	if err != nil {
		return nil, nil, err
	}
	dir, ok := f.(fs.ReadDirFile)
	if !ok {
		return nil, nil, fmt.Errorf("file is not a directory")
	}
	entries, err := dir.ReadDir(0)
	if err != nil {
		panic(err)
	}
	var goFiles []string
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		if !strings.HasSuffix(e.Name(), ".go") {
			// We care about .go files only.
			continue
		}
		goFiles = append(goFiles, path.Join(innerPath, e.Name()))
	}
	if len(goFiles) == 0 {
		// We shouldn't be here, because we only call extratPackage for package units.
		return nil, nil, fmt.Errorf("extractPackage called but no go files: modulePath=%s, pkgPath=%q", modulePath, pkgPath)
	}

	var (
		status error
		errMsg string
	)
	pkg, err := loadPackage(ctx, contentDir, goFiles, innerPath, sourceInfo, modInfo)
	if bpe := (*BadPackageError)(nil); errors.As(err, &bpe) {
		log.Infof(ctx, "Error loading %s: %v", innerPath, err)
		status = derrors.PackageInvalidContents
		errMsg = err.Error()
	} else if err != nil {
		return nil, nil, fmt.Errorf("unexpected error loading package: %v", err)
	}

	if pkg == nil {
		// No package.
		if len(goFiles) > 0 {
			// There were go files, but no build contexts matched them.
			status = derrors.PackageBuildContextNotSupported
		}
	} else {
		if errors.Is(pkg.err, godoc.ErrTooLarge) {
			status = derrors.PackageDocumentationHTMLTooLarge
			errMsg = pkg.err.Error()
		} else if pkg.err != nil {
			// ErrTooLarge is the only valid value of pkg.err.
			return nil, nil, fmt.Errorf("bad package error for %s: %v", pkg.path, pkg.err)
		}
		if d != nil { //  should only be nil for tests
			isRedist, lics := d.PackageInfo(innerPath)
			pkg.isRedistributable = isRedist
			for _, l := range lics {
				pkg.licenseMeta = append(pkg.licenseMeta, l.Metadata)
			}
		}
	}

	pvs := &internal.PackageVersionState{
		ModulePath:  modulePath,
		PackagePath: pkgPath,
		Version:     modInfo.ResolvedVersion,
		Status:      derrors.ToStatus(status),
		Error:       errMsg,
	}

	return pkg, pvs, nil
}

type packageMeta struct {
	path string
	name string
}

// extractPackageMetas returns a slice of packageMetas containing only the information
// needed to produce a UnitMeta from filesystem arranged like a module zip. extractPackage
// does the work to complete the package with all the information needed for a Unit.
// It matches against the given licenses to determine the subset of licenses
// that applies to each package.
// The second return value says whether any packages are "incomplete," meaning
// that they contained .go files but couldn't be processed due to current
// limitations of this site. The limitations are:
// * a maximum file size (MaxFileSize)
// * the particular set of build contexts we consider (goEnvs)
// * whether the import path is valid.
func extractPackageMetas(ctx context.Context, modulePath, resolvedVersion string, contentDir fs.FS) (_ []*packageMeta, _ *godoc.ModuleInfo, _ []*internal.PackageVersionState, err error) {
	defer derrors.Wrap(&err, "extractPackages(ctx, %q, %q, r, d)", modulePath, resolvedVersion)
	ctx, span := trace.StartSpan(ctx, "fetch.extractPackages")
	defer span.End()
	defer func() {
		if e := recover(); e != nil {
			// The package processing code performs some sanity checks along the way.
			// None of the panics should occur, but if they do, we want to log them and
			// be able to find them. So, convert internal panics to internal errors here.
			err = fmt.Errorf("internal panic: %v\n\n%s", e, debug.Stack())
		}
	}()

	// The high-level approach is to split the processing of the zip file
	// into two phases:
	//
	// 	1. loop over all files, looking at file metadata only
	// 	2. process all files by reading their contents
	//
	// During phase 1, we populate the dirs map for each directory
	// that contains at least one .go file.

	var (
		// dirs is the set of directories with at least one .go file,
		// to be populated during phase 1 and used during phase 2.
		//
		// The map key is the directory path, with the modulePrefix trimmed.
		// The map value is a slice of all .go file paths, and no other files.
		dirs = make(map[string][]string)

		// modInfo contains all the module information a package in the module
		// needs to render its documentation, to be populated during phase 1
		// and used during phase 2.
		modInfo = &godoc.ModuleInfo{
			ModulePath:      modulePath,
			ResolvedVersion: resolvedVersion,
			ModulePackages:  make(map[string]bool),
		}

		// incompleteDirs tracks directories for which we have incomplete
		// information, due to a problem processing one of the go files contained
		// therein. We use this so that a single unprocessable package does not
		// prevent processing of other packages in the module.
		incompleteDirs       = make(map[string]bool)
		packageVersionStates = []*internal.PackageVersionState{}
	)

	// Phase 1.
	// Loop over zip files preemptively and check for problems
	// that can be detected by looking at metadata alone.
	// We'll be looking at file contents starting with phase 2 only,
	// only after we're sure this phase passed without errors.
	err = fs.WalkDir(contentDir, ".", func(pathname string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			// Skip directories.
			return nil
		}
		innerPath := path.Dir(pathname)
		if incompleteDirs[innerPath] {
			// We already know this directory cannot be processed, so skip.
			return nil
		}
		importPath := path.Join(modulePath, innerPath)
		if ignoredByGoTool(importPath) || isVendored(importPath) {
			// File is in a directory we're not looking to process at this time, so skip it.
			return nil
		}
		if !strings.HasSuffix(pathname, ".go") {
			// We care about .go files only.
			return nil
		}
		// It's possible to have a Go package in a directory that does not result in a valid import path.
		// That package cannot be imported, but that may be fine if it's a main package, intended to built
		// and run from that directory.
		// Example:  https://github.com/postmannen/go-learning/blob/master/concurrency/01-sending%20numbers%20and%20receving%20numbers%20from%20a%20channel/main.go
		// We're not set up to handle invalid import paths, so skip these packages.
		if err := module.CheckImportPath(importPath); err != nil {
			incompleteDirs[innerPath] = true
			packageVersionStates = append(packageVersionStates, &internal.PackageVersionState{
				ModulePath:  modulePath,
				PackagePath: importPath,
				Version:     resolvedVersion,
				Status:      derrors.ToStatus(derrors.PackageBadImportPath),
				Error:       err.Error(),
			})
			return nil
		}
		info, err := d.Info()
		if err != nil {
			return err
		}
		if info.Size() > MaxFileSize {
			incompleteDirs[innerPath] = true
			status := derrors.ToStatus(derrors.PackageMaxFileSizeLimitExceeded)
			err := fmt.Sprintf("Unable to process %s: file size %d exceeds max limit %d",
				pathname, info.Size(), MaxFileSize)
			packageVersionStates = append(packageVersionStates, &internal.PackageVersionState{
				ModulePath:  modulePath,
				PackagePath: importPath,
				Version:     resolvedVersion,
				Status:      status,
				Error:       err,
			})
			return nil
		}
		dirs[innerPath] = append(dirs[innerPath], pathname)
		if len(dirs) > maxPackagesPerModule {
			return fmt.Errorf("%d packages found in %q; exceeds limit %d for maxPackagePerModule: %w",
				len(dirs), modulePath, maxPackagesPerModule, derrors.ModuleTooLarge)
		}
		return nil
	})
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil, nil, fmt.Errorf("no files: %w", ErrModuleContainsNoPackages)
	}
	if err != nil {
		return nil, nil, nil, err
	}

	for pkgName := range dirs {
		modInfo.ModulePackages[path.Join(modulePath, pkgName)] = true
	}

	// Phase 2.
	// If we got this far, the file metadata was okay.
	// Start reading the file contents now to extract information
	// about Go packages.
	var pkgs []*packageMeta
	var mu sync.Mutex // guards pkgs, incompleteDirs, packageVersionStates
	var errgroup errgroup.Group
	for innerPath, goFiles := range dirs {
		innerPath, goFiles := innerPath, goFiles
		errgroup.Go(func() error {
			var addedPackage bool
			mu.Lock()
			incomplete := incompleteDirs[innerPath]
			mu.Unlock()
			if incomplete {
				// Something went wrong when processing this directory, so we skip.
				log.Infof(ctx, "Skipping %q because it is incomplete", innerPath)
				return nil
			}

			var (
				status error
				errMsg string
			)
			pkg, err := loadPackageMeta(ctx, contentDir, goFiles, innerPath, modInfo)
			if bpe := (*BadPackageError)(nil); errors.As(err, &bpe) {
				log.Infof(ctx, "Error loading %s: %v", innerPath, err)
				mu.Lock()
				incompleteDirs[innerPath] = true
				mu.Unlock()
				status = derrors.PackageInvalidContents
				errMsg = err.Error()
			} else if err != nil {
				return fmt.Errorf("unexpected error loading package: %v", err)
			}
			var pkgPath string
			if pkg == nil {
				// No package.
				if len(goFiles) > 0 {
					// There were go files, but no build contexts matched them.
					mu.Lock()
					incompleteDirs[innerPath] = true
					mu.Unlock()
					status = derrors.PackageBuildContextNotSupported
				}
				pkgPath = path.Join(modulePath, innerPath)
			} else {
				mu.Lock()
				pkgs = append(pkgs, pkg)
				mu.Unlock()
				addedPackage = true
				pkgPath = pkg.path
			}
			if !addedPackage {
				mu.Lock()
				packageVersionStates = append(packageVersionStates, &internal.PackageVersionState{
					ModulePath:  modulePath,
					PackagePath: pkgPath,
					Version:     resolvedVersion,
					Status:      derrors.ToStatus(status),
					Error:       errMsg,
				})
				mu.Unlock()
			}

			return nil
		})
	}

	if err := errgroup.Wait(); err != nil {
		return nil, nil, nil, err
	}

	if len(pkgs) == 0 {
		return nil, nil, packageVersionStates, ErrModuleContainsNoPackages
	}
	return pkgs, modInfo, packageVersionStates, nil
}

// ignoredByGoTool reports whether the given import path corresponds
// to a directory that would be ignored by the go tool.
//
// The logic of the go tool for ignoring directories is documented at
// https://golang.org/cmd/go/#hdr-Package_lists_and_patterns:
//
//	Directory and file names that begin with "." or "_" are ignored
//	by the go tool, as are directories named "testdata".
//
// However, even though `go list` and other commands that take package
// wildcards will ignore these, they can still be imported and used in
// working Go programs. We continue to ignore the "." and "testdata"
// cases, but we've seen valid Go packages with "_", so we accept those.
func ignoredByGoTool(importPath string) bool {
	for _, el := range strings.Split(importPath, "/") {
		if strings.HasPrefix(el, ".") || el == "testdata" {
			return true
		}
	}
	return false
}

// isVendored reports whether the given import path corresponds
// to a Go package that is inside a vendor directory.
//
// The logic for what is considered a vendor directory is documented at
// https://golang.org/cmd/go/#hdr-Vendor_Directories.
func isVendored(importPath string) bool {
	return strings.HasPrefix(importPath, "vendor/") ||
		strings.Contains(importPath, "/vendor/")
}
