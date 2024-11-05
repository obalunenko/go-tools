[![Latest release artifacts](https://img.shields.io/github/v/release/obalunenko/go-tools)](https://github.com/obalunenko/go-tools/releases/latest)

# Go Tools

This repository, **go-tools**,
includes a compilation of common tools used across various projects by the repository owner.

These accompany a Docker image which you can swiftly pull using the following command: 

```shell
docker pull ghcr.io/obalunenko/go-tools:latest
```

## Contents

### Go base image

The base image is built on top of the official image [golang:1.23-alpine3.20](https://hub.docker.com/_/golang) and includes the following tools:


##### Below is a full manifest of the tools available


| Tool                                                       | Version                            | Description                                                                                                                                                                       |
|------------------------------------------------------------|------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [go-enum](https://github.com/abice/go-enum)                | v0.6.0                             | An enum generator for go                                                                                                                                                          |
| [enumer](https://github.com/alvaroloes/enumer)             | v1.1.2                             | A Go tool to auto generate methods for your enums                                                                                                                                 |
| [gocov](https://github.com/axw/gocov)                      | v1.1.0                             | Coverage testing tool for The Go Programming Language                                                                                                                             |
| [svu](https://github.com/caarlos0/svu)                     | v2.2.0                             | Semantic Version Util                                                                                                                                                             |
| [swagger](https://github.com/go-swagger/go-swagger)        | v0.31.0                            | Client/Server from OpenAPI docs generation tool                                                                                                                                   |
| [golangci-lint](https://github.com/golangci/golangci-lint) | v1.60.3                            | Fast linters Runner for Go                                                                                                                                                        |
| [goreleaser](https://github.com/goreleaser/goreleaser)     | v2.2.0                             | Deliver Go binaries as fast and easily as possible                                                                                                                                |
| [gocov-html](https://github.com/matm/gocov-html)           | v1.4.0                             | Make pretty HTML output from gocov, a coverage testing tool for Go                                                                                                                |
| [goveralls](https://github.com/mattn/goveralls)            | v0.0.12                            | Go integration for Coveralls.io continuous code coverage tracking system.                                                                                                         |
| [tparse](https://github.com/mfridman/tparse)               | v0.14.0                            | CLI tool for summarizing go test output. Pipe friendly. CI/CD friendly.                                                                                                           |
| [coverbadger](https://github.com/obalunenko/coverbadger)   | v1.4.0                             | Generate coverage badge images for Markdown files using Go                                                                                                                        |
| [goreadme](https://github.com/posener/goreadme)            | v1.4.2                             | Generate readme file from Go doc                                                                                                                                                  |
| [goose](https://github.com/pressly/goose/v3/cmd/goose)     | v3.22.0                            | A database migration tool                                                                                                                                                         |
| [swag](github.com/swaggo/swag/cmd/swag)                    | v1.16.3                            | Automatically generate RESTful API documentation with Swagger 2.0 for Go                                                                                                          |
| [pkgsite](https://golang.org/x/pkgsite/cmd/pkgsite)        | v0.0.0-20240905030440-6b577b411ef5 | Pkgsite extracts and generates documentation for Go programs. It runs as a web server and presents the documentation as a web page.                                               |
| [fiximports](https://golang.org/x/tools/cmd/fiximports)    | v0.24.0                            | The fiximports command fixes import declarations to use the canonical import path for packages that have an "import comment" as defined by https://golang.org/s/go14customimport. |
| [goimports](https://golang.org/x/tools/cmd/goimports)      | v0.24.0                            | Command goimports updates your Go import lines, adding missing ones and removing unreferenced ones                                                                                |
| [stringer](https://golang.org/x/tools/cmd/stringer)        | v0.24.0                            | Stringer is a tool to automate the creation of methods that satisfy the fmt.Stringer interface.                                                                                   |
| [govulncheck](https://golang.org/x/vuln/cmd/govulncheck)   | v1.1.3                             | Govulncheck reports known vulnerabilities that affect Go code.                                                                                                                    |
| [gotestsum](https://gotest.tools/gotestsum)                | v1.12.0                            | 'go test' runner with output optimized for humans, JUnit XML for CI integration, and a summary of the test results.                                                               |
| [gofumpt](https://mvdan.cc/gofumpt)                        | v0.7.0                             | A stricter gofmt                                                                                                                                                                  |


## How to Use

The go-tools Docker image is engineered to allow for seamless implementation with your code.

 - Pull the image
 - Run the image with your code mounted as a volume
 - Execute the desired commands.

This container is available at 
[go-tools container on GitHub Packages](https://github.com/obalunenko/go-tools/pkgs/container/go-tools)

