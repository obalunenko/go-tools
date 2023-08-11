//go:build tools

package go_tools

import (
	_ "github.com/abice/go-enum"
	_ "github.com/alvaroloes/enumer"
	_ "github.com/axw/gocov/gocov"
	_ "github.com/caarlos0/svu"
	_ "github.com/go-swagger/go-swagger/cmd/swagger"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/goreleaser/goreleaser"
	_ "github.com/matm/gocov-html/cmd/gocov-html"
	_ "github.com/mattn/goveralls"
	_ "github.com/mfridman/tparse"
	_ "github.com/obalunenko/coverbadger/cmd/coverbadger"
	_ "github.com/posener/goreadme/cmd/goreadme"
	_ "github.com/pressly/goose/v3/cmd/goose"
	_ "golang.org/x/pkgsite/cmd/pkgsite"
	_ "golang.org/x/tools/cmd/fiximports"
	_ "golang.org/x/tools/cmd/goimports"
	_ "golang.org/x/tools/cmd/stringer"
	_ "golang.org/x/vuln/cmd/govulncheck"
	_ "gotest.tools/gotestsum"
	_ "mvdan.cc/gofumpt"
)
