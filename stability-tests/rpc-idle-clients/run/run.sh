#!/bin/bash
rm -rf /tmp/nexepad-temp

NUM_CLIENTS=128
nexepad --devnet --appdir=/tmp/nexepad-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
nexePAD_PID=$!
nexePAD_KILLED=0
function killnexepadIfNotKilled() {
  if [ $nexePAD_KILLED -eq 0 ]; then
    kill $nexePAD_PID
  fi
}
trap "killnexepadIfNotKilled" EXIT

sleep 1

rpc-idle-clients --devnet --profile=7000 -n=$NUM_CLIENTS
TEST_EXIT_CODE=$?

kill $nexePAD_PID

wait $nexePAD_PID
nexePAD_EXIT_CODE=$?
nexePAD_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "nexepad exit code: $nexePAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $nexePAD_EXIT_CODE -eq 0 ]; then
  echo "rpc-idle-clients test: PASSED"
  exit 0
fi
echo "rpc-idle-clients test: FAILED"
exit 1
