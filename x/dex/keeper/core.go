package keeper

import (
	"context"
	"fmt"
	"math"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// func (k Keeper) addEdges(goCtx context.Context, token0Index int64, token1Index int64) {
// 	ctx := sdk.UnwrapSDKContext(goCtx)

// 	if k.GetTokensCount(ctx) <  2 * k.GetEdgeRowCount(ctx) {

// 		id := k.AppendEdgeRow(ctx, types.EdgeRow{k.GetEdgeRowCount(ctx), false})
// 		k.AppendAdjanceyMatrix(ctx, types.AdjanceyMatrix{id, make([]types.EdgeRow, k.GetEdgeRowCount8(ctx))})
// 	}
// }

func (k Keeper) DepositPairHelper(goCtx context.Context, token0 string, token1 string, price_index int64, feeIndex int64) error {

	ctx := sdk.UnwrapSDKContext(goCtx)

	token0Index, token0Found := k.GetTokenMap(ctx, token0)
	tokenLength := k.GetTokensCount(ctx)

	if !token0Found {
		k.SetTokenMap(ctx, types.TokenMap{Address: token0, Index: int64(tokenLength)})
		newTokenLength := tokenLength + 1
		token0Index.Index = int64(tokenLength)
		k.SetTokensCount(ctx, newTokenLength)
		k.AppendTokens(ctx, types.Tokens{Id: tokenLength, Address: token0})
	}

	token1Index, token1Found := k.GetTokenMap(ctx, token1)

	if !token1Found {
		k.SetTokenMap(ctx, types.TokenMap{Address: token1, Index: int64(tokenLength)})
		newTokenLength := tokenLength + 1
		token1Index.Index = int64(tokenLength)
		k.SetTokensCount(ctx, newTokenLength)
		k.AppendTokens(ctx, types.Tokens{Id: tokenLength, Address: token1})
	}

	pairId := k.CreatePairId(token0, token1)
	_, PairFound := k.GetPairMap(ctx, pairId)

	if !PairFound {

		feeValue, _ := k.GetFeeList(ctx, uint64(feeIndex))

		// addEdges(goCtx, token0Index.Index, token1Index.Index)

		k.SetPairMap(ctx, types.PairMap{
			PairId: pairId,
			TokenPair: &types.TokenPairType{
				CurrentTick0To1: price_index + feeValue.Fee,
				CurrentTick1To0: price_index - feeValue.Fee,
			},
		})
	}
	return nil

}

func (k Keeper) Calc_price(price_Index int64) (sdk.Dec, error) {
	floatPrice := math.Pow(1.0001, float64(price_Index))
	sPrice := fmt.Sprintf("%f", floatPrice)

	price, err := sdk.NewDecFromStr(sPrice)

	if err != nil {
		return sdk.ZeroDec(), err
	} else {
		return price, nil
	}

}

func (k Keeper) Min(a, b sdk.Dec) sdk.Dec {
	if a.LT(b) {
		return a
	}
	return b
}

func (k Keeper) SingleDeposit(goCtx context.Context, msg *types.MsgDeposit, token0 string, token1 string, callerAddr sdk.AccAddress, amount0 sdk.Dec, amount1 sdk.Dec) error {

	ctx := sdk.UnwrapSDKContext(goCtx)
	feeValue, _ := k.GetFeeList(ctx, uint64(msg.FeeIndex))
	fee := feeValue.Fee
	k.DepositPairHelper(goCtx, token0, token1, msg.PriceIndex, fee)

	pairId := k.CreatePairId(token0, token1)

	pair, pairFound := k.GetPairMap(ctx, pairId)

	if !pairFound {
		return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Pair not found")
	}

	fmt.Println("TopTick Index")
	fmt.Println(msg.PriceIndex + fee)
	bottomTick, bottomTickFound := k.GetTickMap(ctx, pairId, msg.PriceIndex-int64(fee))
	topTick, topTickFound := k.GetTickMap(ctx, pairId, msg.PriceIndex+int64(fee))

	trueAmount0 := amount0
	trueAmount1 := amount1
	var sharesMinted sdk.Dec
	var oldReserve0 sdk.Dec
	var oldReserve1 sdk.Dec

	if amount0.GT(sdk.ZeroDec()) && (msg.PriceIndex-int64(msg.FeeIndex)) < pair.TokenPair.CurrentTick1To0 {
		return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Cannot depsoit amount 0 at a tick less than the CurrentTick1to0")
	}

	if amount1.GT(sdk.ZeroDec()) && (msg.PriceIndex+int64(msg.FeeIndex)) > pair.TokenPair.CurrentTick0To1 {
		return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Cannot depsoit amount 0 at a tick greater than the CurrentTick0to1")
	}

	price, err := k.Calc_price(msg.PriceIndex)

	if err != nil {
		return err
	}

	// TODO add support for adding n+1 tick given we have a fee_list of size n
	// check if tick array is < than fee_list if so append to equal that size.

	if !bottomTickFound || !topTickFound || topTick.TickData.Reserve0AndShares[msg.FeeIndex].TotalShares.Equal(sdk.ZeroDec()) {
		sharesMinted = trueAmount0.Add(amount1.Mul(price))

		feeSize := k.GetFeeListCount(ctx)
		if !bottomTickFound {

			bottomTick = types.TickMap{
				TickIndex: msg.PriceIndex - int64(fee),
				TickData: &types.TickDataType{
					Reserve0AndShares: make([]*types.Reserve0AndSharesType, feeSize),
					Reserve1:          make([]sdk.Dec, feeSize),
				},
			}

			for i, _ := range bottomTick.TickData.Reserve0AndShares {
				bottomTick.TickData.Reserve0AndShares[i] = &types.Reserve0AndSharesType{sdk.ZeroDec(), sdk.ZeroDec()}

			}
			for i, _ := range bottomTick.TickData.Reserve1 {
				bottomTick.TickData.Reserve1[i] = sdk.ZeroDec()
			}
			//bottomTick = NewTick
		}

		if !topTickFound {
			topTick = types.TickMap{
				TickIndex: msg.PriceIndex + int64(fee),
				TickData: &types.TickDataType{
					Reserve0AndShares: make([]*types.Reserve0AndSharesType, feeSize),
					Reserve1:          make([]sdk.Dec, feeSize),
				},
			}

			for i, _ := range topTick.TickData.Reserve0AndShares {
				topTick.TickData.Reserve0AndShares[i] = &types.Reserve0AndSharesType{sdk.ZeroDec(), sdk.ZeroDec()}

			}
			for i, _ := range topTick.TickData.Reserve1 {
				topTick.TickData.Reserve1[i] = sdk.ZeroDec()
			}
			//topTick = NewTick
		}

		oldReserve0 = sdk.ZeroDec()
		oldReserve1 = sdk.ZeroDec()

		NewReserve0andShares := &types.Reserve0AndSharesType{
			Reserve0:    trueAmount0,
			TotalShares: sharesMinted,
		}

		topTick.TickData.Reserve0AndShares[msg.FeeIndex] = NewReserve0andShares

		bottomTick.TickData.Reserve1[msg.FeeIndex] = trueAmount1

	} else {

		if uint64(len(topTick.TickData.Reserve1)) < k.GetFeeListCount(ctx) {
			topTick.TickData.Reserve1 = append(topTick.TickData.Reserve1, sdk.ZeroDec())
			topTick.TickData.Reserve0AndShares = append(topTick.TickData.Reserve0AndShares, &types.Reserve0AndSharesType{})
		}

		if uint64(len(bottomTick.TickData.Reserve1)) < k.GetFeeListCount(ctx) {
			bottomTick.TickData.Reserve1 = append(bottomTick.TickData.Reserve1, sdk.ZeroDec())
			bottomTick.TickData.Reserve0AndShares = append(bottomTick.TickData.Reserve0AndShares, &types.Reserve0AndSharesType{})
		}

		if topTick.TickData.Reserve0AndShares[msg.FeeIndex].Reserve0.GT(sdk.ZeroDec()) {
			trueAmount1 = k.Min(amount1, bottomTick.TickData.Reserve1[msg.FeeIndex].Mul(amount0).Quo(topTick.TickData.Reserve0AndShares[msg.FeeIndex].Reserve0))
			// trueAmount1 = min(amt1 , (reserve1 * amt0)/reserve0 )
		}

		if bottomTick.TickData.Reserve1[msg.FeeIndex].GT(sdk.ZeroDec()) {
			trueAmount0 = k.Min(amount0, topTick.TickData.Reserve0AndShares[msg.FeeIndex].Reserve0.Mul(amount1).Quo(bottomTick.TickData.Reserve1[msg.FeeIndex]))
		}

		// if amount0 is 0 amt1/reserve1 * totalShares = sharesMinted
		// else if amt0/reserve0 * totalShares
		if trueAmount0.GT(sdk.ZeroDec()) {
			sharesMinted = (trueAmount0.Quo(topTick.TickData.Reserve0AndShares[msg.FeeIndex].Reserve0).Mul(topTick.TickData.Reserve0AndShares[msg.FeeIndex].TotalShares))
		} else {
			sharesMinted = (trueAmount1.Quo(bottomTick.TickData.Reserve1[msg.FeeIndex]).Mul(topTick.TickData.Reserve0AndShares[msg.FeeIndex].TotalShares))
		}

		// ((amt0 / reserve0 )* totalShares) + ((amt1 / reserve1) * totalShares)
		oldReserve0 = topTick.TickData.Reserve0AndShares[msg.FeeIndex].Reserve0
		oldReserve1 = bottomTick.TickData.Reserve1[msg.FeeIndex]
		topTick.TickData.Reserve0AndShares[msg.FeeIndex].Reserve0 = topTick.TickData.Reserve0AndShares[msg.FeeIndex].Reserve0.Add(trueAmount0)
		topTick.TickData.Reserve0AndShares[msg.FeeIndex].TotalShares = topTick.TickData.Reserve0AndShares[msg.FeeIndex].TotalShares.Add(sharesMinted)

		bottomTick.TickData.Reserve1[msg.FeeIndex] = bottomTick.TickData.Reserve1[msg.FeeIndex].Add(trueAmount1)

	}

	if trueAmount0.GT(sdk.ZeroDec()) && (msg.PriceIndex-int64(msg.FeeIndex) > pair.TokenPair.CurrentTick1To0 && (msg.PriceIndex-int64(msg.FeeIndex) < pair.TokenPair.CurrentTick0To1)) {
		pair.TokenPair.CurrentTick0To1 = msg.PriceIndex - int64(msg.FeeIndex)
	}

	if trueAmount1.GT(sdk.ZeroDec()) && (msg.PriceIndex+int64(msg.FeeIndex) > pair.TokenPair.CurrentTick1To0 && (msg.PriceIndex+int64(msg.FeeIndex) < pair.TokenPair.CurrentTick0To1)) {
		pair.TokenPair.CurrentTick1To0 = msg.PriceIndex + int64(msg.FeeIndex)
	}

	shares, sharesFound := k.GetShares(ctx, msg.Creator, pairId, msg.PriceIndex, msg.FeeIndex)

	if !sharesFound {
		shares = types.Shares{
			Address:     msg.Creator,
			PairId:      pairId,
			PriceIndex:  msg.PriceIndex,
			FeeIndex:    msg.FeeIndex,
			SharesOwned: sharesMinted,
		}
	} else {
		shares.SharesOwned = shares.SharesOwned.Add(sharesMinted)
	}

	if trueAmount0.GT(sdk.ZeroDec()) {
		coin0 := sdk.NewCoin(token0, sdk.NewIntFromBigInt(trueAmount0.BigInt()))
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin0}); err != nil {
			return err
		}
	}

	if trueAmount1.GT(sdk.ZeroDec()) {
		coin1 := sdk.NewCoin(token1, sdk.NewIntFromBigInt(trueAmount1.BigInt()))
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin1}); err != nil {
			return err
		}
	}

	fmt.Println(shares)
	k.SetTickMap(ctx, pairId, bottomTick)
	k.SetTickMap(ctx, pairId, topTick)
	k.SetShares(ctx, shares)

	ctx.EventManager().EmitEvent(types.CreateDepositEvent(msg.Creator,
		token0, token1, fmt.Sprint(msg.PriceIndex), fmt.Sprint(msg.FeeIndex),
		oldReserve0.String(), oldReserve1.String(), bottomTick.TickData.Reserve1[msg.FeeIndex].String(), topTick.TickData.Reserve0AndShares[msg.FeeIndex].Reserve0.String(),
		sharesMinted.String()),
	)

	_ = goCtx
	return nil
}

func (k Keeper) MultiDeposit(goCtx context.Context, msg *types.MsgDeposit) error {

	_ = goCtx
	return nil
}

func (k Keeper) SingleWithdrawl(goCtx context.Context, msg *types.MsgWithdrawl, token0 string, token1 string, callerAddr sdk.AccAddress, sharesToRemove sdk.Dec) error {
	ctx := sdk.UnwrapSDKContext(goCtx)
	feeValue, _ := k.GetFeeList(ctx, uint64(msg.FeeIndex))
	fee := feeValue.Fee

	pairId := k.CreatePairId(token0, token1)

	pair, pairFound := k.GetPairMap(ctx, pairId)

	if !pairFound {
		return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Pair not found")
	}

	addtick, addtickFound := k.GetTickMap(ctx, pairId, msg.PriceIndex+int64(fee))
	subtick, subTickFound := k.GetTickMap(ctx, pairId, msg.PriceIndex-int64(fee))

	var OldReserve0 sdk.Dec
	var OldReserve1 sdk.Dec
	if !addtickFound || !subTickFound {
		return sdkerrors.Wrapf(types.ErrValidTickNotFound, "No tick found at the requested index")
	}

	if addtick.TickData.Reserve0AndShares[msg.FeeIndex].TotalShares.Equal(sdk.ZeroDec()) {
		return sdkerrors.Wrapf(types.ErrValidTickNotFound, "No tick found at the requested index")
	}

	reserve0ToRemove := (sharesToRemove.Quo(addtick.TickData.Reserve0AndShares[msg.FeeIndex].TotalShares)).Mul(addtick.TickData.Reserve0AndShares[msg.FeeIndex].Reserve0)
	reserve1ToRemove := (sharesToRemove.Quo(addtick.TickData.Reserve0AndShares[msg.FeeIndex].TotalShares)).Mul(subtick.TickData.Reserve1[msg.FeeIndex])

	OldReserve0 = addtick.TickData.Reserve0AndShares[msg.FeeIndex].Reserve0
	OldReserve1 = subtick.TickData.Reserve1[msg.FeeIndex]

	addtick.TickData.Reserve0AndShares[msg.FeeIndex].Reserve0 = addtick.TickData.Reserve0AndShares[msg.FeeIndex].Reserve0.Sub(reserve0ToRemove)
	subtick.TickData.Reserve1[msg.FeeIndex] = subtick.TickData.Reserve1[msg.FeeIndex].Sub(reserve1ToRemove)
	addtick.TickData.Reserve0AndShares[msg.FeeIndex].TotalShares = addtick.TickData.Reserve0AndShares[msg.FeeIndex].TotalShares.Sub(sharesToRemove)

	shareOwner, shareOwnerFound := k.GetShares(ctx, msg.Creator, pairId, msg.PriceIndex, msg.FeeIndex)

	if !shareOwnerFound {
		return sdkerrors.Wrapf(types.ErrValidShareNotFound, "No valid share owner fonnd")
	}

	shareOwner.SharesOwned = shareOwner.SharesOwned.Sub(sharesToRemove)

	removeTick := true
	if addtick.TickData.Reserve0AndShares[msg.FeeIndex].TotalShares.Equal(sdk.ZeroDec()) {
		for _, s := range addtick.TickData.Reserve0AndShares {
			if s.TotalShares.GT(sdk.ZeroDec()) {
				removeTick = false
			}
		}
	}
	if removeTick && (msg.PriceIndex+int64(fee) == pair.TokenPair.CurrentTick1To0) {

		tickFound := false
		c := 0
		for tickFound != true {
			c++
			_, tickFound = k.GetTickMap(ctx, pairId, (msg.PriceIndex + int64(msg.FeeIndex) + int64(c)))

		}

		pair.TokenPair.CurrentTick1To0 = (msg.PriceIndex + int64(msg.FeeIndex) + int64(c))
	}

	if removeTick && (msg.PriceIndex-int64(fee) == pair.TokenPair.CurrentTick0To1) {

		tickFound := false
		c := 0
		for tickFound != true {
			c++
			_, tickFound = k.GetTickMap(ctx, pairId, (msg.PriceIndex - int64(msg.FeeIndex) - int64(c)))

		}

		pair.TokenPair.CurrentTick1To0 = (msg.PriceIndex - int64(msg.FeeIndex) - int64(c))
	}

	k.SetShares(ctx, shareOwner)
	k.SetTickMap(ctx, pairId, addtick)
	k.SetTickMap(ctx, pairId, subtick)

	if reserve0ToRemove.GT(sdk.ZeroDec()) {
		coin0 := sdk.NewCoin(token0, sdk.NewIntFromBigInt(reserve0ToRemove.BigInt()))
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, callerAddr, sdk.Coins{coin0}); err != nil {
			return err
		}
	}

	if reserve1ToRemove.GT(sdk.ZeroDec()) {
		coin1 := sdk.NewCoin(token1, sdk.NewIntFromBigInt(reserve1ToRemove.BigInt()))
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, callerAddr, sdk.Coins{coin1}); err != nil {
			return err
		}
	}

	ctx.EventManager().EmitEvent(types.CreateWithdrawEvent(msg.Creator,
		token0, token1, fmt.Sprint(msg.PriceIndex), fmt.Sprint(msg.FeeIndex), OldReserve0.String(), OldReserve1.String(),
		addtick.TickData.Reserve0AndShares[msg.FeeIndex].Reserve0.String(), subtick.TickData.Reserve1[msg.FeeIndex].String(),
		sdk.NewAttribute(types.WithdrawEventSharesRemoved, sharesToRemove.String()),
	))
	_ = ctx
	return nil
}

////// Swap Functions

func (k Keeper) Swap0to1(goCtx context.Context, msg *types.MsgSwap, token0 string, token1 string, callerAddr sdk.AccAddress, amountIn sdk.Dec, minOut sdk.Dec) error {

	ctx := sdk.UnwrapSDKContext(goCtx)

	pairId := k.CreatePairId(token0, token1)

	feeSize := k.GetFeeListCount(ctx)
	pair, pairFound := k.GetPairMap(ctx, pairId)

	if !pairFound {
		return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Pair not found")
	}

	amount_left := amountIn
	amount_out := sdk.ZeroDec()

	for amount_left.Neg().Equal(sdk.ZeroDec()) {
		Current1Data, _ := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick0To1)
		var i uint64
		i = 0
		for i < feeSize || amount_left.Neg().Equal(sdk.ZeroDec()) {
			fee, _ := k.GetFeeList(ctx, i)
			feeIndex := fee.Fee
			Current0Data, _ := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick0To1+2*feeIndex)
			//Current0Datam := Current0Data.TickData.Reserve1[i]

			price, err := k.Calc_price(pair.TokenPair.CurrentTick0To1)

			if err != nil {
				return err
			}

			if price.Mul(Current1Data.TickData.Reserve1[i]).LT(amount_left) {
				amount_out = amount_out.Add(Current1Data.TickData.Reserve1[i])
				amount_left = amount_left.Sub(price.Mul(Current1Data.TickData.Reserve1[i]))
				Current0Data.TickData.Reserve0AndShares[i].Reserve0 = Current0Data.TickData.Reserve0AndShares[i].Reserve0.Add(price.Mul(Current1Data.TickData.Reserve1[i]))
				Current1Data.TickData.Reserve1[i] = sdk.ZeroDec()

				i++

			} else {
				amount_out = amount_out.Add(amount_left.Mul(price))
				Current0Data.TickData.Reserve0AndShares[i].Reserve0 = Current0Data.TickData.Reserve0AndShares[i].Reserve0.Add(amount_left)
				Current1Data.TickData.Reserve1[i] = Current1Data.TickData.Reserve1[i].Sub(amount_left.Mul(price))
				amount_left = sdk.ZeroDec()
			}

			k.SetTickMap(ctx, pairId, Current0Data)
			k.SetTickMap(ctx, pairId, Current1Data)
		}

		if i == feeSize-1 {
			pair.TokenPair.CurrentTick0To1 = pair.TokenPair.CurrentTick0To1 + 1
		}
	}

	k.SetPairMap(ctx, pair)
	// Todo add slippage tolerance

	if amount_out.LT(minOut) {
		return sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Amount Out is less than minium amount out specified: swap failed")
	}

	if amountIn.GT(sdk.ZeroDec()) {
		coinIn := sdk.NewCoin(token1, sdk.NewIntFromBigInt(amountIn.BigInt()))
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coinIn}); err != nil {
			return err
		}
	} else {
		return sdkerrors.Wrapf(types.ErrNotEnoughCoins, "AmountIn cannot be zero")
	}

	if amount_out.GT(sdk.ZeroDec()) {
		coinOut := sdk.NewCoin(token1, sdk.NewIntFromBigInt(amount_out.BigInt()))
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, callerAddr, sdk.Coins{coinOut}); err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) Swap1to0(goCtx context.Context, msg *types.MsgSwap, token0 string, token1 string, callerAddr sdk.AccAddress, amountIn sdk.Dec, minOut sdk.Dec) error {

	ctx := sdk.UnwrapSDKContext(goCtx)

	pairId := k.CreatePairId(token0, token1)

	feeSize := k.GetFeeListCount(ctx)
	pair, pairFound := k.GetPairMap(ctx, pairId)

	if !pairFound {
		return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Pair not found")
	}

	amount_left := amountIn
	amount_out := sdk.ZeroDec()

	for amount_left.Neg().Equal(sdk.ZeroDec()) {
		Current0Data, _ := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick1To0)
		var i uint64
		i = 0
		for i < feeSize || amount_left.Neg().Equal(sdk.ZeroDec()) {
			fee, _ := k.GetFeeList(ctx, i)
			feeIndex := fee.Fee
			Current1Data, _ := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick1To0-2*feeIndex)
			//Current0Datam := Current0Data.TickData.Reserve1[i]

			price, err := k.Calc_price(pair.TokenPair.CurrentTick1To0)

			if err != nil {
				return err
			}

			if price.Mul(Current0Data.TickData.Reserve0AndShares[i].Reserve0).LT(amount_left) {
				amount_out = amount_out.Add(Current0Data.TickData.Reserve0AndShares[i].Reserve0)
				amount_left = amount_left.Sub(price.Mul(Current0Data.TickData.Reserve0AndShares[i].Reserve0))
				Current1Data.TickData.Reserve1[i] = Current1Data.TickData.Reserve1[i].Add(price.Mul(Current0Data.TickData.Reserve0AndShares[i].Reserve0))
				Current0Data.TickData.Reserve0AndShares[i].Reserve0 = sdk.ZeroDec()

				i++

			} else {
				amount_out = amount_out.Add(amount_left.Mul(price))
				Current1Data.TickData.Reserve1[i] = Current1Data.TickData.Reserve1[i].Add(amount_left)
				Current0Data.TickData.Reserve0AndShares[i].Reserve0 = Current0Data.TickData.Reserve0AndShares[i].Reserve0.Sub(amount_left.Mul(price))
				amount_left = sdk.ZeroDec()
			}

			k.SetTickMap(ctx, pairId, Current0Data)
			k.SetTickMap(ctx, pairId, Current1Data)
		}

		if i == feeSize-1 {
			pair.TokenPair.CurrentTick0To1 = pair.TokenPair.CurrentTick0To1 - 1
		}
	}

	k.SetPairMap(ctx, pair)
	// Todo add slippage tolerance

	if amount_out.LT(minOut) {
		return sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Amount Out is less than minium amount out specified: swap failed")
	}

	if amountIn.GT(sdk.ZeroDec()) {
		coinIn := sdk.NewCoin(token1, sdk.NewIntFromBigInt(amountIn.BigInt()))
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coinIn}); err != nil {
			return err
		}
	} else {
		return sdkerrors.Wrapf(types.ErrNotEnoughCoins, "AmountIn cannot be zero")
	}

	if amount_out.GT(sdk.ZeroDec()) {
		coinOut := sdk.NewCoin(token1, sdk.NewIntFromBigInt(amount_out.BigInt()))
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, callerAddr, sdk.Coins{coinOut}); err != nil {
			return err
		}
	}

	return nil
}
