#!/usr/bin/env bash
# Fast parallel tools builder (docker/memory/disk friendly)

set -Eeuo pipefail

SCRIPT_NAME="$(basename "$0")"
REPO_ROOT="${REPO_ROOT:-$(pwd)}"
TOOLS_DIR="${TOOLS_DIR:-${REPO_ROOT}/tools}"

echo "[$SCRIPT_NAME]: starting…"

REQUIRE_VENDOR="${REQUIRE_VENDOR:-1}"
export REQUIRE_VENDOR

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

# ---- helpers -----------------------------------------------------------------
ensure_vendor_dir() {
  local tool_dir="$1"
  if [[ -d "${tool_dir}/vendor" ]]; then
    return 0
  fi

  if [[ "$REQUIRE_VENDOR" == "1" ]]; then
    echo "[FATAL]: $(basename "$tool_dir"): vendor directory not found. Run 'make sync-vendor' before installing tools."
    exit 3
  fi

  echo "[WARN]: $(basename "$tool_dir"): vendor directory missing; falling back to module proxy."
}

mod_flag_for_dir() {
  local tool_dir="$1"
  if [[ -d "${tool_dir}/vendor" ]]; then
    echo "-mod=vendor"
  else
    echo "-mod=mod"
  fi
}

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

# ---- helpers -----------------------------------------------------------------
bin_name() {
  local path="$1"
  awk -F/ '{
    for (i=NF; i>0; i--) {
      if ($i !~ /^v[0-9]+$/) { print $i; exit }
    }
  }' <<<"$path"
}

build_one() {
  local tool_dir="$1"
  local dep="$2"

  (
    cd "$tool_dir"
    local mod_flag
    mod_flag="$(mod_flag_for_dir "$tool_dir")"

    if [[ "$mod_flag" != "-mod=vendor" && "$REQUIRE_VENDOR" == "1" ]]; then
      echo "[ERROR]: $(basename "$tool_dir"): vendor directory missing while REQUIRE_VENDOR=1"
      exit 3
    fi

    local name out
    name="$(bin_name "$dep")"
    out="${BIN_DIR}/${name}"

    local modver
    modver="$(go list $mod_flag -f '{{with .Module}}{{.Path}}@{{.Version}}{{end}}' "$dep" 2>/dev/null || true)"

    echo "[INFO]: $(basename "$tool_dir"): building ${dep}${modver:+ (${modver})} -> ${out}"

    if ! go build $mod_flag -p "${GO_BUILD_P}" -o "$out" "$dep"; then
      echo "[ERROR]: $(basename "$tool_dir"): build failed for ${dep}"
      exit 255
    fi

    echo "[OK]:   $(basename "$tool_dir"): ${dep} -> ${out}"
  )
}

export -f build_one bin_name
export -f mod_flag_for_dir
export BIN_DIR GO_BUILD_P

# ---- collect jobs ------------------------------------------------------------
tmp_jobs="$(mktemp)"
trap 'rm -f "$tmp_jobs"' EXIT

while IFS= read -r -d '' d; do
  [[ -f "$d/go.mod" ]] || { echo "[WARN]: skipping $(basename "$d") (no go.mod)"; continue; }
  ensure_vendor_dir "$d"
  mod_flag="$(mod_flag_for_dir "$d")"

  if imports="$(cd "$d" && go list "$mod_flag" -e -f '{{ join .Imports "\n" }}' -tags=tools . 2>/dev/null | sed '/^$/d' | sort -u)"; then
    while IFS= read -r dep; do
      printf '%s\t%s\n' "$d" "$dep" >> "$tmp_jobs"
    done <<< "$imports"
  else
    echo "[WARN]: $(basename "$d"): failed to enumerate imports; skipping"
  fi
done < <(find "$TOOLS_DIR" -mindepth 1 -maxdepth 1 -type d -print0)

if [[ ! -s "$tmp_jobs" ]]; then
  echo "[INFO]: No tools to build."
  exit 0
fi

# ---- run builds in parallel --------------------------------------------------
cat "$tmp_jobs" | xargs -r -n 2 -P "$JOBS" bash -c 'build_one "$1" "$2"' _

echo "[SUCCESS]: All tools built."
