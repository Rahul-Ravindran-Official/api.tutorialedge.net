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
    pushd resources
        # curl https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz -o go1.14.2.linux-amd64.tar.gz
        # cp go1.14.2.linux-amd64.tar.gz ../code/go.tar.gz
        # tar -C ../go-bin -xzf go1.14.2.linux-amd64.tar.gz
    popd 

    pushd go-bin/go
        rm AUTHORS CONTRIBUTORS PATENTS SECURITY.md robots.txt CONTRIBUTING.md LICENSE README.md VERSION favicon.ico
        rm -rf doc
        rm -rf test
        rm -rf api
        rm bin/gofmt
        rm -rf misc
    popd

    pushd go-bin
        curl https://images.tutorialedge.net/go.zip -o go.zip

        tar -zcf go.tar.gz go
        ls
    popd
    
    # cp go-bin/go/bin/go bin/go
    # chmod +x bin/go
    # echo "downloaded go"

    export AUTH0_SIGNING_KEY=$(curl https://tutorialedge.eu.auth0.com/pem)

    echo "Deploying Test API..."
    serverless deploy --stage=development --force
    echo "Successfully Deployed Test Stage"
}


main