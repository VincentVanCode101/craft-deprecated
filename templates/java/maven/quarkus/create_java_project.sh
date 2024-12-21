#!/bin/bash

set -e

if [ -z "$1" ]; then
  echo "Usage: $0 <ARTIFACT_ID>"
  exit 1
fi

ARTIFACT_ID=$1
U_ID=$(id -u)
GID=$(id -g)

DOCKERFILE="build.Dockerfile"
DOCKER_IMAGE_NAME="quarkus-project-generator"

docker build \
  -f $DOCKERFILE \
  --build-arg UID=$U_ID \
  --build-arg GID=$GID \
  --build-arg ARTIFACT_ID=$ARTIFACT_ID \
  -t $DOCKER_IMAGE_NAME .

docker run --rm \
  -v "$(pwd):/workspace" \
  $DOCKER_IMAGE_NAME \
  /bin/bash -c "cp -p -r /build-space/* /workspace"
