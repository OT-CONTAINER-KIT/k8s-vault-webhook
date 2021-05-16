#!/bin/bash

install_gosec() {
    curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s latest
}

execute_gosec() {
    ./bin/gosec -fmt=junit-xml -out=./bin/results.xml ./... || true
}

main() {
    install_gosec
    execute_gosec
}

main
