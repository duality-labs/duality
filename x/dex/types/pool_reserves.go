package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (*PoolReserves) LiquidityType() string {
	return LiquidityTypeLP
}

func (p PoolReserves) HasToken() bool {
	return p.Reserves.GT(sdk.ZeroInt())
}

func (p PoolReserves) TickIndexVal() int64 {
	return p.TickIndex
}

func (p PoolReserves) ToLimitOrderTranche() *LimitOrderTranche {
	panic("Cannot convert PoolReserves to LimitOrderTranche")
}

func (p PoolReserves) ToPoolReserves() *PoolReserves {
	return &p
}
