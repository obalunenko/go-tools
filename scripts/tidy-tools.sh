#!/usr/bin/env bash
# Run go mod tidy for all tool modules under tools/

set -Eeuo pipefail

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd -P)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd -P)"
TOOLS_DIR="${TOOLS_DIR:-${REPO_ROOT}/tools}"
JOBS="${JOBS:-$(getconf _NPROCESSORS_ONLN 2>/dev/null || echo 4)}"

echo "[$SCRIPT_NAME]: starting..."

if [[ ! -d "$TOOLS_DIR" ]]; then
  echo "[FATAL]: tools directory not found: $TOOLS_DIR"
  exit 2
fi

for cmd in go find xargs; do
  if ! command -v "$cmd" >/dev/null 2>&1; then
    echo "[FATAL]: required command '$cmd' not found"
    exit 2
  fi
done

RAW_GO_VERSION="${GOVERSION:-}"
if [[ -z "$RAW_GO_VERSION" ]]; then
  RAW_GO_VERSION="$(go env GOVERSION 2>/dev/null || true)"
fi
if [[ -z "$RAW_GO_VERSION" ]]; then
  echo "[FATAL]: unable to determine Go version. Set GOVERSION, e.g. GOVERSION=1.26"
  exit 2
fi

if [[ "$RAW_GO_VERSION" =~ ^go([0-9]+\.[0-9]+(\.[0-9]+)?)$ ]]; then
  TIDY_GO_VERSION="${BASH_REMATCH[1]}"
elif [[ "$RAW_GO_VERSION" =~ ^([0-9]+\.[0-9]+(\.[0-9]+)?)$ ]]; then
  TIDY_GO_VERSION="${BASH_REMATCH[1]}"
else
  echo "[FATAL]: unsupported GOVERSION format: ${RAW_GO_VERSION}"
  exit 2
fi

tidy_one() {
  local tool_dir="$1"

  (
    set -Eeuo pipefail
    cd "$tool_dir"

    if [[ ! -f go.mod ]]; then
      echo "[WARN]: $(basename "$tool_dir"): no go.mod, skipping"
      exit 0
    fi

    local module
    module="$(awk '/^module[[:space:]]+/ {print $2; exit}' go.mod || true)"
    echo "[INFO]: $(basename "$tool_dir"): tidying ${module:-<unknown>}"

    go mod tidy -go="${TIDY_GO_VERSION}"

    echo "[OK]:   $(basename "$tool_dir"): tidy completed"
  ) || {
    echo "[ERROR]: $(basename "$tool_dir"): tidy failed"
    exit 255
  }
}

export -f tidy_one
export TIDY_GO_VERSION

echo "[INFO]: TOOLS_DIR=${TOOLS_DIR}"
echo "[INFO]: Using ${JOBS} parallel jobs"
echo "[INFO]: go mod tidy -go=${TIDY_GO_VERSION}"

find "$TOOLS_DIR" -mindepth 1 -maxdepth 1 -type d -print0 \
  | xargs -0 -r -n 1 -P "$JOBS" bash -c 'tidy_one "$1"' _

echo "[SUCCESS]: All tool modules were tidied."
