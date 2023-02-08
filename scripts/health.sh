#!/usr/bin/env bash

set -euo pipefail

$HTTP get "$INSTANCE_HOST/users/health"
