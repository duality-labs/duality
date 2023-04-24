package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/x/dex/types"
)

func SortTokens(tokenA, tokenB string) (string, string, error) {
	switch {
	case tokenA == tokenB:
		return "", "", sdkerrors.Wrapf(types.ErrInvalidTradingPair, "%s<>%s", tokenA, tokenB)
	case tokenA < tokenB:
		return tokenA, tokenB, nil
	default:
		return tokenB, tokenA, nil
	}
}

func SortAmounts(tokenA, token0 string, amountsA, amountsB []sdk.Int) ([]sdk.Int, []sdk.Int) {
	if tokenA == token0 {
		return amountsA, amountsB
	}

	return amountsB, amountsA
}

func CreatePairID(token0, token1 string) (pairID *types.PairID) {
	return &types.PairID{
		Token0: token0,
		Token1: token1,
	}
}

func CreatePairIDFromUnsorted(tokenA, tokenB string) (*types.PairID, error) {
	token0, token1, err := SortTokens(tokenA, tokenB)
	if err != nil {
		return nil, err
	}

	return CreatePairID(token0, token1), nil
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
