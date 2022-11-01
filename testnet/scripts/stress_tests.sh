#!/bin/bash

# Run tests split across multiple validators

# todo:
# - collect list of test files
# - divide test file list by number of nodes running in testnet
# - have each node take its logical fraction of tests (e.g. node0 takes first 25%, node1 takes next 25% ...)
# - each test should mark in the chain that it has succeeded so that we know when it is finished.

bash /root/.duality/tests/dex-deposit.sh

# todo: add script to wait until all completion markers are in the chain:
# eg. wget -q -O - http://dualitynode0:1317/cosmos/tx/v1beta1/txs?events=tx.height%3E%3D0 | jq .txs[].body.memo
# read text out of tx memos
