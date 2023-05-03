// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/lib/pq"
	"go.opencensus.io/trace"
	"golang.org/x/pkgsite/internal"
	"golang.org/x/pkgsite/internal/database"
	"golang.org/x/pkgsite/internal/derrors"
	"golang.org/x/pkgsite/internal/version"
)

// InsertIndexVersions inserts new versions into the module_version_states
// table with a status of zero.
func (db *DB) InsertIndexVersions(ctx context.Context, versions []*internal.IndexVersion) (err error) {
	defer derrors.WrapStack(&err, "InsertIndexVersions(ctx, %v)", versions)
	conflictAction := `
		ON CONFLICT
			(module_path, version)
		DO UPDATE SET
			index_timestamp=excluded.index_timestamp,
			next_processed_after=CURRENT_TIMESTAMP`
	return insertIndexVersions(ctx, db.db, versions, conflictAction)
}

// InsertNewModuleVersionFromFrontendFetch inserts a new module version into
// the module_version_states table with a status of zero that was requested
// from frontend fetch.
func (db *DB) InsertNewModuleVersionFromFrontendFetch(ctx context.Context, modulePath, resolvedVersion string) (err error) {
	defer derrors.WrapStack(&err, "InsertIndexVersion(ctx, %v)", resolvedVersion)
	conflictAction := `ON CONFLICT (module_path, version) DO NOTHING`
	return insertIndexVersions(ctx, db.db, []*internal.IndexVersion{{Path: modulePath, Version: resolvedVersion}}, conflictAction)
}

func insertIndexVersions(ctx context.Context, ddb *database.DB, versions []*internal.IndexVersion, conflictAction string) (err error) {
	var vals []any
	for _, v := range versions {
		vals = append(vals,
			v.Path,
			v.Version,
			version.ForSorting(v.Version),
			0,
			"",
			"",
			version.IsIncompatible(v.Version),
			v.Timestamp,
		)
	}
	cols := []string{
		"module_path",
		"version",
		"sort_version",
		"status",
		"error",
		"go_mod_path",
		"incompatible",
		"index_timestamp",
	}
	return ddb.Transact(ctx, sql.LevelDefault, func(tx *database.DB) error {
		var updates [][2]string // (module_path, version) to update status
		err := tx.BulkInsertReturning(ctx, "module_version_states", cols, vals, conflictAction,
			[]string{"module_path", "version", "status"},
			func(rows *sql.Rows) error {
				var (
					mod, ver string
					status   int
				)
				if err := rows.Scan(&mod, &ver, &status); err != nil {
					return err
				}
				// Update a module's status to 0 if it wasn't found previously.
				// See https://golang.org/issue/46117.
				if status == http.StatusNotFound {
					updates = append(updates, [2]string{mod, ver})
				}
				return nil
			})
		if err != nil {
			return err
		}
		// We don't have a BulkUpdate function that works for us here
		// (database.BulkUpdate can use only one column as a key). But we expect
		// very few of these, so it's fine to run them individually.
		for _, mv := range updates {
			_, err = tx.Exec(ctx, `
				UPDATE module_version_states
				SET status = 0
				WHERE module_path = $1 AND version = $2
			`, mv[0], mv[1])
			if err != nil {
				return err
			}
		}
		return nil
	})
}

type ModuleVersionStateForUpdate struct {
	ModulePath           string
	Version              string
	AppVersion           string
	Timestamp            time.Time
	Status               int
	HasGoMod             bool
	GoModPath            string
	FetchErr             error
	PackageVersionStates []*internal.PackageVersionState
}

// UpdateModuleVersionState inserts or updates the module_version_state table with
// the results of a fetch operation for a given module version.
func (db *DB) UpdateModuleVersionState(ctx context.Context, mvs *ModuleVersionStateForUpdate) (err error) {
	defer derrors.WrapStack(&err, "UpsertModuleVersionState(ctx, %s@%s)", mvs.ModulePath, mvs.Version)
	ctx, span := trace.StartSpan(ctx, "UpsertModuleVersionState")
	defer span.End()

	var numPackages *int
	if !(mvs.Status >= http.StatusBadRequest && mvs.Status <= http.StatusNotFound) {
		// If a module was fetched a 40x error in this range, we won't know how
		// many packages it has.
		n := len(mvs.PackageVersionStates)
		numPackages = &n
	}
	return db.db.Transact(ctx, sql.LevelDefault, func(tx *database.DB) error {
		if err := updateModuleVersionState(ctx, tx, numPackages, mvs); err != nil {
			return err
		}
		// Sync modules.status if the module exists in the modules table.
		if err := updateModulesStatus(ctx, tx, mvs.ModulePath, mvs.Version, mvs.Status); err != nil {
			return err
		}
		if len(mvs.PackageVersionStates) == 0 {
			return nil
		}
		return upsertPackageVersionStates(ctx, tx, mvs.PackageVersionStates)
	})
}

func updateModuleVersionState(ctx context.Context, db *database.DB, numPackages *int, mvs *ModuleVersionStateForUpdate) (err error) {
	defer derrors.WrapStack(&err, "upsertModuleVersionState(%q, %q, ...)", mvs.ModulePath, mvs.Version)
	ctx, span := trace.StartSpan(ctx, "upsertModuleVersionState")
	defer span.End()

	var sqlErrorMsg string
	if mvs.FetchErr != nil {
		sqlErrorMsg = mvs.FetchErr.Error()
	}

	affected, err := db.Exec(ctx, `
		UPDATE module_version_states
		SET app_version=$1,
			status=$2,
			has_go_mod=$3,
			go_mod_path=$4,
			error=$5,
			num_packages=$6,
			try_count=try_count+1,
			last_processed_at=CURRENT_TIMESTAMP,
			-- back off exponentially until 1 hour, then at constant 1-hour intervals
			next_processed_after=CASE
				WHEN last_processed_at IS NULL THEN
					CURRENT_TIMESTAMP + INTERVAL '1 minute'
				WHEN 2*(next_processed_after - last_processed_at) < INTERVAL '1 hour' THEN
					CURRENT_TIMESTAMP + 2*(next_processed_after - last_processed_at)
				ELSE
					CURRENT_TIMESTAMP + INTERVAL '1 hour'
				END
		WHERE
			module_path=$7
			AND version=$8`,
		mvs.AppVersion,
		mvs.Status,
		mvs.HasGoMod,
		mvs.GoModPath,
		sqlErrorMsg,
		numPackages,
		mvs.ModulePath,
		mvs.Version)
	if err != nil {
		return err
	}
	if affected != 1 {
		return fmt.Errorf("module version state update affected %d rows, expected exactly 1", affected)
	}
	return nil
}

// updateModulesStatus updates the status of the module with the given modulePath
// and version, if it exists, in the modules table.
func updateModulesStatus(ctx context.Context, db *database.DB, modulePath, resolvedVersion string, status int) (err error) {
	defer derrors.WrapStack(&err, "updateModulesStatus(%q, %q, %d)", modulePath, resolvedVersion, status)

	query := `UPDATE modules
			SET
				status = $1
			WHERE
				module_path = $2
				AND version = $3;`
	affected, err := db.Exec(ctx, query, status, modulePath, resolvedVersion)
	if err != nil {
		return err
	}
	if affected > 1 {
		return fmt.Errorf("module status update affected %d rows, expected at most 1", affected)
	}
	return nil
}

// UpdateModuleVersionStatus updates the status and error fields of a module version.
func (db *DB) UpdateModuleVersionStatus(ctx context.Context, modulePath, version string, status int, error string) (err error) {
	defer derrors.WrapStack(&err, "UpdateModuleVersionStatus(%q, %q, %d)", modulePath, version, status)

	query := `
		UPDATE module_version_states
		SET status = $3, error = $4
		WHERE module_path = $1 AND version = $2
	`
	_, err = db.db.Exec(ctx, query, modulePath, version, status, error)
	return err
}

func upsertPackageVersionStates(ctx context.Context, db *database.DB, packageVersionStates []*internal.PackageVersionState) (err error) {
	defer derrors.WrapStack(&err, "upsertPackageVersionStates")
	ctx, span := trace.StartSpan(ctx, "upsertPackageVersionStates")
	defer span.End()

	sort.Slice(packageVersionStates, func(i, j int) bool {
		return packageVersionStates[i].PackagePath < packageVersionStates[j].PackagePath
	})
	var vals []any
	for _, pvs := range packageVersionStates {
		vals = append(vals, pvs.PackagePath, pvs.ModulePath, pvs.Version, pvs.Status, pvs.Error)
	}
	return db.BulkInsert(ctx, "package_version_states",
		[]string{
			"package_path",
			"module_path",
			"version",
			"status",
			"error",
		},
		vals,
		`ON CONFLICT (module_path, package_path, version)
				DO UPDATE
				SET
					package_path=excluded.package_path,
					module_path=excluded.module_path,
					version=excluded.version,
					status=excluded.status,
					error=excluded.error`)
}

// LatestIndexTimestamp returns the last timestamp successfully inserted into
// the module_version_states table.
func (db *DB) LatestIndexTimestamp(ctx context.Context) (_ time.Time, err error) {
	defer derrors.WrapStack(&err, "LatestIndexTimestamp(ctx)")

	query := `SELECT index_timestamp
		FROM module_version_states
		ORDER BY index_timestamp DESC
		LIMIT 1`

	var ts time.Time
	row := db.db.QueryRow(ctx, query)
	switch err := row.Scan(&ts); err {
	case sql.ErrNoRows:
		return time.Time{}, nil
	case nil:
		return ts, nil
	default:
		return time.Time{}, err
	}
}

const moduleVersionStateColumns = `
			module_path,
			version,
			index_timestamp,
			created_at,
			status,
			error,
			try_count,
			last_processed_at,
			next_processed_after,
			app_version,
			has_go_mod,
			go_mod_path,
			num_packages`

// scanModuleVersionState constructs an *internal.ModuleModuleVersionState from the given
// scanner. It expects columns to be in the order of moduleVersionStateColumns.
func scanModuleVersionState(scan func(dest ...any) error) (*internal.ModuleVersionState, error) {
	var (
		v               internal.ModuleVersionState
		indexTimestamp  pq.NullTime
		lastProcessedAt pq.NullTime
		numPackages     sql.NullInt64
		hasGoMod        sql.NullBool
	)
	if err := scan(&v.ModulePath, &v.Version, &indexTimestamp, &v.CreatedAt, &v.Status, &v.Error,
		&v.TryCount, &v.LastProcessedAt, &v.NextProcessedAfter, &v.AppVersion, &hasGoMod, &v.GoModPath,
		&numPackages); err != nil {
		return nil, err
	}
	if indexTimestamp.Valid {
		it := indexTimestamp.Time
		v.IndexTimestamp = &it
	}
	if lastProcessedAt.Valid {
		lp := lastProcessedAt.Time
		v.LastProcessedAt = &lp
	}
	if hasGoMod.Valid {
		v.HasGoMod = hasGoMod.Bool
	}
	if numPackages.Valid {
		n := int(numPackages.Int64)
		v.NumPackages = &n
	}
	return &v, nil
}

// queryModuleVersionStates executes a query for ModuleModuleVersionState rows. It expects the
// given queryFormat be a format specifier with exactly one argument: a %s verb
// for the query columns.
func (db *DB) queryModuleVersionStates(ctx context.Context, queryFormat string, args ...any) ([]*internal.ModuleVersionState, error) {
	query := fmt.Sprintf(queryFormat, moduleVersionStateColumns)
	rows, err := db.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var versions []*internal.ModuleVersionState
	for rows.Next() {
		v, err := scanModuleVersionState(rows.Scan)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan(): %v", err)
		}
		versions = append(versions, v)
	}

	return versions, nil
}

// GetRecentFailedVersions returns versions that have most recently failed.
func (db *DB) GetRecentFailedVersions(ctx context.Context, limit int) (_ []*internal.ModuleVersionState, err error) {
	defer derrors.WrapStack(&err, "GetRecentFailedVersions(ctx, %d)", limit)

	queryFormat := `
		SELECT %s
		FROM
			module_version_states
		WHERE status=500
		ORDER BY last_processed_at DESC
		LIMIT $1`
	return db.queryModuleVersionStates(ctx, queryFormat, limit)
}

// GetRecentVersions returns recent versions that have been processed.
func (db *DB) GetRecentVersions(ctx context.Context, limit int) (_ []*internal.ModuleVersionState, err error) {
	defer derrors.WrapStack(&err, "GetRecentVersions(ctx, %d)", limit)

	queryFormat := `
		SELECT %s
		FROM
			module_version_states
		ORDER BY created_at DESC
		LIMIT $1`
	return db.queryModuleVersionStates(ctx, queryFormat, limit)
}

// GetModuleVersionState returns the current module version state for
// modulePath and version.
func (db *DB) GetModuleVersionState(ctx context.Context, modulePath, resolvedVersion string) (_ *internal.ModuleVersionState, err error) {
	defer derrors.WrapStack(&err, "GetModuleVersionState(ctx, %q, %q)", modulePath, resolvedVersion)

	query := fmt.Sprintf(`
		SELECT %s
		FROM
			module_version_states
		WHERE
			module_path = $1
			AND version = $2;`, moduleVersionStateColumns)

	row := db.db.QueryRow(ctx, query, modulePath, resolvedVersion)
	v, err := scanModuleVersionState(row.Scan)
	switch err {
	case nil:
		return v, nil
	case sql.ErrNoRows:
		return nil, derrors.NotFound
	default:
		return nil, fmt.Errorf("row.Scan(): %v", err)
	}
}

// GetPackageVersionStatesForModule returns the current package version states
// for modulePath and version.
func (db *DB) GetPackageVersionStatesForModule(ctx context.Context, modulePath, resolvedVersion string) (_ []*internal.PackageVersionState, err error) {
	defer derrors.WrapStack(&err, "GetPackageVersionState(ctx, %q, %q)", modulePath, resolvedVersion)

	query := `
		SELECT
			package_path,
			module_path,
			version,
			status,
			error
		FROM
			package_version_states
		WHERE
			module_path = $1
			AND version = $2;`

	var states []*internal.PackageVersionState
	collect := func(rows *sql.Rows) error {
		var s internal.PackageVersionState
		if err := rows.Scan(&s.PackagePath, &s.ModulePath, &s.Version,
			&s.Status, &s.Error); err != nil {
			return fmt.Errorf("rows.Scan(): %v", err)
		}
		states = append(states, &s)
		return nil
	}
	if err := db.db.RunQuery(ctx, query, collect, modulePath, resolvedVersion); err != nil {
		return nil, err
	}
	return states, nil
}

// GetPackageVersionState returns the current package version state for
// pkgPath, modulePath and version.
func (db *DB) GetPackageVersionState(ctx context.Context, pkgPath, modulePath, resolvedVersion string) (_ *internal.PackageVersionState, err error) {
	defer derrors.WrapStack(&err, "GetPackageVersionState(ctx, %q, %q, %q)", pkgPath, modulePath, resolvedVersion)

	query := `
		SELECT
			package_path,
			module_path,
			version,
			status,
			error
		FROM
			package_version_states
		WHERE
			package_path = $1
			AND module_path = $2
			AND version = $3;`

	var pvs internal.PackageVersionState
	err = db.db.QueryRow(ctx, query, pkgPath, modulePath, resolvedVersion).Scan(
		&pvs.PackagePath, &pvs.ModulePath, &pvs.Version,
		&pvs.Status, &pvs.Error)
	switch err {
	case nil:
		return &pvs, nil
	case sql.ErrNoRows:
		return nil, derrors.NotFound
	default:
		return nil, fmt.Errorf("row.Scan(): %v", err)
	}
}

// VersionStats holds statistics extracted from the module_version_states
// table.
type VersionStats struct {
	LatestTimestamp time.Time
	VersionCounts   map[int]int // from status to number of rows
}

// GetVersionStats queries the module_version_states table for aggregate
// information about the current state of module versions, grouping them by
// their current status code.
func (db *DB) GetVersionStats(ctx context.Context) (_ *VersionStats, err error) {
	defer derrors.WrapStack(&err, "GetVersionStats(ctx)")

	query := `
		SELECT
			status,
			max(index_timestamp),
			count(*)
		FROM
			module_version_states
		GROUP BY status;`

	stats := &VersionStats{
		VersionCounts: make(map[int]int),
	}
	err = db.db.RunQuery(ctx, query, func(rows *sql.Rows) error {
		var (
			status         sql.NullInt64
			indexTimestamp time.Time
			count          int
		)
		if err := rows.Scan(&status, &indexTimestamp, &count); err != nil {
			return fmt.Errorf("row.Scan(): %v", err)
		}
		if indexTimestamp.After(stats.LatestTimestamp) {
			stats.LatestTimestamp = indexTimestamp
		}
		stats.VersionCounts[int(status.Int64)] = count
		return nil
	})
	if err != nil {
		return nil, err
	}
	return stats, nil
}

// HasGoMod reports whether a given module version has a go.mod file.
// It returns a NotFound error if it can't find any information.
func (db *DB) HasGoMod(ctx context.Context, modulePath, version string) (has bool, err error) {
	defer derrors.WrapStack(&err, "HasGoMod(%q, %q)", modulePath, version)

	// Check the module_version_states table. It has information about
	// every module we've seen. Ignore rows with status == 0 because
	// they haven't been processed yet.
	var hasP *bool
	err = db.db.QueryRow(ctx, `
		SELECT has_go_mod
		FROM module_version_states
		WHERE module_path = $1
		AND version = $2
		AND status != 0
	`, modulePath, version).Scan(&hasP)
	if err == sql.ErrNoRows {
		return false, derrors.NotFound
	}
	if err != nil {
		return false, err
	}
	if hasP != nil {
		return *hasP, nil
	}
	// the has_go_mod column hasn't been populated yet.
	// Fall back to the modules table.
	// This can be removed when all rows have been populated and
	// module_version_states.has_go_mod is migrated to NOT NULL.
	err = db.db.QueryRow(ctx, `
		SELECT has_go_mod
		FROM modules
		WHERE module_path = $1 AND version = $2
	`, modulePath, version).Scan(&has)
	if err == sql.ErrNoRows {
		return false, derrors.NotFound
	}
	if err != nil {
		return false, err
	}
	return has, nil
}
