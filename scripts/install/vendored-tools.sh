#!/bin/bash

set -euo pipefail

# Enable job control
set -m

# Define metadata variables and paths
SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
REPO_ROOT="$(git -C "${SCRIPT_DIR}" rev-parse --show-toplevel)"
TOOLS_DIR="${REPO_ROOT}/tools"

echo "${SCRIPT_NAME} is running... "

# Check and set GOBIN
if [ -z "${GOBIN:-}" ]; then
  GOBIN="$(go env GOPATH)/bin"
  echo "[INFO]: GOBIN is not set. Using default: ${GOBIN}"
fi

mkdir -p "${GOBIN}"

# Function to get the number of CPU cores
function get_cpu_cores() {
  if command -v nproc >/dev/null 2>&1; then
    nproc
  elif [[ "$(uname)" == "Darwin" ]]; then
    sysctl -n hw.ncpu
  else
    # Attempt to use getconf
    cores=$(getconf _NPROCESSORS_ONLN 2>/dev/null || echo 1)
    echo "${cores}"
  fi
}

# Function to install a module
function install_module() {
  local module_dir="$1"

  if [ ! -d "${module_dir}" ]; then
    echo "[WARN]: ${module_dir} is not a directory. Skipping."
    return
  fi

  if [ ! -f "${module_dir}/go.mod" ]; then
    echo "[WARN]: go.mod not found in ${module_dir}. Skipping."
    return
  fi

  echo "[INFO]: Processing module in directory ${module_dir}"
  cd "${module_dir}" || exit 1

  local module_name
  module_name=$(grep "^module " go.mod | awk '{print $2}')
  echo "[INFO]: Installing dependencies for module ${module_name}"

  local deps
  deps=$(go list -e -f '{{ join .Imports "\n" }}' -tags=tools "${module_name}")

  if [ -z "${deps}" ]; then
    echo "[INFO]: No dependencies found for module ${module_name}."
    return
  fi

  while IFS= read -r dep; do
    local bin_name
    bin_name=$(echo $dep | awk 'BEGIN { FS="/" } {for (i=NF; i>0; i--) if ($i !~ /^v[0-9]/) {print $i;exit}}')
    local bin_out="${GOBIN}/${bin_name}"
    echo "[INFO]: Building ${dep} -> ${bin_out}"
    if go build -mod=readonly -o "${bin_out}" "${dep}"; then
      echo "[SUCCESS]: ${dep} built successfully."
    else
      echo "[ERROR]: Failed to build ${dep}"
      return 1
    fi
  done <<< "${deps}"
}

# Main function to install tools
function install_tools() {
  echo "[INFO]: Starting tool installation"

  local max_jobs
  max_jobs=$(get_cpu_cores)
  echo "[INFO]: Using up to ${max_jobs} parallel tasks"

  # Collect module directories
  modules=()
  for dir in "${TOOLS_DIR}"/*; do
    if [ -d "${dir}" ]; then
      modules+=("${dir}")
    fi
  done

  # Launch module installations with parallelism control
  local i=0
  local total_modules=${#modules[@]}
  while [ $i -lt $total_modules ]; do
    # Check the number of running background jobs
    while [ "$(jobs -r -p | wc -l)" -ge "$max_jobs" ]; do
      sleep 0.1  # Wait before rechecking
    done

    # Start module installation in the background
    install_module "${modules[$i]}" &
    ((i++))
  done

  wait  # Wait for all background jobs to finish
}

# Run the main function
install_tools

echo "[INFO]: ${SCRIPT_NAME} completed successfully."
