# Overview 
This package provides a standalone web server that acts as a fuacet for funding new accounts. In order to function it must be able to recover an existing funded account using a mnemonic. 

# Build Docker Image
The docker image must be built from the top level directory so that the dependencies for dualityd can be included.
Docker build command: docker build -f ./testnet-faucet/Dockerfile -t faucet .

# Running 

The Go binary can be run directly with the following arguments:
* `-denoms` (string)
    Denoms to send (default "token,stake")
* `-faucet-account` (string)
    Account to use for faucet 
* `-node` (string)
    <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
* `-port` (string)
    Port to listen on (default "9000")
* `-token-amount` (int)
    Amount of token to send (default 1000)
        
The web server can also be initialized by running ./scripts/startup.sh
This startup script will also handle the initialization of the faucet wallet account. 
When running the startup script the following ENV variables can be set:

* `MNEMONIC`: Recovery seed for an already funded to be used as the faucet account
* `DENOMS`: see `-denoms' above`
* `TOKEN_AMOUNT`:  see `-token-amount` above
* `RPC_NODE`: See `-node` above
* `PORT`: See `-port` above

# Usage

Once running the faucet can be used by making a POST request to the faucet server nad passing in `address` as an argument:
```
curl -X POST -d '{"address": "[WALLET ADDRESS]"}' [FAUCET-ADDRESS]:[PORT]
```
ie. 
```
curl -X POST -d '{"address": "cosmos1xd2vqcglcp5l8epssskd7tjkwe34rh5tna2ttl"}' 44.234.34.140:9000
```

`
