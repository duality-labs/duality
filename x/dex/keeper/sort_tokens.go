package keeper

import (
	//"sort"
	"strconv"
	//"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	//sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) SortTokens(ctx sdk.Context, token0 string, token1 string) (string, string, error) {

	token0Int, err := strconv.ParseInt(token0, 10, 64)

	if err != nil {
		return "", "", err
	}

	token1Int, err := strconv.ParseInt(token1, 10, 64)

	if err != nil {
		return "", "", err
	}

	if token0Int < token1Int {
		return token0, token1, nil
	}
	return token1, token0, nil

}

func (k Keeper) SortTokensDeposit(ctx sdk.Context, token0 string, token1 string, amounts0 []sdk.Dec, amounts1 []sdk.Dec) (string, string, []sdk.Dec, []sdk.Dec, error) {

	token0Int, err := strconv.ParseInt(token0, 10, 64)

	if err != nil {
		return "", "", nil, nil, err
	}

	token1Int, err := strconv.ParseInt(token1, 10, 64)

	if err != nil {
		return "", "", nil, nil, err
	}

	if token0Int < token1Int {
		return token0, token1, amounts0, amounts1, nil
	}
	return token1, token0, amounts1, amounts0, nil

}
