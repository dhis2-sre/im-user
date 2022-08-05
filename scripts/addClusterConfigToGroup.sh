#!/usr/bin/env bash

set -euo pipefail

GROUP=$1
CONFIG_FILE=$2

$HTTP --ignore-stdin --form post "$INSTANCE_HOST/groups/$GROUP/cluster-configuration" "kubernetesConfiguration@$CONFIG_FILE" "Authorization: Bearer $ACCESS_TOKEN"
