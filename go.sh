#!/bin/sh

PORT="9999"
APP_NAME="playing-with-go"
GOPATH="$HOME/code/go"
BINARY="$GOPATH/bin/$APP_NAME"

set -e

echo "Building"
GOPATH=$GOPATH go get

echo "Running"
PORT=$PORT $BINARY

