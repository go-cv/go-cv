#!/usr/bin/env bash

set -euo pipefail

if [ ! -f ".vars" ]; then
  cp vars.example .vars
fi

if [ ! -f "config.yaml" ]; then
  cp config.yaml.example config.yaml
fi

if [ ! -f "go.mod" ]; then
  module_name=$(git remote get-url origin | sed 's|^https://||')
  go mod init "$module_name"
fi
