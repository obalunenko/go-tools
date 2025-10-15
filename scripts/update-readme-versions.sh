#!/usr/bin/env bash
set -euo pipefail

# Update README.md tool versions and base image versions from tools/*/go.mod and Dockerfile.
# Compatible with macOS default bash (3.2) by relying on awk for maps.

root_dir=$(cd "$(dirname "$0")/.." && pwd)
cd "$root_dir"

readme="README.md"
dockerfile="Dockerfile"
tools_dir="tools"

versions_map=$(mktemp)
trap 'rm -f "$versions_map"' EXIT

trim_major_suffix() {
  local path="$1"
  echo "$path" | sed -E 's#/v([2-9]|[0-9]{2,})$##'
}

parse_first_direct_require() {
  awk '$1=="require" && NF==3 { print $2, $3; exit }' "$1"
}

# Build versions map as TSV: key\tversion
find "$tools_dir" -mindepth 2 -maxdepth 2 -name go.mod | while read -r gomod; do
  read -r mod ver <<EOF
$(parse_first_direct_require "$gomod" || true)
EOF
  if [ -z "${mod:-}" ] || [ -z "${ver:-}" ]; then
    continue
  fi
  base=$(trim_major_suffix "$mod")
  {
    echo "$mod	$ver"
    echo "https://$mod	$ver"
    echo "http://$mod	$ver"
    echo "$base	$ver"
    echo "https://$base	$ver"
    echo "http://$base	$ver"
  } >> "$versions_map"

  # Special cases where README link differs from module path
  case "$mod" in
    github.com/pressly/goose/v3/cmd/goose)
      echo "https://github.com/pressly/goose/v3/cmd/goose	$ver" >> "$versions_map"
      ;;
    github.com/maruel/panicparse/v2)
      echo "https://github.com/maruel/panicparse/v2	$ver" >> "$versions_map"
      echo "https://github.com/maruel/panicparse	$ver" >> "$versions_map"
      ;;
    gotest.tools/gotestsum)
      echo "https://gotest.tools/gotestsum	$ver" >> "$versions_map"
      ;;
    golang.org/x/tools/cmd/fiximports)
      echo "https://golang.org/x/tools/cmd/fiximports	$ver" >> "$versions_map"
      ;;
    golang.org/x/tools/cmd/goimports)
      echo "https://golang.org/x/tools/cmd/goimports	$ver" >> "$versions_map"
      ;;
    golang.org/x/tools/cmd/stringer)
      echo "https://golang.org/x/tools/cmd/stringer	$ver" >> "$versions_map"
      ;;
    github.com/golangci/golangci-lint/v2)
      echo "https://github.com/golangci/golangci-lint	$ver" >> "$versions_map"
      echo "github.com/golangci/golangci-lint/v2	$ver" >> "$versions_map"
      ;;
    github.com/oapi-codegen/oapi-codegen)
      echo "https://github.com/oapi-codegen/oapi-codegen	$ver" >> "$versions_map"
      ;;
    github.com/k1LoW/tbls)
      echo "https://github.com/k1LoW/tbls	$ver" >> "$versions_map"
      ;;
    github.com/swaggo/swag/cmd/swag)
      echo "https://github.com/swaggo/swag/cmd/swag	$ver" >> "$versions_map"
      echo "github.com/swaggo/swag/cmd/swag	$ver" >> "$versions_map"
      ;;
    github.com/go-swagger/go-swagger)
      echo "https://github.com/go-swagger/go-swagger	$ver" >> "$versions_map"
      ;;
    github.com/sqlc-dev/sqlc)
      echo "https://github.com/sqlc-dev/sqlc	$ver" >> "$versions_map"
      ;;
    github.com/abice/go-enum)
      echo "https://github.com/abice/go-enum	$ver" >> "$versions_map"
      ;;
    github.com/vektra/mockery/v3)
      echo "https://github.com/vektra/mockery	$ver" >> "$versions_map"
      ;;
    go.uber.org/mock)
      echo "https://github.com/uber-go/mock	$ver" >> "$versions_map"
      ;;
    github.com/axw/gocov/gocov)
      echo "https://github.com/axw/gocov	$ver" >> "$versions_map"
      ;;
    github.com/matm/gocov-html)
      echo "https://github.com/matm/gocov-html	$ver" >> "$versions_map"
      ;;
    mvdan.cc/gofumpt)
      echo "https://mvdan.cc/gofumpt	$ver" >> "$versions_map"
      ;;
    github.com/mattn/goveralls)
      echo "https://github.com/mattn/goveralls	$ver" >> "$versions_map"
      ;;
    golang.org/x/vuln/cmd/govulncheck)
      echo "https://golang.org/x/vuln/cmd/govulncheck	$ver" >> "$versions_map"
      ;;
    github.com/hexdigest/gowrap)
      echo "https://github.com/hexdigest/gowrap	$ver" >> "$versions_map"
      ;;
    github.com/obalunenko/coverbadger)
      echo "https://github.com/obalunenko/coverbadger	$ver" >> "$versions_map"
      ;;
    github.com/alvaroloes/enumer)
      echo "https://github.com/alvaroloes/enumer	$ver" >> "$versions_map"
      ;;
    github.com/g4s8/envdoc)
      echo "https://github.com/g4s8/envdoc	$ver" >> "$versions_map"
      ;;
    github.com/bufbuild/buf)
      echo "https://github.com/bufbuild/buf	$ver" >> "$versions_map"
      ;;
  esac
done

# Extract GO_VERSION and ALPINE_VERSION from Dockerfile
GO_VERSION=$(sed -nE 's/^ARG[[:space:]]+GO_VERSION=([^[:space:]]+)$/\1/p' "$dockerfile")
ALPINE_VERSION=$(sed -nE 's/^ARG[[:space:]]+ALPINE_VERSION=([^[:space:]]+)$/\1/p' "$dockerfile")

tmp=$(mktemp)
trap 'rm -f "$tmp" "$versions_map"' EXIT

awk -v gover="$GO_VERSION" -v alpinever="$ALPINE_VERSION" -v mapfile="$versions_map" '
  BEGIN {
    FS="\n"; OFS="\n";
    # Load versions map
    while ((getline < mapfile) > 0) {
      split($0, kv, "\t");
      if (length(kv) >= 2) {
        ver[kv[1]] = kv[2];
      }
    }
    close(mapfile)
  }
  function update_base_image(line,   r) {
    if (index(line, "official image") && index(line, "golang:") && index(line, "alpine")) {
      gsub(/golang:[0-9]+\.[0-9]+(\.[0-9]+)?-alpine[0-9]+\.[0-9]+/, "golang:" gover "-alpine" alpinever, line)
    }
    return line
  }
  function extract_url(line,   p, rest, rparen, url) {
    # Find the first "](" then read until the next ")"
    p = index(line, "](")
    if (!p) return ""
    rest = substr(line, p+2)
    rparen = index(rest, ")")
    if (!rparen) return ""
    url = substr(rest, 1, rparen-1)
    gsub(/^ +| +$/, "", url)
    if (url ~ /^https?:\/\//) return url
    return ""
  }
  function lookup_version(url,   key) {
    if (url in ver) return ver[url]
    sub(/^https:\/\//, "", url)
    sub(/^http:\/\//, "", url)
    if (url in ver) return ver[url]
    return ""
  }
  function update_table_row(line,   url, v, n, i, parts) {
    url = extract_url(line)
    if (url == "") return line
    v = lookup_version(url)
    if (v == "") return line
    n = split(line, parts, "|")
    if (n < 4) return line
    parts[3] = " " v " "
    line = parts[1]
    for (i=2; i<=n; i++) line = line "|" parts[i]
    return line
  }
  {
    if ($0 ~ /^\|/) {
      print update_base_image(update_table_row($0))
    } else {
      print update_base_image($0)
    }
  }
' "$readme" > "$tmp"

if ! cmp -s "$readme" "$tmp"; then
  mv "$tmp" "$readme"
  echo "README.md updated"
else
  echo "README.md is up-to-date"
fi
