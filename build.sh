#!/bin/bash
GOOS=js GOARCH=wasm go build -o wasm/xrpwallet.wasm ./
cp -r ./wasm ./demo