# Setup group genesis

Run with:

```
./create-group-with-policy.sh > group_genesis.json
```

I've tested with various members.json addresses and each time the first group
policy address is
`cosmos1afk9zr2hn2jsac63h4hm60vl9z3e5u69gndzf7c99cqge3vzwjzsfwkgpd`. This means
we can hardcode that address into the dualityd binary before we have our final
group genesis generated.

To update the genesis.json with the final group_genesis.json run the following.

```
group_genesis=$(dasel -f "group_genesis.json" '.')
dasel put object -f "old_genesis.json" -s "$group_genesis" '.app_state.group' > new_genesis.json
```
