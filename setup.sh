#!/bin/bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

echo "Building the craft Docker image..."
docker build -t craft-image "$SCRIPT_DIR"

echo "Adding the craft script to your PATH..."
chmod +x "$SCRIPT_DIR/craft"
sudo ln -sf "$SCRIPT_DIR/craft" /usr/local/bin/craft
sudo ln -sf "$SCRIPT_DIR/craft" ~/dotfiles/bin/.local/scripts

echo "Setup complete! You can now run 'craft --operation=new --language=rust'."
