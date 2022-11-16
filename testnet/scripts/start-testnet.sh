#!/bin/bash
set -e

dualityd --home /testnet/dualityd start --moniker "node$ID" --log_level ${LOG_LEVEL:-warn} & \
# wait a while for chain to be ready
sleep 10;

# run tests for this node
bash /root/.duality/scripts/run_tests_of_indexes.sh;

# keep container running
tail -f /dev/null;
