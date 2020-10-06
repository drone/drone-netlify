#!/bin/sh

# force go modules
export GOPATH=""

# disable cgo
export CGO_ENABLED=0

# force linux amd64 platform
export GOOS=linux
export GOARCH=amd64 

set -e
set -x

# build the binary
go build -o release/linux/amd64/plugin

# build the docker image
docker build -f docker/Dockerfile -t plugins/netlify .
