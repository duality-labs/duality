#!/bin/sh

# reset existing chain
dualityd tendermint unsafe-reset-all
# init chain
dualityd init $MONIKER --chain-id $CHAIN_ID
