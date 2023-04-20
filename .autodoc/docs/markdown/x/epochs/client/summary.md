[View code on GitHub](https://github.com/duality-labs/duality/utodoc/docs/json/x/epochs/client)

The `query.go` file in the `epochs/client/cli` folder is responsible for providing a set of Command Line Interface (CLI) query commands for the epochs module in the duality project. These commands enable developers and users to obtain information about the current epoch and running epoch information, which can be helpful for debugging issues related to epoch transitions or monitoring the progress of the current epoch.

The primary function in this file is `GetQueryCmd`, which returns a `cobra.Command` object that can be used to execute the CLI commands. This function first calls `osmocli.QueryIndexCmd` to create a new `cobra.Command` object with the name of the epochs module. It then adds two query commands to the command object using `osmocli.AddQueryCmd`. The first command queries running epoch information using the `EpochInfos` function from the `types` package. The second command queries the current epoch by specified identifier using the `QueryCurrentEpochRequest` function from the `types` package.

In addition to the `GetQueryCmd` function, there are two other functions in this file that return `osmocli.QueryDescriptor` objects that describe the CLI commands for querying epoch information. The `GetCmdEpochInfos` function returns a descriptor for querying running epoch information, while the `GetCmdCurrentEpoch` function returns a descriptor for querying the current epoch by specified identifier.

Here's an example of how these CLI commands might be used:

```sh
# Query running epoch information
dualitycli query epochs epoch-infos

# Query the current epoch by specified identifier
dualitycli query epochs current-epoch --identifier "example-identifier"
```

In summary, the `query.go` file in the `epochs/client/cli` folder offers a set of CLI commands for querying epoch information in the duality project. These commands can be utilized by developers and users to retrieve information about the current epoch and running epoch information, which can be valuable for debugging issues related to epoch transitions or monitoring the progress of the current epoch.
