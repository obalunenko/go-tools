//
// Copyright 2021, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gitlab

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

type (
	SnippetsServiceInterface interface {
		ListSnippets(opt *ListSnippetsOptions, options ...RequestOptionFunc) ([]*Snippet, *Response, error)
		GetSnippet(snippet int, options ...RequestOptionFunc) (*Snippet, *Response, error)
		SnippetContent(snippet int, options ...RequestOptionFunc) ([]byte, *Response, error)
		SnippetFileContent(snippet int, ref, filename string, options ...RequestOptionFunc) ([]byte, *Response, error)
		CreateSnippet(opt *CreateSnippetOptions, options ...RequestOptionFunc) (*Snippet, *Response, error)
		UpdateSnippet(snippet int, opt *UpdateSnippetOptions, options ...RequestOptionFunc) (*Snippet, *Response, error)
		DeleteSnippet(snippet int, options ...RequestOptionFunc) (*Response, error)
		ExploreSnippets(opt *ExploreSnippetsOptions, options ...RequestOptionFunc) ([]*Snippet, *Response, error)
		ListAllSnippets(opt *ListAllSnippetsOptions, options ...RequestOptionFunc) ([]*Snippet, *Response, error)
	}

	// SnippetsService handles communication with the snippets
	// related methods of the GitLab API.
	//
	// GitLab API docs: https://docs.gitlab.com/api/snippets/
	SnippetsService struct {
		client *Client
	}
)

var _ SnippetsServiceInterface = (*SnippetsService)(nil)

// Snippet represents a GitLab snippet.
//
// GitLab API docs: https://docs.gitlab.com/api/snippets/
type Snippet struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	FileName    string `json:"file_name"`
	Description string `json:"description"`
	Visibility  string `json:"visibility"`
	Author      struct {
		ID        int        `json:"id"`
		Username  string     `json:"username"`
		Email     string     `json:"email"`
		Name      string     `json:"name"`
		State     string     `json:"state"`
		CreatedAt *time.Time `json:"created_at"`
	} `json:"author"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedAt *time.Time `json:"created_at"`
	ProjectID int        `json:"project_id"`
	WebURL    string     `json:"web_url"`
	RawURL    string     `json:"raw_url"`
	Files     []struct {
		Path   string `json:"path"`
		RawURL string `json:"raw_url"`
	} `json:"files"`
	RepositoryStorage string `json:"repository_storage"`
}

func (s Snippet) String() string {
	return Stringify(s)
}

// ListSnippetsOptions represents the available ListSnippets() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/snippets/#list-all-snippets-for-current-user
type ListSnippetsOptions ListOptions

// ListSnippets gets a list of snippets.
//
// GitLab API docs:
// https://docs.gitlab.com/api/snippets/#list-all-snippets-for-current-user
func (s *SnippetsService) ListSnippets(opt *ListSnippetsOptions, options ...RequestOptionFunc) ([]*Snippet, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "snippets", opt, options)
	if err != nil {
		return nil, nil, err
	}

	var ps []*Snippet
	resp, err := s.client.Do(req, &ps)
	if err != nil {
		return nil, resp, err
	}

	return ps, resp, nil
}

// GetSnippet gets a single snippet
//
// GitLab API docs:
// https://docs.gitlab.com/api/snippets/#get-a-single-snippet
func (s *SnippetsService) GetSnippet(snippet int, options ...RequestOptionFunc) (*Snippet, *Response, error) {
	u := fmt.Sprintf("snippets/%d", snippet)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	ps := new(Snippet)
	resp, err := s.client.Do(req, ps)
	if err != nil {
		return nil, resp, err
	}

	return ps, resp, nil
}

// SnippetContent gets a single snippet’s raw contents.
//
// GitLab API docs:
// https://docs.gitlab.com/api/snippets/#single-snippet-contents
func (s *SnippetsService) SnippetContent(snippet int, options ...RequestOptionFunc) ([]byte, *Response, error) {
	u := fmt.Sprintf("snippets/%d/raw", snippet)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var b bytes.Buffer
	resp, err := s.client.Do(req, &b)
	if err != nil {
		return nil, resp, err
	}

	return b.Bytes(), resp, err
}

// SnippetFileContent returns the raw file content as plain text.
//
// GitLab API docs:
// https://docs.gitlab.com/api/snippets/#snippet-repository-file-content
func (s *SnippetsService) SnippetFileContent(snippet int, ref, filename string, options ...RequestOptionFunc) ([]byte, *Response, error) {
	filepath := PathEscape(filename)
	u := fmt.Sprintf("snippets/%d/files/%s/%s/raw", snippet, ref, filepath)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var b bytes.Buffer
	resp, err := s.client.Do(req, &b)
	if err != nil {
		return nil, resp, err
	}

	return b.Bytes(), resp, err
}

// CreateSnippetFileOptions represents the create snippet file options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/snippets/#create-new-snippet
type CreateSnippetFileOptions struct {
	FilePath *string `url:"file_path,omitempty" json:"file_path,omitempty"`
	Content  *string `url:"content,omitempty" json:"content,omitempty"`
}

// CreateSnippetOptions represents the available CreateSnippet() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/snippets/#create-new-snippet
type CreateSnippetOptions struct {
	Title       *string                      `url:"title,omitempty" json:"title,omitempty"`
	FileName    *string                      `url:"file_name,omitempty" json:"file_name,omitempty"`
	Description *string                      `url:"description,omitempty" json:"description,omitempty"`
	Content     *string                      `url:"content,omitempty" json:"content,omitempty"`
	Visibility  *VisibilityValue             `url:"visibility,omitempty" json:"visibility,omitempty"`
	Files       *[]*CreateSnippetFileOptions `url:"files,omitempty" json:"files,omitempty"`
}

// CreateSnippet creates a new snippet. The user must have permission
// to create new snippets.
//
// GitLab API docs:
// https://docs.gitlab.com/api/snippets/#create-new-snippet
func (s *SnippetsService) CreateSnippet(opt *CreateSnippetOptions, options ...RequestOptionFunc) (*Snippet, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "snippets", opt, options)
	if err != nil {
		return nil, nil, err
	}

	ps := new(Snippet)
	resp, err := s.client.Do(req, ps)
	if err != nil {
		return nil, resp, err
	}

	return ps, resp, nil
}

// UpdateSnippetFileOptions represents the update snippet file options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/snippets/#update-snippet
type UpdateSnippetFileOptions struct {
	Action       *string `url:"action,omitempty" json:"action,omitempty"`
	FilePath     *string `url:"file_path,omitempty" json:"file_path,omitempty"`
	Content      *string `url:"content,omitempty" json:"content,omitempty"`
	PreviousPath *string `url:"previous_path,omitempty" json:"previous_path,omitempty"`
}

// UpdateSnippetOptions represents the available UpdateSnippet() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/snippets/#update-snippet
type UpdateSnippetOptions struct {
	Title       *string                      `url:"title,omitempty" json:"title,omitempty"`
	FileName    *string                      `url:"file_name,omitempty" json:"file_name,omitempty"`
	Description *string                      `url:"description,omitempty" json:"description,omitempty"`
	Content     *string                      `url:"content,omitempty" json:"content,omitempty"`
	Visibility  *VisibilityValue             `url:"visibility,omitempty" json:"visibility,omitempty"`
	Files       *[]*UpdateSnippetFileOptions `url:"files,omitempty" json:"files,omitempty"`
}

// UpdateSnippet updates an existing snippet. The user must have
// permission to change an existing snippet.
//
// GitLab API docs:
// https://docs.gitlab.com/api/snippets/#update-snippet
func (s *SnippetsService) UpdateSnippet(snippet int, opt *UpdateSnippetOptions, options ...RequestOptionFunc) (*Snippet, *Response, error) {
	u := fmt.Sprintf("snippets/%d", snippet)

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	ps := new(Snippet)
	resp, err := s.client.Do(req, ps)
	if err != nil {
		return nil, resp, err
	}

	return ps, resp, nil
}

// DeleteSnippet deletes an existing snippet. This is an idempotent
// function and deleting a non-existent snippet still returns a 200 OK status
// code.
//
// GitLab API docs:
// https://docs.gitlab.com/api/snippets/#delete-snippet
func (s *SnippetsService) DeleteSnippet(snippet int, options ...RequestOptionFunc) (*Response, error) {
	u := fmt.Sprintf("snippets/%d", snippet)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ExploreSnippetsOptions represents the available ExploreSnippets() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/snippets/#list-all-public-snippets
type ExploreSnippetsOptions ListOptions

// ExploreSnippets gets the list of public snippets.
//
// GitLab API docs:
// https://docs.gitlab.com/api/snippets/#list-all-public-snippets
func (s *SnippetsService) ExploreSnippets(opt *ExploreSnippetsOptions, options ...RequestOptionFunc) ([]*Snippet, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "snippets/public", opt, options)
	if err != nil {
		return nil, nil, err
	}

	var ps []*Snippet
	resp, err := s.client.Do(req, &ps)
	if err != nil {
		return nil, resp, err
	}

	return ps, resp, nil
}

// ListAllSnippetsOptions represents the available ListAllSnippets() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/snippets/#list-all-snippets
type ListAllSnippetsOptions struct {
	ListOptions
	CreatedAfter      *ISOTime `url:"created_after,omitempty" json:"created_after,omitempty"`
	CreatedBefore     *ISOTime `url:"created_before,omitempty" json:"created_before,omitempty"`
	RepositoryStorage *string  `url:"repository_storage,omitempty" json:"repository_storage,omitempty"`
}

// ListAllSnippets gets all snippets the current user has access to.
//
// GitLab API docs:
// https://docs.gitlab.com/api/snippets/#list-all-snippets
func (s *SnippetsService) ListAllSnippets(opt *ListAllSnippetsOptions, options ...RequestOptionFunc) ([]*Snippet, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "snippets/all", opt, options)
	if err != nil {
		return nil, nil, err
	}

	var ps []*Snippet
	resp, err := s.client.Do(req, &ps)
	if err != nil {
		return nil, resp, err
	}

	return ps, resp, nil
}
