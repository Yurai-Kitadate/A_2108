#!/bin/bash

which golangci-lint
if [ $? != 0 ]; then
    echo Please install golangci-lint
    echo https://golangci-lint.run/usage/install/
    exit 1
fi

go test ./...

if [ $? != 0 ]; then
    echo failed to pass the go test
    exit 1
fi

golangci-lint run

if [ $? != 0 ]; then
    echo failed to pass the linter
    exit 1
fi

exit 0
