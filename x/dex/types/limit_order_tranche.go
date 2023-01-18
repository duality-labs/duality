package types

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
