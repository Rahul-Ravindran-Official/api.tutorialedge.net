#!/bin/bash
set -eoux pipefail

function run_unit_tests() {
    echo "Running Unit Tests"
    go test ./...
}

function main() {
    run_unit_tests
}

main