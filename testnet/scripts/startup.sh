
# start or join a chain
if [ -z $RPC_ADDRESS ]
then

    echo "No seed RPC given"
    echo "Starting new chain..."

    if [ ! -z $MONIKER ]
    then
        dualityd start --moniker $MONIKER
    else
        dualityd start
    fi
    exit

else

    echo "Have seed RPC: $RPC_ADDRESS"
    echo "Will attempt to join the network"

fi
