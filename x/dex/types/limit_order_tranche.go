package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (t LimitOrderTranche) IsPlaceTranche() bool {
	return t.ReservesTokenIn.Equal(t.TotalTokenIn)
}

func (t LimitOrderTranche) IsFilled() bool {
	return t.ReservesTokenIn.IsZero()
}

func (t LimitOrderTranche) IsJIT() bool {
	return t.ExpirationTime != nil && *t.ExpirationTime == JITGoodTilTime
}

func (t LimitOrderTranche) IsExpired(ctx sdk.Context) bool {
	return t.ExpirationTime != nil && !t.IsJIT() && !t.ExpirationTime.After(ctx.BlockTime())
}

func (t *LimitOrderTranche) Price() *Price {
	return t.PriceTakerToMaker()
}

func (t LimitOrderTranche) HasTokenIn() bool {
	return t.ReservesTokenIn.GT(sdk.ZeroInt())
}

func (t LimitOrderTranche) HasTokenOut() bool {
	return t.ReservesTokenOut.GT(sdk.ZeroInt())
}

func (t LimitOrderTranche) IsTokenInToken0() bool {
	return t.TokenIn == t.PairID.Token0
}

func (t *LimitOrderTranche) Ref() []byte {
	// returns the KVstore key for a tranche
	return TickLiquidityKey(
		t.PairID,
		t.TokenIn,
		t.TickIndex,
		LiquidityTypeLimitOrder,
		t.TrancheKey,
	)
}

func (t LimitOrderTranche) PriceMakerToTaker() *Price {
	if t.IsTokenInToken0() {
		return MustNewPrice(t.TickIndex)
	}

	return MustNewPrice(-1 * t.TickIndex)
}

func (t LimitOrderTranche) PriceTakerToMaker() *Price {
	if t.IsTokenInToken0() {
		return MustNewPrice(-1 * t.TickIndex)
	}

	return MustNewPrice(t.TickIndex)
}

func (t LimitOrderTranche) RatioFilled() sdk.Dec {
	amountFilled := t.PriceTakerToMaker().MulInt(t.TotalTokenOut)
	ratioFilled := amountFilled.QuoInt(t.TotalTokenIn)
	return ratioFilled
}

func (t LimitOrderTranche) AmountUnfilled() sdk.Dec {
	amountFilled := t.PriceTakerToMaker().MulInt(t.TotalTokenOut)
	return t.TotalTokenIn.ToDec().Sub(amountFilled)
}

func (t LimitOrderTranche) HasLiquidity() bool {
	return t.ReservesTokenIn.GT(sdk.ZeroInt())
}

func (t *LimitOrderTranche) RemoveTokenIn(trancheUser LimitOrderTrancheUser) (amountToRemove sdk.Int) {
	amountUnfilled := t.AmountUnfilled()
	maxAmountToRemove := amountUnfilled.MulInt(trancheUser.SharesOwned).QuoInt(t.TotalTokenIn).TruncateInt()
	amountToRemove = maxAmountToRemove.Sub(trancheUser.SharesCancelled)
	t.ReservesTokenIn = t.ReservesTokenIn.Sub(amountToRemove)

	return amountToRemove
}

func (t *LimitOrderTranche) Withdraw(trancheUser LimitOrderTrancheUser) (sdk.Int, sdk.Dec) {
	reservesTokenOutDec := sdk.NewDecFromInt(t.ReservesTokenOut)

	ratioFilled := t.RatioFilled()
	maxAllowedToWithdraw := ratioFilled.MulInt(trancheUser.SharesOwned).TruncateInt()
	amountOutTokenIn := maxAllowedToWithdraw.Sub(trancheUser.SharesWithdrawn)
	amountOutTokenOut := t.PriceMakerToTaker().MulInt(amountOutTokenIn)
	t.ReservesTokenOut = reservesTokenOutDec.Sub(amountOutTokenOut).TruncateInt()

	return amountOutTokenIn, amountOutTokenOut
}

func (t *LimitOrderTranche) Swap(maxAmountTaker sdk.Int) (
	inAmount sdk.Int,
	outAmount sdk.Int,
) {
	reservesTokenOut := &t.ReservesTokenIn
	fillTokenIn := &t.ReservesTokenOut
	totalTokenIn := &t.TotalTokenOut
	amountFilledTokenOut := t.PriceTakerToMaker().MulInt(maxAmountTaker).TruncateInt()
	if reservesTokenOut.LTE(amountFilledTokenOut) {
		inAmount = t.PriceMakerToTaker().MulInt(*reservesTokenOut).Ceil().TruncateInt()
		outAmount = *reservesTokenOut
		*reservesTokenOut = sdk.ZeroInt()
		*fillTokenIn = fillTokenIn.Add(inAmount)
		*totalTokenIn = totalTokenIn.Add(inAmount)
	} else {
		inAmount = maxAmountTaker
		outAmount = amountFilledTokenOut
		*fillTokenIn = fillTokenIn.Add(maxAmountTaker)
		*totalTokenIn = totalTokenIn.Add(maxAmountTaker)
		*reservesTokenOut = reservesTokenOut.Sub(amountFilledTokenOut)
	}

	return inAmount, outAmount
}

func (t *LimitOrderTranche) PlaceMakerLimitOrder(amountIn sdk.Int) {
	t.ReservesTokenIn = t.ReservesTokenIn.Add(amountIn)
	t.TotalTokenIn = t.TotalTokenIn.Add(amountIn)
}
