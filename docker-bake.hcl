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

target "gotools-latest" {
    dockerfile = "Dockerfile"
    context= "."
    platforms = ["linux/amd64", "linux/arm64"]
    output = ["type=registry"]
}