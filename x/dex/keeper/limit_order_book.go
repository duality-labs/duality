package keeper

import (
	"context"

	"github.com/duality-labs/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type LimitOrderBook struct {
	Pair                *types.TradingPair
	TokenIn             string
	Tick                *types.Tick
	LimitTrancheIndexes *types.LimitTrancheIndexes
	Tranches            map[uint64]*LimitOrderTranche
}

func (k Keeper) GetLimitOrderBook1To0(
	ctx context.Context,
	pair *types.TradingPair,
	tick *types.Tick,
) *LimitOrderBook {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	priceTakerToMaker := tick.Price0To1
	priceMakerToTaker := sdk.OneDec().Quo(*priceTakerToMaker)

	fillIndex := tick.LimitOrderTranche1To0.FillTrancheIndex
	placeIndex := tick.LimitOrderTranche1To0.PlaceTrancheIndex

	_, token1 := pair.ToTokens()
	fillTrancheDTO, found := k.GetLimitOrderTranche(
		sdkCtx,
		pair.PairId,
		tick.TickIndex,
		token1,
		fillIndex,
	)
	if !found {
		panic("There should be a tranche here!")
	}
	fillTranche := NewLimitOrderTranche(
		tick,
		&fillTrancheDTO,
		priceMakerToTaker,
		*priceTakerToMaker,
	)

	var placeTranche *LimitOrderTranche
	if fillIndex != placeIndex {
		placeTrancheDTO, found := k.GetLimitOrderTranche(
			sdkCtx,
			pair.PairId,
			tick.TickIndex,
			token1,
			placeIndex,
		)
		if !found {
			panic("There should be a tranche here!")
		}
		placeTranche = NewLimitOrderTranche(
			tick,
			&placeTrancheDTO,
			priceMakerToTaker,
			*priceTakerToMaker,
		)
	} else {
		placeTranche = fillTranche
	}

	return NewLimitOrderBook0To1(
		pair,
		tick,
		fillTranche,
		placeTranche,
	)
}

func (k Keeper) GetLimitOrderBook0To1(
	ctx context.Context,
	pair *types.TradingPair,
	tick *types.Tick,
) *LimitOrderBook {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	priceMakerToTaker := tick.Price0To1
	priceTakerToMaker := sdk.OneDec().Quo(*priceMakerToTaker)

	fillIndex := tick.LimitOrderTranche0To1.FillTrancheIndex
	placeIndex := tick.LimitOrderTranche0To1.PlaceTrancheIndex

	token0, _ := pair.ToTokens()
	fillTrancheDTO, found := k.GetLimitOrderTranche(
		sdkCtx,
		pair.PairId,
		tick.TickIndex,
		token0,
		fillIndex,
	)
	if !found {
		panic("There should be a tranche here!")
	}
	fillTranche := NewLimitOrderTranche(
		tick,
		&fillTrancheDTO,
		*priceMakerToTaker,
		priceTakerToMaker,
	)

	placeTranche := fillTranche
	if fillIndex != placeIndex {
		placeTrancheDTO, found := k.GetLimitOrderTranche(
			sdkCtx,
			pair.PairId,
			tick.TickIndex,
			token0,
			placeIndex,
		)
		if !found {
			panic("There should be a tranche here!")
		}
		placeTranche = NewLimitOrderTranche(
			tick,
			&placeTrancheDTO,
			*priceMakerToTaker,
			priceTakerToMaker,
		)
	}

	return NewLimitOrderBook0To1(
		pair,
		tick,
		fillTranche,
		placeTranche,
	)
}

func NewLimitOrderBook1To0(
	pair *types.TradingPair,
	tick *types.Tick,
	fillTranche *LimitOrderTranche,
	placeTranche *LimitOrderTranche,
) *LimitOrderBook {
	return &LimitOrderBook{
		Pair:                pair,
		Tick:                tick,
		LimitTrancheIndexes: tick.LimitOrderTranche1To0,
		Tranches: map[uint64]*LimitOrderTranche{
			tick.LimitOrderTranche1To0.FillTrancheIndex:  fillTranche,
			tick.LimitOrderTranche1To0.PlaceTrancheIndex: placeTranche,
		},
	}
}

func NewLimitOrderBook0To1(
	pair *types.TradingPair,
	tick *types.Tick,
	fillTranche *LimitOrderTranche,
	placeTranche *LimitOrderTranche,
) *LimitOrderBook {
	return &LimitOrderBook{
		Pair:                pair,
		Tick:                tick,
		LimitTrancheIndexes: tick.LimitOrderTranche0To1,
		Tranches: map[uint64]*LimitOrderTranche{
			tick.LimitOrderTranche0To1.FillTrancheIndex:  fillTranche,
			tick.LimitOrderTranche0To1.PlaceTrancheIndex: placeTranche,
		},
	}
}

func (l *LimitOrderBook) NewTranche(trancheIndex uint64) *LimitOrderTranche {
	token0, _ := l.Pair.ToTokens()
	var priceMakerToTaker, priceTakerToMaker sdk.Dec
	if l.TokenIn == token0 {
		priceMakerToTaker = *l.Tick.Price0To1
	} else {
		priceTakerToMaker = sdk.OneDec().Quo(*l.Tick.Price0To1)
	}
	return NewLimitOrderTranche(
		l.Tick,
		&types.LimitOrderTranche{
			PairId:           l.Pair.PairId,
			TokenIn:          l.TokenIn,
			TickIndex:        l.Tick.TickIndex,
			TrancheIndex:     trancheIndex,
			ReservesTokenIn:  sdk.ZeroInt(),
			ReservesTokenOut: sdk.ZeroInt(),
			TotalTokenIn:     sdk.ZeroInt(),
			TotalTokenOut:    sdk.ZeroInt(),
		},
		priceMakerToTaker,
		priceTakerToMaker,
	)
}

func (l LimitOrderBook) Swap(maxAmount sdk.Int) (
	inAmount sdk.Int,
	outAmount sdk.Int,
	initedTick *types.Tick,
	deinitedTick *types.Tick,
) {
	inAmount = sdk.ZeroInt()
	outAmount = sdk.ZeroInt()

	fillTranche := &l.LimitTrancheIndexes.FillTrancheIndex
	placeTranche := &l.LimitTrancheIndexes.PlaceTrancheIndex

	for maxAmount.GT(sdk.ZeroInt()) && *fillTranche < *placeTranche {
		var curInAmount sdk.Int
		var curOutAmount sdk.Int
		curInAmount, curOutAmount, _, deinitedTick = l.Tranches[*fillTranche].Swap(maxAmount)
		maxAmount = maxAmount.Sub(curInAmount)
		inAmount = inAmount.Add(curInAmount)
		outAmount = outAmount.Add(curOutAmount)
		if deinitedTick != nil {
			*fillTranche++
		}
	}

	if maxAmount.GT(sdk.ZeroInt()) {
		var curInAmount sdk.Int
		var curOutAmount sdk.Int
		curInAmount, curOutAmount, _, deinitedTick = l.Tranches[*fillTranche].Swap(maxAmount)
		inAmount = inAmount.Add(curInAmount)
		outAmount = outAmount.Add(curOutAmount)
		if deinitedTick != nil {
			*fillTranche++
			*placeTranche++
		} else {
			*placeTranche++
		}
		l.Tranches[*placeTranche] = l.NewTranche(*placeTranche)
	}

	return inAmount, outAmount, nil, deinitedTick
}

func (l LimitOrderBook) Save(ctx context.Context, keeper Keeper) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	keeper.SetTradingPair(sdkCtx, *l.Pair)
	for _, tranche := range l.Tranches {
		tranche.Save(ctx, keeper)
	}
}

func (l LimitOrderBook) Price() sdk.Dec {
	tranche := l.Tranches[l.LimitTrancheIndexes.FillTrancheIndex]
	return tranche.PriceTakerToMaker
}

func (l LimitOrderBook) HasLiquidity() bool {
	tranche := l.Tranches[l.LimitTrancheIndexes.FillTrancheIndex]
	return tranche.HasLiquidity()
}
