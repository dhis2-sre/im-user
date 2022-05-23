#!/usr/bin/env bash

set -euo pipefail

USER_ID=$1
GROUP_NAME=$2

$HTTP post "$INSTANCE_HOST/groups/$GROUP_NAME/users/$USER_ID" "Authorization: Bearer $ACCESS_TOKEN"
