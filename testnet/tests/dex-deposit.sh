#!/bin/bash
set -e

test_name="dex: can make deposit"

# todo: check state before? (may require test in serious)
# bash /root/.duality/scripts/expect_state_value_length_to_be.sh ".app_state.dex.sharesList" 0

echo "$test_name: setup"

# create new person with funds
person=$(bash /root/.duality/scripts/test_helpers.sh createAndFundUser 1000tokenA,1000tokenB)

echo "$test_name: start test"
# deposit tokens
# wait for tx to be processed and return an exit code
# (amounts 0.0000000000000001 are measured in token main denom, whereas when using the bank send function (1000) that is in the small denom which is 1e18x smaller)
tx_result=$(dualityd tx dex deposit $(dualityd keys show $person --output json | jq -r .address) tokenA tokenB 0.0000000000000001 0.0000000000000001 1 0 --from "$person" --yes --output json --broadcast-mode block)

# assert that result has no errors
bash /root/.duality/scripts/test_helpers.sh throwOnTxError "$test_name" "$tx_result"

echo "$test_name: Deposited coins to ticks"

# todo: check state after? (requires tests to be done in series)
# bash /root/.duality/scripts/expect_state_value_length_to_be.sh ".app_state.dex.sharesList" 1
