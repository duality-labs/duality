package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/utils"
)

type Pool struct {
	LowerTick0 *PoolReserves
	UpperTick1 *PoolReserves
}

func (p *Pool) CenterTickIndex() int64 {
	feeInt64 := utils.MustSafeUint64(p.Fee())
	return p.UpperTick1.Key.TickIndexTakerToMaker - feeInt64
}

func (p *Pool) Fee() uint64 {
	return p.UpperTick1.Key.Fee
}

func (p *Pool) GetLowerReserve0() sdk.Int {
	return p.LowerTick0.ReservesMakerDenom
}

func (p *Pool) GetUpperReserve1() sdk.Int {
	return p.UpperTick1.ReservesMakerDenom
}

func (p *Pool) Swap(
	tradePairID *TradePairID,
	maxAmountTakerIn sdk.Int,
	maxAmountMakerOut *sdk.Int,
) (amountTakerIn, amountMakerOut sdk.Int) {
	var takerReserves, makerReserves *PoolReserves
	if tradePairID.IsMakerDenomToken0() {
		makerReserves = p.LowerTick0
		takerReserves = p.UpperTick1
	} else {
		makerReserves = p.UpperTick1
		takerReserves = p.LowerTick0
	}

	if maxAmountTakerIn.Equal(sdk.ZeroInt()) ||
		makerReserves.ReservesMakerDenom.Equal(sdk.ZeroInt()) {
		return sdk.ZeroInt(), sdk.ZeroInt()
	}

	maxOutGivenTakerIn := makerReserves.PriceTakerToMaker.MulInt(maxAmountTakerIn).TruncateInt()
	possibleAmountsMakerOut := []sdk.Int{makerReserves.ReservesMakerDenom, maxOutGivenTakerIn}
	if maxAmountMakerOut != nil {
		possibleAmountsMakerOut = append(possibleAmountsMakerOut, *maxAmountMakerOut)
	}

	// outAmount will be the smallest value of:
	// a.) The available reserves1
	// b.) The most the user could get out given maxAmountIn0 (maxOutGivenIn1)
	// c.) The maximum amount the user wants out (maxAmountOut1)
	amountMakerOut = utils.MinIntArr(possibleAmountsMakerOut)
	amountTakerIn = sdk.NewDecFromInt(amountMakerOut).
		Quo(makerReserves.PriceTakerToMaker).
		Ceil().
		TruncateInt()

	takerReserves.ReservesMakerDenom = takerReserves.ReservesMakerDenom.Add(amountTakerIn)
	makerReserves.ReservesMakerDenom = makerReserves.ReservesMakerDenom.Sub(amountMakerOut)

	return amountTakerIn, amountMakerOut
}

// Mutates the Pool object and returns relevant change variables. Deposit is not committed until
// pool.save() is called or the underlying ticks are saved; this method does not use any keeper methods.
func (p *Pool) Deposit(
	maxAmount0,
	maxAmount1,
	existingShares sdk.Int,
	autoswap bool,
) (inAmount0, inAmount1 sdk.Int, outShares sdk.Coin) {
	lowerReserve0 := &p.LowerTick0.ReservesMakerDenom
	upperReserve1 := &p.UpperTick1.ReservesMakerDenom

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
	return NewDepositDenom(
		p.UpperTick1.Key.TradePairID.MustPairID(),
		p.CenterTickIndex(),
		p.Fee(),
	).String()
}

func (p *Pool) Price(tradePairID *TradePairID) sdk.Dec {
	if tradePairID.IsTakerDenomToken0() {
		return p.UpperTick1.PriceTakerToMaker
	}

	return p.LowerTick0.PriceTakerToMaker
}

func (p *Pool) MustCalcPrice1To0Center() sdk.Dec {
	// NOTE: We can safely call the error-less version of CalcPrice here because the pool object
	// has already been initialized with an upper and lower tick which satisfy a check for IsTickOutOfRange
	return MustCalcPrice(-1 * p.CenterTickIndex())
}

func (p *Pool) CalcSharesMinted(
	amount0 sdk.Int,
	amount1 sdk.Int,
	existingShares sdk.Int,
) (sharesMinted sdk.Coin) {
	price1To0Center := p.MustCalcPrice1To0Center()
	valueMintedToken0 := CalcAmountAsToken0(amount0, amount1, price1To0Center)

	valueExistingToken0 := CalcAmountAsToken0(
		p.LowerTick0.ReservesMakerDenom,
		p.UpperTick1.ReservesMakerDenom,
		price1To0Center,
	)
	var sharesMintedAmount sdk.Int
	if valueExistingToken0.GT(sdk.ZeroDec()) {
		sharesMintedAmount = valueMintedToken0.MulInt(existingShares).
			Quo(valueExistingToken0).
			TruncateInt()
	} else {
		sharesMintedAmount = valueMintedToken0.TruncateInt()
	}

	return sdk.Coin{Denom: p.GetDepositDenom(), Amount: sharesMintedAmount}
}

func (p *Pool) CalcResidualSharesMinted(
	residualAmount0 sdk.Int,
	residualAmount1 sdk.Int,
) (sharesMinted sdk.Coin, err error) {
	fee := CalcFee(p.UpperTick1.Key.TickIndexTakerToMaker, p.LowerTick0.Key.TickIndexTakerToMaker)
	valueMintedToken0, err := CalcResidualValue(
		residualAmount0,
		residualAmount1,
		p.LowerTick0.PriceTakerToMaker,
		fee,
	)
	if err != nil {
		return sdk.Coin{Denom: p.GetDepositDenom()}, err
	}

	return sdk.Coin{Denom: p.GetDepositDenom(), Amount: valueMintedToken0.TruncateInt()}, nil
}

func (p *Pool) RedeemValue(sharesToRemove, totalShares sdk.Int) (outAmount0, outAmount1 sdk.Int) {
	reserves0 := &p.LowerTick0.ReservesMakerDenom
	reserves1 := &p.UpperTick1.ReservesMakerDenom
	// outAmount1 = ownershipRatio * reserves1
	//            = (sharesToRemove / totalShares) * reserves1
	//            = (reserves1 * sharesToRemove ) / totalShares
	outAmount1 = sdk.NewDecFromInt(reserves1.Mul(sharesToRemove)).QuoInt(totalShares).TruncateInt()
	// outAmount0 = ownershipRatio * reserves1
	//            = (sharesToRemove / totalShares) * reserves1
	//            = (reserves1 * sharesToRemove ) / totalShares
	outAmount0 = sdk.NewDecFromInt(reserves0.Mul(sharesToRemove)).QuoInt(totalShares).TruncateInt()

	return outAmount0, outAmount1
}

func (p *Pool) Withdraw(sharesToRemove, totalShares sdk.Int) (outAmount0, outAmount1 sdk.Int) {
	reserves0 := &p.LowerTick0.ReservesMakerDenom
	reserves1 := &p.UpperTick1.ReservesMakerDenom
	outAmount0, outAmount1 = p.RedeemValue(sharesToRemove, totalShares)
	*reserves0 = reserves0.Sub(outAmount0)
	*reserves1 = reserves1.Sub(outAmount1)

	return outAmount0, outAmount1
}

// Balance trueAmount1 to the pool ratio
func CalcGreatestMatchingRatio(
	targetAmount0 sdk.Int,
	targetAmount1 sdk.Int,
	amount0 sdk.Int,
	amount1 sdk.Int,
) (resultAmount0, resultAmount1 sdk.Int) {
	targetAmount0Dec := sdk.NewDecFromInt(targetAmount0)
	targetAmount1Dec := sdk.NewDecFromInt(targetAmount1)

	// See spec: https://www.notion.so/dualityxyz/Autoswap-Spec-e856fa7b2438403c95147010d479b98c
	if targetAmount1.GT(sdk.ZeroInt()) {
		resultAmount0 = sdk.MinInt(
			amount0,
			sdk.NewDecFromInt(amount1).Mul(targetAmount0Dec).Quo(targetAmount1Dec).TruncateInt())
	} else {
		resultAmount0 = amount0
	}

	if targetAmount0.GT(sdk.ZeroInt()) {
		resultAmount1 = sdk.MinInt(
			amount1,
			sdk.NewDecFromInt(amount0).Mul(targetAmount1Dec).Quo(targetAmount0Dec).TruncateInt())
	} else {
		resultAmount1 = amount1
	}

	return resultAmount0, resultAmount1
}

func CalcResidualValue(
	amount0, amount1 sdk.Int,
	priceLowerTakerToMaker sdk.Dec,
	fee int64,
) (sdk.Dec, error) {
	// ResidualValue = Amount0 * (Price1to0Center / Price1to0Upper) + Amount1 * Price1to0Lower
	amount0Discount, err := CalcPrice(-fee)
	if err != nil {
		return sdk.ZeroDec(), err
	}

	return amount0Discount.MulInt(amount0).Add(priceLowerTakerToMaker.MulInt(amount1)), nil
}

func CalcFee(upperTickIndex, lowerTickIndex int64) int64 {
	return (upperTickIndex - lowerTickIndex) / 2
}

func CalcAmountAsToken0(amount0, amount1 sdk.Int, price1To0 sdk.Dec) sdk.Dec {
	amount0Dec := sdk.NewDecFromInt(amount0)

	return amount0Dec.Add(price1To0.MulInt(amount1))
}
