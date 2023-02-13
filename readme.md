# Startup Instructions

## Setup Celestia Node

Install celestia node 

```
cd $HOME 
rm -rf celestia-node 
git clone https://github.com/celestiaorg/celestia-node.git 
cd celestia-node/ 
git checkout tags/v0.6.4 
make go-install 
make cel-key 
```

Create & fund wallet

```
./cel-key add <key_name> --keyring-backend test --node.type light --p2p.network mocha

```

request funds in celestia discord 

```$request <Wallet-Address>```

Initialize node 

```celestia light init ```



Start node


```
celestia light start --core.ip https://rpc-mocha.pops.one  --keyring.accname <key-name> --gateway --gateway.addr localhost --gateway.port 26659 --p2p.network mocha

```

For more detailed instructions on light node setup see [celestia docs](https://docs.celestia.org/nodes/light-node/)

## Run Duality 

init.sh in root directory should handle build and startup for duality on celestia 

```./init.sh```
