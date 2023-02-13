#!/bin/sh

VALIDATOR_NAME=validator1
CHAIN_ID=duality
KEY_NAME=jcp
CHAINFLAG="--chain-id ${CHAIN_ID}"
TOKEN_AMOUNT="10000000000000000000000000stake"
STAKING_AMOUNT="1000000000stake"
NODEIP="--node http://127.0.0.1:26657"

NAMESPACE_ID=$(echo $RANDOM | md5sum | head -c 16; echo;)
echo $NAMESPACE_ID
DA_BLOCK_HEIGHT=$(curl https://rpc-mocha.pops.one/block | jq -r '.result.block.header.height')
echo $DA_BLOCK_HEIGHT


ignite chain build
dualityd tendermint unsafe-reset-all
dualityd init $VALIDATOR_NAME --chain-id $CHAIN_ID

dualityd keys add $KEY_NAME --keyring-backend test
dualityd add-genesis-account $KEY_NAME $TOKEN_AMOUNT --keyring-backend test
dualityd gentx $KEY_NAME $STAKING_AMOUNT --chain-id $CHAIN_ID --keyring-backend test
dualityd collect-gentxs
dualityd start --rollkit.aggregator true --rollkit.da_layer celestia --rollkit.da_config='{"base_url":"http://localhost:26659","timeout":60000000000,"fee":6000,"gas_limit":6000000}' --rollkit.namespace_id $NAMESPACE_ID --rollkit.da_start_height $DA_BLOCK_HEIGHT
