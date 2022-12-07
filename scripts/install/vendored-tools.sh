#!/bin/bash

set -eu

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(dirname "$0")"
REPO_ROOT="$(cd "${SCRIPT_DIR}" && git rev-parse --show-toplevel)"
TOOLS_DIR="${REPO_ROOT}/tools"

echo "${SCRIPT_NAME} is running... "
#echo "${GOOS}/${GOARCH}"

cd "${TOOLS_DIR}" || exit 1

function check_status() {
  # first param is error message to print in case of error
  if [ $? -ne 0 ]; then
    if [ -n "$1" ]; then
      echo "$1"
    fi

    # Exit 255 to pass signal to xargs to abort process with code 1, in other cases xargs will complete with 0.
    exit 255
  fi
}

function install_dep() {
  dep=$1

  bin_out=$GOBIN/$(echo $dep | awk 'BEGIN { FS="/" } {print $NF}')

  echo "[INFO]: Going to build ${dep} - ${bin_out}"

  go build -mod vendor -o "${bin_out}" "${dep}"

  check_status "[FAIL]: build [${dep}] failed!"

  echo "[SUCCESS]: build [${dep}] finished."
}

export -f install_dep
export -f check_status

function install_deps() {
  tools_module="$(go list -m)"
  
  go list -f '{{ join .Imports "\n" }}' -tags="tools" "${tools_module}" |
   xargs -n 1 -P 0 -I {} bash -c 'install_dep "$@"' _ {}
}

install_deps
