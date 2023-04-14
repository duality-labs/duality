package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/incentives module sentinel errors.
var (
	ErrNotStakeOwner   = sdkerrors.Register(ModuleName, 1, "msg sender is not the owner of specified stake")
	ErrStakeupNotFound = sdkerrors.Register(ModuleName, 2, "stakeup not found")
	ErrGaugeNotActive  = sdkerrors.Register(ModuleName, 3, "cannot distribute from gauges when it is not active")
)
