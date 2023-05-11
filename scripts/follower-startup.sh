#!/bin/sh

if [ -z "$RPC_ADDRESS" ]; then
    echo "No RPC address provided"
    exit 1
fi
echo "using: $RPC_ADDRESS"
# check we are on the correct network and can get information from the current network
node_status_json=$( wget --tries 30 --retry-connrefused -O - $RPC_ADDRESS/status )
found_network=$( echo $node_status_json | jq -r ".result.node_info.network" )
if [[ "$found_network" == "$CHAIN_ID" ]]
then
    echo "Found Duality chain: $found_network"
else
    echo "Could not establish connection to Duality chain, found network: ${found_network:-none}"
    echo "Exiting..."
    exit 1
fi

# If $GENESIS_FILE_URL is not supplied and used in startup.sh 
# then get genesis file from the running node
if [ -z "$GENESIS_FILE_URL" ]; then
    if $(wget -O - $RPC_ADDRESS/genesis | jq .result.genesis > $HOME/.duality/config/genesis.json); then
        echo "Loaded genesis.json from original chain"
    else
        echo "Cannot load genesis.json from original chain"
        exit 1
    fi
fi

# Add persistent peers
if [ -z "$PERSISTENT_PEERS" ]; then
    echo "manually setting persistent peers"
    peer_root_address=$(echo $RPC_ADDRESS | sed -E 's/(https?:\/\/|:\/\/)?([^:/]+).*/\2/')
    peer_id=$( wget -O - $RPC_ADDRESS/status | jq -r '.result.node_info.id')
    PERSISTENT_PEERS="$peer_id@$peer_root_address:26656"
fi
echo "Adding persistent peers:$PERSISTENT_PEERS"
dasel put string -f ${HOME}/.duality/config/config.toml -s ".p2p.persistent_peers" $PERSISTENT_PEERS


echo "Starting fullnode..."
dualityd --log_level ${LOG_LEVEL:-info} start & :
echo "Started fullnode"
