package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Pool struct {
	PairId     string
	TickIndex  int64
	FeeIndex   uint64
	LowerTick0 *types.Tick
	UpperTick1 *types.Tick
}

func NewPool(
	pairId string,
	tickIndex int64,
	feeIndex uint64,
	lowerTick0 *types.Tick,
	upperTick1 *types.Tick,
) Pool {
	return Pool{
		PairId:     pairId,
		TickIndex:  tickIndex,
		FeeIndex:   feeIndex,
		LowerTick0: lowerTick0,
		UpperTick1: upperTick1,
	}
}

func (p *Pool) GetLowerReserve0() sdk.Int {
	return p.LowerTick0.TickData.Reserve0AndShares[p.FeeIndex].Reserve0
}

func (p *Pool) GetUpperReserve1() sdk.Int {
	return p.UpperTick1.TickData.Reserve1[p.FeeIndex]
}

func (p *Pool) GetTotalShares() sdk.Int {
	return p.LowerTick0.TickData.Reserve0AndShares[p.FeeIndex].TotalShares
}

func (p *Pool) Swap0To1(maxAmount0Dec sdk.Dec) (inAmount0 sdk.Dec, outAmount1 sdk.Int) {
	price0To1 := CalcPrice0To1(p.UpperTick1.TickIndex)
	price1To0 := sdk.OneDec().Quo(price0To1)
	reserves1 := &p.UpperTick1.TickData.Reserve1[p.FeeIndex]
	reserves0 := &p.LowerTick0.TickData.Reserve0AndShares[p.FeeIndex].Reserve0
	maxAmount1 := maxAmount0Dec.Mul(price0To1).TruncateInt()
	if reserves1.LT(maxAmount1) {
		outAmount1 = *reserves1
		inAmount0 = reserves1.ToDec().Mul(price1To0)
		*reserves0 = reserves0.Add(inAmount0.TruncateInt())
		*reserves1 = sdk.ZeroInt()
	} else {
		outAmount1Dec := maxAmount0Dec.Mul(price0To1)
		outAmount1 = outAmount1Dec.TruncateInt()
		*reserves0 = reserves0.Add(maxAmount0Dec.TruncateInt())
		*reserves1 = reserves1.Sub(outAmount1)
		inAmount0 = maxAmount0Dec
	}
	return inAmount0, outAmount1
}

func (p *Pool) Swap1To0(maxAmount1Dec sdk.Dec) (inAmount1 sdk.Dec, outAmount0 sdk.Int) {
	price1To0 := CalcPrice1To0(p.LowerTick0.TickIndex)
	price0To1 := sdk.OneDec().Quo(price1To0)
	reserves1 := &p.UpperTick1.TickData.Reserve1[p.FeeIndex]
	reserves0 := &p.LowerTick0.TickData.Reserve0AndShares[p.FeeIndex].Reserve0
	maxAmount0 := maxAmount1Dec.Mul(price1To0).TruncateInt()
	if reserves0.LT(maxAmount0) {
		outAmount0 = *reserves0
		inAmount1 = reserves0.ToDec().Mul(price0To1)
		*reserves1 = reserves1.Add(inAmount1.TruncateInt())
		*reserves0 = sdk.ZeroInt()
	} else {
		amountOutDec := maxAmount1Dec.Mul(price1To0)
		outAmount0 = amountOutDec.TruncateInt()
		*reserves1 = reserves1.Add(maxAmount1Dec.TruncateInt())
		*reserves0 = reserves0.Sub(outAmount0)
		inAmount1 = maxAmount1Dec
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
func (p *Pool) Deposit(maxAmount0 sdk.Int, maxAmount1 sdk.Int) (inAmount0 sdk.Int, inAmount1 sdk.Int, outShares sdk.Int) {
	lowerReserve0 := &p.LowerTick0.TickData.Reserve0AndShares[p.FeeIndex].Reserve0
	lowerTotalShares := &p.LowerTick0.TickData.Reserve0AndShares[p.FeeIndex].TotalShares
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
		*lowerTotalShares,
		inAmount0,
		inAmount1,
	)

	*lowerReserve0 = lowerReserve0.Add(inAmount0)
	*upperReserve1 = upperReserve1.Add(inAmount1)
	*lowerTotalShares = lowerTotalShares.Add(outShares)
	return inAmount0, inAmount1, outShares
}

// functionality that must happen outside of pool deposit copied below:
//
//
// shares, sharesFound := k.GetShares(ctx, msg.Receiver, pairId, tickIndex, feeIndex)
// if !sharesFound {
// 	shares = types.Shares{
// 		Address:     msg.Receiver,
// 		PairId:      pairId,
// 		TickIndex:   tickIndex,
// 		FeeIndex:    feeIndex,
// 		SharesOwned: sharesMinted,
// 	}
// } else {
// 	shares.SharesOwned = shares.SharesOwned.Add(sharesMinted)
// }

// k.SetShares(ctx, shares)

// totalAmountReserve0 = totalAmountReserve0.Add(trueAmount0)
// totalAmountReserve1 = totalAmountReserve1.Add(trueAmount1)

// ctx.EventManager().EmitEvent(types.CreateDepositEvent(
// 	msg.Creator,
// 	msg.Receiver,
// 	token0,
// 	token1,
// 	fmt.Sprint(msg.TickIndexes[i]),
// 	fmt.Sprint(msg.FeeIndexes[i]),
// 	lowerReserve0.Sub(trueAmount0).String(),
// 	upperReserve1.Sub(trueAmount1).String(),
// 	lowerReserve0.String(),
// 	upperReserve1.String(),
// 	sharesMinted.String(),
// ))

func (p *Pool) CalcSharesMinted(
	reserve0 sdk.Int,
	reserve1 sdk.Int,
	totalShares sdk.Int,
	amount0 sdk.Int,
	amount1 sdk.Int,
) (sharesMinted sdk.Int) {
	price1To0 := CalcPrice1To0(p.TickIndex) // center tick
	valueMintedToken0 := CalcShares(amount0, amount1, price1To0)
	valueExistingToken0 := CalcShares(reserve0, reserve1, price1To0)
	if valueExistingToken0.GT(sdk.ZeroDec()) {
		sharesMinted = valueMintedToken0.Quo(valueExistingToken0).Mul(totalShares.ToDec()).TruncateInt()
	} else {
		sharesMinted = valueMintedToken0.TruncateInt()
	}
	return sharesMinted
}

func (p *Pool) Withdraw(sharesToRemove sdk.Int) (outAmount0 sdk.Int, outAmount1 sdk.Int, err error) {
	totalShares := &p.LowerTick0.TickData.Reserve0AndShares[p.FeeIndex].TotalShares
	reserves0 := &p.LowerTick0.TickData.Reserve0AndShares[p.FeeIndex].Reserve0
	reserves1 := &p.UpperTick1.TickData.Reserve1[p.FeeIndex]
	if totalShares.LT(sharesToRemove) {
		return sdk.ZeroInt(), sdk.ZeroInt(), types.ErrNotEnoughShares
	}
	ownershipRatio := sharesToRemove.ToDec().Quo(totalShares.ToDec())
	outAmount1 = ownershipRatio.Mul(reserves1.ToDec()).TruncateInt()
	outAmount0 = ownershipRatio.Mul(reserves0.ToDec()).TruncateInt()
	*reserves0 = reserves0.Sub(outAmount0)
	*reserves1 = reserves1.Sub(outAmount1)
	*totalShares = totalShares.Sub(sharesToRemove)
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
