#!/usr/bin/env bash

# Convert deployment paths into array
ENVIRONMENTS=($DEPLOYMENT_PATHS)

# Check if the DEPLOYMENT_PATH is not already set
if [ -z "${DEPLOYMENT_PATH}" ]; then
  # Print and ask for deployment environment (if more than one)
  if [ "${#ENVIRONMENTS[@]}" -gt 1 ]; then
    for i in "${!ENVIRONMENTS[@]}"; do
      echo "$i: ${ENVIRONMENTS[$i]}"
    done
    read -p "Deployment environment: " DEPLOYMENT_ENVIRONMENT
  fi
  if [ -z "${DEPLOYMENT_ENVIRONMENT}" ]; then
    DEPLOYMENT_ENVIRONMENT=0
  fi
  # Select correct path
  DEPLOYMENT_PATH="${ENVIRONMENTS[$DEPLOYMENT_ENVIRONMENT]}"
fi

# Check if the DEPLOYMENT_VERSION is not already set
if [ -z "${DEPLOYMENT_VERSION}" ]; then
  # Ask for deployment version
  read -p "Version [latest]: " DEPLOYMENT_VERSION
  if [ -z "${DEPLOYMENT_VERSION}" ]; then
    DEPLOYMENT_VERSION=latest
  fi
fi

echo "${DEPLOYMENT_PATH}"
echo "${DEPLOYMENT_VERSION}"

ssh $DEPLOYMENT_HOST \
  "cd ${DEPLOYMENT_PATH} && \
  git pull && \
  sed -i "s/VERSION=.*/VERSION=${DEPLOYMENT_VERSION}/" .env && \
  docker compose pull && \
  docker compose up -d"
