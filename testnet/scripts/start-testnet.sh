#!/bin/bash
set -e

# wait a bit for the leader to start the chain
sleep 10;

# start chain and become a node
KEEP_RUNNING=false sh ./scripts/startup.sh

# wait a while for chain to be ready
sleep 10;

# run tests for this node
bash /root/.duality/scripts/run_tests_of_indexes.sh;

# keep container running
tail -f /dev/null;
