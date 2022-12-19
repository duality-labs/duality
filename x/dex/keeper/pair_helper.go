package keeper

import (
	"fmt"
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

func (k Keeper) CreatePairId(token0 string, token1 string) (pairId string) {
	return (token0 + "<>" + token1)
}

func (k Keeper) CreateSharesId(token0 string, token1 string, tickIndex int64, fee int64) (denom string) {
	t0 := strings.ReplaceAll(token0, "-", "")
	t1 := strings.ReplaceAll(token1, "-", "")
	return fmt.Sprintf("%s-%s-t%d-f%d", t0, t1, tickIndex, fee)
}
