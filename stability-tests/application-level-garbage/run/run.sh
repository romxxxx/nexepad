#!/bin/bash
rm -rf /tmp/nexepad-temp

nexepad --devnet --appdir=/tmp/nexepad-temp --profile=6061 --loglevel=debug &
nexePAD_PID=$!
nexePAD_KILLED=0
function killnexepadIfNotKilled() {
    if [ $nexePAD_KILLED -eq 0 ]; then
      kill $nexePAD_PID
    fi
}
trap "killnexepadIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:16611 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $nexePAD_PID

wait $nexePAD_PID
nexePAD_KILLED=1
nexePAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "nexepad exit code: $nexePAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $nexePAD_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1
