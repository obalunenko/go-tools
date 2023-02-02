#!/bin/bash

set -eu

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(dirname "$0")"
REPO_ROOT="$(cd "${SCRIPT_DIR}" && git rev-parse --show-toplevel)"
TOOLS_DIR="${REPO_ROOT}/tools"

echo "${SCRIPT_NAME} is running... "

echo "Gocov check: $(gocov --help)"

echo "golangci-lint check: $(golangci-lint version)"

echo "gocov-htm check: $(gocov-html -v)"

echo "goveralls check: $(goveralls --help)"

echo "coverbadger check: $(coverbadger --help)"

echo "cover check: $(cover --help)"

echo "goimports check: $(goimports --help)"

echo "stringer check: $(stringer --help)"

echo "fiximports check: $(fiximports --help)"

echo "gotestsum check: $(gotestsum --help)"

echo "${SCRIPT_NAME} done "