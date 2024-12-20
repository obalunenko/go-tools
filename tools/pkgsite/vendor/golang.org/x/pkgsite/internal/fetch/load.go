// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fetch provides a way to fetch modules from a proxy.
package fetch

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path"
	"sort"
	"strings"

	"golang.org/x/pkgsite/internal"
	"golang.org/x/pkgsite/internal/derrors"
	"golang.org/x/pkgsite/internal/godoc"
	"golang.org/x/pkgsite/internal/source"
	"golang.org/x/pkgsite/internal/stdlib"
	"golang.org/x/pkgsite/internal/trace"
)

// BadPackageError represents an error loading a package
// because its contents do not make up a valid package.
//
// This can happen, for example, if the .go files fail
// to parse or declare different package names.
type BadPackageError struct {
	Err error // Not nil.
}

func (bpe *BadPackageError) Error() string { return bpe.Err.Error() }

// loadPackage loads a Go package by calling loadPackageWithBuildContext, trying
// several build contexts in turn. It returns a goPackage with documentation
// information for each build context that results in a valid package, in the
// same order that the build contexts are listed. If none of them result in a
// package, then loadPackage returns nil, nil.
//
// If a package is fine except that its documentation is too large, loadPackage
// returns a goPackage whose err field is a non-nil error with godoc.ErrTooLarge in its chain.
func loadPackage(ctx context.Context, contentDir fs.FS, goFilePaths []string, innerPath string,
	sourceInfo *source.Info, modInfo *godoc.ModuleInfo) (_ *goPackage, err error) {
	defer derrors.Wrap(&err, "loadPackage(ctx, zipGoFiles, %q, sourceInfo, modInfo)", innerPath)
	ctx, span := trace.StartSpan(ctx, "fetch.loadPackage")
	defer span.End()

	// Make a map with all the zip file contents.
	files := make(map[string][]byte)
	for _, p := range goFilePaths {
		_, name := path.Split(p)
		b, err := readFSFile(contentDir, p, MaxFileSize)
		if err != nil {
			return nil, err
		}
		files[name] = b
	}

	modulePath := modInfo.ModulePath
	importPath := path.Join(modulePath, innerPath)
	if modulePath == stdlib.ModulePath {
		importPath = innerPath
	}
	v1path := internal.V1Path(importPath, modulePath)

	var pkg *goPackage
	// Parse the package for each build context.
	// The documentation is determined by the set of matching files, so keep
	// track of those to avoid duplication.
	docsByFiles := map[string]*internal.Documentation{}
	for _, bc := range internal.BuildContexts {
		mfiles, err := matchingFiles(bc.GOOS, bc.GOARCH, files)
		if err != nil {
			return nil, err
		}
		filesKey := mapKeyForFiles(mfiles)
		if doc := docsByFiles[filesKey]; doc != nil {
			// We have seen this set of files before.
			// loadPackageWithBuildContext will produce the same outputs,
			// so don't bother calling it. Just copy the doc.
			doc2 := *doc
			doc2.GOOS = bc.GOOS
			doc2.GOARCH = bc.GOARCH
			doc2.API = nil
			for _, s := range doc.API {
				s2 := *s
				s2.Children = nil
				s2.GOOS = bc.GOOS
				s2.GOARCH = bc.GOARCH
				s2.Children = append(s2.Children, s.Children...)
				doc2.API = append(doc2.API, &s2)
			}
			pkg.docs = append(pkg.docs, &doc2)
			continue
		}
		name, imports, synopsis, source, api, err := loadPackageForBuildContext(ctx,
			mfiles, innerPath, sourceInfo, modInfo)
		for _, s := range api {
			s.GOOS = bc.GOOS
			s.GOARCH = bc.GOARCH
		}

		switch {
		case errors.Is(err, derrors.NotFound):
			// No package for this build context.
			continue
		case errors.As(err, new(*BadPackageError)):
			// This build context was bad, but maybe others aren't.
			continue
		case errors.Is(err, godoc.ErrTooLarge):
			// The doc for this build context is too large. To keep things
			// simple, return a single package with this error that will be used
			// for all build contexts, and ignore the others.
			return &goPackage{
				err:     err,
				path:    importPath,
				v1path:  v1path,
				name:    name,
				imports: imports,
				docs: []*internal.Documentation{{
					GOOS:     internal.All,
					GOARCH:   internal.All,
					Synopsis: synopsis,
					Source:   source,
					API:      api,
				}},
			}, nil
		case err != nil:
			// Serious error. Fail.
			return nil, err
		default:
			// No error.
			if pkg == nil {
				pkg = &goPackage{
					path:    importPath,
					v1path:  v1path,
					name:    name,
					imports: imports, // Use the imports from the first successful build context.
				}
			}
			// All the build contexts should use the same package name. Although
			// it's technically legal for different build tags to result in different
			// package names, it's not something we support.
			if name != pkg.name {
				return nil, &BadPackageError{
					Err: fmt.Errorf("more than one package name (%q and %q)", pkg.name, name),
				}
			}
			doc := &internal.Documentation{
				GOOS:     bc.GOOS,
				GOARCH:   bc.GOARCH,
				Synopsis: synopsis,
				Source:   source,
				API:      api,
			}
			docsByFiles[filesKey] = doc
			pkg.docs = append(pkg.docs, doc)
		}
	}
	// If all the build contexts succeeded and had the same set of files, then
	// assume that the package doc is valid for all build contexts. Represent
	// this with a single Documentation whose GOOS and GOARCH are both "all".
	if len(docsByFiles) == 1 && len(pkg.docs) == len(internal.BuildContexts) {
		pkg.docs = pkg.docs[:1]
		pkg.docs[0].GOOS = internal.All
		pkg.docs[0].GOARCH = internal.All
		for _, s := range pkg.docs[0].API {
			s.GOOS = internal.All
			s.GOARCH = internal.All
		}
	}
	return pkg, nil
}

// loadPackageMeta loads only the parts of a package that are needed to load a
// packageMeta.
func loadPackageMeta(ctx context.Context, contentDir fs.FS, goFilePaths []string, innerPath string, modInfo *godoc.ModuleInfo) (_ *packageMeta, err error) {
	defer derrors.Wrap(&err, "loadPackageMeta(ctx, zipGoFiles, %q, sourceInfo, modInfo)", innerPath)

	// Make a map with all the zip file contents.
	files := make(map[string][]byte)
	for _, p := range goFilePaths {
		name := path.Base(p)
		b, err := readFSFile(contentDir, p, MaxFileSize)
		if err != nil {
			return nil, err
		}
		files[name] = b
	}

	modulePath := modInfo.ModulePath
	importPath := path.Join(modulePath, innerPath)
	if modulePath == stdlib.ModulePath {
		importPath = innerPath
	}

	var pkg *packageMeta
	// Try to load the package name for each build context. We're okay
	// as long as all the build contexts that successfully loadPackageName agree
	// on the package name.
	// TODO(matloob): See if we can rewrite this so each file needs to be loaded
	// only once. What we probably want to do is map each file to the package name
	// in the file and then run the logic in loadPackageName on the collection of
	// package name values.
	for _, bc := range internal.BuildContexts {
		mfiles, err := matchingFiles(bc.GOOS, bc.GOARCH, files)
		if err != nil {
			return nil, err
		}
		name, err := loadPackageName(innerPath, mfiles)
		switch {
		case errors.Is(err, derrors.NotFound):
			// No package for this build context.
			continue
		case errors.As(err, new(*BadPackageError)):
			// This build context was bad, but maybe others aren't.
			continue
		case err != nil:
			// Serious error. Fail.
			return nil, err
		default:
			// No error.
			if pkg == nil {
				pkg = &packageMeta{
					path: importPath,
					name: name,
				}
			}
			// All the build contexts should use the same package name. Although
			// it's technically legal for different build tags to result in different
			// package names, it's not something we support.
			if name != pkg.name {
				return nil, &BadPackageError{
					Err: fmt.Errorf("more than one package name (%q and %q)", pkg.name, name),
				}
			}
		}
	}

	return pkg, nil
}

// mapKeyForFiles generates a value that corresponds to the given set of file
// names and can be used as a map key.
// It assumes the filenames do not contain spaces.
func mapKeyForFiles(files map[string][]byte) string {
	var names []string
	for n := range files {
		names = append(names, n)
	}
	sort.Strings(names)
	return strings.Join(names, " ")
}

// httpPost allows package fetch tests to stub out playground URL fetches.
var httpPost = http.Post

// loadPackageForBuildContext loads a Go package made of .go files in
// files, which should match some build context.
// modulePath is stdlib.ModulePath for the Go standard library and the
// module path for all other modules. innerPath is the path of the Go package
// directory relative to the module root. The files argument must contain only
// .go files that have been verified to be of reasonable size and that match
// the build context.
//
// It returns the package name, list of imports, the package synopsis, and the
// serialized source (AST) for the package.
//
// It returns an error with NotFound in its chain if the directory doesn't
// contain a Go package or all .go files have been excluded by constraints. A
// *BadPackageError error is returned if the directory contains .go files but do
// not make up a valid package.
//
// If it returns an error with ErrTooLarge in its chain, the other return values
// are still valid.
func loadPackageForBuildContext(ctx context.Context, files map[string][]byte, innerPath string, sourceInfo *source.Info, modInfo *godoc.ModuleInfo) (
	name string, imports []string, synopsis string, source []byte, api []*internal.Symbol, err error) {
	modulePath := modInfo.ModulePath
	defer derrors.Wrap(&err, "loadPackageWithBuildContext(files, %q, %q, %+v)", innerPath, modulePath, sourceInfo)

	packageName, goFiles, fset, err := loadFilesWithBuildContext(innerPath, files)
	if err != nil {
		return "", nil, "", nil, nil, err
	}
	docPkg := godoc.NewPackage(fset, modInfo.ModulePackages)
	for _, pf := range goFiles {
		removeNodes := true
		// Don't strip the seemingly unexported functions from the builtin package;
		// they are actually Go builtins like make, new, etc.
		if modulePath == stdlib.ModulePath && innerPath == "builtin" {
			removeNodes = false
		}
		docPkg.AddFile(pf, removeNodes)
	}

	// Encode first, because Render messes with the AST.
	src, err := docPkg.Encode(ctx)
	if err != nil {
		return "", nil, "", nil, nil, err
	}

	synopsis, imports, api, err = docPkg.DocInfo(ctx, innerPath, sourceInfo, modInfo)
	if err != nil {
		return "", nil, "", nil, nil, err
	}
	return packageName, imports, synopsis, src, api, err
}

// loadFilesWithBuildContext loads all the given Go files at innerPath. It
// returns the package name as it occurs in the source, a map of the ASTs of all
// the Go files, and the token.FileSet used for parsing.
// If there are no non-test Go files, it returns a NotFound error.
func loadFilesWithBuildContext(innerPath string, files map[string][]byte) (pkgName string, fileMap map[string]*ast.File, _ *token.FileSet, _ error) {
	// Parse .go files and add them to the goFiles slice.
	var (
		fset            = token.NewFileSet()
		goFiles         = make(map[string]*ast.File)
		numNonTestFiles int
		packageName     string
		packageNameFile string // Name of file where packageName came from.
	)
	for name, b := range files {
		pf, err := parser.ParseFile(fset, name, b, parser.ParseComments)
		if err != nil {
			if pf == nil {
				return "", nil, nil, fmt.Errorf("internal error: the source couldn't be read: %v", err)
			}
			return "", nil, nil, &BadPackageError{Err: err}
		}
		// Remember all files, including test files for their examples.
		goFiles[name] = pf
		if strings.HasSuffix(name, "_test.go") {
			continue
		}
		// Keep track of the number of non-test files to check that the package name is the same.
		// and also because a directory with only test files doesn't count as a
		// Go package.
		numNonTestFiles++
		if numNonTestFiles == 1 {
			packageName = pf.Name.Name
			packageNameFile = name
		} else if pf.Name.Name != packageName {
			return "", nil, nil, &BadPackageError{Err: &build.MultiplePackageError{
				Dir:      innerPath,
				Packages: []string{packageName, pf.Name.Name},
				Files:    []string{packageNameFile, name},
			}}
		}
	}
	if numNonTestFiles == 0 {
		// This directory doesn't contain a package, or at least not one
		// that matches this build context.
		return "", nil, nil, derrors.NotFound
	}
	return packageName, goFiles, fset, nil
}

// loadPackageName returns the package name from the files as it occurs in the source.
// If there are no non-test Go files, it returns a NotFound error.
func loadPackageName(innerPath string, files map[string][]byte) (pkgName string, _ error) {
	// Parse .go files and add them to the goFiles slice.
	var (
		fset            = token.NewFileSet()
		numNonTestFiles int
		packageName     string
		packageNameFile string // Name of file where packageName came from.
	)
	for name, b := range files {
		if strings.HasSuffix(name, "_test.go") {
			continue
		}
		pf, err := parser.ParseFile(fset, name, b, parser.PackageClauseOnly)
		if err != nil {
			if pf == nil {
				return "", fmt.Errorf("internal error: the source couldn't be read: %v", err)
			}
			return "", &BadPackageError{Err: err}
		}
		numNonTestFiles++
		if numNonTestFiles == 1 {
			packageName = pf.Name.Name
			packageNameFile = name
		} else if pf.Name.Name != packageName {
			return "", &BadPackageError{Err: &build.MultiplePackageError{
				Dir:      innerPath,
				Packages: []string{packageName, pf.Name.Name},
				Files:    []string{packageNameFile, name},
			}}
		}
	}
	if numNonTestFiles == 0 {
		// This directory doesn't contain a package, or at least not one
		// that matches this build context.
		return "", derrors.NotFound
	}
	return packageName, nil
}

// matchingFiles returns a map from file names to their contents, read from zipGoFiles.
// It includes only those files that match the build context determined by goos and goarch.
func matchingFiles(goos, goarch string, allFiles map[string][]byte) (matchedFiles map[string][]byte, err error) {
	defer derrors.Wrap(&err, "matchingFiles(%q, %q, zipGoFiles)", goos, goarch)

	// bctx is used to make decisions about which of the .go files are included
	// by build constraints.
	bctx := &build.Context{
		GOOS:        goos,
		GOARCH:      goarch,
		CgoEnabled:  true,
		Compiler:    build.Default.Compiler,
		ReleaseTags: build.Default.ReleaseTags,

		JoinPath: path.Join,
		OpenFile: func(name string) (io.ReadCloser, error) {
			return io.NopCloser(bytes.NewReader(allFiles[name])), nil
		},

		// If left nil, the default implementations of these read from disk,
		// which we do not want. None of these functions should be used
		// inside this function; it would be an internal error if they are.
		// Set them to non-nil values to catch if that happens.
		SplitPathList: func(string) []string { panic("internal error: unexpected call to SplitPathList") },
		IsAbsPath:     func(string) bool { panic("internal error: unexpected call to IsAbsPath") },
		IsDir:         func(string) bool { panic("internal error: unexpected call to IsDir") },
		HasSubdir:     func(string, string) (string, bool) { panic("internal error: unexpected call to HasSubdir") },
		ReadDir:       func(string) ([]os.FileInfo, error) { panic("internal error: unexpected call to ReadDir") },
	}

	// Copy the input map so we don't modify it.
	matchedFiles = map[string][]byte{}
	for n, c := range allFiles {
		matchedFiles[n] = c
	}
	for name := range allFiles {
		match, err := bctx.MatchFile(".", name) // This will access the file we just added to files map above.
		if err != nil {
			return nil, &BadPackageError{Err: fmt.Errorf(`bctx.MatchFile(".", %q): %w`, name, err)}
		}
		if !match {
			delete(matchedFiles, name)
		}
	}
	return matchedFiles, nil
}

// readFSFile reads up to limit bytes from path in fsys.
func readFSFile(fsys fs.FS, path string, limit int64) (_ []byte, err error) {
	defer derrors.Add(&err, "readFSFile(%q)", path)
	f, err := fsys.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(io.LimitReader(f, limit))
}
