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
    popd 

    pushd go-bin/go
        rm -rf bin/gofmt
        rm -rf lib/
        rm -rf test/
        rm -rf api/
        rm -rf doc/
        rm -rf pkg/linux_amd64_race/
        rm -rf src/database/
        rm -rf src/net/
        rm -rf src/archive/
        rm -rf src/syscall/
        rm -rf src/cmd/
        rm -rf src/image/
        rm -rf src/debug/
        rm -rf src/compress/
        rm -rf src/cmd/vendor/
        rm -rf misc/

        rm -rf pkg/tool/linux_amd64/doc
        rm -rf pkg/tool/linux_amd64/asm
        rm -rf pkg/tool/linux_amd64/addr2line
        rm -rf pkg/tool/linux_amd64/test2json
        rm -rf pkg/tool/linux_amd64/buildid
        rm -rf pkg/tool/linux_amd64/pprof
        rm -rf pkg/tool/linux_amd64/cgo
        rm -rf pkg/tool/linux_amd64/trace
        rm -rf pkg/tool/linux_amd64/cover
        rm -rf pkg/tool/linux_amd64/pack
        rm -rf pkg/tool/linux_amd64/vet
        rm -rf pkg/tool/linux_amd64/dist
        rm -rf pkg/tool/linux_amd64/link
        rm -rf pkg/tool/linux_amd64/objdump
        rm -rf pkg/tool/linux_amd64/nm
        rm -rf pkg/tool/linux_amd64/fix
    popd   
    
    pushd go-bin
        zip -r ../code/go.zip go
    popd

    # final artefact = ./code/go.tar.gz

    cp go-bin/go/bin/go bin/go
    chmod +x bin/go
    echo "downloaded go"

    export AUTH0_SIGNING_KEY=$(curl https://tutorialedge.eu.auth0.com/pem)

    echo "Deploying Test API..."
    serverless deploy --stage=development --force
    echo "Successfully Deployed Test Stage"
}


main