#!/bin/bash
set -eoux pipefail

echo "Deploying Test"

function build() {
    GOOS=linux 
    pushd api
        for d in */; do
            echo $d

            pushd $d
                go build -o ../../bin/${d%/}
            popd
        done
    popd
}

function main() {
    
    go version

    echo "Building The Serverless Binaries..."
    build
    echo "Successfully build binaries..."

    export AUTH0_SIGNING_KEY=$(curl https://tutorialedge.eu.auth0.com/pem)

    echo "Deploying Test API..."
    serverless deploy --stage=production
    echo "Successfully Deployed Test Stage"
}

main