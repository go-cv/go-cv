#!/usr/bin/env bash

if [ -z ${APP_VERSION} ]; then
  read -p "Version [latest]: " VERSIONINPUT
  if [ -z ${VERSIONINPUT} ]; then
    APP_VERSION="latest"
  else
    APP_VERSION=${VERSIONINPUT}
  fi
fi

# Get docker push option from user
if [ -z ${DOCKERPUSH} ]; then
  read -p "Docker push? [n]: " DOCKERPUSH
  if [ -z ${DOCKERPUSH} ]; then
    DOCKERPUSH=n
  fi
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
