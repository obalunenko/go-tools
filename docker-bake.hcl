group "default" {
    targets = [
        "gotools-latest",
        ]
}

variable "REGISTRY" {
    default = "ghcr.io"
}

variable "IMAGE_NAME" {
    default = "obalunenko/go-tools"
}

variable "IMAGE_WITH_REGISTRY" {
    default = notequal("",REGISTRY) ? "${REGISTRY}/${IMAGE_NAME}": "${IMAGE_NAME}"
}

variable "IMAGE_TITLE" {
    default = "Go Tools"
}

variable "IMAGE_DESCRIPTION" {
    default = ""
}

# Layer cache, one tag per platform: buildcache-linux-amd64, buildcache-linux-arm64.
# CI pushes it from master and release builds; every build reads it.
variable "CACHE_IMAGE" {
    default = "${IMAGE_WITH_REGISTRY}:buildcache"
}

target "docker-metadata-action" {}

target "_common" {
    inherits = ["docker-metadata-action"]
    dockerfile = "Dockerfile"
    context    = "."
    args = {
        GO_VERSION = "1.27.4"
        ALPINE_VERSION = "3.24"
    }
    labels = {
        "org.opencontainers.image.title"       = "${IMAGE_TITLE}"
        "org.opencontainers.image.description" = "${IMAGE_DESCRIPTION}"
    }
    cache-from = [
        "type=registry,ref=${CACHE_IMAGE}-linux-amd64",
        "type=registry,ref=${CACHE_IMAGE}-linux-arm64",
    ]
}

target "gotools-latest" {
    inherits = ["_common"]
    platforms = [
        "linux/amd64",
        "linux/arm64"
    ]
}

# Native platform only, loaded into the local image store.
target "gotools-local" {
    inherits = ["_common"]
    tags     = ["${IMAGE_WITH_REGISTRY}:local"]
    output   = ["type=docker"]
}
