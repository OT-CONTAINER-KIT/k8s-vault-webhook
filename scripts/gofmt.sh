#!/bin/bash

gofmt_files=$(go fmt ./... | wc -l)

if [[ ${gofmt_files} > 0 ]]
then
    echo "Please format golang files using:- go fmt ./..."
    exit 1
else
    echo "All files are formated using gofmt"
fi
