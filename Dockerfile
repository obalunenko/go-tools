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

ENV GLIBC_VERSION 2.35-r0

# Download and install glibc
RUN apk add --update curl && \
  curl -Lo /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub && \
  curl -Lo glibc.apk "https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/glibc-${GLIBC_VERSION}.apk" && \
  curl -Lo glibc-bin.apk "https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/glibc-bin-${GLIBC_VERSION}.apk" && \
  apk add --force-overwrite glibc-bin.apk glibc.apk && \
  /usr/glibc-compat/sbin/ldconfig /lib /usr/glibc-compat/lib && \
  echo 'hosts: files mdns4_minimal [NOTFOUND=return] dns mdns4' >> /etc/nsswitch.conf && \
  apk del curl && \
  rm -rf glibc.apk glibc-bin.apk /var/cache/apk/*

FROM base

ENV GOROOT /usr/local/go

# don't place it into $GOPATH/bin because Drone mounts $GOPATH as volume
COPY --from=builder /src/github.com/obalunenko/common-go-projects-scripts/bin/. /usr/bin/
COPY --from=builder /src/github.com/obalunenko/common-go-projects-scripts/log_build.txt /usr/bin/log_build.txt
COPY --from=builder /src/github.com/obalunenko/common-go-projects-scripts/scripts/test/installed-tools.sh /usr/bin/installed-tools.sh

RUN /usr/bin/installed-tools.sh
