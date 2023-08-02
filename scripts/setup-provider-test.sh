#!/bin/bash -ex

# install gaiad tag v11.0.0-rc0
# download genesis.json from here https://github.com/cosmos/testnets/blob/master/replicated-security/provider/provider-genesis.json 

rm -rf ~/.gaia
gaiad keys add admin --keyring-backend test
gaiad init local-node
curl https://raw.githubusercontent.com/cosmos/testnets/master/replicated-security/provider/provider-genesis.json | jq > ~/.gaia/config/genesis.json
# gaiad add-genesis-account admin 1000000000000000uatom --keyring-backend test
