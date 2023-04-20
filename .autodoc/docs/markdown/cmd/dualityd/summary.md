[View code on GitHub](https://github.com/duality-labs/duality/utodoc/docs/json/cmd/dualityd)

The `dualityd` folder contains code for initializing and executing the duality blockchain node, as well as modifying the genesis state of the network for testing purposes. The code is organized into two main files: `consumer.go` and `main.go`.

`consumer.go` focuses on adding a consumer section to the genesis state of the blockchain network. This is used for testing the interchain security module, which verifies the validity of transactions between different blockchain networks. The `AddConsumerSectionCmd` function returns a Cobra command that can be executed from the command line. When executed, it reads the current genesis state, modifies the consumer section using a callback function, and writes the updated genesis state back to disk. The code utilizes external packages such as `cosmos-sdk`, `tendermint`, and `interchain-security`, as well as a custom package called `duality` for testing purposes.

Example usage of `AddConsumerSectionCmd`:

```sh
$ dualityd add-consumer-section
```

`main.go` initializes and executes the duality blockchain node using the Cosmos SDK-based framework. The `main` function initializes the root command for the duality blockchain node using the `cosmoscmd.NewRootCmd` function, which takes several arguments such as the application name, account address prefix, default node home directory, module basics, and a function that creates a new instance of the application. The `rootCmd` variable is then used to add the `AddConsumerSectionCmd` to the root command. Finally, the `svrcmd.Execute` function is called to execute the root command and start the duality blockchain node.

Example usage of `main.go`:

```sh
$ dualityd start
```

Developers can build custom blockchain applications on top of the duality blockchain by adding new commands to the root command, providing additional functionality to the blockchain node.

In summary, the `dualityd` folder contains code for initializing and executing the duality blockchain node and modifying the genesis state for testing purposes. The code demonstrates how to work with the Cosmos SDK framework and Tendermint consensus engine to build custom blockchain applications.
