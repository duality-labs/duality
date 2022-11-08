#!/bin/sh


echo $MNEMONIC | dualityd keys add faucet --recover --no-backup --keyring-backend test

faucet_address=$(dualityd keys show faucet -a --keyring-backend test)

if [ -z $faucet_address  ]
then
    echo "Failed to import faucet account"
    exit 1
fi

faucet  ${TOKEN_AMOUNT:+-token-amount $TOKEN_AMOUNT} \
        ${PORT:+-port $PORT} \
        ${DENOMS:+-denoms $DENOMS} \
        -faucet-account $faucet_address \
        -node $RPC_NODE



