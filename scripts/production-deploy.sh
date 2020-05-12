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

    echo "downloading go"
    mkdir -p resources
    mkdir -p go-bin
    pushd resources
        curl https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz -o go1.14.2.linux-amd64.tar.gz
        cp go1.14.2.linux-amd64.tar.gz ../code/go.tar.gz
        tar -C ../go-bin -xzf go1.14.2.linux-amd64.tar.gz
    popd 

    cp go-bin/go/bin/go bin/go
    chmod +x bin/go
    echo "downloaded go"

    echo "Deploying Production API..."
    serverless deploy --stage=production
    echo "Successfully Deployed Production Stage"
}

main