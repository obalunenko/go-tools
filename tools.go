//go:build tools

package go_tools

import (
	_ "github.com/axw/gocov/gocov"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/matm/gocov-html/cmd/gocov-html"
	_ "github.com/mattn/goveralls"
	_ "github.com/obalunenko/coverbadger/cmd/coverbadger"
	_ "github.com/posener/goreadme/cmd/goreadme"
	_ "golang.org/x/tools/cmd/cover"
	_ "golang.org/x/tools/cmd/fiximports"
	_ "golang.org/x/tools/cmd/goimports"
	_ "golang.org/x/tools/cmd/stringer"
	_ "gotest.tools/gotestsum"
)