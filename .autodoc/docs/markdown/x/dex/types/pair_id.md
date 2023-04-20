[View code on GitHub](https://github.com/duality-labs/duality/dex/types/pair_id.go)

The `types` package contains code related to data types used in the duality project. This particular file defines a type called `PairID` and provides methods to manipulate it.

The `PairID` type represents a pair of tokens in the duality project. It has two fields, `Token0` and `Token1`, which are strings representing the two tokens in the pair.

The `Stringify` method takes a `PairID` object and returns a string representation of it. The string is constructed by concatenating the two token strings with the `<>` separator. For example:

```
pair := &PairID{Token0: "BTC", Token1: "ETH"}
str := pair.Stringify() // "BTC<>ETH"
```

The `OppositeToken` method takes a `PairID` object and a token string, and returns the opposite token in the pair. If the supplied token matches neither side of the pair, it returns an empty string and a `false` boolean value. For example:

```
pair := &PairID{Token0: "BTC", Token1: "ETH"}
oppToken, ok := pair.OppositeToken("BTC") // oppToken = "ETH", ok = true
```

The `MustOppositeToken` method is similar to `OppositeToken`, but it panics if the supplied token matches neither side of the pair. This method is useful when the caller is certain that the supplied token is valid. For example:

```
pair := &PairID{Token0: "BTC", Token1: "ETH"}
oppToken := pair.MustOppositeToken("BTC") // oppToken = "ETH"
```

The `StringToPairID` function takes a string representation of a `PairID` object and returns a pointer to a `PairID` object. If the string is in the correct format (i.e. contains two token strings separated by `<>`), it constructs a new `PairID` object with the token strings and returns a pointer to it. If the string is not in the correct format, it returns an error wrapped in a `sdkerrors` object. For example:

```
pairStr := "BTC<>ETH"
pair, err := StringToPairID(pairStr) // pair = &PairID{Token0: "BTC", Token1: "ETH"}, err = nil

pairStr = "BTC-ETH"
pair, err = StringToPairID(pairStr) // pair = &PairID{}, err = ErrInvalidPairIDStr
```
## Questions: 
 1. What is the purpose of the `PairID` struct and its associated methods?
- The `PairID` struct represents a pair of tokens and its methods allow for stringifying the pair, getting the opposite token of a given token, and getting the opposite token of a given token with a panic if the token is not part of the pair.

2. What is the `sdkerrors` package used for?
- The `sdkerrors` package is used to define and handle errors specific to the Cosmos SDK.

3. What is the purpose of the `StringToPairID` function?
- The `StringToPairID` function takes a string in the format of "token0<>token1" and returns a `PairID` struct with the corresponding tokens. If the input string is not in the correct format, it returns an error wrapped in an `sdkerrors` error.