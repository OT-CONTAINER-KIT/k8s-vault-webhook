#!/bin/bash

install_spellcheck() {
    sudo apt-get update -y
    sudo apt-get install -y aspell
}

run_spellcheck() {
    aspell ../README.md
    aspell ../CHANGELOG.md
    aspell ../DEVELOPMENT.md
}

main() {
    install_spellcheck
    run_spellcheck
}

main
