// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package proxy provides a client for interacting with a proxy.
package proxy

import (
	"archive/zip"
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"golang.org/x/mod/module"
	"golang.org/x/net/context/ctxhttp"
	"golang.org/x/pkgsite/internal/derrors"
	"golang.org/x/pkgsite/internal/version"
)

// A Client is used by the fetch service to communicate with a module
// proxy. It handles all methods defined by go help goproxy.
type Client struct {
	// URL of the module proxy web server
	url string

	// Client used for HTTP requests. It is mutable for testing purposes.
	HTTPClient *http.Client

	// Whether fetch should be disabled.
	disableFetch bool

	cache *cache
}

// A VersionInfo contains metadata about a given version of a module.
type VersionInfo struct {
	Version string
	Time    time.Time
}

// Setting this header to true prevents the proxy from fetching uncached
// modules.
const DisableFetchHeader = "Disable-Module-Fetch"

// New constructs a *Client using the provided url, which is expected to
// be an absolute URI that can be directly passed to http.Get.
// The optional transport parameter is used by the underlying http client.
func New(u string, transport http.RoundTripper) (_ *Client, err error) {
	defer derrors.WrapStack(&err, "proxy.New(%q)", u)
	return &Client{
		url:          strings.TrimRight(u, "/"),
		HTTPClient:   &http.Client{Transport: transport},
		disableFetch: false,
	}, nil
}

// WithFetchDisabled returns a new client that sets the Disable-Module-Fetch
// header so that the proxy does not fetch a module it doesn't already know
// about.
func (c *Client) WithFetchDisabled() *Client {
	c2 := *c
	c2.disableFetch = true
	return &c2
}

// FetchDisabled reports whether proxy fetch is disabled.
func (c *Client) FetchDisabled() bool {
	return c.disableFetch
}

// WithCache returns a new client that caches some RPCs.
func (c *Client) WithCache() *Client {
	c2 := *c
	c2.cache = &cache{}
	return &c2
}

// Info makes a request to $GOPROXY/<module>/@v/<requestedVersion>.info and
// transforms that data into a *VersionInfo.
// If requestedVersion is internal.LatestVersion, it uses the proxy's @latest
// endpoint instead.
func (c *Client) Info(ctx context.Context, modulePath, requestedVersion string) (_ *VersionInfo, err error) {
	defer func() {
		// Don't report NotFetched, because it is the normal result of fetching
		// an uncached module when fetch is disabled.
		// Don't report timeouts, because they are relatively frequent and not actionable.
		wrap := derrors.Wrap
		if !errors.Is(err, derrors.NotFetched) && !errors.Is(err, derrors.ProxyTimedOut) && !errors.Is(err, derrors.NotFound) {
			wrap = derrors.WrapAndReport
		}
		wrap(&err, "proxy.Client.Info(%q, %q)", modulePath, requestedVersion)
	}()

	if v := c.cache.getInfo(modulePath, requestedVersion); v != nil {
		return v, nil
	}
	data, err := c.readBody(ctx, modulePath, requestedVersion, "info")
	if err != nil {
		return nil, err
	}
	var v VersionInfo
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	c.cache.putInfo(modulePath, requestedVersion, &v)
	return &v, nil
}

// Mod makes a request to $GOPROXY/<module>/@v/<resolvedVersion>.mod and returns the raw data.
func (c *Client) Mod(ctx context.Context, modulePath, resolvedVersion string) (_ []byte, err error) {
	defer derrors.WrapStack(&err, "proxy.Client.Mod(%q, %q)", modulePath, resolvedVersion)

	if b := c.cache.getMod(modulePath, resolvedVersion); b != nil {
		return b, nil
	}
	b, err := c.readBody(ctx, modulePath, resolvedVersion, "mod")
	if err != nil {
		return nil, err
	}
	c.cache.putMod(modulePath, resolvedVersion, b)
	return b, nil
}

// Zip makes a request to $GOPROXY/<modulePath>/@v/<resolvedVersion>.zip and
// transforms that data into a *zip.Reader. <resolvedVersion> must have already
// been resolved by first making a request to
// $GOPROXY/<modulePath>/@v/<requestedVersion>.info to obtained the valid
// semantic version.
func (c *Client) Zip(ctx context.Context, modulePath, resolvedVersion string) (_ *zip.Reader, err error) {
	defer derrors.WrapStack(&err, "proxy.Client.Zip(ctx, %q, %q)", modulePath, resolvedVersion)

	if r := c.cache.getZip(modulePath, resolvedVersion); r != nil {
		return r, nil
	}
	bodyBytes, err := c.readBody(ctx, modulePath, resolvedVersion, "zip")
	if err != nil {
		return nil, err
	}
	zipReader, err := zip.NewReader(bytes.NewReader(bodyBytes), int64(len(bodyBytes)))
	if err != nil {
		return nil, fmt.Errorf("zip.NewReader: %v: %w", err, derrors.BadModule)
	}
	c.cache.putZip(modulePath, resolvedVersion, zipReader)
	return zipReader, nil
}

// ZipSize gets the size in bytes of the zip from the proxy, without downloading it.
// The version must be resolved, as by a call to Client.Info.
func (c *Client) ZipSize(ctx context.Context, modulePath, resolvedVersion string) (_ int64, err error) {
	defer derrors.WrapStack(&err, "proxy.Client.ZipSize(ctx, %q, %q)", modulePath, resolvedVersion)

	url, err := c.EscapedURL(modulePath, resolvedVersion, "zip")
	if err != nil {
		return 0, err
	}
	res, err := ctxhttp.Head(ctx, c.HTTPClient, url)
	if err != nil {
		return 0, fmt.Errorf("ctxhttp.Head(ctx, client, %q): %v", url, err)
	}
	defer res.Body.Close()
	if err := responseError(res, false); err != nil {
		return 0, err
	}
	if res.ContentLength < 0 {
		return 0, errors.New("unknown content length")
	}
	return res.ContentLength, nil
}

func (c *Client) EscapedURL(modulePath, requestedVersion, suffix string) (_ string, err error) {
	defer derrors.WrapStack(&err, "Client.escapedURL(%q, %q, %q)", modulePath, requestedVersion, suffix)

	if suffix != "info" && suffix != "mod" && suffix != "zip" {
		return "", errors.New(`suffix must be "info", "mod" or "zip"`)
	}
	escapedPath, err := module.EscapePath(modulePath)
	if err != nil {
		return "", fmt.Errorf("path: %v: %w", err, derrors.InvalidArgument)
	}
	if requestedVersion == version.Latest {
		if suffix != "info" {
			return "", fmt.Errorf("cannot ask for latest with suffix %q", suffix)
		}
		return fmt.Sprintf("%s/%s/@latest", c.url, escapedPath), nil
	}
	escapedVersion, err := module.EscapeVersion(requestedVersion)
	if err != nil {
		return "", fmt.Errorf("version: %v: %w", err, derrors.InvalidArgument)
	}
	return fmt.Sprintf("%s/%s/@v/%s.%s", c.url, escapedPath, escapedVersion, suffix), nil
}

func (c *Client) readBody(ctx context.Context, modulePath, requestedVersion, suffix string) (_ []byte, err error) {
	defer derrors.WrapStack(&err, "Client.readBody(%q, %q, %q)", modulePath, requestedVersion, suffix)

	u, err := c.EscapedURL(modulePath, requestedVersion, suffix)
	if err != nil {
		return nil, err
	}
	var data []byte
	err = c.executeRequest(ctx, u, func(body io.Reader) error {
		var err error
		data, err = io.ReadAll(body)
		return err
	})
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Versions makes a request to $GOPROXY/<path>/@v/list and returns the
// resulting version strings.
func (c *Client) Versions(ctx context.Context, modulePath string) (_ []string, err error) {
	defer derrors.Wrap(&err, "Versions(ctx, %q)", modulePath)
	escapedPath, err := module.EscapePath(modulePath)
	if err != nil {
		return nil, fmt.Errorf("module.EscapePath(%q): %w", modulePath, derrors.InvalidArgument)
	}
	u := fmt.Sprintf("%s/%s/@v/list", c.url, escapedPath)
	var versions []string
	collect := func(body io.Reader) error {
		scanner := bufio.NewScanner(body)
		for scanner.Scan() {
			versions = append(versions, scanner.Text())
		}
		return scanner.Err()
	}
	if err := c.executeRequest(ctx, u, collect); err != nil {
		return nil, err
	}
	return versions, nil
}

// executeRequest executes an HTTP GET request for u, then calls the bodyFunc
// on the response body, if no error occurred.
func (c *Client) executeRequest(ctx context.Context, u string, bodyFunc func(body io.Reader) error) (err error) {
	defer func() {
		if ctx.Err() != nil {
			err = fmt.Errorf("%v: %w", err, derrors.ProxyTimedOut)
		}
		derrors.WrapStack(&err, "executeRequest(ctx, %q)", u)
	}()

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return err
	}
	if c.disableFetch {
		req.Header.Set(DisableFetchHeader, "true")
	}
	r, err := ctxhttp.Do(ctx, c.HTTPClient, req)
	if err != nil {
		return fmt.Errorf("ctxhttp.Do(ctx, client, %q): %v", u, err)
	}
	defer r.Body.Close()
	if err := responseError(r, c.disableFetch); err != nil {
		return err
	}
	return bodyFunc(r.Body)
}

// responseError translates the response status code to an appropriate error.
func responseError(r *http.Response, fetchDisabled bool) error {
	switch {
	case 200 <= r.StatusCode && r.StatusCode < 300:
		return nil
	case 500 <= r.StatusCode:
		return derrors.ProxyError
	case r.StatusCode == http.StatusNotFound,
		r.StatusCode == http.StatusGone:
		// Treat both 404 Not Found and 410 Gone responses
		// from the proxy as a "not found" error category.
		// If the response body contains "fetch timed out", treat this
		// as a 504 response so that we retry fetching the module version again
		// later.
		//
		// If the Disable-Module-Fetch header was set, use a different
		// error code so we can tell the difference.
		data, err := io.ReadAll(r.Body)
		if err != nil {
			return fmt.Errorf("io.ReadAll: %v", err)
		}
		d := string(data)
		switch {
		case strings.Contains(d, "fetch timed out"):
			err = derrors.ProxyTimedOut
		case fetchDisabled:
			err = derrors.NotFetched
		default:
			err = derrors.NotFound
		}
		return fmt.Errorf("%q: %w", d, err)
	default:
		return fmt.Errorf("unexpected status %d %s", r.StatusCode, r.Status)
	}
}
