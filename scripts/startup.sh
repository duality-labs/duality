#!/bin/sh

# set variable defaults
STARTUP_MODE="${MODE:-fullnode}"
NODE_MONIKER="${MONIKER:-$( head /dev/urandom | tr -dc 0-9a-f | head -c12 )}"

# check current working directorys
if [[ ! -e scripts/startup.sh ]]; then
    echo >&2 "Please run this script from the base repo directory"
    exit 1
fi

echo "Startup mode: $STARTUP_MODE"

echo "Initializing chain..."
dualityd init duality

# replace moniker in the config
sed -i 's#moniker = ""#moniker = "'"$MONIKER"'"#' /root/.duality/config/config.toml

# start or join a chain
if [ $STARTUP_MODE == "new" ]
then

    # add new test accounts
    echo "Creating test accounts..."
    mkdir /root/.duality/testkeys

    # alice
    dualityd keys add alice --keyring-backend test
    dualityd add-genesis-account $(dualityd keys show alice -a --keyring-backend test) 1000000000token,1000000000stake --keyring-backend test
    # bob
    dualityd keys add bob --keyring-backend test
    dualityd add-genesis-account $(dualityd keys show bob -a --keyring-backend test) 1000000000token,1000000000stake --keyring-backend test

    # Add gentxs to the genesis file
    dualityd gentx alice 1000000stake --chain-id duality --keyring-backend test
    dualityd collect-gentxs

    echo "Starting new chain..."
    dualityd --log_level ${LOG_LEVEL:-info} start
    exit

else

    echo "Will attempt to join the network using chain.json information"
    echo "Contacting network..."

    # find an RPC address to check the live chain with
    rpc_address="$( jq -r .apis.rpc[0].address networks/duality-1/chain.json )"

    # check if we can get information from the current network
    abci_info_json=$( wget --tries 30 -q -O - $rpc_address/abci_info )
    if [[ "$( echo $abci_info_json | jq -r ".result.response.data" )" == "duality" ]]
    then
        echo "Found Duality chain!"
    else
        echo "Could not establish connection to Duality chain"
        echo "Exiting..."
        exit 1
    fi

    # read out peers from chain.json
    persistent_peers_array="$( jq .peers.persistent_peers networks/duality-1/chain.json )"
    persistent_peers="$( echo $persistent_peers_array | jq -r 'map(.id + "@" + .address) | join(",")' )"

    # set chain settings
    sed -i 's#persistent_peers = ""#persistent_peers = "'"$persistent_peers"'"#' /root/.duality/config/config.toml
    mv networks/duality-1/genesis.json /root/.duality/config/genesis.json

    # check if this node intends to become a validator
    if [[ "$STARTUP_MODE" == "validator" && ! -z "$MNEMONIC" ]]
    then

        echo "Starting future validator fullnode..."
        dualityd --log_level ${LOG_LEVEL:-info} start & :

        # wait for node to finish catching up to the chain's current height
        sleep 5
        chain_block_height=$(echo $abci_info_json | jq -r ".result.response.last_block_height")
        node_status_json=$( dualityd status )
        while [[ $( echo $node_status_json | jq .SyncInfo.catching_up ) == true ]]
        do
            node_block_height=$( echo $node_status_json | jq -r .SyncInfo.latest_block_height )
            echo "Node is catching up to chain height... ~$( printf '%.1f\n' $( echo "100*$node_block_height/$chain_block_height" | bc -l ) )% done"
            sleep 10
            node_status_json=$( dualityd status )
        done
        echo "Node has caught up to chain height"

        # add validator key (--no-backup ensures the terminal from seeing/logging the MNEMONIC value)
        echo $MNEMONIC | dualityd keys add validator --recover --no-backup

        # sent request to become a validator (to the first RPC address defined)
        dualityd tx staking create-validator \
            --node-id `dualityd tendermint show-node-id` \
            --pubkey `dualityd tendermint show-validator` \
            --commission-rate="${VALIDATOR_COMMISSION_RATE:-1.0}" \
            --commission-max-rate="${VALIDATOR_COMMISSION_MAX_RATE:-1.0}" \
            --commission-max-change-rate="${VALIDATOR_COMMISSION_MAX_CHANGE_RATE:-1.0}" \
            --min-self-delegation="${VALIDATOR_MIN_SELF_DELEGATION:-1}" \
            --gas="${VALIDATOR_GAS:-auto}" \
            --amount "${VALIDATOR_AMOUNT:-1000000stake}" \
            --fees "${VALIDATOR_FEES:-0token}" \
            --from validator -y

        # wait to check the node's validator status
        voting_power=0
        pub_key=$( dualityd tendermint show-validator | jq .key )
        while [[ ! $voting_power -gt 0 ]]
        do
            validator_set=$( wget --tries 3 -q -O - $rpc_address/validators );
            new_voting_power=$( echo $validator_set | jq -r ".result.validators[] | select(.pub_key.value == $pub_key).voting_power" )
            if [[ "$new_voting_power" == "" ]]
            then
                echo "Validator is in not validator set yet..."
                sleep 10
            elif [[ "$new_voting_power" != "$voting_power" ]]
            then
                voting_power=$new_voting_power;
                echo "Validator is in validator set, voting power: $voting_power"
            fi
        done
        echo "Validator node setup complete."

        # keep container running
        tail -f /dev/null;

    else
        # start as not a validator
        echo "Starting fullnode..."
        dualityd --log_level ${LOG_LEVEL:-info} start
    fi
fi