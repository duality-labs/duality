package keeper

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) depositHelperAdd(pool *types.Pool, amount0, amount1 sdk.Dec, Fee, Price string) (sdk.Dec, sdk.Dec, sdk.Dec, error) {

	var trueAmounts0 sdk.Dec
	var trueAmounts1 sdk.Dec


	if pool.Reserve0.GT(sdk.ZeroDec()) {
		trueAmounts1 = k.min(amount1, (pool.Reserve1.Mul(amount0)).Quo(pool.Reserve0))
	}

	if pool.Reserve1.GT(sdk.ZeroDec()) {
		trueAmounts0 = k.min(amount0, (pool.Reserve0.Mul(amount1)).Quo(pool.Reserve1))
	}

	if trueAmounts0 == amount0 && trueAmounts1 != amount1 {
		trueAmounts1 = amount1.Add(((amount1.Sub(trueAmounts1)).Mul(pool.Fee)).Quo(sdk.NewDec(10000).Sub(pool.Fee)))
	} else if trueAmounts1 == amount1 && trueAmounts0 != amount0 {
		trueAmounts0 = amount0.Add(((amount0.Add(trueAmounts0)).Mul(pool.Fee)).Quo(sdk.NewDec(10000).Sub(pool.Fee)))
	}

	// ((TotalShares * (Amt0 + Amt1 * Price)) / Reserve0 + Reserve1 * Price
	SharesMinted := (pool.TotalShares.Mul(amount0.Add(amount1.Mul(pool.Price)))).Quo(pool.Reserve0.Add(pool.Reserve1.Mul(pool.Price)))

	return trueAmounts0, trueAmounts1, SharesMinted, nil

}

func (k Keeper) depositHelperSub(pool *types.Pool, amount0, amount1 sdk.Dec, Fee, Price string) (sdk.Dec, sdk.Dec, sdk.Dec, error) {

	var trueAmounts0 sdk.Dec
	var trueAmounts1 sdk.Dec


	if pool.Reserve0.GT(sdk.ZeroDec()) {
		trueAmounts1 = k.min(amount1, (pool.Reserve1.Mul(amount0)).Quo(pool.Reserve0))
	}

	if pool.Reserve1.GT(sdk.ZeroDec()) {
		trueAmounts0 = k.min(amount0, (pool.Reserve0.Mul(amount1)).Quo(pool.Reserve1))
	}

	if trueAmounts0 == amount0 && trueAmounts1 != amount1 {
		trueAmounts1 = amount1.Sub(((amount1.Sub(trueAmounts1)).Mul(pool.Fee)).Quo(sdk.NewDec(10000).Sub(pool.Fee)))
	} else if trueAmounts1 == amount1 && trueAmounts0 != amount0 {
		trueAmounts0 = amount0.Sub(((amount0.Add(trueAmounts0)).Mul(pool.Fee)).Quo(sdk.NewDec(10000).Sub(pool.Fee)))
	}

	// ((TotalShares * (Amt0 + Amt1 * Price)) / Reserve0 + Reserve1 * Price
	SharesMinted := (pool.TotalShares.Mul(trueAmounts0.Add(trueAmounts1.Mul(pool.Price)))).Quo(pool.Reserve0.Add(pool.Reserve1.Mul(pool.Price)))

	return trueAmounts0, trueAmounts1, SharesMinted, nil

}
