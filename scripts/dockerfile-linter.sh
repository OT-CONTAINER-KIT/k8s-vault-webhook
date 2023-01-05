#!/bin/bash

download_hadolint() {
    wget https://github.com/hadolint/hadolint/releases/download/v2.12.0/hadolint-Linux-x86_64
    chmod +x hadolint-Linux-x86_64
}

execute_hadolint() {
    ./hadolint-Linux-x86_64 Dockerfile --ignore DL3007 --ignore DL3018
}

main() {
    download_hadolint
    execute_hadolint
}

main
