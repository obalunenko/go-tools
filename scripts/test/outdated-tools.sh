#!/usr/bin/env bash

set -Eeuo pipefail

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd -P)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/../.." && pwd -P)"
TOOLS_DIR="${TOOLS_DIR:-${REPO_ROOT}/tools}"

echo "[$SCRIPT_NAME]: checking tools for updates..."

if [[ ! -d "$TOOLS_DIR" ]]; then
  echo "[FATAL]: tools directory not found: $TOOLS_DIR"
  exit 2
fi

for cmd in go find sort sed; do
  if ! command -v "$cmd" >/dev/null 2>&1; then
    echo "[FATAL]: required command '$cmd' not found"
    exit 2
  fi
done

# Use writable temp caches so the script can run in restricted environments
# without polluting the repository with cache files.
GO_CACHE_ROOT="${GO_CACHE_ROOT:-${TMPDIR:-/tmp}/go-tools-cache}"
GOMODCACHE="${GOMODCACHE:-${GO_CACHE_ROOT}/go-mod}"
GOCACHE="${GOCACHE:-${GO_CACHE_ROOT}/go-build}"
mkdir -p "$GOMODCACHE" "$GOCACHE"
export GOMODCACHE GOCACHE

tmp_results="$(mktemp)"
trap 'rm -f "$tmp_results"' EXIT

failed_checks=0
checked_modules=0

while IFS= read -r -d '' dir; do
  [[ -f "$dir/go.mod" ]] || continue

  tool_name="$(basename "$dir")"
  if ! imports="$(cd "$dir" && go list -mod=mod -e -f '{{ join .Imports "\n" }}' -tags=tools . 2>/dev/null | sed '/^$/d' | sort -u)"; then
    echo "[WARN]: ${tool_name}: failed to enumerate tool imports"
    failed_checks=1
    continue
  fi
  [[ -n "$imports" ]] || continue

  while IFS= read -r dep; do
    [[ -n "$dep" ]] || continue

    if ! resolved_module="$(cd "$dir" && go list -mod=mod -f '{{with .Module}}{{.Path}}|{{.Version}}{{end}}' "$dep")"; then
      echo "[WARN]: ${tool_name}: failed to resolve module for ${dep}"
      failed_checks=1
      continue
    fi
    [[ -n "$resolved_module" ]] || continue

    IFS='|' read -r module_path current_version <<< "$resolved_module"
    [[ -n "${module_path:-}" ]] || continue

    if ! module_info="$(cd "$dir" && go list -mod=mod -u -m -f '{{.Path}}|{{.Version}}|{{if .Update}}{{.Update.Version}}{{end}}' "$module_path")"; then
      echo "[WARN]: ${tool_name}: failed to check updates for module ${module_path}"
      failed_checks=1
      continue
    fi
    [[ -n "$module_info" ]] || continue
    checked_modules=$((checked_modules + 1))

    IFS='|' read -r _ _ latest_version <<< "$module_info"
    if [[ -n "${latest_version:-}" && "$latest_version" != "$current_version" ]]; then
      printf '%s\t%s\t%s\t%s\n' "$tool_name" "$module_path" "$current_version" "$latest_version" >> "$tmp_results"
    fi
  done <<< "$imports"
done < <(find "$TOOLS_DIR" -mindepth 1 -maxdepth 1 -type d -print0)

if [[ ! -s "$tmp_results" ]]; then
  if [[ "$failed_checks" -ne 0 ]]; then
    if [[ "$checked_modules" -eq 0 ]]; then
      echo "[ERROR]: no modules were checked successfully."
    else
      echo "[ERROR]: some tools could not be checked; results are incomplete."
    fi
    exit 1
  fi
  echo "[OK]: all tools are up to date."
  exit 0
fi

echo "[INFO]: tools with available updates:"
printf '%-20s %-45s %-18s %s\n' "tool" "module" "current" "latest"
printf '%-20s %-45s %-18s %s\n' "----" "------" "-------" "------"
sort -u "$tmp_results" | while IFS=$'\t' read -r tool module current latest; do
  printf '%-20s %-45s %-18s %s\n' "$tool" "$module" "$current" "$latest"
done

if [[ "$failed_checks" -ne 0 ]]; then
  echo "[ERROR]: some tools could not be checked; list above is partial."
  exit 1
fi
