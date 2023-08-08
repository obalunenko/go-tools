// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stdlib

import (
	"bytes"
	"context"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"

	"golang.org/x/pkgsite/internal/derrors"
	"golang.org/x/pkgsite/internal/testing/testhelper"
	"golang.org/x/pkgsite/internal/version"
)

// A goRepo represents a git repo holding the Go standard library.
type goRepo interface {
	// Clone the repo at the given version to the directory.
	clone(ctx context.Context, version string, toDirectory string) (refName string, err error)

	// Return all the refs of the repo.
	refs(ctx context.Context) ([]ref, error)
}

type remoteGoRepo struct{}

func (remoteGoRepo) clone(ctx context.Context, v, directory string) (refName string, err error) {
	defer derrors.Wrap(&err, "remoteGoRepo.clone(%q)", v)

	refName, err = refNameForVersion(v)
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(directory, 0777); err != nil {
		return "", err
	}
	cmd := exec.CommandContext(ctx, "git", "init")
	cmd.Dir = directory
	if err := cmd.Run(); err != nil {
		return "", err
	}
	cmd = exec.CommandContext(ctx, "git", "fetch", "--depth=1", "--", GoRepoURL, refName+":main")
	cmd.Dir = directory
	if b, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("running git fetch: %v: %s", err, b)
	}
	cmd = exec.CommandContext(ctx, "git", "checkout", "main")
	cmd.Dir = directory
	if b, err := cmd.CombinedOutput(); err != nil {
		return "", fmt.Errorf("running git checkout: %v: %s", err, b)
	}
	return refName, nil
}

type ref struct {
	hash string
	name string
}

func (remoteGoRepo) refs(ctx context.Context) ([]ref, error) {
	cmd := exec.CommandContext(ctx, "git", "ls-remote", "--", GoRepoURL)
	b, err := cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("running git ls-remote: %v: %s", err, ee.Stderr)
		}
		return nil, fmt.Errorf("running git ls-remote: %v", err)
	}
	return gitOutputToRefs(b)
}

func gitOutputToRefs(b []byte) ([]ref, error) {
	var refs []ref
	b = bytes.TrimSpace(b)
	for _, line := range bytes.Split(b, []byte("\n")) {
		fields := bytes.Fields(line)
		if len(fields) != 2 {
			return nil, fmt.Errorf("invalid line in output from git ls-remote: %q: should have two fields", line)
		}
		refs = append(refs, ref{hash: string(fields[0]), name: string(fields[1])})
	}
	return refs, nil
}

type localGoRepo struct {
	path string
}

func newLocalGoRepo(path string) *localGoRepo {
	return &localGoRepo{
		path: path,
	}
}

func (g *localGoRepo) refs(ctx context.Context) (refs []ref, err error) {
	defer derrors.Wrap(&err, "localGoRepo(%s).refs", g.path)

	cmd := exec.CommandContext(ctx, "git", "show-ref")
	cmd.Dir = g.path
	b, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("running git show-ref: %v", err)
	}
	return gitOutputToRefs(b)
}

func (g *localGoRepo) clone(ctx context.Context, v, directory string) (refName string, err error) {
	return "", nil
}

type testGoRepo struct {
}

func (t *testGoRepo) clone(ctx context.Context, v, directory string) (refName string, err error) {
	defer derrors.Wrap(&err, "testGoRepo.clone(%q)", v)
	if v == TestMasterVersion {
		v = version.Master
	}
	if v == TestDevFuzzVersion {
		v = DevFuzz
	}
	cmd := exec.CommandContext(ctx, "git", "init")
	cmd.Dir = directory
	if err := cmd.Run(); err != nil {
		return "", err
	}
	testdatadir := filepath.Join(testhelper.TestDataPath("testdata"), v)
	err = filepath.Walk(testdatadir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(testdatadir, path)
		if err != nil {
			return err
		}
		dstpath := filepath.Join(directory, rel)
		if info.Mode().IsDir() {
			os.MkdirAll(dstpath, 0777)
			return nil
		}
		b, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("reading %q: %v", path, err)
		}
		os.WriteFile(dstpath, b, 0666)
		cmd := exec.CommandContext(ctx, "git", "add", "--", dstpath)
		cmd.Dir = directory
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("running git add: %v", err)
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	cmd = exec.CommandContext(ctx, "git", "commit", "--allow-empty-message", "--author=Joe Random <joe@example.com>",
		"--message=")
	cmd.Dir = directory
	commitTime := fmt.Sprintf("%v +0000", TestCommitTime.Unix())
	name := "Joe Random"
	email := "joe@example.com"
	cmd.Env = append(cmd.Environ(), []string{
		"GIT_COMMITTER_NAME=" + name, "GIT_AUTHOR_NAME=" + name,
		"GIT_COMMITTER_EMAIL=" + email, "GIT_AUTHOR_EMAIL=" + email,
		"GIT_COMMITTER_DATE=" + commitTime, "GIT_AUTHOR_DATE=" + commitTime}...)
	if err := cmd.Run(); err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			return "", fmt.Errorf("running git commit: %v: %s", err, ee.Stderr)
		}
		return "", fmt.Errorf("running git commit: %v", err)
	}
	return "HEAD", nil
}

// References used for Versions during testing.
var testRefs = []string{
	// stdlib versions
	"refs/tags/go1.2.1",
	"refs/tags/go1.3.2",
	"refs/tags/go1.4.2",
	"refs/tags/go1.4.3",
	"refs/tags/go1.6",
	"refs/tags/go1.6.3",
	"refs/tags/go1.6beta1",
	"refs/tags/go1.8",
	"refs/tags/go1.8rc2",
	"refs/tags/go1.9rc1",
	"refs/tags/go1.11",
	"refs/tags/go1.12",
	"refs/tags/go1.12.1",
	"refs/tags/go1.12.5",
	"refs/tags/go1.12.9",
	"refs/tags/go1.13",
	"refs/tags/go1.13beta1",
	"refs/tags/go1.14.6",
	"refs/tags/go1.21.0",
	"refs/heads/dev.fuzz",
	"refs/heads/master",
	// other tags
	"refs/changes/56/93156/13",
	"refs/tags/release.r59",
	"refs/tags/weekly.2011-04-13",
}

func (t *testGoRepo) refs(ctx context.Context) ([]ref, error) {
	var rs []ref
	for _, r := range testRefs {
		// Only the name is ever used, so the referent can be empty.
		rs = append(rs, ref{name: r})
	}
	return rs, nil
}
