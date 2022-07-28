package keeper

import (
	//"fmt"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


/* This function implements an autoswap feature by default to ensure that the value deposited
	 in the tick matches the intended amount as closely as possible. If amount of tokens in the tick
	 are such that the deposit amounts would not match the intended amounts, then perform an internal
    swap to reach the intended amounts. Autoswap is subject to fees just like a normal swap
*/
func (k Keeper) DepositHelperAdd(pool *types.Pool, amount0, amount1 sdk.Dec) (sdk.Dec, sdk.Dec, sdk.Dec, error) {

	var trueAmounts0 sdk.Dec
	var trueAmounts1 sdk.Dec

	 // Check to see if input amount of Token 0 follows tick ratio
	if pool.Reserve0.GT(sdk.ZeroDec()) {
		trueAmounts1 = k.Min(amount1, (pool.Reserve1.Mul(amount0)).Quo(pool.Reserve0))
	}

	 // Check to see if input amount of Token 1 follows tick ratio
	if pool.Reserve1.GT(sdk.ZeroDec()) {
		trueAmounts0 = k.Min(amount0, (pool.Reserve0.Mul(amount1)).Quo(pool.Reserve1))
	}
	// autoswap if token 0 needs to reach target
	if trueAmounts0 == amount0 && trueAmounts1 != amount1 {
		trueAmounts1 = amount1.Add(((amount1.Sub(trueAmounts1)).Mul(pool.Fee)).Quo(sdk.NewDec(10000).Sub(pool.Fee)))
		
	// autoswap if token 1 needs to reach target
	} else if trueAmounts1 == amount1 && trueAmounts0 != amount0 {
		trueAmounts0 = amount0.Add(((amount0.Add(trueAmounts0)).Mul(pool.Fee)).Quo(sdk.NewDec(10000).Sub(pool.Fee)))
	}

	// ((TotalShares * (Amt0 + Amt1 * Price)) / Reserve0 + Reserve1 * Price
	SharesMinted := (pool.TotalShares.Mul(amount0.Add(amount1.Mul(pool.Price)))).Quo(pool.Reserve0.Add(pool.Reserve1.Mul(pool.Price)))

	return trueAmounts0, trueAmounts1, SharesMinted, nil

}

/* This function implements an autoswap feature by default to ensure that the value deposited
	 in the tick matches the intended amount as closely as possible. If amount of tokens in the tick
	 are such that the deposit amounts would not match the intended amounts, then perform an internal
    swap to reach the intended amounts. Autoswap is subject to fees just like a normal swap
*/

func (k Keeper) DepositHelperSub(pool *types.Pool, amount0, amount1 sdk.Dec) (sdk.Dec, sdk.Dec, sdk.Dec, error) {

	var trueAmounts0 sdk.Dec
	var trueAmounts1 sdk.Dec

	// Check to see if input amount of Token 0 follows tick ratio
	if pool.Reserve0.GT(sdk.ZeroDec()) {
		trueAmounts1 = k.Min(amount1, (pool.Reserve1.Mul(amount0)).Quo(pool.Reserve0))
	}

	 // Check to see if input amount of Token 1 follows tick ratio
	if pool.Reserve1.GT(sdk.ZeroDec()) {
		trueAmounts0 = k.Min(amount0, (pool.Reserve0.Mul(amount1)).Quo(pool.Reserve1))
	}

	// autoswap if token 0 needs to reach target
	if trueAmounts0 == amount0 && trueAmounts1 != amount1 {
		trueAmounts1 = amount1.Sub(((amount1.Sub(trueAmounts1)).Mul(pool.Fee)).Quo(sdk.NewDec(10000).Sub(pool.Fee)))

	// autoswap if token 1 needs to reach target
	} else if trueAmounts1 == amount1 && trueAmounts0 != amount0 {
		trueAmounts0 = amount0.Sub(((amount0.Add(trueAmounts0)).Mul(pool.Fee)).Quo(sdk.NewDec(10000).Sub(pool.Fee)))
	}

	// ((TotalShares * (Amt0 + Amt1 * Price)) / Reserve0 + Reserve1 * Price
	// As the tickexist, add the reserves and shares.
	SharesMinted := (pool.TotalShares.Mul(trueAmounts0.Add(trueAmounts1.Mul(pool.Price)))).Quo(pool.Reserve0.Add(pool.Reserve1.Mul(pool.Price)))

	return trueAmounts0, trueAmounts1, SharesMinted, nil

}
