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

func NormalizeTickIndex(takerDenom, token0 string, tickIndexTakerToMaker int64) int64 {
	if takerDenom != token0 {
		return tickIndexTakerToMaker * -1
	}

	return tickIndexTakerToMaker
}

func DenormalizeTickIndex(takerDenom, token0 string, tickIndexNormalized int64) int64 {
	if takerDenom != token0 {
		return tickIndexNormalized * -1
	}

	return tickIndexNormalized
}

func NormalizeAllTickIndexes(takerDenom, token0 string, tickIndexes []int64) []int64 {
	if takerDenom != token0 {
		result := make([]int64, len(tickIndexes))
		for i, idx := range tickIndexes {
			result[i] = idx * -1
		}
		return result
	}

	// NB: does not return a different slice because no change
	return tickIndexes
}
