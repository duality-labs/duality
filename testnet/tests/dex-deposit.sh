#!/bin/bash
set -e

test_name="dex: can make deposit"

# todo: check state before? (may require test in serious)
# bash /root/.duality/scripts/expect_state_value_length_to_be.sh ".app_state.dex.sharesList" 0

echo "$test_name: setup"
person=$(openssl rand -hex 12)
dualityd keys add $person <<< $'asdfasdf\nn'
dualityd tx bank send $(dualityd keys show fred --output json | jq -r .address) $(dualityd keys show $person --output json | jq -r .address) 1000tokenA,1000tokenB -y --broadcast-mode block --output json

echo "$test_name: start test"
# deposit tokens
# wait for tx to be processed and return an exit code
tx_result=$(dualityd tx dex deposit $(dualityd keys show $person --output json | jq -r .address) tokenA tokenB 0.0000000000000001 0.0000000000000001 1 0 --from "$person" --yes --output json --broadcast-mode block)
tx_code=$(echo $tx_result | jq -r .code)
if [[ "$tx_code" != "0" ]]
then
    echo "$test_name error ($tx_code) at $(echo $tx_result | jq -r .txhash): $(echo $tx_result | jq -r .raw_log)"
    exit $tx_code
fi

echo "$test_name: Deposited coins to ticks"

# todo: time the above function, it is a rough calculation of round trip processing time of the function

# todo: somehow time how long it takes just the function to be processed
# we could use --broadcast-mode async to send just the transaction and then time the response time
# this could be quite inaccurate due to the chain may be waiting for other Msgs to appear before finalizing the transaction.
# bash /root/.duality/scripts/wait_for_transaction_to_equal_code.sh $txhash 0

# todo: check state after? (requires tests to be done in series)
# bash /root/.duality/scripts/expect_state_value_length_to_be.sh ".app_state.dex.sharesList" 1
