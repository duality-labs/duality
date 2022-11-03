#!/bin/sh


rm -r ~/.duality


dualityd init duality

dualityd keys add alice --keyring-backend test
dualityd add-genesis-account $(dualityd keys show alice -a --keyring-backend test) 1000000000token,1000000000stake 
# bob
dualityd keys add bob --keyring-backend test
dualityd add-genesis-account $(dualityd keys show bob -a --keyring-backend test) 1000000000token,1000000000stake --keyring-backend test

# Add gentxs to the genesis file
dualityd gentx alice 1000000stake --chain-id duality --keyring-backend test
dualityd collect-gentxs

dualityd start
