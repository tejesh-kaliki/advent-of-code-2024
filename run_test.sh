#!/bin/bash

# Continuously run tests

# Function to show an error and exit
err() {
    echo "Error: $1" >&2
    exit 1
}

# Check if fswatch is installed
command -v fswatch >/dev/null || err "You need 'fswatch' installed. Try 'brew install fswatch'."

# Watch for changes in Go files and run tests
fswatch -o ./*.go ./**/*.go | while read -r _; do
    clear
    go test -v "$@" | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
done