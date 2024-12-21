#!/bin/bash

set -e

if [ -z "$1" ]; then
  echo "Usage: $0 <PROJECT_NAME>"
  exit 1
fi

PROJECT_NAME=$1
U_ID=$(id -u)
GID=$(id -g)

DOCKERFILE="build.Dockerfile"
DOCKER_IMAGE_NAME="rust-project-generator"

docker build \
  -f $DOCKERFILE \
  --build-arg UID=$U_ID \
  --build-arg GID=$GID \
  --build-arg PROJECT_NAME=$PROJECT_NAME \
  -t $DOCKER_IMAGE_NAME .

docker run --rm \
  -v "$(pwd):/workspace" \
  $DOCKER_IMAGE_NAME \
  /bin/sh -c "cp -p -r /build-space/* /workspace"
