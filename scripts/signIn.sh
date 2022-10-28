#!/usr/bin/env bash

set -euo pipefail

$HTTP --auth "$USER_EMAIL:$PASSWORD" post "$INSTANCE_HOST/tokens"
