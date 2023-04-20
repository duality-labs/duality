[View code on GitHub](https://github.com/duality-labs/duality/keeper/pair_helper.go)

The `keeper` package in the Duality project provides utility functions for handling token pairs and their associated data. These functions are essential for managing trading pairs and their related information in the decentralized exchange (DEX) module of the project.

`SortTokens(tokenA, tokenB string)` takes two token strings as input and returns them in a sorted order. If the tokens are the same, it returns an error indicating an invalid trading pair.

`SortAmounts(tokenA, token0 string, amountsA, amountsB []sdk.Int)` takes two token strings and two corresponding lists of amounts. It returns the lists of amounts in the same order as the sorted tokens.

`CreatePairID(token0, token1 string)` creates a `PairID` struct with the given sorted tokens.

`CreatePairIDFromUnsorted(tokenA, tokenB string)` takes two unsorted tokens, sorts them, and creates a `PairID` struct.

`GetInOutTokens(tokenIn, tokenA, tokenB string)` takes an input token and two other tokens. It returns the input token and the other token that is not the input token.

`NormalizeTickIndex(baseToken, token0 string, tickIndex int64)` takes a base token, a sorted token, and a tick index. It returns the tick index with the correct sign based on the order of the base token and the sorted token.

`NormalizeAllTickIndexes(baseToken, token0 string, tickIndexes []int64)` takes a base token, a sorted token, and a list of tick indexes. It returns the list of tick indexes with the correct signs based on the order of the base token and the sorted token.

These utility functions are used throughout the DEX module to manage trading pairs, their associated amounts, and tick indexes. They ensure that the data is consistently sorted and formatted, which is crucial for the correct functioning of the DEX.
## Questions: 
 1. **Question**: What is the purpose of the `SortTokens` function and what are the possible return values?
   **Answer**: The `SortTokens` function is used to sort two input tokens (tokenA and tokenB) in alphabetical order. It returns the sorted tokens (either as tokenA, tokenB or tokenB, tokenA) and a nil error if successful, or an empty string and an error if the input tokens are the same.

2. **Question**: How does the `CreatePairIDFromUnsorted` function work and when should it be used?
   **Answer**: The `CreatePairIDFromUnsorted` function takes two unsorted tokens (tokenA and tokenB) as input, sorts them using the `SortTokens` function, and then creates a PairID using the sorted tokens with the `CreatePairID` function. It should be used when you have two unsorted tokens and need to create a PairID for them.

3. **Question**: What is the purpose of the `NormalizeTickIndex` and `NormalizeAllTickIndexes` functions, and how do they differ?
   **Answer**: The `NormalizeTickIndex` function takes a baseToken, token0, and a tickIndex as input, and returns the normalized tickIndex based on whether the baseToken is equal to token0 or not. The `NormalizeAllTickIndexes` function takes the same baseToken and token0, but instead of a single tickIndex, it takes a slice of tickIndexes and normalizes all of them using the `NormalizeTickIndex` function. The main difference is that `NormalizeTickIndex` works on a single tickIndex, while `NormalizeAllTickIndexes` works on a slice of tickIndexes.