package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/dex/utils"
)

type Pool struct {
	CenterTickIndex int64
	Fee             uint64
	LowerTick0      *types.PoolReserves
	UpperTick1      *types.PoolReserves
	Price1To0Lower  *types.Price
	Price0To1Upper  *types.Price
}

// TODO: Accept a PairID here so that GetDepositDenom() can reference it directly
func NewPool(
	centerTickIndex int64,
	lowerTick0 *types.PoolReserves,
	upperTick1 *types.PoolReserves,
) Pool {
	// TODO: maybe store this somewhere so we don't have to recalculate
	price0To1Upper := types.MustNewPrice(upperTick1.TickIndex)
	price1To0Lower := types.MustNewPrice(-1 * lowerTick0.TickIndex)

	return Pool{
		CenterTickIndex: centerTickIndex,
		Fee:             lowerTick0.Fee,
		LowerTick0:      lowerTick0,
		UpperTick1:      upperTick1,
		Price0To1Upper:  price0To1Upper,
		Price1To0Lower:  price1To0Lower,
	}
}

func (k Keeper) GetOrInitPool(ctx sdk.Context, pairID *types.PairID, centerTickIndex int64, fee uint64) (Pool, error) {
	feeUint := utils.MustSafeUint64(fee)
	lowertick, err := k.GetOrInitPoolReserves(ctx, pairID, pairID.Token0, centerTickIndex-feeUint, fee)
	if err != nil {
		return Pool{}, sdkerrors.Wrapf(err, "Error for lower tick")
	}

	upperTick, err := k.GetOrInitPoolReserves(ctx, pairID, pairID.Token1, centerTickIndex+feeUint, fee)
	if err != nil {
		return Pool{}, sdkerrors.Wrapf(err, "Error for upper tick")
	}

	return NewPool(centerTickIndex, lowertick, upperTick), nil
}

func (p *Pool) GetLowerReserve0() sdk.Int {
	return p.LowerTick0.Reserves
}

func (p *Pool) GetUpperReserve1() sdk.Int {
	return p.UpperTick1.Reserves
}

func (p *Pool) Swap0To1(maxAmount0 sdk.Int, maxAmountOut1 *sdk.Int) (inAmount0, outAmount1 sdk.Int) {
	reserves1 := &p.UpperTick1.Reserves
	if maxAmount0.Equal(sdk.ZeroInt()) || reserves1.Equal(sdk.ZeroInt()) {
		return sdk.ZeroInt(), sdk.ZeroInt()
	}

	reserves0 := &p.LowerTick0.Reserves

	maxAmount1 := p.Price0To1Upper.MulInt(maxAmount0).TruncateInt()
	possibleOutAmounts := []sdk.Int{*reserves1, maxAmount1}
	if maxAmountOut1 != nil {
		possibleOutAmounts = append(possibleOutAmounts, *maxAmountOut1)
	}

	// outAmount will be the smallest value of:
	// a.) The available reserve1,
	// b.) The most the user could get out given maxAmount0 (maxAmount1),
	// c.) The maximum amount the user wants out (amountOutCap1)
	outAmount1 = utils.MinIntArr(possibleOutAmounts)

	// we can skip price calc if we are using maxAmount1, since we already know it
	if outAmount1 == maxAmount1 {
		inAmount0 = maxAmount0
	} else {
		inAmount0 = p.Price0To1Upper.Inv().MulInt(outAmount1).Ceil().TruncateInt()
	}

	*reserves0 = reserves0.Add(inAmount0)
	*reserves1 = reserves1.Sub(outAmount1)

	return inAmount0, outAmount1
}

func (p *Pool) Swap1To0(maxAmount1 sdk.Int, maxAmountOut0 *sdk.Int) (inAmount1, outAmount0 sdk.Int) {
	reserves0 := &p.LowerTick0.Reserves
	if maxAmount1.Equal(sdk.ZeroInt()) || reserves0.Equal(sdk.ZeroInt()) {
		return sdk.ZeroInt(), sdk.ZeroInt()
	}

	reserves1 := &p.UpperTick1.Reserves

	maxAmount0 := p.Price1To0Lower.MulInt(maxAmount1).TruncateInt()
	possibleOutAmounts := []sdk.Int{*reserves0, maxAmount0}
	if maxAmountOut0 != nil {
		possibleOutAmounts = append(possibleOutAmounts, *maxAmountOut0)
	}

	// outAmount will be the smallest value of:
	// a.) The available reserve1,
	// b.) The most the user could get out given maxAmount1 (maxAmount0),
	// c.) The maximum amount the user wants out (amountOutCap0)
	outAmount0 = utils.MinIntArr(possibleOutAmounts)

	// we can skip price calc if we are using maxAmount1, since we already know it
	if outAmount0 == maxAmount1 {
		inAmount1 = maxAmount0
	} else {
		inAmount1 = p.Price1To0Lower.Inv().MulInt(outAmount0).Ceil().TruncateInt()
	}

	*reserves1 = reserves1.Add(inAmount1)
	*reserves0 = reserves0.Sub(outAmount0)

	return inAmount1, outAmount0
}

// Balance trueAmount1 to the pool ratio
func CalcGreatestMatchingRatio(
	targetAmount0 sdk.Int,
	targetAmount1 sdk.Int,
	amount0 sdk.Int,
	amount1 sdk.Int,
) (resultAmount0, resultAmount1 sdk.Int) {
	targetAmount0Dec := targetAmount0.ToDec()
	targetAmount1Dec := targetAmount1.ToDec()

	// See spec: https://www.notion.so/dualityxyz/Autoswap-Spec-e856fa7b2438403c95147010d479b98c
	if targetAmount1.GT(sdk.ZeroInt()) {
		resultAmount0 = sdk.MinInt(
			amount0,
			amount1.ToDec().Mul(targetAmount0Dec).Quo(targetAmount1Dec).TruncateInt())
	} else {
		resultAmount0 = amount0
	}

	if targetAmount0.GT(sdk.ZeroInt()) {
		resultAmount1 = sdk.MinInt(
			amount1,
			amount0.ToDec().Mul(targetAmount1Dec).Quo(targetAmount0Dec).TruncateInt())
	} else {
		resultAmount1 = amount1
	}

	return resultAmount0, resultAmount1
}

// Mutates the Pool object and returns relevant change variables. Deposit is not committed until
// pool.save() is called or the underlying ticks are saved; this method does not use any keeper methods.
func (p *Pool) Deposit(
	maxAmount0,
	maxAmount1,
	existingShares sdk.Int,
	autoswap bool,
) (inAmount0, inAmount1 sdk.Int, outShares sdk.Coin) {
	lowerReserve0 := &p.LowerTick0.Reserves
	upperReserve1 := &p.UpperTick1.Reserves

	inAmount0, inAmount1 = CalcGreatestMatchingRatio(
		*lowerReserve0,
		*upperReserve1,
		maxAmount0,
		maxAmount1,
	)

	if inAmount0.Equal(sdk.ZeroInt()) && inAmount1.Equal(sdk.ZeroInt()) {
		return sdk.ZeroInt(), sdk.ZeroInt(), sdk.Coin{Denom: p.GetDepositDenom()}
	}

	outShares = p.CalcSharesMinted(inAmount0, inAmount1, existingShares)

	if autoswap {
		residualAmount0 := maxAmount0.Sub(inAmount0)
		residualAmount1 := maxAmount1.Sub(inAmount1)

		// NOTE: Currently not doing anything with the error,
		// but added error handling to all of the new functions for autoswap.
		// Open to changing it however.
		residualShares, _ := p.CalcResidualSharesMinted(residualAmount0, residualAmount1)
		// TODO: Fix

		outShares = outShares.Add(residualShares)

		inAmount0 = maxAmount0
		inAmount1 = maxAmount1
	}

	*lowerReserve0 = lowerReserve0.Add(inAmount0)
	*upperReserve1 = upperReserve1.Add(inAmount1)

	return inAmount0, inAmount1, outShares
}

func (p *Pool) GetDepositDenom() string {
	return types.NewDepositDenom(p.LowerTick0.PairID, p.CenterTickIndex, p.Fee).String()
}

func (p *Pool) MustCalcPrice1To0Center() *types.Price {
	// NOTE: We can safely call the error-less version of CalcPrice here because the pool object
	// has already been initialized with an upper and lower tick which satisfy a check for IsTickOutOfRange
	return types.MustNewPrice(-1 * p.CenterTickIndex)
}

func (p *Pool) CalcSharesMinted(
	amount0 sdk.Int,
	amount1 sdk.Int,
	existingShares sdk.Int,
) (sharesMinted sdk.Coin) {
	price1To0Center := p.MustCalcPrice1To0Center()
	valueMintedToken0 := CalcAmountAsToken0(amount0, amount1, *price1To0Center)

	valueExistingToken0 := CalcAmountAsToken0(p.LowerTick0.Reserves, p.UpperTick1.Reserves, *price1To0Center)
	var sharesMintedAmount sdk.Int
	if valueExistingToken0.GT(sdk.ZeroDec()) {
		sharesMintedAmount = valueMintedToken0.MulInt(existingShares).Quo(valueExistingToken0).TruncateInt()
	} else {
		sharesMintedAmount = valueMintedToken0.TruncateInt()
	}

	return sdk.Coin{Denom: p.GetDepositDenom(), Amount: sharesMintedAmount}
}

func (p *Pool) CalcResidualSharesMinted(
	residualAmount0 sdk.Int,
	residualAmount1 sdk.Int,
) (sharesMinted sdk.Coin, err error) {
	fee := CalcFee(p.UpperTick1.TickIndex, p.LowerTick0.TickIndex)
	valueMintedToken0, err := CalcResidualValue(residualAmount0, residualAmount1, p.Price1To0Lower, fee)
	if err != nil {
		return sdk.Coin{Denom: p.GetDepositDenom()}, err
	}

	return sdk.Coin{Denom: p.GetDepositDenom(), Amount: valueMintedToken0.TruncateInt()}, nil
}

func (p *Pool) RedeemValue(sharesToRemove, totalShares sdk.Int) (outAmount0, outAmount1 sdk.Int) {
	reserves0 := &p.LowerTick0.Reserves
	reserves1 := &p.UpperTick1.Reserves
	// outAmount1 = ownershipRatio * reserves1
	//            = (sharesToRemove / totalShares) * reserves1
	//            = (reserves1 * sharesToRemove ) / totalShares
	outAmount1 = reserves1.Mul(sharesToRemove).ToDec().QuoInt(totalShares).TruncateInt()
	// outAmount0 = ownershipRatio * reserves1
	//            = (sharesToRemove / totalShares) * reserves1
	//            = (reserves1 * sharesToRemove ) / totalShares
	outAmount0 = reserves0.Mul(sharesToRemove).ToDec().QuoInt(totalShares).TruncateInt()

	return outAmount0, outAmount1
}

func (p *Pool) Withdraw(sharesToRemove, totalShares sdk.Int) (outAmount0, outAmount1 sdk.Int) {
	reserves0 := &p.LowerTick0.Reserves
	reserves1 := &p.UpperTick1.Reserves
	outAmount0, outAmount1 = p.RedeemValue(sharesToRemove, totalShares)
	*reserves0 = reserves0.Sub(outAmount0)
	*reserves1 = reserves1.Sub(outAmount1)

	return outAmount0, outAmount1
}

func CalcResidualValue(amount0, amount1 sdk.Int, priceLower1To0 *types.Price, fee int64) (sdk.Dec, error) {
	// ResidualValue = Amount0 * (Price1to0Center / Price1to0Upper) + Amount1 * Price1to0Lower
	amount0Discount, err := types.NewPrice(-fee)
	if err != nil {
		return sdk.ZeroDec(), err
	}

	return amount0Discount.MulInt(amount0).Add(priceLower1To0.MulInt(amount1)), nil
}

func CalcFee(upperTickIndex, lowerTickIndex int64) int64 {
	return (upperTickIndex - lowerTickIndex) / 2
}

func (k Keeper) SavePool(ctx sdk.Context, pool Pool) {
	if pool.LowerTick0.HasToken() {
		k.SetPoolReserves(ctx, *pool.LowerTick0)
	} else {
		k.RemovePoolReserves(ctx, *pool.LowerTick0)
	}
	if pool.UpperTick1.HasToken() {
		k.SetPoolReserves(ctx, *pool.UpperTick1)
	} else {
		k.RemovePoolReserves(ctx, *pool.UpperTick1)
	}

	// TODO: this will create a bit of extra noise since not every Save is updating both ticks
	// This should be solved upstream by better tracking of dirty ticks
	ctx.EventManager().EmitEvent(types.CreateTickUpdatePoolReserves(*pool.LowerTick0))
	ctx.EventManager().EmitEvent(types.CreateTickUpdatePoolReserves(*pool.UpperTick1))
}
