# Go Tools

[![Latest release artifacts](https://img.shields.io/github/v/release/obalunenko/go-tools)](https://github.com/obalunenko/go-tools/releases/latest)

This repository, **go-tools**,
includes a compilation of common tools used across various projects by the repository owner.

These accompany a Docker image, which you can swiftly pull using the following command:

```shell
docker pull ghcr.io/obalunenko/go-tools:latest
```

## Contents

### Go base image

The base image is built on top of the official image [golang:1.27.4-alpine3.24](https://hub.docker.com/_/golang) and includes the following tools:

#### Below is a full manifest of the tools available

| Tool                                                         | Version                            | Description                                                                                                                                                                       |
|--------------------------------------------------------------|------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| [buf](https://github.com/bufbuild/buf)                       | v1.75.0 | The buf CLI is the best tool for working with Protocol Buffers                                                                                                                    |
| [gocov](https://github.com/axw/gocov)                        | v1.2.1 | Coverage testing tool for The Go Programming Language                                                                                                                             |
| [gocov-html](https://github.com/matm/gocov-html)             | v1.4.0 | Make pretty HTML output from gocov, a coverage testing tool for Go                                                                                                                |
| [goimports](https://golang.org/x/tools/cmd/goimports)        | v0.36.0                            | Command goimports updates your Go import lines, adding missing ones and removing unreferenced ones                                                                                |
| [golangci-lint](https://github.com/golangci/golangci-lint)   | v2.15.0 | Fast linters Runner for Go                                                                                                                                                        |
| [goose](https://github.com/pressly/goose/v3/cmd/goose)       | v3.24.2                            | A database migration tool                                                                                                                                                         |
| [goreleaser](https://github.com/goreleaser/goreleaser)       | v2.20.2 | Deliver Go binaries as fast and easily as possible                                                                                                                                |
| [gotestsum](https://gotest.tools/gotestsum)                  | v1.13.0 | 'go test' runner with output optimized for humans, JUnit XML for CI integration, and a summary of the test results.                                                               |
| [govulncheck](https://golang.org/x/vuln/cmd/govulncheck)     | v1.1.4                             | Govulncheck reports known vulnerabilities that affect Go code.                                                                                                                    |
| [gowrap](https://github.com/hexdigest/gowrap)                | v1.5.0 | GoWrap is a command line tool for generating decorators for Go interfaces                                                                                                         |
| [mockgen](https://github.com/uber-go/mock)                   | v0.6.0 | GoMock is a mocking framework for the Go programming language.                                                                                                                    |
| [pp](https://github.com/maruel/panicparse/v2)          | v2.5.0 | Panicparse Parses panic stack traces, densifies and deduplicates goroutines with similar stack traces. Helps debugging crashes and deadlocks in heavily parallelized processes.                                               |
| [sqlc](https://github.com/sqlc-dev/sqlc)                     | v1.31.1 | Generate type-safe code from SQL                                                                                                                                                  |
| [stringer](https://golang.org/x/tools/cmd/stringer)          | v0.36.0                            | Stringer is a tool to automate the creation of methods that satisfy the fmt.Stringer interface.                                                                                   |
| [svu](https://github.com/caarlos0/svu)                       | v3.4.4 | Semantic Version Util                                                                                                                                                             |
| [swag](github.com/swaggo/swag/cmd/swag)                      | v2.0.0-rc4                         | Automatically generate RESTful API documentation with Swagger 2.0 for Go                                                                                                          |
| [swagger](https://github.com/go-swagger/go-swagger)          | v0.36.7 | Client/Server from OpenAPI docs generation tool                                                                                                                                   |
| [tbls](https://github.com/k1LoW/tbls)                 | v1.96.1 | tbls is a CI-Friendly tool to document a database, written in Go.                                                                                                           |

## How to Use

The go-tools Docker image is engineered to allow for seamless implementation with your code.

- Pull the image
- Run the image with your code mounted as a volume
- Execute the desired commands.

This container is available at
[go-tools container on GitHub Packages](https://github.com/obalunenko/go-tools/pkgs/container/go-tools)

## Development

Tool modules under `tools/` are built in module mode. Use `make install-tools` to build and install tool binaries, and `make list-tool-updates` to check whether dependencies have newer versions available.

Use `make build-local` to build the image for your platform and load it as `ghcr.io/obalunenko/go-tools:local`. Builds reuse the layer cache that CI publishes from master, so only tools that differ from master are compiled. Each tool has its own stage in the `Dockerfile`; a new tool needs a stage and a `COPY` line there.
