package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/utils"
	"github.com/duality-labs/duality/x/dex/types"
)

type Pool struct {
	TickIndex      int64
	FeeIndex       uint64
	LowerTick0     *types.PoolReserves
	UpperTick1     *types.PoolReserves
	Price1To0Lower sdk.Dec
	Price0To1Upper sdk.Dec
}

func NewPool(
	tickIndex int64,
	// feeIndex uint64,
	lowerTick0 *types.PoolReserves,
	upperTick1 *types.PoolReserves,
) Pool {
	// TODO: maybe store this somewhere so we don't have to recalculate
	price0To1 := utils.MustCalcPrice0To1(tickIndex)
	return Pool{
		TickIndex:      tickIndex,
		LowerTick0:     lowerTick0,
		UpperTick1:     upperTick1,
		Price0To1Upper: price0To1,
		Price1To0Lower: sdk.OneDec().Quo(price0To1),
	}
}

func (k Keeper) GetOrInitPool(ctx sdk.Context, pairId *types.PairId, tickIndex int64, feeTier types.FeeTier) (Pool, error) {
	fee := feeTier.Fee
	lowertick, err := k.GetOrInitPoolReserves(ctx, pairId, pairId.Token0, tickIndex-int64(fee), fee)
	if err != nil {
		return Pool{}, sdkerrors.Wrapf(err, "Error for lower tick")
	}

	upperTick, err := k.GetOrInitPoolReserves(ctx, pairId, pairId.Token1, tickIndex+int64(fee), fee)
	if err != nil {
		return Pool{}, sdkerrors.Wrapf(err, "Error for upper tick")
	}

	return NewPool(tickIndex, lowertick, upperTick), nil
}
func (p *Pool) GetLowerReserve0() sdk.Int {
	return p.LowerTick0.Reserves
}

func (p *Pool) GetUpperReserve1() sdk.Int {
	return p.UpperTick1.Reserves
}

func (p *Pool) Swap0To1(maxAmount0 sdk.Int) (inAmount0 sdk.Int, outAmount1 sdk.Int) {
	reserves1 := &p.UpperTick1.Reserves
	if maxAmount0.Equal(sdk.ZeroInt()) || reserves1.Equal(sdk.ZeroInt()) {
		return sdk.ZeroInt(), sdk.ZeroInt()
	}

	reserves0 := &p.LowerTick0.Reserves

	price1To0Upper := sdk.OneDec().Quo(p.Price0To1Upper)
	maxAmount1 := maxAmount0.ToDec().Mul(p.Price0To1Upper).TruncateInt()
	if reserves1.LTE(maxAmount1) {
		outAmount1 = *reserves1
		inAmount0 = reserves1.ToDec().Mul(price1To0Upper).TruncateInt()
		*reserves0 = reserves0.Add(inAmount0)
		*reserves1 = sdk.ZeroInt()
	} else {
		outAmount1 = maxAmount0.ToDec().Mul(p.Price0To1Upper).TruncateInt()
		*reserves0 = reserves0.Add(maxAmount0)
		*reserves1 = reserves1.Sub(outAmount1)
		inAmount0 = maxAmount0
	}
	return inAmount0, outAmount1
}

func (p *Pool) Swap1To0(maxAmount1 sdk.Int) (inAmount1 sdk.Int, outAmount0 sdk.Int) {
	reserves0 := &p.LowerTick0.Reserves
	if maxAmount1.Equal(sdk.ZeroInt()) || reserves0.Equal(sdk.ZeroInt()) {
		return sdk.ZeroInt(), sdk.ZeroInt()
	}

	reserves1 := &p.UpperTick1.Reserves

	price0To1Lower := sdk.OneDec().Quo(p.Price1To0Lower)
	maxAmount0 := maxAmount1.ToDec().Mul(p.Price1To0Lower).TruncateInt()
	if reserves0.LTE(maxAmount0) {
		outAmount0 = *reserves0
		inAmount1 = reserves0.ToDec().Mul(price0To1Lower).TruncateInt()
		*reserves1 = reserves1.Add(inAmount1)
		*reserves0 = sdk.ZeroInt()
	} else {
		outAmount0 = maxAmount1.ToDec().Mul(p.Price1To0Lower).TruncateInt()
		*reserves1 = reserves1.Add(maxAmount1)
		*reserves0 = reserves0.Sub(outAmount0)
		inAmount1 = maxAmount1
	}
	return inAmount1, outAmount0
}

// Balance trueAmount1 to the pool ratio
func CalcGreatestMatchingRatio(
	targetAmount0 sdk.Int,
	targetAmount1 sdk.Int,
	amount0 sdk.Int,
	amount1 sdk.Int,
) (resultAmount0 sdk.Int, resultAmount1 sdk.Int) {
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

// Mutates the Pool object and returns relevant change variables. Deposit is not commited until
// pool.save() is called or the underlying ticks are saved; this method does not use any keeper methods.
func (p *Pool) Deposit(maxAmount0 sdk.Int, maxAmount1 sdk.Int, totalShares sdk.Int, autoswap bool) (inAmount0 sdk.Int, inAmount1 sdk.Int, outShares sdk.Int) {

	lowerReserve0 := &p.LowerTick0.Reserves
	upperReserve1 := &p.UpperTick1.Reserves

	inAmount0, inAmount1 = CalcGreatestMatchingRatio(
		*lowerReserve0,
		*upperReserve1,
		maxAmount0,
		maxAmount1,
	)

	if inAmount0.Equal(sdk.ZeroInt()) && inAmount1.Equal(sdk.ZeroInt()) {
		return sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt()
	}

	outShares = p.CalcSharesMinted(
		*lowerReserve0,
		*upperReserve1,
		totalShares,
		inAmount0,
		inAmount1,
	)

	if autoswap {
		residualAmount0 := maxAmount0.Sub(inAmount0)
		residualAmount1 := maxAmount1.Sub(inAmount1)

		// NOTE: Currently not doing anything with the error, but added error handling to all of the new functions for autoswap.
		// Open to changing it however.
		residualShares, _ := p.CalcResidualSharesMinted(
			*lowerReserve0,
			*upperReserve1,
			totalShares,
			residualAmount0,
			residualAmount1,
		)

		outShares = outShares.Add(residualShares)

		inAmount0 = maxAmount0
		inAmount1 = maxAmount1
	}

	*lowerReserve0 = lowerReserve0.Add(inAmount0)
	*upperReserve1 = upperReserve1.Add(inAmount1)
	return inAmount0, inAmount1, outShares
}

func (p *Pool) MustCalcPrice1To0Center() sdk.Dec {
	// NOTE: We can safely call the error-less version of CalcPrice here because the pool object
	// has already been initialized with an upper and lower tick which satisfy a check for IsTickOutOfRange
	return utils.MustCalcPrice1To0(p.TickIndex)
}
func (p *Pool) CalcSharesMinted(
	reserve0 sdk.Int,
	reserve1 sdk.Int,
	totalShares sdk.Int,
	amount0 sdk.Int,
	amount1 sdk.Int,
) (sharesMinted sdk.Int) {
	price1To0Center := p.MustCalcPrice1To0Center()
	valueMintedToken0 := CalcShares(amount0, amount1, price1To0Center)
	valueExistingToken0 := CalcShares(reserve0, reserve1, price1To0Center)
	if valueExistingToken0.GT(sdk.ZeroDec()) {
		sharesMinted = valueMintedToken0.Quo(valueExistingToken0).Mul(totalShares.ToDec()).TruncateInt()
	} else {
		sharesMinted = valueMintedToken0.TruncateInt()
	}
	return sharesMinted
}

func (p *Pool) CalcResidualSharesMinted(
	reserve0 sdk.Int,
	reserve1 sdk.Int,
	totalShares sdk.Int,
	residualAmount0 sdk.Int,
	residualAmount1 sdk.Int,
) (sharesMinted sdk.Int, err error) {
	fee := CalcFee(p.UpperTick1.TickIndex, p.LowerTick0.TickIndex)
	price1To0Center := p.MustCalcPrice1To0Center()
	valueMintedToken0, err := CalcResidualValue(residualAmount0, residualAmount1, p.Price1To0Lower, fee)
	if err != nil {
		return sdk.ZeroInt(), err
	}
	valueExistingToken0 := CalcShares(reserve0, reserve1, price1To0Center)
	if valueExistingToken0.GT(sdk.ZeroDec()) {
		sharesMinted = valueMintedToken0.Quo(valueExistingToken0).Mul(totalShares.ToDec()).TruncateInt()
	} else {
		sharesMinted = valueMintedToken0.TruncateInt()
	}
	return sharesMinted, nil
}

func (p *Pool) Withdraw(sharesToRemove sdk.Int, totalShares sdk.Int) (outAmount0 sdk.Int, outAmount1 sdk.Int) {
	reserves0 := &p.LowerTick0.Reserves
	reserves1 := &p.UpperTick1.Reserves
	ownershipRatio := sharesToRemove.ToDec().Quo(totalShares.ToDec())
	outAmount1 = ownershipRatio.Mul(reserves1.ToDec()).TruncateInt()
	outAmount0 = ownershipRatio.Mul(reserves0.ToDec()).TruncateInt()
	*reserves0 = reserves0.Sub(outAmount0)
	*reserves1 = reserves1.Sub(outAmount1)
	return outAmount0, outAmount1
}

func CalcShares(amount0 sdk.Int, amount1 sdk.Int, priceCenter1To0 sdk.Dec) sdk.Dec {
	amount0Dec := amount0.ToDec()
	amount1Dec := amount1.ToDec()
	return amount0Dec.Add(amount1Dec.Mul(priceCenter1To0))
}
func CalcResidualValue(amount0 sdk.Int, amount1 sdk.Int, priceLower1To0 sdk.Dec, fee int64) (sdk.Dec, error) {
	amount0Dec := amount0.ToDec()
	amount1Dec := amount1.ToDec()
	// ResidualValue = Amount0 * (Price1to0Center / Price1to0Upper) + Amount1 * Price1to0Lower
	amount0Discount, err := utils.CalcPrice0To1(-fee)
	if err != nil {
		return sdk.ZeroDec(), err
	}
	return (amount0Dec.Mul(amount0Discount)).Add(amount1Dec.Mul(priceLower1To0)), nil
}

func CalcFee(upperTickIndex int64, lowerTickIndex int64) int64 {
	return (upperTickIndex - lowerTickIndex) / 2
}

func (p *Pool) Save(sdkCtx sdk.Context, keeper Keeper) {
	if p.LowerTick0.HasToken() {
		keeper.SetPoolReserves(sdkCtx, *p.LowerTick0)
	} else {
		keeper.RemovePoolReserves(sdkCtx, *p.LowerTick0)
	}

	if p.UpperTick1.HasToken() {
		keeper.SetPoolReserves(sdkCtx, *p.UpperTick1)
	} else {
		keeper.RemovePoolReserves(sdkCtx, *p.UpperTick1)
	}
}
