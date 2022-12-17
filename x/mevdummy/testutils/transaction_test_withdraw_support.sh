#!/bin/bash


ALICE_ADDRESS=$(cat alice.txt | grep "address:" | awk '{print $2}')
BOB_ADDRESS=$(cat bob.txt | grep "address:" | awk '{print $2}')
AMOUNT=$((100))
COIN=$((${AMOUNT} + stake ))
# Print the value of the ALICE_ADDRESS and BOB_ADDRESS variables
echo "Alice's address: ${ALICE_ADDRESS}"
echo "Bob's address: ${BOB_ADDRESS}"

dualityd tx mevdummy send ${AMOUNT} "stake" --from ${ALICE_ADDRESS} -y
sleep 7
#Query for all account balances
dualityd query auth accounts > chainAccounts.txt

#return the mevdummy module address
MEV_ADDRESS=$(grep -B 3 -o "mevdummy" chainAccounts.txt  | grep -m 1 "address: .*" | cut -d':' -f2 )
echo "${MEV_ADDRESS}"

MODULE_ORIGINAL_BALANCE_SNIPPET=$(dualityd query bank balances ${MEV_ADDRESS} )
MODULE_ORIGINAL_BALANCE=$(grep -B1 'denom: stake' <<< "$MODULE_ORIGINAL_BALANCE_SNIPPET" | grep -m 1 "amount: .*" | cut -d':' -f2 | grep -o -E '[^"]+')
echo "Module's balance pre-send: ${MODULE_ORIGINAL_BALANCE_SNIPPET}"

ALICE_ORIGINAL_BALANCE_SNIPPET=$(dualityd query bank balances ${ALICE_ADDRESS} )
ALICE_ORIGINAL_BALANCE=$(grep -B1 'denom: stake' <<< "$ALICE_ORIGINAL_BALANCE_SNIPPET" | grep -m 1 "amount: .*" | cut -d':' -f2 | grep -o -E '[^"]+')
ALICE_ORIGINAL_BALANCE_INT=$((ALICE_ORIGINAL_BALANCE))

echo "Alice's balances pre-send: ${ALICE_ORIGINAL_BALANCE}"


# Alice sends 100Stake to MEVDummy Module
dualityd tx mevdummy send ${AMOUNT} "stake" --from ${ALICE_ADDRESS} -y
sleep 7

MODULE_NEW_BALANCE_SNIPPET=$(dualityd query bank balances ${MEV_ADDRESS} )
MODULE_NEW_BALANCE=$(grep -B1 'denom: stake' <<< "$MODULE_NEW_BALANCE_SNIPPET" | grep -m 1 "amount: .*" | cut -d':' -f2 | grep -o -E '[^"]+')
MODULE_NEW_BALANCE_INT=$((MODULE_NEW_BALANCE))

echo "Module's balance post-send: ${MODULE_NEW_BALANCE_SNIPPET}"

ALICE_NEW_BALANCE=$(dualityd query bank balances ${ALICE_ADDRESS} )
echo "Alice's balances post-send: ${ALICE_NEW_BALANCE}"

EXPECTED_VALUE=$((MODULE_ORIGINAL_BALANCE + $AMOUNT ))

if [ $MODULE_NEW_BALANCE_INT != $EXPECTED_VALUE  ]; then
    echo "Error: Module's new balance does not equal the expected value" >&2
    exit 1
fi

dualityd tx mevdummy withdraw-funds $AMOUNT "stake" --from ${ALICE_ADDRESS} -y
sleep 7

MODULE_POST_WITHDRAW_BALANCE_SNIPPET=$(dualityd query bank balances ${MEV_ADDRESS} )
MODULE_POST_WITHDRAW_BALANCE=$(grep -B1 'denom: stake' <<< "$MODULE_POST_WITHDRAW_BALANCE_SNIPPET" | grep -m 1 "amount: .*" | cut -d':' -f2 | grep -o -E '[^"]+')
MODULE_POST_WITHDRAW_BALANCE_INT=$((MODULE_POST_WITHDRAW_BALANCE))

if [ $MODULE_POST_WITHDRAW_BALANCE_INT != $(($MODULE_NEW_BALANCE_INT - $AMOUNT )) ]; then
    echo "Error: Module's new balance does not equal the expected value" >&2
    exit 1
fi 

ALICE_POST_WITHDRAW_BALANCE_SNIPPET=$(dualityd query bank balances ${ALICE_ADDRESS} )
ALICE_POST_WITHDRAW_BALANCE=$(grep -B1 'denom: stake' <<< "$ALICE_POST_WITHDRAW_BALANCE_SNIPPET" | grep -m 1 "amount: .*" | cut -d':' -f2 | grep -o -E '[^"]+')
ALICE_POST_WITHDRAW_BALANCE_INT=$((ALICE_POST_WITHDRAW_BALANCE))

if [ $ALICE_ORIGINAL_BALANCE_INT != $ALICE_POST_WITHDRAW_BALANCE_INT ]; then
    echo "Error: Alice's new balance does not equal the expected value" >&2
    exit 1
fi

#Verifies that a user cannot send tokens from the module address
if ! dualityd tx bank send $MEV_ADDRESS $ALICE_ADDRESS $COIN; then
    echo "Bank Send fails as expected"
else 
    echo "Error: Bank Send did not fail" >&2
    exit
fi




