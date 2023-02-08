#!/usr/bin/env bash

set -euo pipefail

$HTTP get "$INSTANCE_HOST/me" "Authorization: Bearer $ACCESS_TOKEN"
