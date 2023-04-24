package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
)

type PoolLiquidity struct {
	pool   *Pool
	is0To1 bool
}

func (pl *PoolLiquidity) Swap(maxAmountIn sdk.Int, maxAmountOut sdk.Int) (inAmount, outAmount sdk.Int) {
	if pl.is0To1 {
		return pl.pool.Swap0To1(maxAmountIn, maxAmountOut)
	}

	return pl.pool.Swap1To0(maxAmountIn, maxAmountOut)
}

func (pl *PoolLiquidity) Price() *types.Price {
	if pl.is0To1 {
		return pl.pool.Price0To1Upper
	}

	return pl.pool.Price1To0Lower
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
