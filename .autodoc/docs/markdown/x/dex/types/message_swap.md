[View code on GitHub](https://github.com/duality-labs/duality/types/message_swap.go)

The code in this file defines a `MsgSwap` struct and its associated methods, which are part of the `types` package. The purpose of this code is to facilitate token swapping functionality within the larger project.

`MsgSwap` struct contains fields such as `Creator`, `TokenIn`, `TokenOut`, `AmountIn`, `MaxAmountOut`, and `Receiver`. These fields store information about the user initiating the swap, the input and output tokens, the input amount, the maximum output amount, and the receiver of the swapped tokens.

The `NewMsgSwap` function is a constructor that initializes a new `MsgSwap` instance with the provided parameters. This function can be used to create a new swap message with the desired token swap details.

The `Route`, `Type`, `GetSigners`, `GetSignBytes`, and `ValidateBasic` methods implement the `sdk.Msg` interface for the `MsgSwap` struct. These methods are used by the Cosmos SDK to handle and process the swap message.

- `Route` returns the router key, which is used to route the message to the appropriate module.
- `Type` returns the message type, which is "swap" in this case.
- `GetSigners` returns the account address of the creator, who is required to sign the message.
- `GetSignBytes` returns the JSON-encoded message in a canonical form, which is used for signing.
- `ValidateBasic` checks the validity of the message, such as ensuring that the creator and receiver addresses are valid, the input amount is positive, and the maximum output amount is non-negative.

Here's an example of how to create a new `MsgSwap` instance:

```go
msg := NewMsgSwap("cosmos1...", "tokenA", "tokenB", sdk.NewInt(100), sdk.NewInt(200), "cosmos2...")
```

This code creates a new swap message with the specified creator, input and output tokens, input amount, maximum output amount, and receiver. The message can then be processed by the Cosmos SDK to perform the token swap.
## Questions: 
 1. **What is the purpose of the `NewMsgSwap` function?**

   The `NewMsgSwap` function is a constructor that initializes and returns a new `MsgSwap` struct with the provided parameters, such as creator, tokenIn, tokenOut, amountIn, maxAmountOut, and receiver.

2. **How does the `GetSigners` function work and what does it return?**

   The `GetSigners` function converts the `msg.Creator` string into an `sdk.AccAddress` type using the `sdk.AccAddressFromBech32` function. If there is an error during the conversion, it panics. Otherwise, it returns a slice containing the creator's `sdk.AccAddress`.

3. **What does the `ValidateBasic` function do and what are the possible error cases?**

   The `ValidateBasic` function checks if the provided creator and receiver addresses are valid by converting them using `sdk.AccAddressFromBech32`. It also checks if the `MaxAmountIn` is positive and if the `MaxAmountOut` is not negative. If any of these conditions are not met, it returns an appropriate error.