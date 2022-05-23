#!/usr/bin/env bash

set -euo pipefail

$HTTP "$INSTANCE_HOST/users/health"
