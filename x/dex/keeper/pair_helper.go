package keeper

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/x/dex/types"
)

func SortTokens(tokenA, tokenB string) (string, string, error) {
	relativeOrder := tokenA < tokenB

	equalCheck := tokenA == tokenB
	if equalCheck {
		return "", "", sdkerrors.Wrapf(types.ErrInvalidTradingPair, "%s<>%s", tokenA, tokenB)
	} else if relativeOrder {
		return tokenA, tokenB, nil
	} else {
		return tokenB, tokenA, nil
	}
}

func SortAmounts(tokenA, token0 string, amountsA, amountsB []sdk.Int) ([]sdk.Int, []sdk.Int) {
	if tokenA == token0 {
		return amountsA, amountsB
	} else {
		return amountsB, amountsA
	}
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

func GetInOutTokens(tokenIn_, tokenA, tokenB string) (tokenIn, tokenOut string) {
	if tokenIn_ == tokenA {
		return tokenA, tokenB
	} else {
		return tokenB, tokenA
	}
}

func StringToPairID(pairIDStr string) (*types.PairID, error) {
	tokens := strings.Split(pairIDStr, "<>")

	if len(tokens) == 2 {
		return &types.PairID{
			Token0: tokens[0],
			Token1: tokens[1],
		}, nil
	} else {
		return &types.PairID{}, sdkerrors.Wrapf(types.ErrInvalidPairIDStr, pairIDStr)
	}
}

func NormalizeTickIndex(baseToken, token0 string, tickIndex int64) int64 {
	if baseToken != token0 {
		tickIndex = tickIndex * -1
	}
	return tickIndex
}

func NormalizeAllTickIndexes(baseToken, token0 string, tickIndexes []int64) []int64 {
	for i, idx := range tickIndexes {
		tickIndexes[i] = NormalizeTickIndex(baseToken, token0, idx)
	}
	return tickIndexes
}
