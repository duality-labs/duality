package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/utils"
)

func (t LimitOrderTranche) IsPlaceTranche() bool {
	return t.ReservesTokenIn.Equal(t.TotalTokenIn)
}

func NewFromFilledTranche(t FilledLimitOrderTranche) LimitOrderTranche {
	return LimitOrderTranche{
		TrancheKey:       t.TrancheKey,
		TickIndex:        t.TickIndex,
		TokenIn:          t.TokenIn,
		PairId:           t.PairId,
		TotalTokenOut:    t.TotalTokenOut,
		TotalTokenIn:     t.TotalTokenIn,
		ReservesTokenOut: t.ReservesTokenOut,
	}
}

func (t LimitOrderTranche) CreateFilledTranche() FilledLimitOrderTranche {
	return FilledLimitOrderTranche{
		TrancheKey:       t.TrancheKey,
		TickIndex:        t.TickIndex,
		TokenIn:          t.TokenIn,
		PairId:           t.PairId,
		TotalTokenIn:     t.TotalTokenIn,
		TotalTokenOut:    t.TotalTokenOut,
		ReservesTokenOut: t.ReservesTokenOut,
	}
}

func (t LimitOrderTranche) IsFilled() bool {
	return t.ReservesTokenIn.IsZero()
}

func (t *LimitOrderTranche) Price() sdk.Dec {
	return t.PriceTakerToMaker()
}

func (t LimitOrderTranche) HasToken() bool {
	return t.ReservesTokenIn.GT(sdk.ZeroInt())
}

func (t LimitOrderTranche) IsTokenInToken0() bool {
	return t.TokenIn == t.PairId.Token0
}

func (t LimitOrderTranche) PriceMakerToTaker() sdk.Dec {
	if t.IsTokenInToken0() {
		return utils.MustCalcPrice0To1(t.TickIndex)
	} else {
		return utils.MustCalcPrice1To0(t.TickIndex)
	}
}

func (t LimitOrderTranche) PriceTakerToMaker() sdk.Dec {
	if t.IsTokenInToken0() {
		return utils.MustCalcPrice1To0(t.TickIndex)
	} else {
		return utils.MustCalcPrice0To1(t.TickIndex)
	}
}

func (t LimitOrderTranche) RatioFilled() sdk.Dec {
	amountFilled := t.PriceTakerToMaker().MulInt(t.TotalTokenOut)
	ratioFilled := amountFilled.QuoInt(t.TotalTokenIn)
	return ratioFilled
}

func (t LimitOrderTranche) HasLiquidity() bool {
	return t.ReservesTokenIn.GT(sdk.ZeroInt())
}

func (t *LimitOrderTranche) Cancel(trancheUser LimitOrderTrancheUser) (amountToCancel sdk.Int) {
	ratioNotFilled := sdk.OneDec().Sub(t.RatioFilled())

	amountToCancel = trancheUser.SharesOwned.ToDec().Mul(ratioNotFilled).TruncateInt()
	t.ReservesTokenIn = t.ReservesTokenIn.Sub(amountToCancel)

	return amountToCancel

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
	amountFilledTokenOut := maxAmountTaker.ToDec().Mul(t.PriceTakerToMaker()).TruncateInt()
	if reservesTokenOut.LTE(amountFilledTokenOut) {
		inAmount = reservesTokenOut.ToDec().Mul(t.PriceMakerToTaker()).TruncateInt()
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
