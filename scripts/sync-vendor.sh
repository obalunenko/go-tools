#!/usr/bin/env bash
# Fast parallel vendor sync for tools modules

set -Eeuo pipefail

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd -P)"
REPO_ROOT="$(cd "${SCRIPT_DIR}" && git rev-parse --show-toplevel)"
TOOLS_DIR="${TOOLS_DIR:-${REPO_ROOT}/tools}"

echo "[$SCRIPT_NAME]: starting…"

# ---- sanity checks -----------------------------------------------------------
if [[ ! -d "$TOOLS_DIR" ]]; then
  echo "[FATAL]: tools directory not found: $TOOLS_DIR"
  exit 2
fi

for cmd in go git xargs find awk sed; do
  if ! command -v "$cmd" >/dev/null 2>&1; then
    echo "[FATAL]: required command '$cmd' not found"
    exit 2
  fi
done

# Normalize Go version to MAJOR.MINOR (what 'go mod tidy -go' expects)
# Examples: go1.22.1 -> 1.22 ; go1.21.3 -> 1.21
export GO_VERSION=$(go version | sed -E 's/.*go([0-9]+\.[0-9]+\.[0-9]+).*/\1/')

# Concurrency
CPU_CORES="$(getconf _NPROCESSORS_ONLN 2>/dev/null || nproc 2>/dev/null || sysctl -n hw.ncpu 2>/dev/null || echo 4)"
JOBS="${JOBS:-$CPU_CORES}"
echo "[INFO]: Using ${JOBS} parallel jobs. TOOLS_DIR=${TOOLS_DIR}  GO=${GO_VERSION}"

# ---- per-tool sync -----------------------------------------------------------
sync_one() {
  local tool_dir="$1"

  # Allow spaces in directory names: we always pass through xargs -0 and "$1"
  if [[ ! -d "$tool_dir" ]]; then
    echo "[WARN]: not a directory: $tool_dir"
    return 0
  fi
  if [[ ! -f "$tool_dir/go.mod" ]]; then
    echo "[WARN]: skipping $(basename "$tool_dir") (no go.mod)"
    return 0
  fi

  (
    set -Eeuo pipefail
    cd "$tool_dir"

    local module
    module="$(awk '/^module[[:space:]]+/ {print $2; exit}' go.mod || true)"
    echo "[INFO]: $(basename "$tool_dir"): syncing vendor for ${module:-<unknown>}"

    # Clean vendor to avoid stale content, then re‑vendor under the target Go version
    rm -rf vendor
    go mod tidy -go="${GO_VERSION}"
    go mod download

    echo "[OK]:   $(basename "$tool_dir"): ${module:-<unknown>} vendor synced."
  ) || {
    echo "[ERROR]: $(basename "$tool_dir"): vendor sync failed."
    # Return 255 to make xargs stop scheduling new jobs (running ones will finish)
    exit 255
  }
}

export -f sync_one

# ---- enumerate tools and run in parallel ------------------------------------
# Find immediate subdirectories of TOOLS_DIR and process each in parallel.
# Using -print0 / -0 keeps paths with spaces safe.
find "$TOOLS_DIR" -mindepth 1 -maxdepth 1 -type d -print0 \
  | xargs -0 -r -n 1 -P "$JOBS" bash -c 'sync_one "$1"' _

echo "[$SCRIPT_NAME]: finished successfully."
