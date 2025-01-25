#!/bin/sh
set -e

GITHUB_REPO="shammianand/rtt"
BINARY_NAME="rtt"

echo "Installing $BINARY_NAME..."

LATEST_TAG=$(curl -s https://api.github.com/repos/$GITHUB_REPO/releases/latest | grep -Po '"tag_name": "\K.*?(?=")')
LATEST_RELEASE_URL="https://github.com/$GITHUB_REPO/releases/download/$LATEST_TAG/$BINARY_NAME"

sudo curl -fL "$LATEST_RELEASE_URL" -o "/usr/local/bin/$BINARY_NAME"
sudo chmod +x "/usr/local/bin/$BINARY_NAME"

echo "$BINARY_NAME installed successfully"
