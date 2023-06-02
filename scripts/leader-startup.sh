#!/bin/sh

# check current working directory
if [[ ! -e scripts/startup.sh ]]; then
    echo >&2 "Please run this script from the base repo directory"
    exit 1
fi

export NETWORK=${NETWORK:-duality-devnet}
export CHAIN_ID="${CHAIN_ID:-$NETWORK}"
echo "NETWORK: $NETWORK \n CHAIN_ID: $CHAIN_ID"

# define a million, billion, Carl Sagan's worth of minimum denomination to save space


dualityd init $NODE_MONIKER --chain-id $CHAIN_ID

# Add consumer section to the ICS chain
dualityd add-consumer-section

# setup accounts from mnemonic file

B=1000000000000000000000000
mnemonic_file="networks/${NETWORK}/mnemonics.txt"
if [ ! -f "$mnemonic_file" ]; then
    echo "File '$mnemonic_file' does not exist."
    exit 1
fi

# setup normal accounts
sed \$d ${mnemonic_file} | nl | while read line; do
    num=$(echo "$line" | awk '{print $1}')
    mnemonic=$(echo "$line" | awk {'$1=""; print $0'})
    acct_name="user${num}"
    echo "$mnemonic" | dualityd keys add $acct_name --recover --keyring-backend test
    dualityd add-genesis-account $acct_name ${B}token,${B}stake --keyring-backend test 

done

# Add faucet account
faucet_mnemonic=$(tail -n 1 $mnemonic_file)
echo $faucet_mnemonic | dualityd keys add faucet --recover --keyring-backend test
dualityd add-genesis-account faucet "${B}token,${B}stake,${B}tokenA,${B}tokenB,${B}tokenC,${B}tokenD,${B}tokenE,${B}tokenF,${B}tokenG,${B}tokenH,${B}tokenI,${B}tokenJ,${B}tokenK,${B}tokenL,${B}tokenM,${B}tokenN,${B}tokenO,${B}tokenP,${B}tokenQ,${B}tokenR,${B}tokenS,${B}tokenT,${B}tokenU,${B}tokenV,${B}tokenW,${B}tokenX,${B}tokenY,${B}tokenZ" --keyring-backend test



./scripts/config_setup.sh

echo "Starting new chain..."
dualityd --log_level ${LOG_LEVEL:-info} start


tail -f /dev/null;
