#!/usr/bin/env bash
# Fast parallel tools builder (docker/memory friendly)

set -Eeuo pipefail

SCRIPT_NAME="$(basename "$0")"
REPO_ROOT="${REPO_ROOT:-$(pwd)}"
TOOLS_DIR="${TOOLS_DIR:-${REPO_ROOT}/tools}"

echo "[$SCRIPT_NAME]: starting…"

# ---- sanity checks -----------------------------------------------------------
if [[ ! -d "$TOOLS_DIR" ]]; then
  echo "[FATAL]: tools directory not found: $TOOLS_DIR"
  exit 2
fi

for cmd in go xargs awk; do
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

# ---- memory-aware concurrency -----------------------------------------------
detect_mem_bytes() {
  # cgroup v2
  if [[ -r /sys/fs/cgroup/memory.max ]]; then
    local m
    m="$(cat /sys/fs/cgroup/memory.max)"
    if [[ "$m" != "max" ]]; then
      echo "$m"; return
    fi
  fi
  # cgroup v1
  if [[ -r /sys/fs/cgroup/memory/memory.limit_in_bytes ]]; then
    local m
    m="$(cat /sys/fs/cgroup/memory/memory.limit_in_bytes)"
    # large 'dummy' numbers indicate no memory limit
    if (( m < 1<<60 )); then
      echo "$m"; return
    fi
  fi
  # fallback: доступная память хоста/контейнера
  awk '/MemAvailable:/ {print $2*1024}' /proc/meminfo
}

CPU_CORES="$(getconf _NPROCESSORS_ONLN 2>/dev/null || nproc 2>/dev/null || sysctl -n hw.ncpu 2>/dev/null || echo 4)"
MEM_BYTES="$(detect_mem_bytes || echo $((1024*1024*1024*32)))"
MEM_MB=$(( MEM_BYTES / 1048576 ))

# Amount of memory allocated per parallel build job
PER_JOB_MB="${PER_JOB_MB:-900}"   # adjust if necessary (500..900)
AUTO_JOBS=$(( MEM_MB / PER_JOB_MB ))
(( AUTO_JOBS < 1 )) && AUTO_JOBS=1
# не больше, чем CPU
(( AUTO_JOBS > CPU_CORES )) && AUTO_JOBS=$CPU_CORES

# Пользователь может переопределить JOBS из окружения
JOBS="${JOBS:-$AUTO_JOBS}"

# Internal parallelism of the Go compiler (default 1 to save RAM)
GO_BUILD_P="${GO_BUILD_P:-1}"

echo "[INFO]: CPU=${CPU_CORES}, Mem=${MEM_MB} MiB, PER_JOB_MB=${PER_JOB_MB}, auto-JOBS=${AUTO_JOBS}"
echo "[INFO]: Using ${JOBS} parallel jobs. go build -p ${GO_BUILD_P}. Binaries -> ${BIN_DIR}"

# ---- helpers ----------------------------------------------------------------
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

    local mod_flag="-mod=mod"
    [[ -d vendor ]] && mod_flag="-mod=vendor"

    local name out
    name="$(bin_name "$dep")"
    out="${BIN_DIR}/${name}"

    local modver
    modver="$(go list $mod_flag -f '{{with .Module}}{{.Path}}@{{.Version}}{{end}}' "$dep" 2>/dev/null || true)"

    echo "[INFO]: $(basename "$tool_dir"): building ${dep}${modver:+ (${modver})} -> ${out}"

    # Ограничиваем внутреннюю параллельность компилятора
    if ! go build $mod_flag -p "${GO_BUILD_P}" -o "$out" "$dep"; then
      echo "[ERROR]: $(basename "$tool_dir"): build failed for ${dep}"
      # 255 просим xargs завершить остальные задачи
      exit 255
    fi

    echo "[OK]:   $(basename "$tool_dir"): ${dep} -> ${out}"
  )
}

export -f build_one bin_name
export BIN_DIR GO_BUILD_P

# ---- collect jobs (tool_dir, dep) -------------------------------------------
tmp_jobs="$(mktemp)"
trap 'rm -f "$tmp_jobs"' EXIT

while IFS= read -r -d '' d; do
  [[ -f "$d/go.mod" ]] || { echo "[WARN]: skipping $(basename "$d") (no go.mod)"; continue; }
  if imports="$(cd "$d" && go list -mod=mod -e -f '{{ join .Imports "\n" }}' -tags=tools . 2>/dev/null | sed '/^$/d' | sort -u)"; then
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
# Примечание: пути с пробелами не поддерживаются.
cat "$tmp_jobs" \
  | xargs -r -n 2 -P "$JOBS" bash -c 'build_one "$1" "$2"' _

echo "[SUCCESS]: All tools built."
