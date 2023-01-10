#!/bin/bash
set -e

test_name="dex: can make deposit"

# todo: check state before? (may require test in serious)
# bash /root/.duality/scripts/expect_state_value_length_to_be.sh ".app_state.dex.sharesList" 0

echo "$test_name: setup"

# create new person with funds
# (token amounts here are measured in utoken denom)
person=$(bash /root/.duality/scripts/test_helpers.sh createAndFundUser 1000tokenA,1000tokenB)

echo "$test_name: start test"
# deposit tokens
# wait for tx to be processed and return an exit code
# (token amounts here are measured in utoken denom)
tx_result=$(dualityd tx dex deposit $(dualityd keys show $person --output json | jq -r .address) tokenA tokenB 100 100 1 0 --from "$person" --yes --output json --broadcast-mode block --gas 300000)

# assert that result has no errors
bash /root/.duality/scripts/test_helpers.sh throwOnTxError "$test_name" "$tx_result"

echo "$test_name: Deposited coins to ticks"

# todo: check state after? (requires tests to be done in series)
# bash /root/.duality/scripts/expect_state_value_length_to_be.sh ".app_state.dex.sharesList" 1
