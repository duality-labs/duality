[View code on GitHub](https://github.com/duality-labs/duality/dex/keeper/grpc_query_limit_order_tranche.go)

The `keeper` package contains two functions that are used to query limit order tranches in the Duality project. The `LimitOrderTrancheAll` function returns all active limit order tranches for a given pairID/tokenIn combination. It does not return inactive limit order tranches. The function takes a context and a `QueryAllLimitOrderTrancheRequest` as input and returns a `QueryAllLimitOrderTrancheResponse` and an error. The function first checks if the request is valid and then retrieves the limit order tranches from the KVStore using the `TickLiquidityPrefix` function. It then filters the retrieved tranches to only include limit order tranches and appends them to a slice. Finally, it returns the slice of limit order tranches and a pagination response.

The `LimitOrderTranche` function returns a specific limit order tranche either from the `tickLiquidity` index or from the `FillLimitOrderTranche` index. The function takes a context and a `QueryGetLimitOrderTrancheRequest` as input and returns a `QueryGetLimitOrderTrancheResponse` and an error. The function first checks if the request is valid and then retrieves the limit order tranche from the KVStore using the `FindLimitOrderTranche` function. If the limit order tranche is not found, the function returns an error.

These functions are used to query limit order tranches in the Duality project. The `LimitOrderTrancheAll` function can be used to retrieve all active limit order tranches for a given pairID/tokenIn combination, while the `LimitOrderTranche` function can be used to retrieve a specific limit order tranche. These functions are part of the `keeper` package, which is responsible for managing the state of the Duality blockchain.
## Questions: 
 1. What is the purpose of the `LimitOrderTrancheAll` function?
- The `LimitOrderTrancheAll` function returns all active limit order tranches for a given pairID/tokenIn combination, excluding inactiveLimitOrderTranches.

2. What is the difference between the `LimitOrderTrancheAll` and `LimitOrderTranche` functions?
- The `LimitOrderTrancheAll` function returns all active limit order tranches for a given pairID/tokenIn combination, while the `LimitOrderTranche` function returns a specific limit order tranche either from the tickLiquidity index or from the FillLimitOrderTranche index.

3. What is the purpose of the `prefix.NewStore` function call in the `LimitOrderTrancheAll` function?
- The `prefix.NewStore` function call creates a new prefix store for the given pairID and tokenIn combination, which is used to retrieve the limit order tranches from the KVStore.