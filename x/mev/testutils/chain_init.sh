#!/bin/bash

# Remove the .duality directory
rm -rf ~/.duality

# Initialize a new duality blockchain
dualityd init duality

# Add a consumer section to the duality blockchain
dualityd add-consumer-section

# Generate a new key pair for the user "alice" and write the output to a file named "alice"
dualityd keys add alice > alice.txt

# Generate a new key pair for the user "bob" and write the output to a file named "bob"
dualityd keys add bob > bob.txt

# Define local variables based on the output of the dualityd keys add commands
ALICE_ADDRESS=$(cat alice.txt | grep "address:" | awk '{print $2}')
BOB_ADDRESS=$(cat bob.txt | grep "address:" | awk '{print $2}')

# Print the value of the ALICE_ADDRESS and BOB_ADDRESS variables
echo "Alice's address: ${ALICE_ADDRESS}"
echo "Bob's address: ${BOB_ADDRESS}"

# Use sed to change the "bank" and "auth" sections in the genesis.json file
sed -i '' "s/\"bank\": {/\"bank\": {\n  \"balances\": [\n    {\n      \"address\": \"${ALICE_ADDRESS}\",\n      \"coins\": [\n        {\n          \"denom\": \"stake\",\n          \"amount\": \"1000000000000\"\n        },\n        {\n          \"denom\": \"stake2\",\n          \"amount\": \"1000000000000\"\n        }\n      ]\n    },\n    {\n      \"address\": \"${BOB_ADDRESS}\",\n      \"coins\": [\n        {\n          \"denom\": \"stake\",\n          \"amount\": \"1000000000000\"\n        },\n        {\n          \"denom\": \"stake2\",\n          \"amount\": \"1000000000000\"\n        }\n      ]\n    }\n  ],/g" ~/.duality/config/genesis.json
sed -i '' "s/\"auth\": {/\"auth\": {\n  \"accounts\": [\n    {\n      \"@type\": \"\/cosmos.auth.v1beta1.BaseAccount\",\n      \"address\": \"${ALICE_ADDRESS}\",\n      \"pub_key\": null,\n      \"account_number\": \"0\",\n      \"sequence\": \"0\"\n    },\n    {\n      \"@type\": \"\/cosmos.auth.v1beta1.BaseAccount\",\n      \"address\": \"${BOB_ADDRESS}\",\n      \"pub_key\": null,\n      \"account_number\": \"0\",\n      \"sequence\": \"0\"\n    }\n  ],/g" ~/.duality/config/genesis.json

sed -i '' '52d' ~/.duality/config/genesis.json
sed -i '' '89,91d' ~/.duality/config/genesis.json 
sed -i '' '51 s/.$//' ~/.duality/config/genesis.json
sed -i '' '88 s/.$//' ~/.duality/config/genesis.json  

# Use sed to change the "FeeTierList" and "FeeTierCount" values in the genesis.json file
sed -i '' "s/\"FeeTierList\": \[\],/\"FeeTierList\": [\n    {\n        \"fee\": \"1\",\n        \"id\": \"0\"\n    },\n    {\n        \"fee\": \"3\",\n        \"id\": \"1\"\n    },\n    {\n        \"fee\": \"5\",\n        \"id\": \"2\"\n    },\n    {\n        \"fee\": \"10\",\n        \"id\": \"3\"\n    }\n],/g" ~/.duality/config/genesis.json
sed -i '' "s/\"FeeTierCount\": \"0\",/\"FeeTierCount\": \"4\",/g" ~/.duality/config/genesis.json