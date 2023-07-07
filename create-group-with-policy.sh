#!/bin/bash -e

nodehome=$(mktemp -d)

cleanup() {
    echo $?
    kill $dualityd_pid >/dev/null 2>&1
    rm -rf $nodehome
}
trap cleanup ERR
trap 'exit_code=$?; cleanup; exit $exit_code' EXIT

dualityd init local-node --home $nodehome >/dev/null 2>&1
dualityd add-consumer-section --home $nodehome >/dev/null 2>&1
dualityd keys add user --keyring-backend test --home $nodehome >/dev/null 2>&1
user=$(dualityd keys show -a user --keyring-backend test --home $nodehome)
dualityd add-genesis-account user 100000000000stake,100000000000token \
    --keyring-backend test \
    --home $nodehome >/dev/null 2>&1
dualityd start --home $nodehome >/dev/null 2>&1 &
dualityd_pid=$!
sleep 10
dualityd tx group create-group-with-policy \
    "$user" \
    "" \
    "" \
    members.json \
    policy.json \
    --group-policy-as-admin \
    --from $user \
    --keyring-backend test \
    --home $nodehome \
    --yes >/dev/null 2>&1
sleep 10
kill $dualityd_pid
dualityd export --home $nodehome  > "$nodehome"/genesis.json
dasel -f "$nodehome/genesis.json" ".app_state.group"