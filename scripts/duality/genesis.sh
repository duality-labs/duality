#!/bin/sh

echo "GENESIS"
# duplicate genesis for easier merging and recovery
cp /root/.duality/config/genesis.json /root/.duality/config/genesis-init.json

# add genesis state of modules into genensis
# combine initial genesis data with all found pregenesis parts
# deepmerge from https://stackoverflow.com/questions/53661930/jq-recursively-merge-objects-and-concatenate-arrays#68362041
jq -s 'def deepmerge(a;b):
    reduce b[] as $item (a;
        reduce ($item | keys_unsorted[]) as $key (.;
        $item[$key] as $val | ($val | type) as $type | .[$key] = if ($type == "object") then
            deepmerge({}; [if .[$key] == null then {} else .[$key] end, $val])
        elif ($type == "array") then
            (.[$key] + $val | unique)
        else
            $val
        end)
        );
    deepmerge({}; .)' \
    /root/.duality/config/genesis-init.json $(find networks/$NETWORK/pregenesis | grep .*\.json$) \
    > /root/.duality/config/genesis.json

# add key
(echo -n $KEY_MNEMONIC) | dualityd keys add $KEY_NAME --recover --keyring-backend $KEYRING_BACKEND

# add genesis balances
# define a million, billion, Carl Sagan's worth of minimum denomination to save space
B=1000000000000000000000000

echo $KEY_PASSWD | dualityd add-genesis-account $KEY_NAME  --keyring-backend $KEYRING_BACKEND ${B}token,${B}stake
echo $KEY_PASSWD | dualityd gentx $KEY_NAME 1000000stake --chain-id $CHAIN_ID --keyring-backend $KEYRING_BACKEND

dualityd collect-gentxs
