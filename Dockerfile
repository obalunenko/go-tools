FROM --platform=$BUILDPLATFORM golang:1.21.0-alpine3.18 as builder

ARG APK_BASH_VERSION=~5
ARG APK_GIT_VERSION=~2
ARG APK_MAKE_VERSION=~4
ARG APK_OPENSSH_VERSION=~9
ARG APK_GCC_VERSION=~12
ARG APK_BUILDBASE_VERSION=~0
ARG APK_CA_CERTIFICATES_VERSION=20230506-r0
ARG APK_BINUTILS_VERSION=~2

RUN apk add --no-cache \
    "bash=${APK_BASH_VERSION}" \
	"git=${APK_GIT_VERSION}" \
	"make=${APK_MAKE_VERSION}" \
	"openssh-client=${APK_OPENSSH_VERSION}" \
	"build-base=${APK_BUILDBASE_VERSION}" \
    "gcc=${APK_GCC_VERSION}" \
    "ca-certificates=${APK_CA_CERTIFICATES_VERSION}" \
    "binutils-gold=${APK_BINUTILS_VERSION}"

ARG TARGETPLATFORM
ARG BUILDPLATFORM

ENV PROJECT_DIR="/src/github.com/obalunenko/common-go-projects-scripts"
ENV GOBIN=${PROJECT_DIR}/bin

RUN mkdir -p "${PROJECT_DIR}"

WORKDIR "${PROJECT_DIR}"

RUN echo "I am running on ${BUILDPLATFORM}, building for ${TARGETPLATFORM}" > ./log_build.txt

COPY ./.git ./.git
COPY ./scripts ./scripts
COPY tools.go tools.go
COPY go.mod go.mod
COPY go.sum go.sum
COPY Makefile Makefile

ARG TARGETOS
ARG TARGETARCH
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH make install-tools


FROM golang:1.21.0-alpine3.18 as releaser

ARG APK_BASH_VERSION=~5
ARG APK_BUILDBASE_VERSION=~0
ARG APK_GIT_VERSION=~2
ARG APK_DOCKER_CLI_VERSION=~23
ARG APK_DOCKER_CLI_COMPOSE_VERSION=~2
ARG APK_DOCKER_CLI_BUILDX_VERSION=~0
ARG APK_OPENSSH_VERSION=~9
ARG APK_TINI_VERSION=~0

RUN apk add --no-cache \
    "bash=${APK_BASH_VERSION}" \
    "git=${APK_GIT_VERSION}" \
	"build-base=${APK_BUILDBASE_VERSION}" \
    "docker-cli=${APK_DOCKER_CLI_VERSION}" \
    "docker-cli-buildx=${APK_DOCKER_CLI_BUILDX_VERSION}" \
    "docker-cli-compose=${APK_DOCKER_CLI_COMPOSE_VERSION}" \
    "openssh-client=${APK_OPENSSH_VERSION}" \
    "tini=${APK_TINI_VERSION}"

ENV GOROOT /usr/local/go

# don't place it into $GOPATH/bin because Drone mounts $GOPATH as volume
COPY --from=builder /src/github.com/obalunenko/common-go-projects-scripts/bin/. /usr/bin/

FROM releaser as tester

COPY --from=builder /src/github.com/obalunenko/common-go-projects-scripts/scripts/test/installed-tools.sh /usr/bin/installed-tools.sh

RUN /usr/bin/installed-tools.sh

FROM releaser


