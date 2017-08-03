#!/usr/bin/env bash
set -eux

export PROJECT_ROOT="$(dirname "$0")"
export MAIN_PATH="$PROJECT_ROOT/src/fivestars.com/docker-extended/cmd/main.go"

go get github.com/stretchr/testify/require
go get github.com/Masterminds/sprig

mkdir -p "$PROJECT_ROOT/output"

## What you want to build : https://golang.org/doc/install/source#environment
build(){
    typeset GOOS="$1" GOARCH="$2"
    env GOOS="$1" GOARCH="$2" go build -o "$PROJECT_ROOT/output/main_${GOOS}" "$MAIN_PATH"
}


build windows amd64
build linux amd64
build darwin amd64

