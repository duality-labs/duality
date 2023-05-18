package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func SortAmounts(tokenA, token0 string, amountsA, amountsB []sdk.Int) ([]sdk.Int, []sdk.Int) {
	if tokenA == token0 {
		return amountsA, amountsB
	}

	return amountsB, amountsA
}

func GetInOutTokens(tokenIn, tokenA, tokenB string) (_, tokenOut string) {
	if tokenIn == tokenA {
		return tokenA, tokenB
	}

	return tokenB, tokenA
}

func NormalizeTickIndex(baseToken, token0 string, tickIndex int64) int64 {
	if baseToken != token0 {
		tickIndex *= -1
	}

	return tickIndex
}

func NormalizeAllTickIndexes(baseToken, token0 string, tickIndexes []int64) []int64 {
	for i, idx := range tickIndexes {
		tickIndexes[i] = NormalizeTickIndex(baseToken, token0, idx)
	}

	return tickIndexes
}
