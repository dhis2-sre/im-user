#!/usr/bin/env bash

set -euo pipefail

CONFIG_FILE=$1
GROUP_ID=$2

$HTTP --ignore-stdin --form post "$INSTANCE_HOST/groups/$GROUP_ID/cluster-configuration" "kubernetesConfiguration@$CONFIG_FILE" "Authorization: Bearer $ACCESS_TOKEN"
