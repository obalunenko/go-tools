#!/bin/bash

set -eu

SCRIPT_NAME="$(basename "$0")"

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

echo "goreadme check: $(goreadme -h)"

echo "goreleaser check: $(goreleaser -v)"

echo "govulncheck check: $(govulncheck -h)

echo "${SCRIPT_NAME} done "
