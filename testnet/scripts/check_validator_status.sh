#!/bin/bash

echo " --- Starting validator status checks --- ";

# poll if we are a validator for a minute
timeout 30 bash -c -- '
    voting_power=""
    pub_key=$( dualityd tendermint show-validator | jq .key )
    while [[ ! $voting_power -gt 0 ]]
    do
        validator_set=$( curl -s http://dualitynode2:26657/validators );
        new_voting_power=$( echo $validator_set | jq -r ".result.validators[] | select(.pub_key.value == $pub_key).voting_power" )
        if [[ "$new_voting_power" == "" ]]
        then
            echo "Validator is in not validator set"
            sleep 1
        elif [[ "$new_voting_power" != "$voting_power" ]]
        then
            voting_power=$new_voting_power;
            echo "Validator is in validator set, voting power: $voting_power"
        fi
    done
';

echo " --- Finished validator status checks --- ";
