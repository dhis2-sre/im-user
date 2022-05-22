#!/usr/bin/env bash

set -euo pipefail

GROUP_NAME=$1
GROUP_HOSTNAME=$2

echo "{
  \"name\": \"$GROUP_NAME\",
  \"hostname\": \"$GROUP_HOSTNAME\"
}" | $HTTP post "$INSTANCE_HOST/groups" "Authorization: Bearer $ACCESS_TOKEN"
