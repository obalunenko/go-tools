#!/bin/bash

set -eu

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(dirname "$0")"
REPO_ROOT="$(cd "${SCRIPT_DIR}" && git rev-parse --show-toplevel)"
TOOLS_DIR="${REPO_ROOT}/tools"

echo "${SCRIPT_NAME} is running... "

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

  bin_out=$GOBIN/$(echo $dep | awk 'BEGIN { FS="/" } {for (i=NF; i>0; i--) if ($i !~ /^v[0-9]/) {print $i;exit}}')

  echo "[INFO]: Going to build ${dep} - ${bin_out}"

  go build -mod=readonly -o "${bin_out}" "${dep}"

  check_status "[FAIL]: build [${dep}] failed!"

  echo "[SUCCESS]: build [${dep}] finished."
}

export -f install_dep
export -f check_status

function install_deps() {
  tools_module="$(grep '^module ' go.mod | awk '{print $2}')"

  echo "[INFO]: Running install_deps for ${tools_module}"
  
  go list -e -f '{{ join .Imports "\n" }}' -tags="tools" "${tools_module}" |
   xargs -n 1 -P 0 -I {} bash -c 'install_dep "$@"' _ {}
}

function install_tools() {
  declare -a tools_list

  temp_file=$(mktemp) # создаем временный файл

  go list -m > "$temp_file" # сохраняем вывод команды в файл

  while IFS= read -r t; do
    tools_list+=("$t")
  done < "$temp_file" # читаем файл в массив

  rm "$temp_file" # удаляем временный файл

  for t in "${tools_list[@]}"; do
    echo "In loop - current ${t}"

    cd "${TOOLS_DIR}/${t}" || exit 1
    install_deps
    cd - || exit 1
  done
}


install_tools
