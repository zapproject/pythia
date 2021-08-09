#!/bin/sh
LDFLAGS="-X main.GitHash=$(git rev-parse --short HEAD)"
# LDFLAGS="-X main.GitHash=$(git rev-parse --short HEAD) -s -w"


# #linux
# go run ./pow/generate_opencl.go
# mv kernelSource.go pow/
# go build -ldflags "$LDFLAGS"

#windows
CC="x86_64-8.1.0-win32-seh-rt_v6-rev0" GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "$LDFLAGS"




