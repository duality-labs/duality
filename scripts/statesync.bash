#!/bin/bash
# microtick and bitcanna contributed significantly here.
# Pebbledb state sync script.
# invoke like: bash scripts/statesync.bash

## USAGE RUNDOWN
# Not for use on live nodes
# For use when testing.
# Assumes that ~/.dualityd doesn't exist
# can be modified to suit your purposes if ~/.dualityd does already exist

set -uxe

# Set Golang environment variables.
export GOPATH=~/go
export PATH=$PATH:~/go/bin

# Install
go install ./...

# Initialize chain.
dualityd init test

# Get Genesis
wget https://raw.githubusercontent.com/cosmos/testnets/master/replicated-security/duality-testnet-1/duality-testnet-1-genesis.json
mv duality-testnet-1-genesis.json ~/.duality/config/genesis.json

# Get "trust_hash" and "trust_height".
INTERVAL=1000
LATEST_HEIGHT=$(curl -s https://rpc.testnet-1.duality.xyz/block | jq -r .result.block.header.height)
BLOCK_HEIGHT=$(($LATEST_HEIGHT - $INTERVAL))
TRUST_HASH=$(curl -s "https://rpc.testnet-1.duality.xyz/block?height=$BLOCK_HEIGHT" | jq -r .result.block_id.hash)

# Print out block and transaction hash from which to sync state.
echo "trust_height: $BLOCK_HEIGHT"
echo "trust_hash: $TRUST_HASH"

# Export state sync variables.
export DUALITYD_STATESYNC_ENABLE=true
export DUALITYD_P2P_MAX_NUM_OUTBOUND_PEERS=200
export DUALITYD_STATESYNC_RPC_SERVERS="https://rpc.testnet-1.duality.xyz:443,https://rpc.testnet-1.duality.xyz:443"
export DUALITYD_STATESYNC_TRUST_HEIGHT=$BLOCK_HEIGHT
export DUALITYD_STATESYNC_TRUST_HASH=$TRUST_HASH

# Set seed node
export DUALITYD_P2P_PERSISTENT_PEERS="25da44bb094907c4a476363a5b678a090a4c0140@5.161.179.189:26656,b0e1a54e0be7ff8af3caf457e29d217ca1184129@46.101.195.113:46656"
export DUALITYD_P2P_SEEDS="df5b21498dd5594a609e2e2af41434bbd9297ffd@p2p.testnet-1.duality.xyz:26656"

# Start chain.
dualityd start --x-crisis-skip-assert-invariants