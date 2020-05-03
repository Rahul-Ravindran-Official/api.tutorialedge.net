#!/bin/bash
set -eoux pipefail

function create_pull_request() {
    gh pr create
}

function main() {
    echo "Creating Pull Request into Master"
    create_pull_request
}

main
