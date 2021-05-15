#!/bin/bash

build_docs() {
    cd docs; yarn install
    cd docs; yarn build
}

build_docs
