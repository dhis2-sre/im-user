#!/usr/bin/env bash

set -euo pipefail

CONFIG_FILE=$1
GROUP_ID=$2

ADMIN_ACCESS_TOKEN=$($HTTP --auth "$ADMIN_USER_EMAIL:$ADMIN_USER_PASSWORD" post "$INSTANCE_HOST/tokens" | jq -r '.access_token')

$HTTP --ignore-stdin --form post "$INSTANCE_HOST/groups/$GROUP_ID/cluster-configuration" "kubernetesConfiguration@$CONFIG_FILE" "Authorization: Bearer $ADMIN_ACCESS_TOKEN"
