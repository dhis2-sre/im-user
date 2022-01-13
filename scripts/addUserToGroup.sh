#!/usr/bin/env bash

set -euo pipefail

USER_ID="$1"
GROUP_NAME="$2"

GROUP_ID=$($HTTP "$INSTANCE_HOST/groups-name-to-id/$GROUP_NAME" "Authorization: Bearer $ACCESS_TOKEN")

$HTTP post "$INSTANCE_HOST/groups/$GROUP_ID/users/$USER_ID" "Authorization: Bearer $ACCESS_TOKEN"
