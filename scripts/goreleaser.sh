#!/bin/bash

install_goreleaser() {
    curl -sfL https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | sh
}

release() {
    install_goreleaser
    ./bin/goreleaser release --rm-dist
}

compare_version() {
    version=$(cat VERSION)
    if ! git tag -l | grep "${version}"
    then
        git checkout master
        echo "git tag ${version}"
        git tag "${version}"
        release
    else
        git tag -l
        echo "Latest version is already updated"
    fi
}

compare_version
