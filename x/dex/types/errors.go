package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/dex module sentinel errors
var (
	ErrNoSpendableCoins     = sdkerrors.Register(ModuleName, 1100, "No Spendable Coins found: %s")
	ErrNotEnoughCoins       = sdkerrors.Register(ModuleName, 1101, "Not enough Spendable Coins found: %s")
	ErrInvalidTokenPair     = sdkerrors.Register(ModuleName, 1102, "Invalid Token Pair: (%s, %s)")
	ErrInvalidTokenListSize = sdkerrors.Register(ModuleName, 1103, "Invalid Array: Array Tokens0 size does not equal Array Tokens1")
	ErrNotEnoughShares      = sdkerrors.Register(ModuleName, 1104, "Not enough shares are owned by:  %s")
	ErrValidShareNotFound   = sdkerrors.Register(ModuleName, 1105, "Valid share not found")
	ErrValidTickNotFound    = sdkerrors.Register(ModuleName, 1106, "Valid tick not found")
)
