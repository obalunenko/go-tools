#!/bin/bash

set -eu

# Constants
SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(dirname "$0")"
REPO_ROOT="$(cd "${SCRIPT_DIR}" && git rev-parse --show-toplevel)"
TOOLS_DIR="${REPO_ROOT}/tools"
GO_VERSION=$(go version | sed -E 's/.*go([0-9]+\.[0-9]+\.[0-9]+).*/\1/')

echo "${SCRIPT_NAME} is running..."

# Function to sync the vendor directory
sync_vendor() {
  echo "Syncing vendor..."
  rm -rf ./vendor
  go mod tidy -go=${GO_VERSION}
  go mod vendor
}


# Function to sync dependencies
sync_deps() {
  local tools_module
  tools_module="$(grep '^module ' go.mod | awk '{print $2}')"

  echo "[INFO]: Running install_deps for ${tools_module}"

  sync_vendor
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