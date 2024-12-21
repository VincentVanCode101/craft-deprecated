#!/bin/bash

set -e

if [ -z "$1" ]; then
  echo "Usage: $0 <ARTIFACT_ID>"
  exit 1
fi

ARTIFACT_ID=$1
GROUP_ID="com.main"
OUTPUT_DIR="/build-space/$ARTIFACT_ID"
U_ID=$(id -u)
G_ID=$(id -g)

DOCKERFILE="build.Dockerfile"
DOCKER_IMAGE_NAME="maven-project-generator"

docker build \
  -f $DOCKERFILE \
  --build-arg UID=$U_ID \
  --build-arg G_ID=$G_ID \
  --build-arg GROUP_ID=$GROUP_ID \
  --build-arg ARTIFACT_ID=$ARTIFACT_ID \
  -t $DOCKER_IMAGE_NAME .

docker run --rm \
  -v "$(pwd):/workspace" \
  $DOCKER_IMAGE_NAME \
  /bin/bash -c "cp -p -r /build-space/* /workspace"
