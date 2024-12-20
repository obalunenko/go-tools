#!/bin/bash

set -eu

# Constants
SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(dirname "$0")"
REPO_ROOT="$(cd "${SCRIPT_DIR}" && git rev-parse --show-toplevel)"
TOOLS_DIR="${REPO_ROOT}/tools"

echo "${SCRIPT_NAME} is running..."

# Function to sync the vendor directory
sync_vendor() {
  echo "Syncing vendor..."
  rm -rf vendor
  go mod tidy
  go mod vendor
}

# Function to check the status of the last command
check_status() {
  local err_msg="$1"

  if [ $? -ne 0 ]; then
    if [ -n "${err_msg}" ]; then
      echo "${err_msg}"
    fi
    # Exit 255 to signal xargs to abort the process with code 1
    exit 255
  fi
}

# Function to sync a single dependency
sync_dep() {
  local dep="$1"

  local tools_module
  tools_module="$(grep '^module ' go.mod | awk '{print $2}')"

  local raw_version
  raw_version=$(go list -mod=readonly -m -f '{{if not .Indirect}}{{.Path}} {{.Version}}{{end}}' all | grep -v "^${tools_module}")

  local version
  version=$(echo "${raw_version}" | awk '{print $2}')

  echo "[INFO]: Syncing ${dep}@${version}"

  sync_vendor
  check_status "[FAIL]: Failed to sync [${dep}@${version}]!"

  echo "[SUCCESS]: Synced [${dep}@${version}] successfully."
}

# Export functions for use in subshells
export -f sync_vendor
export -f check_status
export -f sync_dep

# Function to sync dependencies
sync_deps() {
  local tools_module
  tools_module="$(grep '^module ' go.mod | awk '{print $2}')"

  echo "[INFO]: Running install_deps for ${tools_module}"

  go list -e -f '{{ join .Imports "\n" }}' -tags="tools" "${tools_module}" |
    xargs -n 1 -P 0 -I {} bash -c 'sync_dep "$@"' _ {}
}

# Function to sync all tools in the TOOLS_DIR
sync_tools() {
  local tools_list=()

  # Temporary file to store the list of tools
  local temp_file="./tools_list.txt"
  touch "${temp_file}"

  # Save tools directories into the temporary file
  ls -d "${TOOLS_DIR}"/*/ > "${temp_file}"

  # Read tools into an array
  while IFS= read -r tool_dir; do
    tools_list+=("${tool_dir}")
  done < "${temp_file}"

  # Remove the temporary file
  rm "${temp_file}"

  # Loop through each tool directory and sync dependencies
  for tool_path in "${tools_list[@]}"; do
    echo "[INFO]: Processing tool directory: ${tool_path}"

    local tool
    tool=$(basename "${tool_path}")

    cd "${TOOLS_DIR}/${tool}" || exit 1
    sync_deps
    cd - > /dev/null || exit 1
  done
}

# Start the tool synchronization process
sync_tools

echo "${SCRIPT_NAME} finished."