package keeper

import (

	//"strings"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) SortTokens(ctx sdk.Context, tokenA string, tokenB string) (string, string, error) {

	relativeOrder := tokenA < tokenB

	equalCheck := tokenA == tokenB
	if equalCheck {
		return "", "", sdkerrors.Wrapf(types.ErrInvalidTokenPair, "Not a valid Token Pair: tokenA and tokenB cannot be the same")
	} else if relativeOrder {
		return tokenA, tokenB, nil
	} else {
		return tokenB, tokenA, nil
	}
}

func (k Keeper) CreatePairId(token0 string, token1 string) (pairId string) {
	return (token0 + "<>" + token1)
}
