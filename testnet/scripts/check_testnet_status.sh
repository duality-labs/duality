#!/bin/bash

# check if we can get information from the testnet
abci_info=$( curl --retry 30 --retry-connrefused --retry-delay 1 -s http://dualitynode0:26657/abci_info )
if [[ "$( echo $abci_info | jq -r ".result.response.data" )" != "duality" ]]
then
  echo "Could not establish connection to Duality testnet"
  exit
fi

echo "Duality testnet available"
