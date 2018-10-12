#!/bin/bash

VERSION="$(git describe --tags)"
LAST_COMMIT="$(git rev-parse HEAD)"

echo "commitr ${VERSION} (last commit ${LAST_COMMIT:0:7})"

GOOS=linux GOARCH=amd64 go build -ldflags "-X main.versionInfo=${VERSION} -X main.commit=${LAST_COMMIT:0:7}" -o commitr *.go
