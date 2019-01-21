#!/bin/bash

pushd webpage
    npm install
    npm run build
popd

go run main.go
