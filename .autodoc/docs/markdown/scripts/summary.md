[View code on GitHub](https://github.com/duality-labs/duality/utodoc/docs/json/scripts)

The `startup.sh` script in the `.autodoc/docs/json/scripts` folder is responsible for initializing and starting a Duality blockchain node. It automates the process of setting up a new chain or joining an existing chain and can be used to start a node as a full node or a validator node. This script is an essential part of the Duality project as it simplifies the process of starting and joining a Duality blockchain network.

The script begins by setting some default variables and checking if it is being run from the correct directory. It then initializes the chain using `dualityd init` with the specified network and chain ID. A consumer section is added to the ICS chain by running `dualityd add-consumer-section`. If a moniker is provided, the script replaces the moniker in the config file.

Depending on the startup mode, the script either creates a new chain or joins an existing one. For a new chain, it copies the genesis file, adds initial genesis data, creates test accounts, and starts the new chain. To join an existing chain, it uses the provided RPC address or reads it from the chain.json file, checks if the node is on the correct network, sets chain settings, and starts the node as a full node or a validator node.

If the node is started as a validator node, the script waits for the node to catch up to the chain's current height, adds the validator key, sends a request to become a validator, and checks the node's validator status. If the node is not started as a validator node, it starts the node as a full node.

Example usage:

To start a new chain:

```bash
./startup.sh MODE=new NETWORK=duality-1 MONIKER=my-node
```

To join an existing chain:

```bash
./startup.sh MODE=fullnode NETWORK=duality-1 RPC_ADDRESS=http://127.0.0.1:26657
```

To start a validator node:

```bash
./startup.sh MODE=validator NETWORK=duality-1 MNEMONIC="my mnemonic"
```

In the larger project, this script plays a crucial role in setting up and managing Duality blockchain nodes. It interacts with other parts of the project, such as the `dualityd` command-line tool and the chain configuration files, to automate the process of starting and joining a Duality blockchain network. This script is particularly useful for developers who want to quickly set up a node for testing or deployment purposes.
