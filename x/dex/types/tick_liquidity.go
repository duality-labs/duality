package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// NOTE: These methods should be avoided if possible.
// Generally default to dealing with LimitOrderTranche or PoolReserves explicitly

func (t TickLiquidity) Price() sdk.Dec {
	switch liquidity := t.Liquidity.(type) {
	case *TickLiquidity_LimitOrderTranche:
		return liquidity.LimitOrderTranche.PriceTakerToMaker

	case *TickLiquidity_PoolReserves:
		return liquidity.PoolReserves.PriceTakerToMaker
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

func (t TickLiquidity) HasToken() bool {
	switch liquidity := t.Liquidity.(type) {
	case *TickLiquidity_LimitOrderTranche:
		return liquidity.LimitOrderTranche.HasTokenIn()

	case *TickLiquidity_PoolReserves:
		return liquidity.PoolReserves.HasToken()
	default:
		panic("Tick does not contain valid liqudityType")
	}
}
