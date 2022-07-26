package keeper

import (
	"crypto/sha256"
	"bytes"
	//"strings"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) SortTokens(ctx sdk.Context, token0 string, token1 string) (string, string, error) {

	token0Hash := sha256.Sum256([]byte(token0))
	token1Hash := sha256.Sum256([]byte(token1))

	comparisonInt :=  bytes.Compare(token0Hash[:], token1Hash[:])

	if comparisonInt == -1 {
		return token0, token1, nil
	} else if comparisonInt == 0 {
		return "", "",  sdkerrors.Wrapf(types.ErrInvalidTokenPair, "Not a valid Token Pair: tokenA and tokenB cannot be the same")
	} else {
		return token1, token0, nil
	}
	

}

func (k Keeper) SortTokensDeposit(ctx sdk.Context, token0 string, token1 string, amounts0 []sdk.Dec, amounts1 []sdk.Dec) (string, string, []sdk.Dec, []sdk.Dec, error) {

	token0Hash := sha256.Sum256([]byte(token0))
	token1Hash := sha256.Sum256([]byte(token1))

	comparisonInt :=  bytes.Compare(token0Hash[:], token1Hash[:])

	if comparisonInt == -1 {
		return token0, token1, amounts0, amounts1, nil
	} else if comparisonInt == 0 {
		return "", "", nil, nil,  sdkerrors.Wrapf(types.ErrInvalidTokenPair, "Not a valid Token Pair: tokenA and tokenB cannot be the same")
	} else {
		return token1, token0, amounts1, amounts0, nil
	}

}
