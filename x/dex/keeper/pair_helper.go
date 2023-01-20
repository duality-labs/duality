package keeper

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/x/dex/types"
)

func SortTokens(ctx sdk.Context, tokenA string, tokenB string) (string, string, error) {

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

func SortAmounts(tokenA string, token0 string, amountsA []sdk.Int, amountsB []sdk.Int) ([]sdk.Int, []sdk.Int) {
	if tokenA == token0 {
		return amountsA, amountsB
	} else {
		return amountsB, amountsA
	}
}

func CreatePairId(token0 string, token1 string) (pairId *types.PairId) {
	return &types.PairId{
		Token0: token0,
		Token1: token1,
	}
}

func GetInOutTokens(tokenIn_ string, tokenA string, tokenB string) (tokenIn string, tokenOut string) {
	if tokenIn_ == tokenA {
		return tokenA, tokenB
	} else {
		return tokenB, tokenA
	}
}

func StringToPairId(pairIdStr string) (*types.PairId, error) {
	tokens := strings.Split(pairIdStr, "<>")

	if len(tokens) == 2 {
		return &types.PairId{
			Token0: tokens[0],
			Token1: tokens[1],
		}, nil
	} else {
		return &types.PairId{}, sdkerrors.Wrapf(types.ErrInvalidPairIdStr, pairIdStr)
	}
}
