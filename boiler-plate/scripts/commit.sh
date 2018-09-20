#!/bin/sh

cd "$(dirname "$0")/.."

find internal/ -type f -name '*.go' -exec goimports -w -v {} \;
gofmt -r 'α[β:len(α)] -> α[β:]' -w internal;
