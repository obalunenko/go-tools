FROM --platform=$BUILDPLATFORM golang:1.20 as builder
ARG TARGETPLATFORM
ARG BUILDPLATFORM

ENV PROJECT_DIR="/src/github.com/obalunenko/common-go-projects-scripts"
ENV GOBIN=${PROJECT_DIR}/bin

RUN mkdir -p "${PROJECT_DIR}"

WORKDIR "${PROJECT_DIR}"

RUN echo "I am running on ${BUILDPLATFORM}, building for ${TARGETPLATFORM}" > ./log_build.txt

COPY ./.git ./.git
COPY ./scripts ./scripts
COPY ./tools ./tools
COPY Makefile Makefile
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH make install-tools

FROM golang:1.20-alpine as base
ARG APK_GIT_VERSION=~2
ARG APK_NCURSES_VERSION=~6
ARG APK_MAKE_VERSION=~4
ARG APK_GCC_VERSION=~12
ARG APK_BASH_VERSION=~5
ARG APK_CURL_VERSION=~7
ARG APK_MUSL_DEV_VERSION=~1
ARG APK_UNZIP_VERSION=~6
ARG APK_CA_CERTIFICATES_VERSION=~20220614
ARG APK_LIBSTDC_VERSION=~12
ARG APK_BINUTILS_VERSION=~2
RUN apk update && \
    apk add --no-cache \
        "git=${APK_GIT_VERSION}" \
        "make=${APK_MAKE_VERSION}" \
        "gcc=${APK_GCC_VERSION}" \
        "bash=${APK_BASH_VERSION}" \
        "curl=${APK_CURL_VERSION}" \
        "musl-dev=${APK_MUSL_DEV_VERSION}" \
        "unzip=${APK_UNZIP_VERSION}" \
        "ca-certificates=${APK_CA_CERTIFICATES_VERSION}" \
        "libstdc++=${APK_LIBSTDC_VERSION}" \
        "binutils-gold=${APK_BINUTILS_VERSION}" && \
    rm -rf /var/cache/apk/*

# Get and install glibc for alpine
ARG APK_GLIBC_VERSION=2.35-r0
ARG APK_GLIBC_FILE="glibc-${APK_GLIBC_VERSION}.apk"
ARG APK_GLIBC_BIN_FILE="glibc-bin-${APK_GLIBC_VERSION}.apk"
ARG APK_GLIBC_BASE_URL="https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${APK_GLIBC_VERSION}"
# hadolint ignore=DL3018
# hadolint ignore=DL3018,DL3019
RUN wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub && \
    wget -nv "${APK_GLIBC_BASE_URL}/${APK_GLIBC_FILE}" && \
    apk add --force-overwrite "${APK_GLIBC_FILE}" && \
    wget -nv "${APK_GLIBC_BASE_URL}/${APK_GLIBC_BIN_FILE}" && \
    apk --no-cache add "${APK_GLIBC_BIN_FILE}" && \
    apk fix --force-overwrite alpine-baselayout-data && \
    rm glibc-*

FROM base

ENV GOROOT /usr/local/go

# don't place it into $GOPATH/bin because Drone mounts $GOPATH as volume
COPY --from=builder /src/github.com/obalunenko/common-go-projects-scripts/bin/. /usr/bin/
COPY --from=builder /src/github.com/obalunenko/common-go-projects-scripts/log_build.txt /usr/bin/log_build.txt
COPY --from=builder /src/github.com/obalunenko/common-go-projects-scripts/scripts/test/installed-tools.sh /usr/bin/installed-tools.sh

RUN /usr/bin/installed-tools.sh
