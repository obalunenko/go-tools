#!/usr/bin/env bash
set -euo pipefail

readonly CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
readonly ROOT_DIR="$(dirname "$CURRENT_DIR")"

usage() {
  cat <<EOF
Usage:
  $(basename "$0") <go_major_minor> [go_image_version]

Examples:
  $(basename "$0") 1.26            # sets go.mod to 1.26 and Docker image to 1.26
  $(basename "$0") 1.26 1.26.2     # sets go.mod to 1.26 and Docker image to 1.26.2

This script updates Go version across:
  - all go.mod files (go <version>)
  - Dockerfile ARG GO_VERSION
  - docker-bake.hcl args.GO_VERSION
  - GitHub Actions workflows (actions/setup-go go-version)
  - Optionally refreshes README versions via scripts/update-readme-versions.sh if present
EOF
}

main() {
  if [[ ${#} -lt 1 ]]; then
    usage
    exit 1
  fi

  local go_major_minor="$1"            # e.g. 1.26
  local go_image_version="${2:-$1}"    # e.g. 1.26.2 or 1.26

  echo "Updating Go version across the repository:"
  echo " - New go.mod version: ${go_major_minor}"
  echo " - New Docker image version: ${go_image_version}"

  bump_all_go_mods "${go_major_minor}"
  bump_dockerfile_go_version "${go_image_version}"
  bump_bakefile_go_version "${go_image_version}"
  bump_ci_go_versions "${go_major_minor}"

  # If README updater exists, run it to sync versions table
  if [[ -x "${ROOT_DIR}/scripts/update-readme-versions.sh" ]]; then
    echo " - Syncing README versions"
    ( cd "${ROOT_DIR}" && ./scripts/update-readme-versions.sh ) || true
  fi

  echo "Done."
}

# Replace the 'go <any>' line with 'go <new>' across all go.mod files
bump_all_go_mods() {
  local newGo="${1}"
  echo " - Updating go.mod files..."
  # Find all go.mod files excluding vendor and .git
  while IFS= read -r goModFile; do
    # Use a temp file to keep compatibility with macOS sed
    sed -E "s/^go[[:space:]]+[0-9]+\.[0-9]+(\.[0-9]+)?/go ${newGo}/" "${goModFile}" > "${goModFile}.tmp"
    mv "${goModFile}.tmp" "${goModFile}"
  done < <(find "${ROOT_DIR}" -type f -name "go.mod" \
            -not -path "${ROOT_DIR}/vendor/*" \
            -not -path "${ROOT_DIR}/.git/*")
}

# Update ARG GO_VERSION in Dockerfile
bump_dockerfile_go_version() {
  local newImageVer="${1}"
  local dockerfile="${ROOT_DIR}/Dockerfile"
  if [[ -f "${dockerfile}" ]]; then
    echo " - Updating Dockerfile GO_VERSION to ${newImageVer}"
    sed -E "s/^(ARG[[:space:]]+GO_VERSION=).*/\\1${newImageVer}/" "${dockerfile}" > "${dockerfile}.tmp"
    mv "${dockerfile}.tmp" "${dockerfile}"
  fi
}

# Update GO_VERSION in docker-bake.hcl
bump_bakefile_go_version() {
  local newImageVer="${1}"
  local bakefile="${ROOT_DIR}/docker-bake.hcl"
  if [[ -f "${bakefile}" ]]; then
    echo " - Updating docker-bake.hcl GO_VERSION to ${newImageVer}"
    sed -E "s/(GO_VERSION[[:space:]]*=[[:space:]]*\")([^\"]+)(\")/\\1${newImageVer}\\3/" "${bakefile}" > "${bakefile}.tmp"
    mv "${bakefile}.tmp" "${bakefile}"
  fi
}

# Update go-version across GitHub Actions workflows to the new major.minor
bump_ci_go_versions() {
  local newGoMM="${1}" # e.g. 1.26
  local workflows_dir="${ROOT_DIR}/.github/workflows"
  [[ -d "${workflows_dir}" ]] || return 0

  echo " - Updating GitHub Actions go-version to ${newGoMM}"
  for f in "${workflows_dir}"/*.yml; do
    [[ -e "$f" ]] || continue
    # Apply a sequence of transformations into a temp file
    tmp="${f}.tmp"
    cp "$f" "$tmp"
    # 1) Matrix array form: go-version: [1.23, 1.x] -> replace first element
    sed -E "s/(go-version:[[:space:]]*\[)[0-9]+\.[0-9]+/\\1${newGoMM}/g" "$tmp" > "$tmp.1" && mv "$tmp.1" "$tmp"
    # 2) Single value with .x: go-version: 1.23.x or "1.23.x"
    sed -E "s/(go-version:[[:space:]]*['\"]?)\d+\.[0-9]+(\.x)(['\"]?)/\\1${newGoMM}\\2\\3/g" "$tmp" > "$tmp.1" && mv "$tmp.1" "$tmp"
    # 3) Single value exact: go-version: 1.23 or "1.23" or 1.23.4 -> set to new major.minor (keep quotes if any)
    sed -E "s/(go-version:[[:space:]]*['\"]?)\d+\.[0-9]+(\.[0-9]+)?(['\"]?)/\\1${newGoMM}\\3/g" "$tmp" > "$tmp.1" && mv "$tmp.1" "$tmp"
    mv "$tmp" "$f"
  done
}

main "$@"