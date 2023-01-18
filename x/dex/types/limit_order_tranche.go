package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (t LimitOrderTranche) IsPlaceTranche() bool {
	return t.ReservesTokenIn.Equal(t.TotalTokenIn)
}

func NewFromFilledTranche(t FilledLimitOrderTranche) LimitOrderTranche {
	return LimitOrderTranche{
		TrancheIndex:     t.TrancheIndex,
		TickIndex:        t.TickIndex,
		TokenIn:          t.TokenIn,
		PairId:           t.PairId,
		TotalTokenOut:    t.TotalTokenOut,
		TotalTokenIn:     t.TotalTokenIn,
		ReservesTokenOut: t.ReservesTokenOut,
	}
}

func (t LimitOrderTranche) CreateFilledTranche() FilledLimitOrderTranche {
	return FilledLimitOrderTranche{
		TrancheIndex:     t.TrancheIndex,
		TickIndex:        t.TickIndex,
		TokenIn:          t.TokenIn,
		PairId:           t.PairId,
		TotalTokenIn:     t.TotalTokenIn,
		TotalTokenOut:    t.TotalTokenOut,
		ReservesTokenOut: t.ReservesTokenOut,
	}
}

func (LimitOrderTranche) LiquidityType() string {
	return LiquidityTypeLO
}

func (t LimitOrderTranche) HasToken() bool {
	return t.ReservesTokenIn.GT(sdk.ZeroInt())
}

func (t LimitOrderTranche) TickIndexVal() int64 {
	return t.TickIndex
}

func (t LimitOrderTranche) ToLimitOrderTranche() *LimitOrderTranche {
	return &t
}

func (LimitOrderTranche) ToPoolReserves() *PoolReserves {
	panic("Cannot convert LimitOrderTranche to PoolReserves")
}
