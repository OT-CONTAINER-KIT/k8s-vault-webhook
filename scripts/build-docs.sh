#!/bin/bash

build_docs() {
    cd docs; yarn install && \
        yarn add -D vuepress && \
        yarn build
}

build_docs
