#!/usr/bin/env bash

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
IMAGE_NAME="craft-exec:latest"

echo "Building the craft Docker image..."
docker build -t "$IMAGE_NAME" "$SCRIPT_DIR" || { echo "Docker build failed"; exit 1; }

echo "Adding the craft script to your PATH..."
chmod +x "$SCRIPT_DIR/craft"
sudo ln -sf "$SCRIPT_DIR/craft" /usr/local/bin/craft

if [ -d ~/dotfiles/bin/.local/scripts ]; then
    echo "Directory ~/dotfiles/bin/.local/scripts exists. Adding the craft script..."
    sudo ln -sf "$SCRIPT_DIR/craft" ~/dotfiles/bin/.local/scripts/craft
else
    echo "Directory ~/dotfiles/bin/.local/scripts does not exist. Skipping."
fi

echo "Setup complete! You can now run 'craft'."
