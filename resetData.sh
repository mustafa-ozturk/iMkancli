#!/bin/bash

DEFAULT_FILE="data.default.json"
TARGET_FILE="data.json"

if [ ! -f "$DEFAULT_FILE" ]; then
    echo "Error: $DEFAULT_FILE not found!"
    exit 1
fi

cp "$DEFAULT_FILE" "$TARGET_FILE"
echo "Restored $TARGET_FILE from $DEFAULT_FILE"
