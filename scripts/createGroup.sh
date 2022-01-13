#!/usr/bin/env bash

set -euo pipefail

GROUP_NAME="$1"
GROUP_HOSTNAME="$2"

ADMIN_ACCESS_TOKEN=$($HTTP --auth "$ADMIN_USER_EMAIL:$ADMIN_USER_PASSWORD" post "$INSTANCE_HOST/tokens" | jq -r '.access_token')

echo "{
  \"name\": \"$GROUP_NAME\",
  \"hostname\": \"$GROUP_HOSTNAME\"
}" | $HTTP post "$INSTANCE_HOST/groups" "Authorization: Bearer $ADMIN_ACCESS_TOKEN"
