package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
)

type PoolLiquidity struct {
	tradePairID *types.TradePairID
	pool        *Pool
}

func (pl *PoolLiquidity) Swap(maxAmountTakerDenomIn sdk.Int, maxAmountMakerDenomOut sdk.Int) (inAmount, outAmount sdk.Int) {
	if pl.tradePairID.IsTakerDenomToken0() {
		return pl.pool.Swap0To1(maxAmountTakerDenomIn, maxAmountMakerDenomOut)
	}

	return pl.pool.Swap1To0(maxAmountTakerDenomIn, maxAmountMakerDenomOut)
}

func (pl *PoolLiquidity) Price() sdk.Dec {
	if pl.tradePairID.IsTakerDenomToken0() {
		return pl.pool.UpperTick1.PriceTakerToMaker
	}

	return pl.pool.LowerTick0.PriceTakerToMaker
}

func NewPoolLiquidity(tradePairID *types.TradePairID, pool *Pool) Liquidity {
	return &PoolLiquidity{
		tradePairID: tradePairID,
		pool:        pool,
	}
}
