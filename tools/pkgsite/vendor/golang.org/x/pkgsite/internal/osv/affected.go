// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package osv

// AffectedModulesAndPackages returns a list of module paths affected
// by a vuln. If the vuln is in the standard library or toolchain,
// it lists package names instead of modules.
func (e Entry) AffectedModulesAndPackages() []string {
	var affected []string
	for _, a := range e.Affected {
		switch a.Module.Path {
		case GoStdModulePath, GoCmdModulePath:
			// Name specific standard library packages and tools.
			for _, p := range a.EcosystemSpecific.Packages {
				affected = append(affected, p.Path)
			}
		default:
			// Outside the standard library, name the module.
			affected = append(affected, a.Module.Path)
		}
	}
	return affected
}

func (e *Entry) AffectsStandardLibrary() bool {
	for _, a := range e.Affected {
		m := a.Module.Path
		if m == GoStdModulePath || m == GoCmdModulePath {
			return true
		}
	}
	return false
}

func (e *Entry) IsUnreviewed() bool {
	return e.DatabaseSpecific.ReviewStatus == ReviewStatusUnreviewed
}

func (e *Entry) HasCustomVersions() bool {
	for _, a := range e.Affected {
		if len(a.EcosystemSpecific.CustomRanges) != 0 {
			return true
		}
	}
	return false
}
