package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrInvalidSwapMetadata = sdkerrors.Register(ModuleName, 2, "invalid swap metadata")
	ErrSwapFailed          = sdkerrors.Register(ModuleName, 3, "swap failed")
)
