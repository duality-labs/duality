#!/bin/sh


# define a million, billion, Carl Sagan's worth of minimum denomination to save space
B=1000000000000000000000000

# Setup normal accounts
if [ -z "$MNEMONICS" ]; then
   echo "No MNEMONICS provided"
else
    echo "Adding accounts from MNEMONICS"
    i=1
    while mnemonic=$(echo "$MNEMONICS" | cut -d\; -f$i); [ -n "$mnemonic" ]
    do
        echo $mnemonic | dualityd keys add user-$i --recover --keyring-backend test
        dualityd add-genesis-account $(dualityd keys show user-$i -a --keyring-backend test) ${B}token,${B}stake --keyring-backend test
        i=$((i+1))
    done
fi


# Add faucet account
if [ -z "$FAUCET_MNEMONIC" ]; then
   echo "No FAUCET_MNEMONIC"
else
    echo $FAUCET_MNEMONIC | dualityd keys add faucet --recover --keyring-backend test
    dualityd add-genesis-account faucet "${B}token,${B}stake,${B}tokenA,${B}tokenB,${B}tokenC,${B}tokenD,${B}tokenE,${B}tokenF,${B}tokenG,${B}tokenH,${B}tokenI,${B}tokenJ,${B}tokenK,${B}tokenL,${B}tokenM,${B}tokenN,${B}tokenO,${B}tokenP,${B}tokenQ,${B}tokenR,${B}tokenS,${B}tokenT,${B}tokenU,${B}tokenV,${B}tokenW,${B}tokenX,${B}tokenY,${B}tokenZ" --keyring-backend test
fi



echo "Starting new chain..."
dualityd --log_level ${LOG_LEVEL:-info} start & :


