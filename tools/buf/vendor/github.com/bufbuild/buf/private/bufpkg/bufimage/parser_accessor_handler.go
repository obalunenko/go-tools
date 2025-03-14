// Copyright 2020-2025 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bufimage

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"sync"

	"github.com/bufbuild/buf/private/bufpkg/bufmodule"
	"github.com/bufbuild/buf/private/bufpkg/bufparse"
	"github.com/bufbuild/buf/private/gen/data/datawkt"
	"github.com/google/uuid"
)

type parserAccessorHandler struct {
	ctx                context.Context
	moduleReadBucket   bufmodule.ModuleReadBucket
	pathToExternalPath map[string]string
	pathToLocalPath    map[string]string
	nonImportPaths     map[string]struct{}
	pathToFullName     map[string]bufparse.FullName
	pathToCommitID     map[string]uuid.UUID
	lock               sync.RWMutex
}

func newParserAccessorHandler(
	ctx context.Context,
	moduleReadBucket bufmodule.ModuleReadBucket,
) *parserAccessorHandler {
	return &parserAccessorHandler{
		ctx:                ctx,
		moduleReadBucket:   moduleReadBucket,
		pathToExternalPath: make(map[string]string),
		pathToLocalPath:    make(map[string]string),
		nonImportPaths:     make(map[string]struct{}),
		pathToFullName:     make(map[string]bufparse.FullName),
		pathToCommitID:     make(map[string]uuid.UUID),
	}
}

// Open opens the given path, and tracks the external path and import status.
//
// This function can be used as the accessor function for a protocompile.SourceResolver.
func (p *parserAccessorHandler) Open(path string) (_ io.ReadCloser, retErr error) {
	moduleFile, moduleErr := p.moduleReadBucket.GetFile(p.ctx, path)
	if moduleErr != nil {
		if !errors.Is(moduleErr, fs.ErrNotExist) {
			return nil, moduleErr
		}
		if wktModuleFile, wktErr := datawkt.ReadBucket.Get(p.ctx, path); wktErr == nil {
			if wktModuleFile.Path() != path {
				// this should never happen, but just in case
				return nil, fmt.Errorf("parser accessor requested path %q but got %q", path, wktModuleFile.Path())
			}
			if err := p.addPath(path, path, "", nil, uuid.Nil); err != nil {
				return nil, err
			}
			return wktModuleFile, nil
		}
		return nil, moduleErr
	}
	defer func() {
		if retErr != nil {
			retErr = errors.Join(retErr, moduleFile.Close())
		}
	}()
	if moduleFile.Path() != path {
		// this should never happen, but just in case
		return nil, fmt.Errorf("parser accessor requested path %q but got %q", path, moduleFile.Path())
	}
	if err := p.addPath(
		path,
		moduleFile.ExternalPath(),
		moduleFile.LocalPath(),
		moduleFile.Module().FullName(),
		moduleFile.Module().CommitID(),
	); err != nil {
		return nil, err
	}
	return moduleFile, nil
}

// ExternalPath returns the external path for the input path.
//
// Returns the input path if the external path is not known.
func (p *parserAccessorHandler) ExternalPath(path string) string {
	p.lock.RLock()
	defer p.lock.RUnlock()
	if externalPath := p.pathToExternalPath[path]; externalPath != "" {
		return externalPath
	}
	return path
}

// LocalPath returns the local path for the input path if present.
func (p *parserAccessorHandler) LocalPath(path string) string {
	p.lock.RLock()
	defer p.lock.RUnlock()
	return p.pathToLocalPath[path]
}

// FullName returns nil if not available.
func (p *parserAccessorHandler) FullName(path string) bufparse.FullName {
	p.lock.RLock()
	defer p.lock.RUnlock()
	return p.pathToFullName[path] // nil is a valid value.
}

// CommitID returns empty if not available.
func (p *parserAccessorHandler) CommitID(path string) uuid.UUID {
	p.lock.RLock()
	defer p.lock.RUnlock()
	return p.pathToCommitID[path] // empty is a valid value.
}

func (p *parserAccessorHandler) addPath(
	path string,
	externalPath string,
	localPath string,
	moduleFullName bufparse.FullName,
	commitID uuid.UUID,
) error {
	p.lock.Lock()
	defer p.lock.Unlock()
	existingExternalPath, ok := p.pathToExternalPath[path]
	if ok {
		if existingExternalPath != externalPath {
			return fmt.Errorf("parser accessor had external paths %q and %q for path %q", existingExternalPath, externalPath, path)
		}
	} else {
		p.pathToExternalPath[path] = externalPath
	}
	if localPath != "" {
		existingLocalPath, ok := p.pathToLocalPath[path]
		if ok {
			if existingLocalPath != localPath {
				return fmt.Errorf("parser accessor had local paths %q and %q for path %q", existingLocalPath, localPath, path)
			}
		} else {
			p.pathToLocalPath[path] = localPath
		}
	}
	if moduleFullName != nil {
		p.pathToFullName[path] = moduleFullName
	}
	if commitID != uuid.Nil {
		p.pathToCommitID[path] = commitID
	}
	return nil
}
