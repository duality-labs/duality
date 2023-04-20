[View code on GitHub](https://github.com/duality-labs/duality/oc/docs/json/x/incentives/keeper)

The `keeper` package in the `incentives` module of the Duality project is responsible for managing the state of the module, which includes storing and retrieving data from the key-value store. It provides functions for creating, modifying, and retrieving gauges, which are used to distribute rewards to users based on certain conditions. Additionally, it handles the storage and retrieval of data related to incentives and staking.

For example, the `CreateGauge` function creates a gauge and sends coins to it. The `Stake` function stakes tokens, either adding to an existing stake or creating a new one. The `Unstake` function begins the unstaking of the specified stake, which enters the unstaking queue.

The `keeper` package also provides functions for managing references to objects, such as `addRefByKey`, `deleteRefByKey`, and `getRefs`. These functions can be used to manage references to objects in the Duality project, such as adding a reference to an object when it is created and removing the reference when the object is deleted.

Here's an example of how the `keeper` package might be used in the larger project:

```go
// create a new context object
ctx := sdk.NewContext(...)

// create a new keeper object
keeper := NewKeeper(...)

// create a new gauge
gauge := types.Gauge{
    ID:        "1",
    StartTime: time.Now(),
    EndTime:   time.Now().Add(time.Hour * 24),
    Coins:     sdk.NewCoins(sdk.NewCoin("abc", sdk.NewInt(100))),
}

// create the gauge and send coins to it
_, err := keeper.CreateGauge(ctx, &gauge)
if err != nil {
    panic(err)
}

// stake tokens
msgStake := types.MsgStake{
    Owner:     "cosmos1abc...",
    Coins:     sdk.NewCoins(sdk.NewCoin("abc", sdk.NewInt(100))),
    Duration:  time.Hour * 24,
}

_, err = keeper.Stake(ctx, &msgStake)
if err != nil {
    panic(err)
}

// unstake tokens
msgUnstake := types.MsgUnstake{
    Owner: "cosmos1abc...",
    ID:    "1",
}

_, err = keeper.Unstake(ctx, &msgUnstake)
if err != nil {
    panic(err)
}
```

In this example, we create a new context object and a new keeper object. We then create a new gauge, send coins to it, stake tokens, and unstake tokens. This demonstrates how the `keeper` package can be used to manage the state of the incentives module in the Duality project.
