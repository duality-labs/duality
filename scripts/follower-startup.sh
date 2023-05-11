#!/bin/sh

# check current working directory
if [[ ! -e scripts/startup.sh ]]; then
    echo >&2 "Please run this script from the base repo directory"
    exit 1
fi

export NETWORK=${NETWORK:-duality-devnet}
export CHAIN_ID="${CHAIN_ID:-$NETWORK}"
export NODE_MONIKER="${MONIKER:-devnet-follower-node}"

echo "NETWORK: $NETWORK /n CHAIN_ID: $CHAIN_ID"

dualityd init --chain-id $CHAIN_ID duality

./scripts/config_setup.sh

# Add consumer section to the ICS chain
dualityd add-consumer-section

# check we are on the correct network and can get information from the current network
node_status_json=$( wget --tries 30 --retry-connrefused -O - $RPC_ADDRESS/status )
found_network=$( echo $node_status_json | jq -r ".result.node_info.network" )
if [[ "$found_network" == "$CHAIN_ID" ]]
then
    echo "Found Duality chain: $found_network"
else
    echo "Could not establish connection to Duality chain, found network: ${found_network-none}"
    echo "Exiting..."
    exit 1
fi

# Get genesis file from an running node
if [ -z "$RPC_ADDRESS" ]; then
    echo "No RPC address provided"
    exit 1
else
    if $(wget -O - $RPC_ADDRESS/genesis | jq .result.genesis > $HOME/.duality/config/genesis.json); then
        echo "Loaded genesis.json"
    else
        echo "Cannot load genesis.json from original chain"
        exit 1
    fi
fi

# Add persistent peers
if [ -z "$PERSISTENT_PEERS" ]; then
    echo "manually setting persistent peers"
    peer_root_address=$(echo $RPC_ADDRESS |  cut -d ':' -f 1)
    echo "peer root: $peer_root_address"
    peer_id=$( wget -O - $RPC_ADDRESS/status | jq -r '.result.node_info.id')
    PERSISTENT_PEERS="$peer_id@$peer_root_address:26656" 
fi
echo "Adding persistent peers:$PERSISTENT_PEERS"
dasel put string -f ${HOME}/.duality/config/config.toml -s ".p2p.persistent_peers" $PERSISTENT_PEERS



# start as not a validator
echo "Starting fullnode..."
dualityd --log_level ${LOG_LEVEL:-info} start & :
echo "Started fullnode"

# keep container running
if [ "$KEEP_RUNNING" != "false" ]
then
    tail -f /dev/null;
fi
