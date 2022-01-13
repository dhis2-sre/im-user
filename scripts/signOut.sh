#!/usr/bin/env bash

set -euo pipefail

$HTTP delete "$INSTANCE_HOST/users" "Authorization: Bearer $ACCESS_TOKEN"
