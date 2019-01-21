#!/bin/bash
CGO_ENABLED=1 GOOS=linux GOARCH=arm CC=arm-linux-gnueabi-gcc go build -o main.arm -v  -ldflags "-w -s -linkmode external -extldflags -static" main.go
CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC=aarch64-linux-gnu-gcc go build -ldflags "-w" -o main.arm64 main.go
CGO_ENABLED=1 go build -ldflags "-w" -o main main.go
pushd webpage
    npm run build
popd

rm -rf release.arm64.tgz
tar -czvmf release.arm64.tgz main.arm64 main.arm  main webpage/dist/
