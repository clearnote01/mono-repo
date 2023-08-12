#!/bin/bash

# Required parameters:
# @raycast.schemaVersion 1
# @raycast.title gc-cli
# @raycast.mode fullOutput

# Optional parameters:
# @raycast.icon 🤖

# Documentation:
# @raycast.author utkarsh_raj
# @raycast.authorURL https://raycast.com/utkarsh_raj

export GCLOUD_ACCESS_TOKEN=$(gcloud auth print-access-token)
./gc-cli view dashboard
