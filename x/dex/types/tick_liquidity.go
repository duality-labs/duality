package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (t TickLiquidity) HasToken() bool {
	return t.HasLPReserves() || t.HasActiveLimitOrders()
}

func (t TickLiquidity) HasLPReserves() bool {
	return t.LPReserve != nil && t.LPReserve.GT(sdk.ZeroInt())
}

func (t TickLiquidity) HasActiveLimitOrders() bool {
	return t.LimitOrderTranche != nil && t.LimitOrderTranche.ReservesTokenIn.GT(sdk.ZeroInt())
}
