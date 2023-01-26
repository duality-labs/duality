package types

// NOTE: These methods should be avoided if possible.
// Generally default to dealing with LimitOrderTranche or PoolReserves explicitly

func (t TickLiquidity) TickIndex() int64 {
	switch liquidity := t.Liquidity.(type) {
	case *TickLiquidity_LimitOrderTranche:
		return liquidity.LimitOrderTranche.TickIndex

	case *TickLiquidity_PoolReserves:
		return liquidity.PoolReserves.TickIndex
	default:
		panic("Tick does not contain valid liqudityType")
	}
}

func (t TickLiquidity) HasToken() bool {
	switch liquidity := t.Liquidity.(type) {
	case *TickLiquidity_LimitOrderTranche:
		return liquidity.LimitOrderTranche.HasToken()

	case *TickLiquidity_PoolReserves:
		return liquidity.PoolReserves.HasToken()
	default:
		panic("Tick does not contain valid liqudityType")
	}
}
