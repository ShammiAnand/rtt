#!/bin/sh
set -e

GITHUB_REPO="shammianand/rtt"
BINARY_NAME="rtt"

echo "Installing $BINARY_NAME..."

# Construct download URL
LATEST_RELEASE_URL="https://github.com/$GITHUB_REPO/releases/latest/download/$BINARY_NAME"

# Download and install
sudo curl -fL "$LATEST_RELEASE_URL" -o "/usr/local/bin/$BINARY_NAME"
sudo chmod +x "/usr/local/bin/$BINARY_NAME"

echo "$BINARY_NAME installed successfully"
