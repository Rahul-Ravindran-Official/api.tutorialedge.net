#!/bin/bash
set -eoux pipefail

function run_integration_tests() {
    go test -tags=integration ./...
}

function main() {
    echo "Running Integration Test Suite..."
}

main
