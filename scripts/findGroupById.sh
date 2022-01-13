#!/usr/bin/env bash

set -euo pipefail

GROUP_ID="$1"

$HTTP "$INSTANCE_HOST/groups/$GROUP_ID" "Authorization: Bearer $ACCESS_TOKEN"
