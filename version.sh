#!/usr/bin/env bash

# Get version from user
read -p "Version [latest]: " VERSIONINPUT
# If version was not provided, use the latest commit short hash as version
if [ -z ${VERSIONINPUT} ]; then
  APP_VERSION="latest"
else
  APP_VERSION=${VERSIONINPUT}
fi

# Get docker push option from user
read -p "Docker push? [n]: " DOCKERPUSH
if [ -z ${DOCKERPUSH} ]; then
  DOCKERPUSH=n
fi

# Create version tag (if provided)
if [ ! -z ${VERSIONINPUT} ]; then
  git tag ${APP_VERSION}
fi

# Build the app
export APP_VERSION
make docker
# If wanted, push the docker image
if [ ${DOCKERPUSH} = "y" ]; then
  make dockerpush
fi
