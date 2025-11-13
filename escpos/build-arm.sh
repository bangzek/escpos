#!/bin/sh
cd $(dirname $0)
env GOOS=linux GOARCH=arm GOARM=7 \
    go build -trimpath -ldflags="-s -w" && \
    ls -l $(basename $PWD) &&
    file $(basename $PWD) | cut -d, -f1-3
