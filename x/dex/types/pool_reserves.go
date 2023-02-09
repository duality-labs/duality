package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (p PoolReserves) HasToken() bool {
	return p.Reserves.GT(sdk.ZeroInt())
}
