#!/bin/bash
set -e

function format() {
  if [ -n "$(gofmt -l .)" ]; then
    echo "Go code is not formatted:"
    gofmt -d ../.
    exit 1
  fi
}

function lint() {
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.26.0

  golangci-lint run
}

function main() {
  format
  lint
}

main