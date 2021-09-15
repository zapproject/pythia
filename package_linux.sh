#!/bin/sh
LDFLAGS="-X main.GitHash=$(git rev-parse --short HEAD)"
# LDFLAGS="-X main.GitHash=$(git rev-parse --short HEAD) -s -w"


#linux

go build -ldflags "$LDFLAGS"

zip -r pythia_linux.zip pythia indexes.json loggingConfig.json