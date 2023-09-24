#!/bin/sh

# set variable defaults
export NETWORK="${NETWORK:-duality-devnet}"
export CHAIN_ID="${CHAIN_ID:-$NETWORK}"
export STARTUP_MODE="${MODE:-fullnode}" # New network can be created with STARTUP_MODE=new
export NODE_MONIKER="${NODE_MONIKER:-duality-node}"
export IS_MAINNET=${IS_MAINNET:-$([[ "$NETWORK" =~ "^duality-\d+$" ]] && echo "true" || echo "false")}
export genesis_file="./networks/${NETWORK}/genesis.json"

echo -e "NETWORK: $NETWORK \nCHAIN_ID: $CHAIN_ID\nMAINNET: $IS_MAINNET\nMONIKER:$NODE_MONIKER"
# check current working directorys
if [[ ! -e scripts/startup.sh ]]; then
    echo >&2 "Please run this script from the base repo directory"
    exit 1
fi

echo "Startup mode: $STARTUP_MODE"

echo "Initializing chain..."
dualityd init --chain-id $CHAIN_ID $NODE_MONIKER

# Add consumer section to the ICS chain
dualityd add-consumer-section

# Update config files
./scripts/config_setup.sh


# Use genesis file from supplied GENESIS_FILE_URL
if [ ! -z "$GENESIS_FILE_URL" ]; then
    if $(wget -O $HOME/.duality/config/genesis.json $GENESIS_FILE_URL); then
        echo "Loaded genesis from: $GENESIS_FILE_URL"
    else
        echo "Cannot load genesis.json from: $GENESIS_FILE_URL"
        exit 1
    fi

# Use network genesis file if present
elif [ -f $genesis_file ] && [ $STARTUP_MODE == "new" ]; then
    echo "Using network genesis.json: $genesis_file"
    cp $genesis_file ${HOME}/.duality/

# Fallback to auto-generated genesis.json
elif [ $STARTUP_MODE == "new" ]; then
    echo "Using auto-generated genesis file"
fi

# Start or join a chain
# note: "new" should not be used for mainnet chains
# for mainnets a custom genesis file should be curated outside of these scripts
if [ $STARTUP_MODE == "new" ]
then
    ./scripts/leader-startup.sh
elif [ $STARTUP_MODE == "fullnode" ]; then
   ./scripts/follower-startup.sh
else
    echo "Invalid startup mode: $STARTUP_MODE"
fi

if [[ $IS_MAINNET == true ]]; then
   echo "Removing unsafe binaries"

   rm -r /usr/lib/ /usr/bin /bin/busybox
fi

if [ "$KEEP_RUNNING" != "false" ]
then
    tail -f /dev/null;
fi
