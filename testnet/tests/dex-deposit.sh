#!/bin/bash
set -e

test_name="dex: can make deposit"

# todo: check state before? (may require test in serious)
# bash /root/.duality/scripts/expect_state_value_length_to_be.sh ".app_state.dex.sharesList" 0

# deposit tokens
# wait for tx to be processed and return an exit code
tx_result=$(dualityd tx dex deposit --node tcp://dualitynode0:26657 $(dualityd keys show alice --output json | jq -r .address) token stake 0.0000000001 0.0000000001 1 0 --from alice --yes --output json --broadcast-mode block)
tx_code=$(echo $tx_result | jq -r .code)
if [[ "$tx_code" != "0" ]]
then
    echo "Error at $test_name: $tx_result"
    exit $tx_code
fi

# todo: time the above function, it is a rough calculation of round trip processing time of the function

# todo: somehow time how long it takes just the function to be processed
# we could use --broadcast-mode async to send just the transaction and then time the response time
# this could be quite inaccurate due to the chain may be waiting for other Msgs to appear before finalizing the transaction.
# bash /root/.duality/scripts/wait_for_transaction_to_equal_code.sh $txhash 0

# todo: check state after? (requires tests to be done in series)
# bash /root/.duality/scripts/expect_state_value_length_to_be.sh ".app_state.dex.sharesList" 1
