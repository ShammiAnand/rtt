#!/bin/sh
set -e

GITHUB_REPO="shammianand/rtt"
BINARY_NAME="rtt"

echo "Installing $BINARY_NAME..."

OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)
case $ARCH in
x86_64) ARCH="amd64" ;;
aarch64) ARCH="arm64" ;;
esac

# Construct download URL
LATEST_RELEASE_URL="https://github.com/$GITHUB_REPO/releases/latest/download/${BINARY_NAME}_${OS}_${ARCH}.tar.gz"

# Create temp directory
TMP_DIR=$(mktemp -d)
cd "$TMP_DIR"

# Download and extract
curl -sfLo release.tar.gz "$LATEST_RELEASE_URL"
tar xzf release.tar.gz

# Install binary
sudo mv "$BINARY_NAME" /usr/local/bin/

# Cleanup
cd -
rm -rf "$TMP_DIR"

echo "$BINARY_NAME installed successfully"
