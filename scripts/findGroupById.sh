#!/usr/bin/env bash

set -euo pipefail

GROUP_NAME=$1

$HTTP "$INSTANCE_HOST/groups/$GROUP_NAME" "Authorization: Bearer $ACCESS_TOKEN"
