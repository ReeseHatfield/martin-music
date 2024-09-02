#!/bin/bash

if [ $# -eq 0 ]
then
    echo "Missing album file argument"
    exit 1
fi


ALBUM_FILE_PATH=$(realpath $1)

CURRENT_DIR=$(pwd)


mkdir -p "$CURRENT_DIR/out"

mkdir -p "$CURRENT_DIR/bin"

# store files in temp directory
mkdir -p "$CURRENT_DIR/temp"

cd "$CURRENT_DIR/martin"

go build -o "$CURRENT_DIR/bin/martin.exe" main.go


"$CURRENT_DIR/bin/martin.exe" "$ALBUM_FILE_PATH" # execute

# cleanup temp files
rm -rf "$CURRENT_DIR/temp"

cd "$CURRENT_DIR"
