#!/usr/bin/env bash

set -euo pipefail

RESPONSE=$(echo "{\"refreshToken\": \"$REFRESH_TOKEN\"}" | $HTTP post "$INSTANCE_HOST/refresh")

ACCESS_TOKEN=$(echo "$RESPONSE" | jq -r '.access_token')
REFRESH_TOKEN=$(echo "$RESPONSE" | jq -r '.refresh_token')

echo "export ACCESS_TOKEN=$ACCESS_TOKEN"
echo "export REFRESH_TOKEN=$REFRESH_TOKEN"
