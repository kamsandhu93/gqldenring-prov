#!/usr/bin/env bash

set -euo pipefail

if [ ! -x gqldenring/gqldenring ] ; then
  echo "[ERROR] No binary found at gqldenring/gqldenring, try run ./scripts/dl_gqldenring.sh and then retry"
  exit 1
fi

echo "[INFO] Starting GQLdenring server"
./gqldenring/gqldenring &
GQLDENRING_PID=$!
echo "[INFO] GQLdenring server started with PID $GQLDENRING_PID"


function cleanup {
    echo "[INFO] Killing GQLdenring server with PID $GQLDENRING_PID"
    kill "$GQLDENRING_PID"
}
trap cleanup EXIT

echo "[INFO] Running acceptance tests"
TF_ACC=1 go test ./... -v -timeout 120m
echo "[INFO] Acceptance tests completed, status=$?"
