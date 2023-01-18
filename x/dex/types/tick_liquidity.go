package types

type TickLiquidityI interface {
	LiquidityType() string
	HasToken() bool
	TickIndexVal() int64
	ToLimitOrderTranche() *LimitOrderTranche
	ToPoolReserves() *PoolReserves
}

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
		return LiquidityTypeLO

	case *TickLiquidity_PoolReserves:
		return LiquidityTypeLP
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
