#!/bin/bash

CURRENT_DIR=$(pwd)


mkdir -p "$CURRENT_DIR/out"

mkdir -p "$CURRENT_DIR/bin"

# store files in temp directory
mkdir -p "$CURRENT_DIR/temp"

cd "$CURRENT_DIR/martin"

go build -o "$CURRENT_DIR/bin/martin.exe" main.go

"$CURRENT_DIR/bin/martin.exe" # execute

# cleanup temp files
rm -rf "$CURRENT_DIR/temp"

cd "$CURRENT_DIR"
