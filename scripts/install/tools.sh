#!/usr/bin/env bash
# Parallel builder for all tool modules (memory/disk friendly)

set -Eeuo pipefail

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd -P)"
REPO_ROOT="${REPO_ROOT:-$(pwd)}"
TOOLS_DIR="${TOOLS_DIR:-${REPO_ROOT}/tools}"

echo "[$SCRIPT_NAME]: starting…"

# ---- sanity checks -----------------------------------------------------------
if [[ ! -d "$TOOLS_DIR" ]]; then
  echo "[FATAL]: tools directory not found: $TOOLS_DIR"
  exit 2
fi

for cmd in go xargs awk df; do
  if ! command -v "$cmd" >/dev/null 2>&1; then
    echo "[FATAL]: required command '$cmd' not found"
    exit 2
  fi
done

# Resolve binary output dir
BIN_DIR="$(go env GOBIN)"
if [[ -z "$BIN_DIR" ]]; then
  BIN_DIR="$(go env GOPATH)/bin"
fi
mkdir -p "$BIN_DIR"

# ---- ensure Go cache and tmp dirs --------------------------------------------
: "${XDG_CACHE_HOME:=${HOME}/.cache}"
: "${GOCACHE:=${XDG_CACHE_HOME}/go-build}"
: "${GOTMPDIR:=${XDG_CACHE_HOME}/go-build-tmp}"
mkdir -p "$GOCACHE" "$GOTMPDIR"
export XDG_CACHE_HOME GOCACHE GOTMPDIR

# ---- memory-aware concurrency ------------------------------------------------
detect_mem_bytes() {
  if [[ -r /sys/fs/cgroup/memory.max ]]; then
    local m
    m="$(cat /sys/fs/cgroup/memory.max)"
    [[ "$m" != "max" ]] && echo "$m" && return
  fi
  if [[ -r /sys/fs/cgroup/memory/memory.limit_in_bytes ]]; then
    local m
    m="$(cat /sys/fs/cgroup/memory/memory.limit_in_bytes)"
    (( m < 1<<60 )) && echo "$m" && return
  fi
  awk '/MemAvailable:/ {print $2*1024}' /proc/meminfo
}

CPU_CORES="$(getconf _NPROCESSORS_ONLN 2>/dev/null || nproc 2>/dev/null || sysctl -n hw.ncpu 2>/dev/null || echo 4)"
MEM_BYTES="$(detect_mem_bytes || echo $((1024*1024*1024*32)))"
MEM_MB=$(( MEM_BYTES / 1048576 ))

PER_JOB_MB="${PER_JOB_MB:-900}"
AUTO_JOBS=$(( MEM_MB / PER_JOB_MB ))
(( AUTO_JOBS < 1 )) && AUTO_JOBS=1
(( AUTO_JOBS > CPU_CORES )) && AUTO_JOBS=$CPU_CORES

# ---- disk-aware concurrency --------------------------------------------------
disk_free_bytes() {
  df -Pk "$GOTMPDIR" | awk 'NR==2 {print $4*1024}'
}
TMP_BYTES="$(disk_free_bytes || echo $((1024*1024*1024)))"
TMP_MB=$(( TMP_BYTES / 1048576 ))

PER_JOB_TMP_MB="${PER_JOB_TMP_MB:-1500}"  # примерно 1.5 ГБ на тяжелый тул
AUTO_JOBS_TMP=$(( TMP_MB / PER_JOB_TMP_MB ))
(( AUTO_JOBS_TMP < 1 )) && AUTO_JOBS_TMP=1

(( AUTO_JOBS > AUTO_JOBS_TMP )) && AUTO_JOBS=$AUTO_JOBS_TMP

# ---- final job settings ------------------------------------------------------
JOBS="${JOBS:-$AUTO_JOBS}"
GO_BUILD_P="${GO_BUILD_P:-1}"

echo "[INFO]: CPU=${CPU_CORES}, Mem=${MEM_MB} MiB, DiskFree=${TMP_MB} MiB, auto-JOBS=${AUTO_JOBS}"
echo "[INFO]: Using ${JOBS} parallel jobs. go build -p ${GO_BUILD_P}. Binaries -> ${BIN_DIR}"

# ---- collect tool modules ----------------------------------------------------
tool_dirs=()
while IFS= read -r -d '' d; do
  [[ -f "$d/go.mod" ]] || { echo "[WARN]: skipping $(basename "$d") (no go.mod)"; continue; }
  tool_dirs+=("$d")
done < <(find "$TOOLS_DIR" -mindepth 1 -maxdepth 1 -type d -print0)

if [[ ${#tool_dirs[@]} -eq 0 ]]; then
  echo "[INFO]: No tools to build."
  exit 0
fi

# ---- run builds in parallel --------------------------------------------------
export GO_BUILD_P
printf '%s\0' "${tool_dirs[@]}" | xargs -0 -I{} -P "$JOBS" "${SCRIPT_DIR}/build-tool.sh" {} "$BIN_DIR"

echo "[SUCCESS]: All tools built."
