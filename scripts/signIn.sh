#!/usr/bin/env bash

set -euo pipefail

RESPONSE=$($HTTP --auth "$USER_EMAIL:$PASSWORD" post "$INSTANCE_HOST/tokens")

ACCESS_TOKEN=$(echo "$RESPONSE" | jq -r '.access_token')
REFRESH_TOKEN=$(echo "$RESPONSE" | jq -r '.refresh_token')

echo "export ACCESS_TOKEN=$ACCESS_TOKEN"
echo "export REFRESH_TOKEN=$REFRESH_TOKEN"
