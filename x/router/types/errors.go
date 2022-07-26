package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/router module sentinel errors
var (
	ErrValidPathNotFound    = sdkerrors.Register(ModuleName, 1100, "Valid Path not found")
	ErrNotEnoughCoins       = sdkerrors.Register(ModuleName, 1101, "Not enough Spendable Coins found: %s")
)
