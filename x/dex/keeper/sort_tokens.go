package keeper

import (
	"sort"
	"strings"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)
func (k Keeper) sortTokens(ctx sdk.Context, tokens0 []string, tokens1 []string) ([]string, []string, error) {

	if len(tokens0) != len(tokens1) {
		return nil, nil, sdkerrors.Wrapf(types.ErrInvalidTokenListSize, "Invalid Array: Array Tokens0 size does not equal Array Tokens1")
	}
	newTokens0 := make([]string, len(tokens0))
	newTokens1 := make([]string, len(tokens1))


	for i, s := range tokens0 {

		if strings.EqualFold(s, tokens1[i]) {
			return nil, nil, sdkerrors.Wrapf(types.ErrInvalidTokenPair, "invalid token pair (%s, %s)", s, tokens1[i])
		}

		pair := []string{s, tokens1[i]}
		sort.Strings(pair)

		newTokens0[i] = pair[0]
		newTokens1[i] = pair[1]
	}

	return newTokens0, newTokens1, nil
}

func (k Keeper) sortTokensDeposit(ctx sdk.Context, tokens0 []string, tokens1 []string, amounts0 []sdk.Dec, amounts1 []sdk.Dec ) ([]string, []string, []sdk.Dec, []sdk.Dec, error) {

	if len(tokens0) != len(tokens1) {
		return nil, nil,nil, nil, sdkerrors.Wrapf(types.ErrInvalidTokenListSize, "Invalid Array: Array Tokens0 size does not equal Array Tokens1")
	}
	newTokens0 := make([]string, len(tokens0))
	newTokens1 := make([]string, len(tokens1))
	newAmounts0 := make([]sdk.Dec, len(amounts0))
	newAmounts1 := make([]sdk.Dec, len(amounts1))

	for i, s := range tokens0 {

		if strings.EqualFold(s, tokens1[i]) {
			return nil, nil, nil, nil, sdkerrors.Wrapf(types.ErrInvalidTokenPair, "invalid token pair (%s, %s)", s, tokens1[i])
		}

		pair := []string{s, tokens1[i]}
		sort.Strings(pair)

		newTokens0[i] = pair[0]
		newTokens1[i] = pair[1]

		if s == pair[0] {
			newAmounts0[i] = amounts0[i]
			newAmounts1[i] = amounts1[i]
		} else {
			newAmounts0[i] = amounts1[i]
			newAmounts1[i] = amounts0[i]
		}

	}

	return newTokens0, newTokens1, newAmounts0, newAmounts1, nil
}
