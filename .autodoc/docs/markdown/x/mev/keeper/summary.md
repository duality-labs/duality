[View code on GitHub](https://github.com/duality-labs/duality/utodoc/docs/json/x/mev/keeper)

The `keeper` package in the `duality` project plays a crucial role in managing the state of the blockchain and handling various operations, such as querying data, sending coins, and managing parameters. It contains several important files, each with specific functionality.

`grpc_query.go` defines a `Keeper` struct that implements the `QueryServer` interface from the `types` package. This allows the `duality` project to query data from the `keeper` module. For instance, to retrieve data from the `keeper` module, a query can be made using the methods defined in the `QueryServer` interface:

```go
import (
    "github.com/duality-labs/duality/keeper"
    "github.com/duality-labs/duality/x/mev/types"
)

func main() {
    k := keeper.Keeper{}
    query := types.Query{...}
    response := k.Query(query)
}
```

`grpc_query_params.go` contains a `Params` function that retrieves the current parameters of the Duality network. It takes a context and a `QueryParamsRequest` object as arguments and returns a `QueryParamsResponse` object containing the current parameters:

```go
import (
    "context"
    "github.com/duality-labs/duality/x/mev/types"
    "github.com/duality-labs/duality/keeper"
)

func main() {
    ctx := context.Background()
    req := &types.QueryParamsRequest{}
    k := keeper.NewKeeper()
    params, err := k.Params(ctx, req)
    displayParams(params)
}
```

`keeper.go` contains the `Keeper` struct and a `NewKeeper` function that returns an instance of this struct. The `Keeper` struct is responsible for interacting with the state of the `duality` blockchain, including reading and writing data, managing parameters, and handling transactions and events.

`msg_server.go` defines a `msgServer` struct that implements the `types.MsgServer` interface, providing an implementation for the `Keeper` struct. This allows for efficient and organized message handling in the `duality` project:

```go
keeper := NewKeeper(...)
msgServer := NewMsgServerImpl(keeper)
```

`msg_server_send.go` contains a `Send` function that sends coins from a user's account to a module's account. It takes a context and a `MsgSend` message as arguments and returns a `MsgSendResponse` object:

```go
import (
    "context"
    "github.com/duality-labs/duality/x/mev/types"
)

func main() {
    msg := &types.MsgSend{
        Creator: "user1",
        TokenIn: "dual",
        AmountIn: 100,
    }
    ctx := context.Background()
    response, err := Send(ctx, msg)
}
```

`params.go` defines `GetParams` and `SetParams` functions that allow for the retrieval and setting of parameters for the `mev` module, which handles miner-extractable value (MEV) transactions on the `duality` blockchain:

```go
k := Keeper{}
params := k.GetParams(ctx)
k.SetParams(ctx, params)
```

In summary, the `keeper` package provides essential functionality for the `duality` project, enabling interaction with the blockchain state and handling various operations.
