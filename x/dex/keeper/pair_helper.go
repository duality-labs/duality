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

	// Calculate sha256 Checksum for Token0, Token1
	//token0Hash := sha256.Sum256([]byte(tokenA))
	//token1Hash := sha256.Sum256([]byte(tokenB))

	//calculates an comparison integer for the two hashes
	//comparisonInt := bytes.Compare(token0Hash[:], token1Hash[:])

	//if comparisonInt == -1 {
	//	return tokenA, tokenB, nil
	//} else if comparisonInt == 0 {
	//return "", "", sdkerrors.Wrapf(types.ErrInvalidTokenPair, "Not a valid Token Pair: tokenA and tokenB cannot be the same")
	//} else {
	//	return tokenB, tokenA, nil
	//}

}

func (k Keeper) createPairId(token0 string, token1 string) (pairId string) {
	return (token0 + "/" + token1)
}
