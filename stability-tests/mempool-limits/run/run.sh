#!/bin/bash

APPDIR=/tmp/nexepad-temp
nexePAD_RPC_PORT=29587

rm -rf "${APPDIR}"

nexepad --simnet --appdir="${APPDIR}" --rpclisten=0.0.0.0:"${nexePAD_RPC_PORT}" --profile=6061 &
nexePAD_PID=$!

sleep 1

RUN_STABILITY_TESTS=true go test ../ -v -timeout 86400s -- --rpc-address=127.0.0.1:"${nexePAD_RPC_PORT}" --profile=7000
TEST_EXIT_CODE=$?

kill $nexePAD_PID

wait $nexePAD_PID
nexePAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "nexepad exit code: $nexePAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $nexePAD_EXIT_CODE -eq 0 ]; then
  echo "mempool-limits test: PASSED"
  exit 0
fi
echo "mempool-limits test: FAILED"
exit 1
