[![Latest release artifacts](https://img.shields.io/github/v/release/obalunenko/go-tools)](https://github.com/obalunenko/go-tools/releases/latest)

# Go Tools

This repository, **go-tools**,
includes a compilation of common tools used across various projects by the repository owner.

These accompany a Docker image, which you can swiftly pull using the following command: 

```shell
docker pull ghcr.io/obalunenko/go-tools:latest
```

## Contents

### Go base image

The base image is built on top of the official image [golang:1.24.6-alpine3.22](https://hub.docker.com/_/golang) and includes the following tools:


##### Below is a full manifest of the tools available


| Tool                                                       | Version                            | Description                                                                                                                                                                       |
|------------------------------------------------------------|------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [buf](https://github.com/bufbuild/buf)                     | v1.56.0                            | The buf CLI is the best tool for working with Protocol Buffers                                                                                                                    |
| [coverbadger](https://github.com/obalunenko/coverbadger)   | v1.4.0                             | Generate coverage badge images for Markdown files using Go                                                                                                                        |
| [enumer](https://github.com/alvaroloes/enumer)             | v1.1.2                             | A Go tool to auto generate methods for your enums                                                                                                                                 |
| [envdoc](https://github.com/g4s8/envdoc)                   | v1.6.0                             | Go tool to generate documentation for environment variables                                                                                                                       |
| [fiximports](https://golang.org/x/tools/cmd/fiximports)    | v0.35.0                            | The fiximports command fixes import declarations to use the canonical import path for packages that have an "import comment" as defined by https://golang.org/s/go14customimport. |
| [go-enum](https://github.com/abice/go-enum)                | v0.9.1                             | An enum generator for go                                                                                                                                                          |
| [gocov](https://github.com/axw/gocov)                      | v1.2.1                             | Coverage testing tool for The Go Programming Language                                                                                                                             |
| [gocov-html](https://github.com/matm/gocov-html)           | v1.4.0                             | Make pretty HTML output from gocov, a coverage testing tool for Go                                                                                                                |
| [gofumpt](https://mvdan.cc/gofumpt)                        | v0.8.0                             | A stricter gofmt                                                                                                                                                                  |
| [goimports](https://golang.org/x/tools/cmd/goimports)      | v0.35.0                            | Command goimports updates your Go import lines, adding missing ones and removing unreferenced ones                                                                                |
| [golangci-lint](https://github.com/golangci/golangci-lint) | v2.3.1                             | Fast linters Runner for Go                                                                                                                                                        |
| [goose](https://github.com/pressly/goose/v3/cmd/goose)     | v3.24.2                            | A database migration tool                                                                                                                                                         |
| [goreadme](https://github.com/posener/goreadme)            | v1.4.2                             | Generate readme file from Go doc                                                                                                                                                  |
| [goreleaser](https://github.com/goreleaser/goreleaser)     | v2.11.2                             | Deliver Go binaries as fast and easily as possible                                                                                                                                |
| [gotestsum](https://gotest.tools/gotestsum)                | v1.12.3                            | 'go test' runner with output optimized for humans, JUnit XML for CI integration, and a summary of the test results.                                                               |
| [goveralls](https://github.com/mattn/goveralls)            | v0.0.12                            | Go integration for Coveralls.io continuous code coverage tracking system.                                                                                                         |
| [govulncheck](https://golang.org/x/vuln/cmd/govulncheck)   | v1.1.4                             | Govulncheck reports known vulnerabilities that affect Go code.                                                                                                                    |
| [pkgsite](https://golang.org/x/pkgsite/cmd/pkgsite)        | v0.0.0-20250606033525-6805ff32e9c8 | Pkgsite extracts and generates documentation for Go programs. It runs as a web server and presents the documentation as a web page.                                               |
| [stringer](https://golang.org/x/tools/cmd/stringer)        | v0.35.0                            | Stringer is a tool to automate the creation of methods that satisfy the fmt.Stringer interface.                                                                                   |
| [svu](https://github.com/caarlos0/svu)                     | v3.2.3                             | Semantic Version Util                                                                                                                                                             |
| [swag](github.com/swaggo/swag/cmd/swag)                    | v2.0.0-rc4                         | Automatically generate RESTful API documentation with Swagger 2.0 for Go                                                                                                          |
| [swagger](https://github.com/go-swagger/go-swagger)        | v0.32.3                            | Client/Server from OpenAPI docs generation tool                                                                                                                                   |
| [tparse](https://github.com/mfridman/tparse)               | v0.17.0                            | CLI tool for summarizing go test output. Pipe friendly. CI/CD friendly.                                                                                                           |


## How to Use

The go-tools Docker image is engineered to allow for seamless implementation with your code.

 - Pull the image
 - Run the image with your code mounted as a volume
 - Execute the desired commands.

This container is available at 
[go-tools container on GitHub Packages](https://github.com/obalunenko/go-tools/pkgs/container/go-tools)

