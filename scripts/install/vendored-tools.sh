#!/bin/bash

# Enable strict error handling
set -eu

# Script metadata
SCRIPT_NAME="$(basename "$0")"
REPO_ROOT="$(pwd)"
TOOLS_DIR="${REPO_ROOT}/tools"

echo "${SCRIPT_NAME} is running... "

# Change directory to TOOLS_DIR
cd "${TOOLS_DIR}" || exit 1

# Function to check the status of the last executed command
function check_status() {
  local error_message="$1"
  if [ $? -ne 0 ]; then
    if [ -n "$error_message" ]; then
      echo "[ERROR]: $error_message"
    fi
    # Exit with code 255 to signal xargs to abort the process, otherwise it will return 0
    exit 255
  fi
}

# Function to install a dependency
function install_dep() {
  local dep=$1

  # Extract the binary output path
  bin_out="$GOBIN/$(echo "$dep" | awk 'BEGIN { FS="/" } { for (i=NF; i>0; i--) if ($i !~ /^v[0-9]/) { print $i; exit } }')"

  # Determine tools module
  tools_module="$(grep '^module ' go.mod | awk '{print $2}')"

  # Extract the version of the dependency
  raw_version=$(go list -mod=readonly -m -f '{{if not .Indirect}}{{.Path}} {{.Version}}{{end}}' all | grep -v "^${tools_module}")
  version=$(echo "$raw_version" | awk '{print $2}')

  echo "[INFO]: Building ${dep}@${version} - Output: ${bin_out}"

  # Build the dependency
  go build -mod=vendor -o "${bin_out}" "${dep}"
  check_status "Build failed for dependency [${dep}@${version}]!"

  echo "[SUCCESS]: Successfully built [${dep}@${version}]."
}

# Export functions for use in subshells (xargs)
export -f install_dep
export -f check_status

# Function to install dependencies listed in the go.mod file
function install_deps() {
  local tools_module

  # Extract the tools module name
  tools_module="$(grep '^module ' go.mod | awk '{print $2}')"
  echo "[INFO]: Installing dependencies for module: ${tools_module}"

  # List and install dependencies using xargs for parallel execution
  go list -e -f '{{ join .Imports "\n" }}' -tags="tools" "${tools_module}" |
    xargs -n 1 -P 0 -I {} bash -c 'install_dep "$@"' _ {}
}

# Function to iterate over tools and install them
function install_tools() {
  declare -a tools_list
  local temp_file=./tools_list.txt

  # Create a temporary file to store tool directories
  touch "$temp_file"
  ls -d "${TOOLS_DIR}"/*/ > "$temp_file"

  # Read all tool directories into an array
  while IFS= read -r t; do
    tools_list+=("$t")
  done < "$temp_file"

  # Clean up the temporary file
  rm "$temp_file"

  # Loop through each tool and install its dependencies
  for t in "${tools_list[@]}"; do
    echo "[INFO]: Processing tool: ${t}"

    tool=$(basename "${t}")
    cd "${TOOLS_DIR}/${tool}" || exit 1
    install_deps
    cd - > /dev/null || exit 1
  done
}

# Start the tool installation process
install_tools