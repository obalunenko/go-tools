// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package frontend

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/safehtml"
	"github.com/google/safehtml/uncheckedconversions"
	"golang.org/x/pkgsite/internal"
	"golang.org/x/pkgsite/internal/cookie"
	"golang.org/x/pkgsite/internal/derrors"
	"golang.org/x/pkgsite/internal/frontend/page"
	"golang.org/x/pkgsite/internal/frontend/serrors"
	"golang.org/x/pkgsite/internal/frontend/urlinfo"
	"golang.org/x/pkgsite/internal/frontend/versions"
	"golang.org/x/pkgsite/internal/log"
	"golang.org/x/pkgsite/internal/middleware/stats"
	"golang.org/x/pkgsite/internal/stdlib"
	"golang.org/x/pkgsite/internal/version"
	"golang.org/x/pkgsite/internal/vuln"
)

// UnitPage contains data needed to render the unit template.
type UnitPage struct {
	page.BasePage
	// Unit is the unit for this page.
	Unit *internal.UnitMeta

	// Breadcrumb contains data used to render breadcrumb UI elements.
	Breadcrumb breadcrumb

	// Title is the title of the page.
	Title string

	// URLPath is the path suitable for links on the page.
	// See the unitURLPath for details.
	URLPath string

	// CanonicalURLPath is a permanent representation of the URL path for a
	// unit.
	// It uses the resolved module path and version.
	// For example, if the latest version of /my.module/pkg is version v1.5.2,
	// the canonical URL path for that unit would be /my.module@v1.5.2/pkg
	CanonicalURLPath string

	// The version string formatted for display.
	DisplayVersion string

	// LinkVersion is version string suitable for links used to compute
	// latest badges.
	LinkVersion string

	// LatestURL is a url pointing to the latest version of a unit.
	LatestURL string

	// IsLatestMinor is true if the version displayed is the latest minor of the unit.
	// Used to determine the canonical URL for search engines and robots meta directives.
	IsLatestMinor bool

	// LatestMinorClass is the CSS class that describes the current unit's minor
	// version in relationship to the latest version of the unit.
	LatestMinorClass string

	// Information about the latest major version of the module.
	LatestMajorVersion    string
	LatestMajorVersionURL string

	// PageType is the type of page (pkg, cmd, dir, std, or mod).
	PageType string

	// PageLabels are the labels that will be displayed
	// for a given page.
	PageLabels []string

	// CanShowDetails indicates whether details can be shown or must be
	// hidden due to issues like license restrictions.
	CanShowDetails bool

	// Settings contains settings for the selected tab.
	SelectedTab TabSettings

	// RedirectedFromPath is the path that redirected to the current page.
	// If non-empty, a "redirected from" banner will be displayed
	// (see static/frontend/unit/_header.tmpl).
	RedirectedFromPath string

	// Details contains data specific to the type of page being rendered.
	Details any

	// Vulns holds vulnerability information.
	Vulns []vuln.Vuln

	// DepsDevURL holds the full URL to this module version on deps.dev.
	DepsDevURL string

	// IsGoProject is true if the package is from the standard library or a
	// golang.org sub-repository.
	IsGoProject bool
}

// serveUnitPage serves a unit page for a path.
func (s *Server) serveUnitPage(ctx context.Context, w http.ResponseWriter, r *http.Request,
	ds internal.DataSource, info *urlinfo.URLPathInfo) (err error) {
	defer derrors.Wrap(&err, "serveUnitPage(ctx, w, r, ds, %v)", info)
	defer stats.Elapsed(ctx, "serveUnitPage")()

	tab := r.FormValue("tab")
	if tab == "" {
		// Default to details tab when there is no tab param.
		tab = tabMain
	}
	// Redirect to clean URL path when tab param is invalid.
	if _, ok := unitTabLookup[tab]; !ok {
		http.Redirect(w, r, r.URL.Path, http.StatusFound)
		return nil
	}

	um, err := ds.GetUnitMeta(ctx, info.FullPath, info.ModulePath, info.RequestedVersion)
	if err != nil {
		if !errors.Is(err, derrors.NotFound) {
			return err
		}
		db, ok := ds.(internal.PostgresDB)
		if !ok || s.fetchServer == nil {
			return serrors.DatasourceNotSupportedError()
		}
		return s.fetchServer.ServePathNotFoundPage(w, r, db, info.FullPath, info.ModulePath, info.RequestedVersion)
	}

	makeDepsDevURL := depsDevURLGenerator(ctx, s.depsDevHTTPClient, um)

	// Use GOOS and GOARCH query parameters to create a build context, which
	// affects the documentation and synopsis. Omitting both results in an empty
	// build context, which will match the first (and preferred) build context.
	// It's also okay to provide just one (e.g. GOOS=windows), which will select
	// the first doc with that value, ignoring the other one.
	bc := internal.BuildContext{GOOS: r.FormValue("GOOS"), GOARCH: r.FormValue("GOARCH")}
	d, err := fetchDetailsForUnit(ctx, r, tab, ds, um, info.RequestedVersion, bc, s.vulnClient)
	if err != nil {
		return err
	}
	if s.shouldServeJSON(r) {
		return s.serveJSONPage(w, r, d)
	}

	if _, ok := internal.DefaultBranches[info.RequestedVersion]; ok {
		// Since path@master is a moving target, we don't want it to be stale.
		// As a result, we enqueue every request of path@master to the frontend
		// task queue, which will initiate a fetch request depending on the
		// last time we tried to fetch this module version.
		//
		// Use a separate context here to prevent the context from being canceled
		// elsewhere before a task is enqueued.
		if s.queue == nil {
			return &serrors.ServerError{
				Status: http.StatusBadRequest,
				Err:    err,
				Epage: &page.ErrorPage{
					MessageData: fmt.Sprintf(`Default branches like "@%s" are not supported. Omit to get the current version.`,
						info.RequestedVersion),
				},
			}
		}
		go func() {
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
			defer cancel()
			log.Infof(ctx, "serveUnitPage: Scheduling %q@%q to be fetched", um.ModulePath, info.RequestedVersion)
			if _, err := s.queue.ScheduleFetch(ctx, um.ModulePath, info.RequestedVersion, nil); err != nil {
				log.Errorf(ctx, "serveUnitPage(%q): scheduling fetch for %q@%q: %v",
					r.URL.Path, um.ModulePath, info.RequestedVersion, err)
			}
		}()
	}

	if !isValidTabForUnit(tab, um, d) {
		// Redirect to clean URL path when tab param is invalid for the unit
		// type.
		http.Redirect(w, r, r.URL.Path, http.StatusFound)
		return nil
	}

	// If we've already called GetUnitMeta for an unknown module path and the latest version, pass
	// it to GetLatestInfo to avoid a redundant call.
	var latestUnitMeta *internal.UnitMeta
	if info.ModulePath == internal.UnknownModulePath && info.RequestedVersion == version.Latest {
		latestUnitMeta = um
	}
	latestInfo := s.GetLatestInfo(ctx, um.Path, um.ModulePath, latestUnitMeta)
	var redirectPath string
	redirectPath, err = cookie.Extract(w, r, cookie.AlternativeModuleFlash)
	if err != nil {
		// Don't fail, but don't display a banner either.
		log.Errorf(ctx, "extracting AlternativeModuleFlash cookie: %v", err)
	}
	title := pageTitle(um)
	basePage := s.newBasePage(r, title)
	tabSettings := unitTabLookup[tab]
	basePage.AllowWideContent = true
	if tabSettings.Name == "" {
		basePage.UseResponsiveLayout = true
	}
	lv := versions.LinkVersion(um.ModulePath, info.RequestedVersion, um.Version)
	page := UnitPage{
		BasePage:              basePage,
		Unit:                  um,
		Breadcrumb:            displayBreadcrumb(um, info.RequestedVersion),
		Title:                 title,
		SelectedTab:           tabSettings,
		URLPath:               versions.ConstructUnitURL(um.Path, um.ModulePath, info.RequestedVersion),
		CanonicalURLPath:      canonicalURLPath(um.Path, um.ModulePath, info.RequestedVersion, um.Version),
		DisplayVersion:        versions.DisplayVersion(um.ModulePath, info.RequestedVersion, um.Version),
		LinkVersion:           lv,
		LatestURL:             versions.ConstructUnitURL(um.Path, um.ModulePath, version.Latest),
		LatestMinorClass:      latestMinorClass(lv, latestInfo),
		LatestMajorVersionURL: latestInfo.MajorUnitPath,
		PageLabels:            pageLabels(um),
		PageType:              pageType(um),
		RedirectedFromPath:    redirectPath,
		DepsDevURL:            makeDepsDevURL(),
		IsGoProject:           isGoProject(um.ModulePath),
		IsLatestMinor:         lv == latestInfo.MinorVersion,
	}

	// Show the banner if there was no error getting the latest major version,
	// and it is different from the major version of the current module path.
	latestMajor := internal.MajorVersionForModule(latestInfo.MajorModulePath)
	if latestMajor != "" && latestMajor != internal.MajorVersionForModule(um.ModulePath) {
		page.LatestMajorVersion = latestMajor
	}

	page.Details = d
	main, ok := d.(*MainDetails)
	if ok {
		page.MetaDescription = metaDescription(main.DocSynopsis)
	}

	// Get vulnerability information.
	page.Vulns = vuln.VulnsForPackage(ctx, um.ModulePath, um.Version, um.Path, s.vulnClient)

	s.servePage(ctx, w, tabSettings.TemplateName, page)
	return nil
}

func (s *Server) shouldServeJSON(r *http.Request) bool {
	return s.serveStats && r.FormValue("content") == "json"
}

func (s *Server) serveJSONPage(w http.ResponseWriter, r *http.Request, d any) (err error) {
	defer derrors.Wrap(&err, "serveJSONPage(ctx, w, r)")
	if !s.shouldServeJSON(r) {
		return derrors.NotFound
	}
	data, err := json.Marshal(d)
	if err != nil {
		return fmt.Errorf("json.Marshal: %v", err)
	}
	if _, err := w.Write(data); err != nil {
		return fmt.Errorf("w.Write: %v", err)
	}
	return nil
}

func latestMinorClass(version string, latest internal.LatestInfo) string {
	c := "DetailsHeader-badge"
	switch {
	case latest.MinorVersion == "":
		c += "--unknown"
	case latest.MinorVersion == version && !latest.UnitExistsAtMinor:
		c += "--notAtLatest"
	case latest.MinorVersion == version:
		c += "--latest"
	default:
		c += "--goToLatest"
	}
	return c
}

// metaDescription uses a safehtml escape hatch to build HTML used
// to render the <meta name="Description"> for unit pages as a
// workaround for https://github.com/google/safehtml/issues/6.
func metaDescription(synopsis string) safehtml.HTML {
	if synopsis == "" {
		return safehtml.HTML{}
	}
	return safehtml.HTMLConcat(
		uncheckedconversions.HTMLFromStringKnownToSatisfyTypeContract(`<meta name="Description" content="`),
		safehtml.HTMLEscaped(synopsis),
		uncheckedconversions.HTMLFromStringKnownToSatisfyTypeContract(`">`),
	)
}

// isValidTabForUnit reports whether the tab is valid for the given unit.
// It is assumed that tab is a key in unitTabLookup.
func isValidTabForUnit(tab string, um *internal.UnitMeta, details any) bool {
	if tab == tabLicenses && !(details.(*LicensesDetails).IsRedistributable) {
		return false
	}
	if !um.IsPackage() && (tab == tabImports || tab == tabImportedBy) {
		return false
	}
	return true
}

// canonicalURLPath constructs a URL path to the unit that always includes the
// resolved version.
func canonicalURLPath(fullPath, modulePath, requestedVersion, resolvedVersion string) string {
	return versions.ConstructUnitURL(fullPath, modulePath,
		versions.LinkVersion(modulePath, requestedVersion, resolvedVersion))
}

func isGoProject(modulePath string) bool {
	return modulePath == stdlib.ModulePath || strings.HasPrefix(modulePath, "golang.org")
}
