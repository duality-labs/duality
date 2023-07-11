package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type PoolLiquidity struct {
	TradePairID *TradePairID
	Pool        *Pool
}

func (pl *PoolLiquidity) Swap(
	maxAmountTakerDenomIn sdk.Int,
	maxAmountMakerDenomOut *sdk.Int,
) (inAmount, outAmount sdk.Int) {
	return pl.Pool.Swap(
		pl.TradePairID,
		maxAmountTakerDenomIn,
		maxAmountMakerDenomOut,
	)
}

func (pl *PoolLiquidity) Price() sdk.Dec {
	return pl.Pool.Price(pl.TradePairID)
}

func NewPoolLiquidity(tradePairID *TradePairID, pool *Pool) Liquidity {
	return &PoolLiquidity{
		TradePairID: tradePairID,
		Pool:        pool,
	}
}
