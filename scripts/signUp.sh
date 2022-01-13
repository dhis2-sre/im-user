#!/usr/bin/env bash

set -euo pipefail

echo "{
  \"email\": \"$USER_EMAIL\",
  \"password\": \"$PASSWORD\"
}" | $HTTP post "$INSTANCE_HOST/users"
