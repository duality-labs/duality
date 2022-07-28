package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Calculates and returns the minimum element of two sdk.Dec (fixed point integer) values
func (k Keeper) Min(a, b sdk.Dec) sdk.Dec {
	if a.LT(b) {
		return a
	}
	return b
}
