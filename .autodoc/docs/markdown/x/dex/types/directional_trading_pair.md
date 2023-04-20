[View code on GitHub](https://github.com/duality-labs/duality/types/directional_trading_pair.go)

The `duality` code defines a package named `types` which contains a struct and associated methods for handling directional trading pairs in a trading system. The main purpose of this code is to manage and manipulate trading pairs with specific input and output tokens.

The `DirectionalTradingPair` struct consists of three fields:

- `PairID`: A pointer to a `PairID` struct, which presumably contains information about the trading pair's unique identifier.
- `TokenIn`: A string representing the input token for the trading pair.
- `TokenOut`: A string representing the output token for the trading pair.

The `NewDirectionalTradingPair` function is a constructor for creating a new `DirectionalTradingPair` instance. It takes a pointer to a `PairID` struct, an input token string, and an output token string as arguments, and returns a new `DirectionalTradingPair` with the provided values.

```go
pair := NewDirectionalTradingPair(pairID, "ETH", "BTC")
```

Two methods are defined on the `DirectionalTradingPair` struct:

1. `IsTokenInToken0`: This method checks if the input token (`TokenIn`) is equal to the first token in the `PairID` struct (`Token0`). It returns a boolean value, `true` if they are equal, and `false` otherwise.

```go
isTokenInToken0 := pair.IsTokenInToken0() // true if TokenIn is Token0
```

2. `IsTokenOutToken0`: This method checks if the output token (`TokenOut`) is equal to the first token in the `PairID` struct (`Token0`). It does this by calling the `IsTokenInToken0` method and returning the negation of its result. If `IsTokenInToken0` returns `true`, this method will return `false`, and vice versa.

```go
isTokenOutToken0 := pair.IsTokenOutToken0() // true if TokenOut is Token0
```

These methods can be useful in the larger project for determining the direction of a trade, validating trading pairs, or performing calculations based on the input and output tokens.
## Questions: 
 1. **Question:** What is the purpose of the `DirectionalTradingPair` struct and its fields?
   **Answer:** The `DirectionalTradingPair` struct represents a trading pair with a specific direction, containing a `PairID` pointer, and two strings `TokenIn` and `TokenOut` representing the input and output tokens for the trade.

2. **Question:** How does the `NewDirectionalTradingPair` function work and what are its parameters?
   **Answer:** The `NewDirectionalTradingPair` function is a constructor for creating a new `DirectionalTradingPair` instance. It takes a pointer to a `PairID`, and two strings `tokenIn` and `tokenOut` as parameters, and returns a new `DirectionalTradingPair` with the provided values.

3. **Question:** What do the `IsTokenInToken0` and `IsTokenOutToken0` methods do, and how do they relate to each other?
   **Answer:** The `IsTokenInToken0` method checks if the input token (`TokenIn`) is equal to the first token in the pair (`Token0`). The `IsTokenOutToken0` method checks if the output token (`TokenOut`) is equal to the first token in the pair by negating the result of `IsTokenInToken0`. These methods help determine the direction of the trading pair.