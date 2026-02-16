#!/usr/bin/env bash

set -euo pipefail

if [ ! -f ".vars" ]; then
  cp vars.example .vars
fi

if [ ! -f "config.yaml" ]; then
  cp config.yaml.example config.yaml
fi
