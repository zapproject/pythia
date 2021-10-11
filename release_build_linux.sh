#!/bin/sh
LDFLAGS="-X main.GitHash=$(git rev-parse --short HEAD)"
# LDFLAGS="-X main.GitHash=$(git rev-parse --short HEAD) -s -w"


#linux
go run ./pow/generate_opencl.go
mv kernelSource.go pow/
go build -ldflags "$LDFLAGS"