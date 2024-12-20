#!/bin/bash

set -e

SCRIPT="./getBinary.sh"

WATCH_DIRS=("cmd" "internal" "registry" "templates" "main.go")

if ! command -v inotifywait &>/dev/null; then
  echo "Error: inotifywait is not installed. Install it with 'sudo apt install inotify-tools'."
  exit 1
fi

echo "Watching for changes in: ${WATCH_DIRS[*]}"
inotifywait -m -r -e modify --format "%w%f" "${WATCH_DIRS[@]}" | while read -r changed_file; do
  echo "Change detected in $changed_file. Executing script..."
  if ! bash "$SCRIPT"; then
    echo "Error: $SCRIPT encountered an issue but continuing to watch for changes."
  fi
done