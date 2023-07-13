#!/bin/bash -e

# Ensure that dualityd is built and on your path before running.
CHAIN_ID=duality-1
GENESIS_TIME="2023-07-12T15:00:00Z"
MAX_GAS="35000000"

# Get default genesis
nodehome=$(mktemp -d)
dualityd init local-node --home $nodehome >/dev/null 2>&1
cp -f $nodehome/config/genesis.json .
rm -rf $nodehome

# Set individual parameters
dasel put -f ./genesis.json -t string -v "$CHAIN_ID" '.chain_id'
dasel put -f ./genesis.json -t string -v "$GENESIS_TIME" '.genesis_time'
dasel put -f ./genesis.json -t string -v "$MAX_GAS" '.consensus_params.block.max_gas'
dasel put -f ./genesis.json -t string -v "uatom" '.app_state.ccvconsumer.params.provider_reward_denoms.append()'
dasel put -f ./genesis.json -t string -v "864000" '.app_state.slashing.params.signed_blocks_window'

# Configure group genesis
group_genesis=$(./build-group-genesis.sh)
dasel put -f ./genesis.json -t json -v "$group_genesis" '.app_state.group'

# Configure wasm genesis
wasm_genesis=$(dasel -f "./src/wasm_genesis.json" '.')
dasel put -f ./genesis.json -t json -v "$wasm_genesis" '.app_state.wasm'

# Configure auth genesis
accounts=$(dasel -f "./src/accounts.json" '.')
dasel put -f ./genesis.json -t json -v "$accounts" '.app_state.auth.accounts'