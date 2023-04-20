[View code on GitHub](https://github.com/duality-labs/duality/dex/keeper/grpc_query_pool_reserves.go)

The `keeper` package contains the implementation of the `Keeper` struct, which is responsible for managing the state of the duality blockchain. This file contains two methods that allow querying the pool reserves of a given token pair.

The `PoolReservesAll` method takes a `QueryAllPoolReservesRequest` object as input and returns a `QueryAllPoolReservesResponse` object. The request object contains the `PairID` and `TokenIn` fields, which are used to identify the token pair and the input token, respectively. The method retrieves the pool reserves for the specified token pair and token from the state store and returns them in a paginated response. The `FilteredPaginate` function is used to iterate over the pool reserves and filter out any limit order tranches. The resulting pool reserves are returned in the response object.

The `PoolReserves` method takes a `QueryGetPoolReservesRequest` object as input and returns a `QueryGetPoolReservesResponse` object. The request object contains the same fields as the `QueryAllPoolReservesRequest` object, as well as the `TickIndex` and `Fee` fields, which are used to identify the specific pool reserves to retrieve. The method retrieves the pool reserves for the specified token pair, token, tick index, and fee from the state store and returns them in the response object.

These methods are used to query the pool reserves of a token pair in the duality blockchain. The `PoolReservesAll` method can be used to retrieve all the pool reserves for a given token pair, while the `PoolReserves` method can be used to retrieve a specific pool reserve for a given tick index and fee. These methods are likely used by other modules in the duality blockchain to perform various operations, such as executing trades or calculating liquidity.
## Questions: 
 1. What is the purpose of this code?
   
   This code defines two functions `PoolReservesAll` and `PoolReserves` that query pool reserves for a given token pair and tick index in a decentralized exchange (DEX) implemented using the Cosmos SDK.

2. What external dependencies does this code have?
   
   This code imports several packages from the Cosmos SDK, including `sdk`, `query`, and `types`, as well as the `prefix` package for working with key-value stores. It also imports the `status` and `codes` packages from `google.golang.org/grpc` for error handling.

3. What is the expected input and output of the `PoolReservesAll` and `PoolReserves` functions?
   
   Both functions take a context and a request object as input and return a response object and an error as output. `PoolReservesAll` returns a list of pool reserves and pagination information for a given token pair and token in, while `PoolReserves` returns the pool reserves for a specific tick index and fee for the same token pair and token in.