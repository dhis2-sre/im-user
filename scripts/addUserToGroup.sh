#!/usr/bin/env bash

set -euo pipefail

USER_ID="$1"
GROUP_NAME="$2"

ADMIN_ACCESS_TOKEN=$($HTTP --auth "$ADMIN_USER_EMAIL:$ADMIN_USER_PASSWORD" post "$INSTANCE_HOST/tokens" | jq -r '.access_token')

GROUP_ID=$($HTTP "$INSTANCE_HOST/groups-name-to-id/$GROUP_NAME" "Authorization: Bearer $ADMIN_ACCESS_TOKEN")

$HTTP post "$INSTANCE_HOST/groups/$GROUP_ID/users/$USER_ID" "Authorization: Bearer $ADMIN_ACCESS_TOKEN"
