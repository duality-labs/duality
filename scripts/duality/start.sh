#!/bin/sh

echo "START"

NAMESPACE_ID=$(echo $RANDOM | md5sum | head -c 16; echo;)
echo $NAMESPACE_ID
DA_BLOCK_HEIGHT=$(curl $NODE_ADDR/block | jq -r '.result.block.header.height')
echo $DA_BLOCK_HEIGHT

dualityd start \
    --rollmint.aggregator true \
    --rollmint.da_layer celestia \
    --rollmint.da_config='{"base_url":"http://celestia-lc:26659","timeout":60000000000,"fee":6000,"gas_limit":6000000}' \
    --rollmint.namespace_id $NAMESPACE_ID \
    --rollmint.da_start_height $DA_BLOCK_HEIGHT
