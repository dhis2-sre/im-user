#!/usr/bin/env bash

set -euo pipefail

USER_ID=$1

$HTTP get "$INSTANCE_HOST/users/$USER_ID" "Authorization: Bearer $ACCESS_TOKEN"
