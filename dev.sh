#!/bin/sh
# Script to run the Fiber application with Air for hot reloading

# Ensure the script is executable
# chmod +x dev.sh

# Add Go bin to PATH
export PATH="$PATH:$(go env GOPATH)/bin"

# Run Air with the configuration file
air
