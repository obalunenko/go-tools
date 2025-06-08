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

target "docker-metadata-action" {}

target "gotools-latest" {
    inherits = ["docker-metadata-action"]
    dockerfile = "Dockerfile"
    context    = "."
    platforms = [
        "linux/amd64",
        "linux/arm64"
    ]
    args = {
        GO_VERSION = "1.24.4"
        ALPINE_VERSION = "3.22"
    }
    labels = {
        "org.opencontainers.image.title"       = "${IMAGE_TITLE}"
        "org.opencontainers.image.description" = "${IMAGE_DESCRIPTION}"
    }
}