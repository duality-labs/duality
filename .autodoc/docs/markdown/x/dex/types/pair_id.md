[View code on GitHub](https://github.com/duality-labs/duality/types/pair_id.go)

The code in this file is part of the `types` package and primarily deals with the `PairID` struct, which represents a pair of tokens in the duality project. The purpose of this code is to provide utility functions for working with `PairID` instances, such as stringifying, finding the opposite token, and converting a string to a `PairID`.

The `Stringify` method returns a string representation of the `PairID` in the format "Token0<>Token1". This can be useful for displaying the pair in logs or user interfaces. For example, if a `PairID` has tokens "A" and "B", calling `Stringify` would return "A<>B".

The `OppositeToken` method takes a token as input and returns the opposite token in the pair, if it exists. This can be useful when working with pairs and needing to find the other token in the pair. For example, if a `PairID` has tokens "A" and "B", calling `OppositeToken("A")` would return "B" and `true`, while calling `OppositeToken("C")` would return an empty string and `false`.

The `MustOppositeToken` method is similar to `OppositeToken`, but it panics if the supplied token does not match either side of the pair. This can be useful when the code expects the token to be part of the pair and wants to enforce this constraint.

The `StringToPairID` function takes a string in the format "Token0<>Token1" and returns a `PairID` instance with the corresponding tokens. If the input string is not in the correct format, it returns an error. This can be useful when parsing user input or reading data from external sources. For example, calling `StringToPairID("A<>B")` would return a `PairID` with tokens "A" and "B", while calling `StringToPairID("invalid")` would return an error.
## Questions: 
 1. **What is the purpose of the `PairID` struct and its methods?**

   The `PairID` struct represents a pair of tokens, and its methods provide functionality to manipulate and interact with the token pair, such as converting it to a string, finding the opposite token, and converting a string back to a `PairID`.

2. **How does the `OppositeToken` method work and what does it return?**

   The `OppositeToken` method takes a token string as input and checks if it matches either `Token0` or `Token1` of the `PairID`. If it matches, the method returns the opposite token and a boolean value `true`. If it doesn't match, the method returns an empty string and `false`.

3. **What is the purpose of the `MustOppositeToken` method and when should it be used?**

   The `MustOppositeToken` method is a wrapper around the `OppositeToken` method that panics if the supplied token doesn't match either side of the pair. It should be used when the developer is certain that the input token is part of the pair and wants to avoid handling the error case explicitly.