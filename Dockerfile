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
COPY tools.go tools.go
COPY vendor vendor
COPY go.mod go.mod
COPY go.sum go.sum
COPY Makefile Makefile

ARG TARGETOS
ARG TARGETARCH
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH make install-tools


FROM golang:1.20

ENV GOROOT /usr/local/go

# don't place it into $GOPATH/bin because Drone mounts $GOPATH as volume
COPY --from=builder /src/github.com/obalunenko/common-go-projects-scripts/bin/. /usr/bin/
COPY --from=builder /src/github.com/obalunenko/common-go-projects-scripts/log_build.txt /usr/bin/log_build.txt
COPY --from=builder /src/github.com/obalunenko/common-go-projects-scripts/scripts/test/installed-tools.sh /usr/bin/installed-tools.sh

RUN /usr/bin/installed-tools.sh
