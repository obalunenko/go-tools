#!/usr/bin/env bash
set -euo pipefail

readonly CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
readonly ROOT_DIR="$(dirname "$CURRENT_DIR")"

usage() {
  cat <<EOF
Usage:
  $(basename "$0") <alpine_version>

Example:
  $(basename "$0") 3.24

This script updates Alpine version across:
  - Dockerfile ARG ALPINE_VERSION
  - docker-bake.hcl args.ALPINE_VERSION
  - Optionally refreshes README versions via scripts/update-readme-versions.sh if present
EOF
}

main() {
  if [[ ${#} -lt 1 ]]; then
    usage
    exit 1
  fi

  local alpine_version="$1"

  echo "Updating Alpine version across the repository:"
  echo " - New Alpine version: ${alpine_version}"

  bump_dockerfile_alpine_version "${alpine_version}"
  bump_bakefile_alpine_version "${alpine_version}"

  # If README updater exists, run it to sync the base image version.
  if [[ -x "${ROOT_DIR}/scripts/update-readme-versions.sh" ]]; then
    echo " - Syncing README versions"
    ( cd "${ROOT_DIR}" && ./scripts/update-readme-versions.sh ) || true
  fi

  echo "Done."
}

# Update ARG ALPINE_VERSION in Dockerfile.
bump_dockerfile_alpine_version() {
  local new_version="${1}"
  local dockerfile="${ROOT_DIR}/Dockerfile"
  if [[ -f "${dockerfile}" ]]; then
    echo " - Updating Dockerfile ALPINE_VERSION to ${new_version}"
    sed -E "s/^(ARG[[:space:]]+ALPINE_VERSION=).*/\\1${new_version}/" "${dockerfile}" > "${dockerfile}.tmp"
    mv "${dockerfile}.tmp" "${dockerfile}"
  fi
}

# Update ALPINE_VERSION in docker-bake.hcl.
bump_bakefile_alpine_version() {
  local new_version="${1}"
  local bakefile="${ROOT_DIR}/docker-bake.hcl"
  if [[ -f "${bakefile}" ]]; then
    echo " - Updating docker-bake.hcl ALPINE_VERSION to ${new_version}"
    sed -E "s/(ALPINE_VERSION[[:space:]]*=[[:space:]]*\")([^\"]+)(\")/\\1${new_version}\\3/" "${bakefile}" > "${bakefile}.tmp"
    mv "${bakefile}.tmp" "${bakefile}"
  fi
}

main "$@"
