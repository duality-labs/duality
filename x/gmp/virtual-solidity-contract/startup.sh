#!/bin/bash -eu

truffle compile &> truffle-compile.log
echo "ABI: $(jq '.abi' /usr/src/app/build/contracts/*.json)"

# Run ganache-cli in the background and truffle test in the foreground when the docker image is run
ganache-cli &> ganache.log &
truffle test &> truffle-test.log
echo ""
echo "Test encoding: $(cat ./abi-encoded-args.bin)"