#!/usr/bin/env bash

set -euo pipefail

$HTTP "$INSTANCE_HOST/me" "Authorization: Bearer $ACCESS_TOKEN"
