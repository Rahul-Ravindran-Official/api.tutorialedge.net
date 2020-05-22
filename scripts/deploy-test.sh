#!/bin/bash
set -eoux pipefail

echo "Deploying Test"

function setup() {
    npm install -g serverless
    sls plugin install -n serverless-prune-plugin
    serverless version
}

function build() {
    pushd cmd
        for d in */; do
            echo $d

            pushd $d
                go build -o ../../bin/${d%/}
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

    echo "downloading go"
    mkdir -p resources
    mkdir -p go-bin
    mkdir -p go-code
    pushd resources
        curl https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz -o go1.14.2.linux-amd64.tar.gz
        tar -C ../go-bin -xzf go1.14.2.linux-amd64.tar.gz  

        curl https://dl.google.com/go/go1.14.3.src.tar.gz -o go.tar.gz
        # tar -C ../go-code -xzf go.tar.gz
    popd 

    pushd go-code
        mkdir -p go/pkg/tool/linux_amd64/

        cp -r ../go-bin/go/pkg go/pkg
        cp -r ../go-bin/go/src go/src

        tar -zcf ../code/go.tar.gz go
    popd
            
    
    cp go-bin/go/bin/go bin/go
    chmod +x bin/go
    echo "downloaded go"

    export AUTH0_SIGNING_KEY=$(curl https://tutorialedge.eu.auth0.com/pem)

    echo "Deploying Test API..."
    serverless deploy --stage=development --force
    echo "Successfully Deployed Test Stage"
}


main