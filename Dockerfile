ARG GO_VERSION=1.27.4
ARG ALPINE_VERSION=3.24
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS builder

ARG APK_BASH_VERSION=~5
ARG APK_GIT_VERSION=~2
ARG APK_GCC_VERSION=~15
ARG APK_MUSL_DEV_VERSION=~1

RUN apk add --no-cache \
    "bash=${APK_BASH_VERSION}" \
    "git=${APK_GIT_VERSION}" \
    "gcc=${APK_GCC_VERSION}" \
    "musl-dev=${APK_MUSL_DEV_VERSION}"

ARG TARGETOS
ARG TARGETARCH

# Native builds get cgo (tbls needs it for its sqlite driver). Cross builds fall
# back to CGO_ENABLED=0, which is why CI builds each platform on a native runner.
ENV GOOS=${TARGETOS} \
    GOARCH=${TARGETARCH} \
    GOCACHE=/root/.cache/go-build \
    GOMODCACHE=/go/pkg/mod \
    GO_BUILD_P=4

WORKDIR /src

# One stage per tool: a stage only sees its own tools/<name> module, so bumping
# one tool rebuilds only that tool and the rest come from the layer cache.
# Stages sharing a locked build-slot mount run one at a time, which caps a cold
# build at three concurrent go builds so it fits into CI runner and Docker
# Desktop memory. Slots are only balanced by build time; any slot works.

FROM builder AS buf
RUN --mount=type=bind,source=tools/buf,target=. \
    --mount=type=bind,source=scripts/install/build-tool.sh,target=/usr/local/bin/build-tool \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,id=build-slot-2,target=/run/build-slot,sharing=locked \
    build-tool . /out


FROM builder AS gocov
RUN --mount=type=bind,source=tools/gocov,target=. \
    --mount=type=bind,source=scripts/install/build-tool.sh,target=/usr/local/bin/build-tool \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,id=build-slot-2,target=/run/build-slot,sharing=locked \
    build-tool . /out

FROM builder AS gocov-html
RUN --mount=type=bind,source=tools/gocov-html,target=. \
    --mount=type=bind,source=scripts/install/build-tool.sh,target=/usr/local/bin/build-tool \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,id=build-slot-3,target=/run/build-slot,sharing=locked \
    build-tool . /out

FROM builder AS goimports
RUN --mount=type=bind,source=tools/goimports,target=. \
    --mount=type=bind,source=scripts/install/build-tool.sh,target=/usr/local/bin/build-tool \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,id=build-slot-1,target=/run/build-slot,sharing=locked \
    build-tool . /out

FROM builder AS golangci-lint
RUN --mount=type=bind,source=tools/golangci-lint,target=. \
    --mount=type=bind,source=scripts/install/build-tool.sh,target=/usr/local/bin/build-tool \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,id=build-slot-3,target=/run/build-slot,sharing=locked \
    build-tool . /out

FROM builder AS goose
RUN --mount=type=bind,source=tools/goose,target=. \
    --mount=type=bind,source=scripts/install/build-tool.sh,target=/usr/local/bin/build-tool \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,id=build-slot-3,target=/run/build-slot,sharing=locked \
    build-tool . /out


FROM builder AS goreleaser
RUN --mount=type=bind,source=tools/goreleaser,target=. \
    --mount=type=bind,source=scripts/install/build-tool.sh,target=/usr/local/bin/build-tool \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,id=build-slot-1,target=/run/build-slot,sharing=locked \
    build-tool . /out

FROM builder AS gotestsum
RUN --mount=type=bind,source=tools/gotestsum,target=. \
    --mount=type=bind,source=scripts/install/build-tool.sh,target=/usr/local/bin/build-tool \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,id=build-slot-2,target=/run/build-slot,sharing=locked \
    build-tool . /out


FROM builder AS govulncheck
RUN --mount=type=bind,source=tools/govulncheck,target=. \
    --mount=type=bind,source=scripts/install/build-tool.sh,target=/usr/local/bin/build-tool \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,id=build-slot-2,target=/run/build-slot,sharing=locked \
    build-tool . /out

FROM builder AS gowrap
RUN --mount=type=bind,source=tools/gowrap,target=. \
    --mount=type=bind,source=scripts/install/build-tool.sh,target=/usr/local/bin/build-tool \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,id=build-slot-3,target=/run/build-slot,sharing=locked \
    build-tool . /out


FROM builder AS mockgen
RUN --mount=type=bind,source=tools/mockgen,target=. \
    --mount=type=bind,source=scripts/install/build-tool.sh,target=/usr/local/bin/build-tool \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,id=build-slot-3,target=/run/build-slot,sharing=locked \
    build-tool . /out


FROM builder AS panicparse
RUN --mount=type=bind,source=tools/panicparse,target=. \
    --mount=type=bind,source=scripts/install/build-tool.sh,target=/usr/local/bin/build-tool \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,id=build-slot-1,target=/run/build-slot,sharing=locked \
    build-tool . /out

FROM builder AS sqlc
RUN --mount=type=bind,source=tools/sqlc,target=. \
    --mount=type=bind,source=scripts/install/build-tool.sh,target=/usr/local/bin/build-tool \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,id=build-slot-1,target=/run/build-slot,sharing=locked \
    build-tool . /out

FROM builder AS stringer
RUN --mount=type=bind,source=tools/stringer,target=. \
    --mount=type=bind,source=scripts/install/build-tool.sh,target=/usr/local/bin/build-tool \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,id=build-slot-3,target=/run/build-slot,sharing=locked \
    build-tool . /out

FROM builder AS svu
RUN --mount=type=bind,source=tools/svu,target=. \
    --mount=type=bind,source=scripts/install/build-tool.sh,target=/usr/local/bin/build-tool \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,id=build-slot-3,target=/run/build-slot,sharing=locked \
    build-tool . /out

FROM builder AS swag
RUN --mount=type=bind,source=tools/swag,target=. \
    --mount=type=bind,source=scripts/install/build-tool.sh,target=/usr/local/bin/build-tool \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,id=build-slot-3,target=/run/build-slot,sharing=locked \
    build-tool . /out

FROM builder AS swagger
RUN --mount=type=bind,source=tools/swagger,target=. \
    --mount=type=bind,source=scripts/install/build-tool.sh,target=/usr/local/bin/build-tool \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,id=build-slot-2,target=/run/build-slot,sharing=locked \
    build-tool . /out

FROM builder AS tbls
RUN --mount=type=bind,source=tools/tbls,target=. \
    --mount=type=bind,source=scripts/install/build-tool.sh,target=/usr/local/bin/build-tool \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,id=build-slot-2,target=/run/build-slot,sharing=locked \
    build-tool . /out


FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS releaser

ARG APK_BASH_VERSION=~5
ARG APK_GIT_VERSION=~2
ARG APK_GCC_VERSION=~15
ARG APK_MUSL_DEV_VERSION=~1
ARG APK_MAKE_VERSION=~4
ARG APK_OPENSSH_VERSION=~10
ARG APK_TINI_VERSION=~0

# gcc and musl-dev are enough for cgo and go test -race.
RUN apk add --no-cache \
    "bash=${APK_BASH_VERSION}" \
    "git=${APK_GIT_VERSION}" \
    "gcc=${APK_GCC_VERSION}" \
    "musl-dev=${APK_MUSL_DEV_VERSION}" \
    "make=${APK_MAKE_VERSION}" \
    "openssh-client=${APK_OPENSSH_VERSION}" \
    "tini=${APK_TINI_VERSION}"

ENV GOROOT=/usr/local/go

# A layer per tool, independent of the base image: an update only ships the
# layers of the tools that changed.
COPY --link --from=buf /out/ /usr/bin/
COPY --link --from=gocov /out/ /usr/bin/
COPY --link --from=gocov-html /out/ /usr/bin/
COPY --link --from=goimports /out/ /usr/bin/
COPY --link --from=golangci-lint /out/ /usr/bin/
COPY --link --from=goose /out/ /usr/bin/
COPY --link --from=goreleaser /out/ /usr/bin/
COPY --link --from=gotestsum /out/ /usr/bin/
COPY --link --from=govulncheck /out/ /usr/bin/
COPY --link --from=gowrap /out/ /usr/bin/
COPY --link --from=mockgen /out/ /usr/bin/
COPY --link --from=panicparse /out/ /usr/bin/
COPY --link --from=sqlc /out/ /usr/bin/
COPY --link --from=stringer /out/ /usr/bin/
COPY --link --from=svu /out/ /usr/bin/
COPY --link --from=swag /out/ /usr/bin/
COPY --link --from=swagger /out/ /usr/bin/
COPY --link --from=tbls /out/ /usr/bin/
