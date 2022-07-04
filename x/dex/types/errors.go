package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/dex module sentinel errors
var (
	ErrNoSpendableCoins = sdkerrors.Register(ModuleName, 1100, "No Spendable Coins found: %s")
	ErrNotEnoughCoins   = sdkerrors.Register(ModuleName, 1101, "Not enough Spendable Coins found: %s")
)
