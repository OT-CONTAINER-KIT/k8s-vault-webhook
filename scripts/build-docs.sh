#!/bin/bash

build_docs() {
    cd docs; yarn install
    cd docs; npm install -D vuepress
    cd docs; yarn build
}

build_docs
