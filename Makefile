SHELL := env VERSION=$(VERSION) $(SHELL)
VERSION ?= $(shell git describe --tags $(git rev-list --tags --max-count=1))

APP_NAME?=go-tools
SHELL := env APP_NAME=$(APP_NAME) $(SHELL)

RELEASE_BRANCH?=master
SHELL := env RELEASE_BRANCH=$(RELEASE_BRANCH) $(SHELL)

# Go versions
# GOVERSION controls go.mod language version (major.minor)
# GOIMAGEVERSION controls the base image tag (major.minor[.patch]); defaults to GOVERSION
GOVERSION?=1.25
GOIMAGEVERSION?=$(GOVERSION)

TARGET_MAX_CHAR_NUM=20

## Show help
help:
	${call colored, help is running...}
	@echo ''
	@echo 'Usage:'
	@echo '  make <target>'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  %-$(TARGET_MAX_CHAR_NUM)s %s\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

## Build docker images.
build:
	IMAGE_DESCRIPTION="$$(cat README.md)" docker buildx bake
.PHONY: build

## Sync vendor.
sync-vendor:
	./scripts/sync-vendor.sh
.PHONY: sync-vendor

## Install vendored tools.
install-tools:
	./scripts/install/vendored-tools.sh
.PHONY: install-tools

check-tools:
	./scripts/test/installed-tools.sh
.PHONY: check-tools

## Issue new release.
new-version: install-tools build check-releaser
	./scripts/release/new-version.sh
.PHONY: new-release

## Release
release:
	./scripts/release/release.sh
.PHONY: release

## Check goreleaser config.
check-releaser:
	./scripts/release/check.sh
.PHONY: check-releaser

## Bump Go version across the repo (go.mod, Dockerfile, bake, CI)
# Usage examples:
#   make bump-go-version GOVERSION=1.26              # image will use 1.26
#   make bump-go-version GOVERSION=1.26 GOIMAGEVERSION=1.26.2
bump-go-version:
	./scripts/bump-go.sh $(GOVERSION) $(GOIMAGEVERSION)
.PHONY: bump-go-version

## Update versions in README from go.mod and Dockerfile
update-readme-versions:
	./scripts/update-readme-versions.sh
.PHONY: update-readme-versions

.DEFAULT_GOAL := help

