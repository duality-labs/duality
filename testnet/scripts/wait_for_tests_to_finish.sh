#!/bin/bash
set -e

# abort if there are no tests to wait for
if [[ "$1" == "0" ]]
then
    exit 0
fi

# set total wait time in seconds
total_wait_time_all_tests=180

# check for desired state
total_test_files=$(find /root/.duality/tests | grep .*\.sh$ --count)
expected_tests="${1-$total_test_files}" timeout $total_wait_time_all_tests bash -s -c -- '

    echo "waiting for B: node$ID: $expected_tests"

    echo "expected_tests: $expected_tests"

    echo " --- Waiting for $expected_tests tests to have finished --- ";
    completed_tests=0
    while [ $completed_tests -lt $expected_tests ]
    do
        sleep 1;
        completed_tests=$( wget -q -O - http://dualitynode0:1317/cosmos/tx/v1beta1/txs?events=tx.height%3E%3D0 | jq .txs[].body.memo | grep --count completed-test- )
        echo "waited, has: $completed_tests of $expected_tests done"
    done
    echo " --- $expected_tests tests have finished --- ";
';
