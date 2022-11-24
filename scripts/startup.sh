#!/bin/sh

# set variable defaults
MAINNET="duality-1"
NETWORK="${NETWORK:-$MAINNET}"
STARTUP_MODE="${MODE:-fullnode}"
NODE_MONIKER="${MONIKER}"

# check current working directorys
if [[ ! -e scripts/startup.sh ]]; then
    echo >&2 "Please run this script from the base repo directory"
    exit 1
fi

echo "Startup mode: $STARTUP_MODE"

echo "Initializing chain..."
dualityd init --chain-id $NETWORK duality

# Add consumer section to the ICS chain
dualityd add-consumer-section

# replace moniker in the config
if [ ! -z $NODE_MONIKER ]
then
    sed -i 's#moniker = ".*"#moniker = "'"$NODE_MONIKER"'"#' /root/.duality/config/config.toml
    # alternative below if using dasel
    # eg. dasel put string -f /root/.duality/config/config.toml ".moniker" "$NODE_MONIKER";
fi

# start or join a chain
# note: "new" should not be used for mainnet chains
# for mainnets a custom genesis file should be curated outside of these scripts
if [ $STARTUP_MODE == "new" ]
then

    # duplicate genesis for easier merging and recovery
    cp /root/.duality/config/genesis.json /root/.duality/config/genesis-init.json

    # add genesis state of modules into genensis
    # combine initial genesis data with all found pregenesis parts
    # deepmerge from https://stackoverflow.com/questions/53661930/jq-recursively-merge-objects-and-concatenate-arrays#68362041
    jq -s 'def deepmerge(a;b):
        reduce b[] as $item (a;
            reduce ($item | keys_unsorted[]) as $key (.;
            $item[$key] as $val | ($val | type) as $type | .[$key] = if ($type == "object") then
                deepmerge({}; [if .[$key] == null then {} else .[$key] end, $val])
            elif ($type == "array") then
                (.[$key] + $val | unique)
            else
                $val
            end)
            );
        deepmerge({}; .)' \
        /root/.duality/config/genesis-init.json $(find networks/$NETWORK/pregenesis | grep .*\.json$) \
        > /root/.duality/config/genesis.json

    # add new test accounts
    echo "Creating test accounts..."
    mkdir /root/.duality/testkeys
    # define a million, billion, Carl Sagan's worth of minimum denomination to save space
    B=1000000000000000000000000

    # alice
    dualityd keys add alice --keyring-backend test
    dualityd add-genesis-account $(dualityd keys show alice -a --keyring-backend test) ${B}token,${B}stake --keyring-backend test
    # bob
    dualityd keys add bob --keyring-backend test
    dualityd add-genesis-account $(dualityd keys show bob -a --keyring-backend test) ${B}token,${B}stake --keyring-backend test
    # custom account
    if [[ ! -z "$MNEMONIC" ]]
    then
        printf '\n%s\n%s\n' 'add custom account from mnemonic:' "$mnemonic"
        echo -n "$MNEMONIC" | dualityd keys add custom-user --recover --keyring-backend test
        dualityd add-genesis-account $(dualityd keys show custom-user -a --keyring-backend test) ${B}token,${B}stake --keyring-backend test
    fi
    # custom accounts
    if [[ ! -z "$MNEMONICS" ]]
    then
        printf '\n%s\n' 'add custom accounts from mnemonic'
        i=1
        while mnemonic=$(echo "$MNEMONICS" | cut -d\; -f$i | xargs echo -n); [ -n "$mnemonic" ]
        do
            printf '%s\n%s\n' "custom-user-$i from mnemonic:" "$mnemonic"
            echo $mnemonic | dualityd keys add custom-user-$i --recover --keyring-backend test
            dualityd add-genesis-account $(dualityd keys show custom-user-$i -a --keyring-backend test) ${B}token,${B}stake --keyring-backend test
            i=$((i+1))
        done
    fi
    # fred (faucet)
    if [[ ! -z "$FAUCET" ]]
    then
        echo -n "$FAUCET" | dualityd keys add fred --recover --keyring-backend test
    else
        dualityd keys add fred --keyring-backend test
    fi
    dualityd add-genesis-account $(dualityd keys show fred -a --keyring-backend test) "${B}token,${B}stake,${B}tokenA,${B}tokenB,${B}tokenC,${B}tokenD,${B}tokenE,${B}tokenF,${B}tokenG,${B}tokenH,${B}tokenI,${B}tokenJ,${B}tokenK,${B}tokenL,${B}tokenM,${B}tokenN,${B}tokenO,${B}tokenP,${B}tokenQ,${B}tokenR,${B}tokenS,${B}tokenT,${B}tokenU,${B}tokenV,${B}tokenW,${B}tokenX,${B}tokenY,${B}tokenZ" --keyring-backend test

    # do not add a validator gentx here as there is already a leading ICS validator
    # eg. dualityd gentx alice 1000000stake --chain-id $NETWORK --keyring-backend test
    # eg. dualityd collect-gentxs

    echo "Starting new chain..."
    dualityd --log_level ${LOG_LEVEL:-info} start
    exit

else

    if [[ ! -z "$FAUCET" ]]
    then
        echo -n "$FAUCET" | dualityd keys add fred --recover --keyring-backend test
    fi

    # find an RPC address to check the live chain with
    if [ ! -z "$RPC_ADDRESS" ]; then

        # if RPC address was provided directly then use that as the lead node
        echo "Will attempt to join the network using RPC address information"
        echo "Contacting network..."

        rpc_address=$RPC_ADDRESS
        echo "RPC ADDRESS: $rpc_address"

        # assert that address is an IP address
        if [[ ! "$rpc_address" =~ "https?://([0-9]{1,3}\.){3}[0-9]{1,3}\b" ]]; then
            echo "Must provide ENV variable RPC_ADDRESS as an IP address"
            exit 1
        fi

        # add genesis
        if $(wget -q -O - $rpc_address/genesis | jq .result.genesis > /root/.duality/config/genesis.json); then
            echo "Loaded genesis.json"
        else
            echo "Cannot load genesis.json from original chain"
            exit 1
        fi

        # add persistent peers
        genesis_ip=$(echo $rpc_address | grep -oE "\b([0-9]{1,3}\.){3}[0-9]{1,3}\b")
        # TODO: ideally this should parse listen_addr to get the port
        persistent_peers=$( wget -q -O - $rpc_address/status \
                            | jq -r --arg ip $genesis_ip '.result.node_info.id + "@" + $ip + ":26656"' )

    else

        # if RPC_ADDRESS was not provided then read it from the chain.json
        echo "Will attempt to join the network using chain.json information"
        echo "Contacting network..."

        rpc_address="$( jq -r .apis.rpc[0].address networks/$NETWORK/chain.json )"
        echo "RPC ADDRESS: $rpc_address"

        # add genesis
        mv networks/$NETWORK/genesis.json /root/.duality/config/genesis.json

        # add persistent peers
        persistent_peers_array="$( jq .peers.persistent_peers networks/$NETWORK/chain.json )"
        persistent_peers="$( echo $persistent_peers_array | jq -r 'map(.id + "@" + .address) | join(",")' )"

    fi

    # check we are on the correct network and can get information from the current network
    node_status_json=$( wget --tries 30 -q -O - $rpc_address/status )
    found_network=$( echo $node_status_json | jq -r ".result.node_info.network" )
    if [[ "$found_network" == "$NETWORK" ]]
    then
        echo "Found Duality chain!"
    else
        echo "Could not establish connection to Duality chain, found network: ${found_network-none}"
        echo "Exiting..."
        exit 1
    fi

    # set chain settings
    sed -i 's#persistent_peers = ""#persistent_peers = "'"$persistent_peers"'"#' /root/.duality/config/config.toml

    # check if this node intends to become a validator
    if [[ "$STARTUP_MODE" == "validator" && ! -z "$MNEMONIC" ]]
    then

        echo "Starting future validator fullnode..."
        dualityd --log_level ${LOG_LEVEL:-info} start & :

        # wait for node to finish catching up to the chain's current height
        sleep 5
        chain_block_height=$(echo $node_status_json | jq -r ".result.sync_info.latest_block_height")
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
        echo -n "$MNEMONIC" | dualityd keys add validator --recover --no-backup --keyring-backend test

        # sent request to become a validator (to the first RPC address defined)
        dualityd tx staking create-validator \
            --node-id `dualityd tendermint show-node-id` \
            --chain-id $NETWORK \
            --pubkey `dualityd tendermint show-validator` \
            --commission-rate="${VALIDATOR_COMMISSION_RATE:-1.0}" \
            --commission-max-rate="${VALIDATOR_COMMISSION_MAX_RATE:-1.0}" \
            --commission-max-change-rate="${VALIDATOR_COMMISSION_MAX_CHANGE_RATE:-1.0}" \
            --min-self-delegation="${VALIDATOR_MIN_SELF_DELEGATION:-1}" \
            --gas="${VALIDATOR_GAS:-auto}" \
            --amount "${VALIDATOR_AMOUNT:-1000000stake}" \
            --fees "${VALIDATOR_FEES:-0token}" \
            --keyring-backend test \
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

    else
        # start as not a validator
        echo "Starting fullnode..."
        dualityd --log_level ${LOG_LEVEL:-info} start & :
        echo "Started fullnode"
    fi

    # keep container running
    if [ "$KEEP_RUNNING" != "false" ]
    then
        tail -f /dev/null;
    fi
fi
