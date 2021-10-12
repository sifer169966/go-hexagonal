#!/bin/bash
set -e

PROJ="hex-gonal-template"
ORG_PATH="github.com/backend"
REPO_PATH="${ORG_PATH}/${PROJ}"
if ! [ -x "$(command -v go)" ]; then
    echo "go is not installed"
    exit 1
fi

if [ -z "${GOPATH}" ]; then
    echo "set GOPATH"
    exit 1
fi

PATH="${PATH}:${GOPATH}/bin"


export GO111MODULE=on
