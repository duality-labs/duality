package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Pool struct {
	PairId         string
	TickIndex      int64
	FeeIndex       uint64
	LowerTick0     *types.Tick
	UpperTick1     *types.Tick
	price1To0Lower sdk.Dec
	price0To1Upper sdk.Dec
}

func NewPool(
	pairId string,
	tickIndex int64,
	feeIndex uint64,
	fee int64,
	lowerTick0 *types.Tick,
	upperTick1 *types.Tick,
) *Pool {
	return &Pool{
		PairId:         pairId,
		TickIndex:      tickIndex,
		FeeIndex:       feeIndex,
		LowerTick0:     lowerTick0,
		UpperTick1:     upperTick1,
		price0To1Upper: *upperTick1.Price0To1,
		price1To0Lower: sdk.OneDec().Quo(*lowerTick0.Price0To1),
	}
}

func (p *Pool) GetLowerReserve0() sdk.Int {
	return p.LowerTick0.TickData.Reserve0[p.FeeIndex]
}

func (p *Pool) GetUpperReserve1() sdk.Int {
	return p.UpperTick1.TickData.Reserve1[p.FeeIndex]
}

func (p *Pool) Swap0To1(maxAmount0 sdk.Int) (inAmount0 sdk.Int, outAmount1 sdk.Int) {
	price1To0Upper := sdk.OneDec().Quo(p.price0To1Upper)
	reserves1 := &p.UpperTick1.TickData.Reserve1[p.FeeIndex]
	reserves0 := &p.LowerTick0.TickData.Reserve0[p.FeeIndex]
	maxAmount1 := maxAmount0.ToDec().Mul(p.price0To1Upper).TruncateInt()
	if reserves1.LT(maxAmount1) {
		outAmount1 = *reserves1
		inAmount0 = reserves1.ToDec().Mul(price1To0Upper).TruncateInt()
		*reserves0 = reserves0.Add(inAmount0)
		*reserves1 = sdk.ZeroInt()
	} else {
		outAmount1 = maxAmount0.ToDec().Mul(p.price0To1Upper).TruncateInt()
		*reserves0 = reserves0.Add(maxAmount0)
		*reserves1 = reserves1.Sub(outAmount1)
		inAmount0 = maxAmount0
	}
	return inAmount0, outAmount1
}

func (p *Pool) Swap1To0(maxAmount1 sdk.Int) (inAmount1 sdk.Int, outAmount0 sdk.Int) {
	price0To1Lower := sdk.OneDec().Quo(p.price1To0Lower)
	reserves1 := &p.UpperTick1.TickData.Reserve1[p.FeeIndex]
	reserves0 := &p.LowerTick0.TickData.Reserve0[p.FeeIndex]
	maxAmount0 := maxAmount1.ToDec().Mul(p.price1To0Lower).TruncateInt()
	if reserves0.LT(maxAmount0) {
		outAmount0 = *reserves0
		inAmount1 = reserves0.ToDec().Mul(price0To1Lower).TruncateInt()
		*reserves1 = reserves1.Add(inAmount1)
		*reserves0 = sdk.ZeroInt()
	} else {
		outAmount0 = maxAmount1.ToDec().Mul(p.price1To0Lower).TruncateInt()
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
func (p *Pool) Deposit(maxAmount0 sdk.Int, maxAmount1 sdk.Int, totalShares sdk.Int) (inAmount0 sdk.Int, inAmount1 sdk.Int, outShares sdk.Int) {
	lowerReserve0 := &p.LowerTick0.TickData.Reserve0[p.FeeIndex]
	upperReserve1 := &p.UpperTick1.TickData.Reserve1[p.FeeIndex]

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

	*lowerReserve0 = lowerReserve0.Add(inAmount0)
	*upperReserve1 = upperReserve1.Add(inAmount1)
	return inAmount0, inAmount1, outShares
}

func (p *Pool) MustCalcPrice1To0Center() sdk.Dec {
	// NOTE: We can safely call the error-less version of CalcPrice here because the pool object
	// has already been initialized with an upper and lower tick which satisfy a check for IsTickOutOfRange
	return MustCalcPrice1To0(p.TickIndex)
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

func (p *Pool) Withdraw(sharesToRemove sdk.Int, totalShares sdk.Int) (outAmount0 sdk.Int, outAmount1 sdk.Int, err error) {
	reserves0 := &p.LowerTick0.TickData.Reserve0[p.FeeIndex]
	reserves1 := &p.UpperTick1.TickData.Reserve1[p.FeeIndex]
	ownershipRatio := sharesToRemove.ToDec().Quo(totalShares.ToDec())
	outAmount1 = ownershipRatio.Mul(reserves1.ToDec()).TruncateInt()
	outAmount0 = ownershipRatio.Mul(reserves0.ToDec()).TruncateInt()
	*reserves0 = reserves0.Sub(outAmount0)
	*reserves1 = reserves1.Sub(outAmount1)
	return outAmount0, outAmount1, nil
}

func CalcShares(amount0 sdk.Int, amount1 sdk.Int, priceCenter1To0 sdk.Dec) sdk.Dec {
	amount0Dec := amount0.ToDec()
	amount1Dec := amount1.ToDec()
	return amount0Dec.Add(amount1Dec.Mul(priceCenter1To0))
}

func (p *Pool) Save(ctx context.Context, keeper Keeper) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	keeper.SetTick(sdkCtx, p.PairId, *p.LowerTick0)
	keeper.SetTick(sdkCtx, p.PairId, *p.UpperTick1)
}
