[View code on GitHub](https://github.com/duality-labs/duality/oc/docs/json/x/incentives/types)

The `types` package in the `duality` project contains various data types, functions, and interfaces used throughout the project, particularly for the incentives module. This module is responsible for managing the distribution of rewards to users who participate in the network.

The package provides functionality for registering concrete types and interfaces used for Amino JSON serialization and message services. It also defines sentinel errors for specific error conditions that may occur while using the `x/incentives` module. Additionally, it defines event types and attribute keys for the Incentive module, which are used to track and record various actions taken by users.

The package also contains interfaces that are expected to be implemented by other modules in the duality project, such as `BankKeeper`, `EpochKeeper`, `AccountKeeper`, and `DexKeeper`. These interfaces provide a way
