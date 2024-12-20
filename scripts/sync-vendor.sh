#!/bin/bash

set -eu

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(dirname "$0")"
REPO_ROOT="$(cd "${SCRIPT_DIR}" && git rev-parse --show-toplevel)"
TOOLS_DIR="${REPO_ROOT}/tools"

echo "${SCRIPT_NAME} is running... "

function sync_vendor() {
  echo "Syncing vendor..."
  go mod tidy
  go mod vendor
}

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

function sync_dep() {
  dep=$1

  echo "[INFO]: Going to sync ${dep}"

  sync_vendor

  check_status "[FAIL]: sync [${dep}] failed!"

  echo "[SUCCESS]: sync [${dep}] finished."
}

export -f sync_dep
export -f check_status
export -f sync_vendor

function sync_deps() {
  tools_module="$(grep '^module ' go.mod | awk '{print $2}')"

  echo "[INFO]: Running install_deps for ${tools_module}"

  go list -e -f '{{ join .Imports "\n" }}' -tags="tools" "${tools_module}" |
   xargs -n 1 -P 0 -I {} bash -c 'sync_dep "$@"' _ {}
}

function sync_tools() {
  declare -a tools_list

  temp_file=./tools_list.txt # создаем временный файл

  touch "$temp_file" # создаем временный файл

  ls -d ${TOOLS_DIR}/*/ > "$temp_file" # сохраняем вывод команды в файл

  while IFS= read -r t; do
    tools_list+=("$t")
  done < "$temp_file" # читаем файл в массив

  rm "$temp_file" # удаляем временный файл

  for t in "${tools_list[@]}"; do
    echo "In loop - current ${t}"

    tool=$(basename "${t}")
    cd "${TOOLS_DIR}/${tool}" || exit 1
    sync_deps
    cd - || exit 1
  done
}

sync_tools

echo "${SCRIPT_NAME} done."
