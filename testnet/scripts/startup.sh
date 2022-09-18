
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
    echo "Contacting network..."

    # check if we can get information from the current network
    abci_info_json=$( wget --tries 30 -q -O - $RPC_ADDRESS/abci_info )
    if [[ "$( echo $abci_info_json | jq -r ".result.response.data" )" == "duality" ]]
    then
        echo "Found Duality chain!"
    else
        echo "Could not establish connection to Duality chain"
        echo "Exiting..."
        exit 1
    fi

fi
