package keeper

import (
	"strings"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func SortTokens(ctx sdk.Context, tokenA string, tokenB string) (string, string, error) {

	relativeOrder := tokenA < tokenB

	equalCheck := tokenA == tokenB
	if equalCheck {
		return "", "", sdkerrors.Wrapf(types.ErrInvalidTradingPair, "Not a valid Token Pair: tokenA and tokenB cannot be the same")
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

func CreatePairId(token0 string, token1 string) (pairId string) {
	return (token0 + "<>" + token1)
}

func PairToTokens(pairId string) (token0 string, token1 string) {
	tokens := strings.Split(pairId, "<>")

	return tokens[0], tokens[1]
}
