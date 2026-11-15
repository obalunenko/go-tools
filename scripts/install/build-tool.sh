#!/usr/bin/env bash
# Build every tool imported by tools.go of a single tools/<name> module.
#
# Usage: build-tool.sh <tool-module-dir> <out-dir>
#
# Binaries are stripped (-s -w) and built with -trimpath. Set GO_BUILD_P to
# limit go build parallelism.

set -Eeuo pipefail

if [[ $# -ne 2 ]]; then
  echo "usage: $(basename "$0") <tool-module-dir> <out-dir>"
  exit 2
fi

TOOL_DIR="$1"
OUT_DIR="$2"

mkdir -p "$OUT_DIR"
OUT_DIR="$(cd "$OUT_DIR" && pwd -P)"

cd "$TOOL_DIR"

# -e: tools.go imports main packages, which go list otherwise reports as errors.
imports="$(go list -mod=readonly -e -tags=tools -f '{{ join .Imports "\n" }}' .)"

deps=()
while IFS= read -r dep; do
  [[ -n "$dep" ]] && deps+=("$dep")
done <<< "$imports"

if [[ ${#deps[@]} -eq 0 ]]; then
  echo "[ERROR]: ${TOOL_DIR}: tools.go imports no tools"
  exit 255
fi

build_flags=(-mod=readonly -trimpath -ldflags='-s -w')
if [[ -n "${GO_BUILD_P:-}" ]]; then
  build_flags+=(-p "$GO_BUILD_P")
fi

echo "[INFO]: building ${deps[*]} -> ${OUT_DIR}"

# A trailing slash makes go build write one binary per main package, named
# after the import path without its /vN suffix.
if ! go build "${build_flags[@]}" -o "${OUT_DIR}/" "${deps[@]}"; then
  echo "[ERROR]: build failed for ${deps[*]}"
  exit 255
fi

echo "[OK]:   built ${deps[*]}"
