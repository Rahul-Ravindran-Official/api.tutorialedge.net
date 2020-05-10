#!/bin/bash
set -eoux pipefail

echo "Deploying Production"

function setup() {
    npm install -g serverless
    serverless version
}


function build() {
    pushd cmd
        for d in */; do
            echo $d

            pushd $d
                GOOS=linux go build -o ../../bin/${d%/}
                chmod +x ../../bin/${d%/}
            popd
        done
    popd
}

function main() {
    
    go version

    echo "Downloading serverless CLI"
    setup
    echo "Successfull downloaded the serverless cli"

    echo "Building The Serverless Binaries..."
    build
    echo "Successfully build binaries..."

    export AUTH0_SIGNING_KEY=$(curl https://tutorialedge.eu.auth0.com/pem)

    echo "Deploying Production API..."
    serverless deploy --stage=production
    echo "Successfully Deployed Production Stage"
}

main