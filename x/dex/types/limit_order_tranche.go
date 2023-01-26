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
		TrancheIndex:     t.TrancheIndex,
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
		TrancheIndex:     t.TrancheIndex,
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

func (t LimitOrderTranche) HasLiquidity() bool {
	return t.ReservesTokenIn.GT(sdk.ZeroInt())
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
