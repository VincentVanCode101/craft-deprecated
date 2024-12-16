#!/usr/bin/env bash
#
# Script to build and set up the craft Docker image and executable script.
# Adds the craft script to the system PATH and optionally to a custom directory.

set -e  # Exit on any error

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly IMAGE_NAME="craft-exec:latest"

source "$SCRIPT_DIR/utils.sh"

#######################################
# Main function to build Docker image and set up the craft script.
# Globals:
#   SCRIPT_DIR
#   IMAGE_NAME
# Arguments:
#   None
#######################################
main() {
  log "INFO" "Building the craft Docker image..."
  if ! docker build -t "$IMAGE_NAME" "$SCRIPT_DIR"; then
    err "Docker build failed"
    exit 1
  fi

  log "INFO" "Adding the craft script to your PATH..."
  chmod +x "$SCRIPT_DIR/craft"
  if ! sudo ln -sf "$SCRIPT_DIR/craft" /usr/local/bin/craft; then
    err "Failed to add craft script to /usr/local/bin"
    exit 1
  fi

  local custom_dir=~/dotfiles/bin/.local/scripts
  if [[ -d "$custom_dir" ]]; then
    log "INFO" "Directory $custom_dir exists. Adding the craft script..."
    if ! sudo ln -sf "$SCRIPT_DIR/craft" "$custom_dir/craft"; then
      err "Failed to add craft script to $custom_dir"
      exit 1
    fi
  else
    log "WARN" "Directory $custom_dir does not exist. Skipping."
  fi

  log "INFO" "Setup complete! You can now run 'craft'."
}

main "$@"
