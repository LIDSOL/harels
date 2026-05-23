#!/bin/bash
# Build script for the Hare Language Server

# Create bin directory if it doesn't exist
mkdir -p bin

# Build the Go server
echo "Building Hare Language Server..."

if [ "$(basename "$PWD")" != "server" ]; then
    echo "Getting inside $(basename "$PWD")"
    cd server/
fi

go build -o bin/harels-server .

if [ $? -eq 0 ]; then
    echo "Build successful! Binary created at: bin/harels-server"
else
    echo "Build failed!"
    exit 1
fi
