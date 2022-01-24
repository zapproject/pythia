# !/bin/sh
# LDFLAGS="-X main.GitHash=$(git rev-parse --short HEAD)"
LDFLAGS="-H windowsgui"


#linux
# go run ./pow/generate_opencl.go
# mv kernelSource.go pow/
# go build -ldflags "$LDFLAGS"

#windows
CC="x86_64-w64-mingw32-gcc" GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -ldflags "$LDFLAGS"



