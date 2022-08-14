package keeper

import (
	"bytes"
	"crypto/sha256"

	//"strings"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) SortTokens(ctx sdk.Context, tokenA string, tokenB string) (string, string, error) {

	// Calculate sha256 Checksum for Token0, Token1
	token0Hash := sha256.Sum256([]byte(tokenA))
	token1Hash := sha256.Sum256([]byte(tokenB))

	//calculates an comparison integer for the two hashes
	comparisonInt := bytes.Compare(token0Hash[:], token1Hash[:])

	if comparisonInt == -1 {
		return tokenA, tokenB, nil
	} else if comparisonInt == 0 {
		return "", "", sdkerrors.Wrapf(types.ErrInvalidTokenPair, "Not a valid Token Pair: tokenA and tokenB cannot be the same")
	} else {
		return tokenB, tokenA, nil
	}

}

// func (k Keeper) SortTokens(ctx sdk.Context, token0 string, token1 string, price sdk.Dec) (string, string, sdk.Dec, error) {

// 	// Calculate sha256 Checksum for Token0, Token1
// 	token0Hash := sha256.Sum256([]byte(token0))
// 	token1Hash := sha256.Sum256([]byte(token1))

// 	//calculates an comparison integer for the two hashes
// 	comparisonInt := bytes.Compare(token0Hash[:], token1Hash[:])

// 	/* If comparisonInt == -1 : token0Hash < token1Hash
// 	   comparisonInt == 0 token0Hash == token1Hash (return an error)
// 	   comparisonInt == 1 token0Hash > token1Hash (switch elements)
// 	*/
// 	if comparisonInt == -1 {
// 		return token0, token1, price, nil
// 	} else if comparisonInt == 0 {
// 		return "", "", sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrInvalidTokenPair, "Not a valid Token Pair: tokenA and tokenB cannot be the same")
// 	} else {
// 		return token1, token0, sdk.OneDec().Quo(price), nil
// 	}

// }

// func (k Keeper) SortTokensDeposit(ctx sdk.Context, token0 string, token1 string, amounts0 []sdk.Dec, amounts1 []sdk.Dec, price sdk.Dec) (string, string, []sdk.Dec, []sdk.Dec, sdk.Dec, error) {

// 	// Calculate sha256 Checksum for Token0, Token1
// 	token0Hash := sha256.Sum256([]byte(token0))
// 	token1Hash := sha256.Sum256([]byte(token1))

// 	//calculates an comparison integer for the two hashes
// 	comparisonInt := bytes.Compare(token0Hash[:], token1Hash[:])

// 	/* If comparisonInt == -1 : token0Hash < token1Hash (returns parameters as given)
// 	   comparisonInt == 0 token0Hash == token1Hash (return an error)
// 	   comparisonInt == 1 token0Hash > token1Hash (switch elements and amount arrays)
// 	*/
// 	if comparisonInt == -1 {
// 		return token0, token1, amounts0, amounts1, price, nil
// 	} else if comparisonInt == 0 {
// 		return "", "", nil, nil, sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrInvalidTokenPair, "Not a valid Token Pair: tokenA and tokenB cannot be the same")
// 	} else {
// 		return token1, token0, amounts1, amounts0, sdk.OneDec().Quo(price), nil
// 	}

// }
