#!/bin/bash

CURRENT_DIR=$(pwd)

mkdir -p "$CURRENT_DIR/martin/bin"

cd "$CURRENT_DIR/martin"

go build -o "$CURRENT_DIR/martin/bin/martin.exe" main.go

"$CURRENT_DIR/martin/bin/martin.exe" # execute

cd "$CURRENT_DIR"
