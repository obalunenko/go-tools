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

package bufprotosource

import "google.golang.org/protobuf/reflect/protoreflect"

type mergeCommentLocation struct {
	base            Location
	delegate        Location
	baseHasComments bool
}

func newMergeCommentLocation(base Location, delegate Location) *mergeCommentLocation {
	return &mergeCommentLocation{
		base:            base,
		delegate:        delegate,
		baseHasComments: base.LeadingComments() != "" || base.TrailingComments() != "" || len(base.LeadingDetachedComments()) > 0,
	}
}

func (l *mergeCommentLocation) FilePath() string {
	return l.base.FilePath()
}

func (l *mergeCommentLocation) StartLine() int {
	return l.base.StartLine()
}

func (l *mergeCommentLocation) StartColumn() int {
	return l.base.StartColumn()
}

func (l *mergeCommentLocation) EndLine() int {
	return l.base.EndLine()
}

func (l *mergeCommentLocation) EndColumn() int {
	return l.base.EndColumn()
}

func (l *mergeCommentLocation) LeadingComments() string {
	if l.baseHasComments {
		return l.base.LeadingComments()
	}
	return l.delegate.LeadingComments()
}

func (l *mergeCommentLocation) TrailingComments() string {
	if l.baseHasComments {
		return l.base.TrailingComments()
	}
	return l.delegate.TrailingComments()
}

func (l *mergeCommentLocation) LeadingDetachedComments() []string {
	if l.baseHasComments {
		return l.base.LeadingDetachedComments()
	}
	return l.delegate.LeadingDetachedComments()
}

func (l *mergeCommentLocation) SourcePath() protoreflect.SourcePath {
	return l.base.SourcePath()
}
