#!/bin/bash
rm -rf /tmp/nexepad-temp

nexepad --devnet --appdir=/tmp/nexepad-temp --profile=6061 &
nexePAD_PID=$!

sleep 1

infra-level-garbage --devnet -alocalhost:16611 -m messages.dat --profile=7000
TEST_EXIT_CODE=$?

kill $nexePAD_PID

wait $nexePAD_PID
nexePAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "nexepad exit code: $nexePAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $nexePAD_EXIT_CODE -eq 0 ]; then
  echo "infra-level-garbage test: PASSED"
  exit 0
fi
echo "infra-level-garbage test: FAILED"
exit 1
