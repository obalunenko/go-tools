// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package frontend

import (
	"net/http"
	"net/url"
	"strconv"

	"golang.org/x/pkgsite/internal/log"
)

// pagination holds information related to paginated display. It is intended to
// be part of a view model struct.
//
// Given a sequence of results with offsets 0, 1, 2 ... (typically from a
// database query), we paginate it by dividing it into numbered pages
// 1, 2, 3, .... Each page except possibly the last has the same number of results.
type pagination struct {
	baseURL      *url.URL // URL common to all pages
	Limit        int      // the number of results requested on a page
	DefaultLimit int      // the default search limit
	MaxLimit     int      // the maximum number of results allowed for any request
	ResultCount  int      // number of results on this page
	TotalCount   int      // total number of results
	Page         int      // number of the current page
	PrevPage     int      //   "    "   "  previous page, usually Page-1 but zero if Page == 1
	NextPage     int      //   "    "   "  next page, usually Page+1, but zero on the last page
	Offset       int      // offset of the first item on the current page
	Pages        []int    // consecutive page numbers to be displayed for navigation
	Limits       []int    // limits to be displayed
}

// PageURL constructs a URL that displays the given page.
// It adds a "page" query parameter to the base URL.
func (p pagination) PageURL(page int) string {
	newQuery := p.baseURL.Query()
	newQuery.Set("page", strconv.Itoa(page))
	p.baseURL.RawQuery = newQuery.Encode()
	return p.baseURL.String()
}

// URL constructs a URL that adds limit and mode query parameters to the base
// URL. Passing a zero value omits the parameter.
func (p pagination) URL(limit int, mode, q string) string {
	newQuery := p.baseURL.Query()
	if limit != 0 {
		newQuery.Set("limit", strconv.Itoa(limit))
	}
	if mode != "" {
		newQuery.Set("m", mode)
	}
	if q != "" {
		newQuery.Set("q", q)
	}
	p.baseURL.RawQuery = newQuery.Encode()
	return p.baseURL.String()
}

// newPagination constructs a pagination. Call it after some results have been
// obtained.
// resultCount is the number of results in the current page.
// totalCount is the total number of results.
func newPagination(params paginationParams, resultCount, totalCount int) pagination {
	return pagination{
		baseURL:      params.baseURL,
		TotalCount:   totalCount,
		ResultCount:  resultCount,
		Offset:       params.offset(),
		Limit:        params.limit,
		DefaultLimit: defaultSearchLimit,
		MaxLimit:     maxSearchPageSize,
		Page:         params.page,
		PrevPage:     prev(params.page),
		NextPage:     next(params.page, params.limit, totalCount),
		Pages:        pagesToLink(params.page, numPages(params.limit, totalCount), defaultNumPagesToLink),
	}
}

// paginationParams holds pagination parameters extracted from the request.
type paginationParams struct {
	baseURL *url.URL
	page    int // the number of the page to display
	limit   int // the maximum number of results to display on the page
}

// offset returns the offset of the first result on the page.
func (p paginationParams) offset() int {
	return offset(p.page, p.limit)
}

// newPaginationParams extracts pagination params from the request.
func newPaginationParams(r *http.Request, defaultLimit int) paginationParams {
	positiveParam := func(key string, dflt int) (val int) {
		var err error
		if a := r.FormValue(key); a != "" {
			val, err = strconv.Atoi(a)
			if err != nil {
				log.Errorf(r.Context(), "strconv.Atoi(%q) for page: %v", a, err)
			}
		}
		if val < 1 {
			val = dflt
		}
		return val
	}
	return paginationParams{
		baseURL: r.URL,
		page:    positiveParam("page", 1),
		limit:   positiveParam("limit", defaultLimit),
	}
}

const defaultNumPagesToLink = 5

// pagesToLink returns the page numbers that will be displayed. Given a
// page, it returns a slice containing numPagesToLink integers in ascending
// order and optimizes for page to be in the middle of that range. The max
// value of an integer in the return slice will be less than numPages.
func pagesToLink(page, numPages, numPagesToLink int) []int {
	var pages []int
	start := page - (numPagesToLink / 2)
	if (numPages - start) < numPagesToLink {
		start = numPages - numPagesToLink + 1
	}
	if start < 1 {
		start = 1
	}

	for i := start; (i < start+numPagesToLink) && (i <= numPages); i++ {
		pages = append(pages, i)
	}
	return pages
}

// numPages is the total number of pages needed to display all the results,
// given the specified maximum page size and the total number of results.
func numPages(pageSize, totalCount int) int {
	return (totalCount + pageSize - 1) / pageSize
}

// offset returns the offset of the first result on page, assuming all previous
// pages were of size limit.
func offset(page, limit int) int {
	if page <= 1 {
		return 0
	}
	return (page - 1) * limit
}

// prev returns the number of the page before the given page, or zero if the
// given page is 1 or smaller.
func prev(page int) int {
	if page <= 1 {
		return 0
	}
	return page - 1
}

// next returns the number of the page after the given page, or zero if page is the last page or larger.
// limit and totalCount are used to calculate the last page (see numPages).
func next(page, limit, totalCount int) int {
	if page >= numPages(limit, totalCount) {
		return 0
	}
	return page + 1
}
