package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type PoolLiquidity struct {
	pool   *Pool
	is0To1 bool
}

func (pl *PoolLiquidity) Swap(maxAmount sdk.Int) (inAmount sdk.Int, outAmount sdk.Int) {
	if pl.is0To1 {
		return pl.pool.Swap0To1(maxAmount)
	} else {
		return pl.pool.Swap1To0(maxAmount)
	}
}

func (pl *PoolLiquidity) Save(sdkCtx sdk.Context, keeper Keeper) {
	pl.pool.Save(sdkCtx, keeper)
}

func (pl *PoolLiquidity) Price() sdk.Dec {
	if pl.is0To1 {
		return pl.pool.Price0To1Upper
	} else {
		return pl.pool.Price1To0Lower
	}
}

func NewLiquidityFromPool0To1(pool *Pool) Liquidity {
	return &PoolLiquidity{
		pool:   pool,
		is0To1: true,
	}
}

func NewLiquidityFromPool1To0(pool *Pool) Liquidity {
	return &PoolLiquidity{
		pool:   pool,
		is0To1: false,
	}
}
