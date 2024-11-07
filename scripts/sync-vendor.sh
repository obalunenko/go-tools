#!/bin/sh

set -eu

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(dirname "$0")"
REPO_ROOT="$(cd "${SCRIPT_DIR}" && git rev-parse --show-toplevel)"

echo "${SCRIPT_NAME} is running... "

sync_vendor() {
  echo "Syncing vendor..."
  go work sync
  go work vendor

  for dir in $REPO_ROOT/tools/*; do
    if [ -d "$dir" ] && [ -f "$dir/go.mod" ]; then
      echo "Running 'go mod tidy' in $dir"
      (
        cd "$dir"
        go mod tidy
      )
    fi
  done
}

cd "${REPO_ROOT}" || exit 1
pwd
sync_vendor

echo "${SCRIPT_NAME} done."
