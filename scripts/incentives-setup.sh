#!/bin/bash -ex


cleanup() {
    kill $dualityd_pid >/dev/null 2>&1
}
trap cleanup ERR
trap 'exit_code=$?; cleanup; exit $exit_code' EXIT

rm -rf ~/.duality
dualityd init local-node
dualityd add-consumer-section

epochs_genesis='{
    "epochs": [
        {
            "identifier": "day",
            "start_time": "0001-01-01T00:00:00Z",
            "duration": "10s",
            "current_epoch": "0",
            "current_epoch_start_time": "0001-01-01T00:00:00Z",
            "epoch_counting_started": false,
            "current_epoch_start_height": "0"
        },
        {
            "identifier": "hour",
            "start_time": "0001-01-01T00:00:00Z",
            "duration": "3600s",
            "current_epoch": "0",
            "current_epoch_start_time": "0001-01-01T00:00:00Z",
            "epoch_counting_started": false,
            "current_epoch_start_height": "0"
        },
        {
            "identifier": "week",
            "start_time": "0001-01-01T00:00:00Z",
            "duration": "604800s",
            "current_epoch": "0",
            "current_epoch_start_time": "0001-01-01T00:00:00Z",
            "epoch_counting_started": false,
            "current_epoch_start_height": "0"
        }
    ]
}'
dasel put -f ~/.duality/config/genesis.json -t json -v "$epochs_genesis" '.app_state.epochs' 

dualityd keys add user1 --keyring-backend test
user1=$(dualityd keys show -a user1 --keyring-backend test)
dualityd add-genesis-account user1 100000000000stake,100000000000token --keyring-backend test

dualityd keys add user2 --keyring-backend test
user2=$(dualityd keys show -a user2 --keyring-backend test)
dualityd add-genesis-account user2 100000000000stake,100000000000token --keyring-backend test

dualityd start >/dev/null 2>&1 &
dualityd_pid=$!
sleep 10

dualityd tx dex deposit $user1 stake token 100 100 0 1 false \
    --from user1 --keyring-backend test -y --broadcast-mode sync
sleep 10

dualityd tx incentives stake-tokens 200DualityPoolShares-stake-token-t0-f1 \
    --from user1 --keyring-backend test -y --broadcast-mode sync
sleep 10

dualityd tx incentives create-gauge token stake "[-10]" "10" 10000stake 5 0 \
    --from user2 --keyring-backend test -y --broadcast-mode sync
sleep 10

# dualityd q incentives list-gauges "ACTIVE_UPCOMING" ""
# sleep 10

# dualityd q incentives list-gauges "ACTIVE_UPCOMING" ""
# sleep 10

# dualityd q incentives list-gauges "ACTIVE_UPCOMING" ""
# sleep 10

# dualityd q incentives list-gauges "ACTIVE_UPCOMING" ""
# sleep 10

# dualityd q incentives list-gauges "ACTIVE_UPCOMING" ""
# sleep 10

# dualityd q incentives list-gauges "ACTIVE_UPCOMING" ""
# sleep 10