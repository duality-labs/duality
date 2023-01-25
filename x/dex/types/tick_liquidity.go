package types

// NOTE: These type methods should be avoided if possible.
// Generally default to dealing with LimitOrderTranche or PoolReserves explicityly
func (t TickLiquidity) TokenIn() string {
	switch liquidity := t.Liquidity.(type) {
	case *TickLiquidity_LimitOrderTranche:
		return liquidity.LimitOrderTranche.TokenIn

	case *TickLiquidity_PoolReserves:
		return liquidity.PoolReserves.TokenIn
	default:
		panic("Tick does not contain valid liqudityType")
	}
}

func (t TickLiquidity) PairId() *PairId {
	switch liquidity := t.Liquidity.(type) {
	case *TickLiquidity_LimitOrderTranche:
		return liquidity.LimitOrderTranche.PairId

	case *TickLiquidity_PoolReserves:
		return liquidity.PoolReserves.PairId
	default:
		panic("Tick does not contain valid liqudityType")
	}
}

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

func (t TickLiquidity) LiquidityType() string {
	switch t.Liquidity.(type) {
	case *TickLiquidity_LimitOrderTranche:
		return LiquidityTypeLimitOrder

	case *TickLiquidity_PoolReserves:
		return LiquidityTypePoolReserves
	default:
		panic("Tick does not contain valid liqudityType")
	}
}

func (t TickLiquidity) LiquidityIndex() uint64 {
	switch liquidity := t.Liquidity.(type) {
	case *TickLiquidity_LimitOrderTranche:
		return liquidity.LimitOrderTranche.TrancheIndex

	case *TickLiquidity_PoolReserves:
		return liquidity.PoolReserves.Fee
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
