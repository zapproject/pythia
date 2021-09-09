#!/bin/sh
LDFLAGS="-X main.GitHash=$(git rev-parse --short HEAD)"
# LDFLAGS="-X main.GitHash=$(git rev-parse --short HEAD) -s -w"


go run ./pow/generate_opencl.go
mv kernelSource.go pow/

#windows
CC="x86_64-w64-mingw32-gcc-win32" GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -ldflags "$LDFLAGS"


zip -r pythia_win.zip pythia.exe indexes.json loggingConfig.json