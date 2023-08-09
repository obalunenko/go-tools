#!/bin/bash

set -Eeuo pipefail

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(dirname "$0")"
REPO_ROOT="$(cd "${SCRIPT_DIR}" && git rev-parse --show-toplevel)"
SCRIPTS_DIR="${REPO_ROOT}/scripts"

source "${SCRIPTS_DIR}/helpers-source.sh"

APP=${APP_NAME}

echo "${SCRIPT_NAME} is running fo ${APP}... "

checkInstalled 'goreleaser'

# Get new tags from the remote
git fetch --tags -f

COMMIT="$(git rev-parse HEAD)"
SHORTCOMMIT="$(git rev-parse --short HEAD)"
DATE="$(date -u +%Y-%m-%dT%H:%M:%SZ)"
VERSION="$(git tag | sort -V | tail -1)"
GOVERSION="$(go version | awk '{print $3;}')"

if [ -z "${VERSION}" ] || [ "${VERSION}" = "${SHORTCOMMIT}" ]; then
  VERSION="v0.0.0"
fi

VERSION="${VERSION}-local"

goreleaser check

goreleaser build --clean --single-target --snapshot
