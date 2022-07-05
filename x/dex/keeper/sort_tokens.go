package keeper

import (
	"sort"
	"strings"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) sortTokens(ctx sdk.Context, tokens0 []string, tokens1 []string) ([]string, []string, error) {

	if len(tokens0) == len(tokens1) {
		return nil, nil, sdkerrors.Wrapf(types.ErrInvalidTokenListSize, "Invalid Array: Tokens0 size: %s, Tokens1 size: %t ")
	}
	newTokens0 := make([]string, len(tokens0))
	newTokens1 := make([]string, len(tokens1))

	for i, s := range tokens0 {

		if strings.EqualFold(s, tokens1[i]) {
			return nil, nil, sdkerrors.Wrapf(types.ErrInvalidTokenPair, "invalid token pair (%s, %t)", s, tokens1[i])
		}

		pair := []string{s, tokens1[i]}
		sort.Strings(pair)

		newTokens0[i] = pair[0]
		newTokens1[i] = pair[1]

	}

	return newTokens0, newTokens1, nil
}
