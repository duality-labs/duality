package keeper

import (
	"context"
	"fmt"
	"math"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) DepositPairInit(goCtx context.Context, token0 string, token1 string, price_index int64, fee int64) (string, error) {

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

		// addEdges(goCtx, token0Index.Index, token1Index.Index)
		k.SetPairMap(ctx, types.PairMap{
			PairId: pairId,
			TokenPair: &types.TokenPairType{
				CurrentTick0To1: price_index - fee,
				CurrentTick1To0: price_index + fee,
			},
			PairCount: 0,
		})

	}
	return pairId, nil

}

func (k Keeper) DepositHelper(goCtx context.Context, pairId string, pair types.PairMap, tickIndex int64, amount0 sdk.Dec, amount1 sdk.Dec, fee int64, feeIndex uint64) (sdk.Dec, sdk.Dec, sdk.Dec, sdk.Dec, sdk.Dec, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	lowerTick, lowerTickFound := k.GetTickMap(ctx, pairId, tickIndex-int64(fee))
	upperTick, upperTickFound := k.GetTickMap(ctx, pairId, tickIndex+int64(fee))

	trueAmount0 := amount0
	trueAmount1 := amount1
	var sharesMinted sdk.Dec

	price, err := k.Calc_price(tickIndex)

	if err != nil {
		return sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec(), err
	}

	if !lowerTickFound || !upperTickFound || upperTick.TickData.Reserve0AndShares[feeIndex].TotalShares.Equal(sdk.ZeroDec()) {
		sharesMinted = trueAmount0.Add(amount1.Mul(price))

		feeSize := k.GetFeeListCount(ctx)
		if !lowerTickFound || !upperTickFound {
			pair.PairCount = pair.PairCount + 1
		}
		if !lowerTickFound {

			lowerTick = types.TickMap{
				TickIndex: tickIndex - fee,
				TickData: &types.TickDataType{
					Reserve0AndShares: make([]*types.Reserve0AndSharesType, feeSize),
					Reserve1:          make([]sdk.Dec, feeSize),
				},
			}

			for i, _ := range lowerTick.TickData.Reserve0AndShares {
				lowerTick.TickData.Reserve0AndShares[i] = &types.Reserve0AndSharesType{sdk.ZeroDec(), sdk.ZeroDec()}

			}
			for i, _ := range lowerTick.TickData.Reserve1 {
				lowerTick.TickData.Reserve1[i] = sdk.ZeroDec()
			}
			//lowerTick = NewTick
		}

		if !upperTickFound {
			upperTick = types.TickMap{
				TickIndex: tickIndex + fee,
				TickData: &types.TickDataType{
					Reserve0AndShares: make([]*types.Reserve0AndSharesType, feeSize),
					Reserve1:          make([]sdk.Dec, feeSize),
				},
			}

			for i, _ := range upperTick.TickData.Reserve0AndShares {
				upperTick.TickData.Reserve0AndShares[i] = &types.Reserve0AndSharesType{sdk.ZeroDec(), sdk.ZeroDec()}

			}
			for i, _ := range upperTick.TickData.Reserve1 {
				upperTick.TickData.Reserve1[i] = sdk.ZeroDec()
			}
			//upperTick = NewTick
		}

		NewReserve0andShares := &types.Reserve0AndSharesType{
			Reserve0:    trueAmount0,
			TotalShares: sharesMinted,
		}

		upperTick.TickData.Reserve0AndShares[feeIndex] = NewReserve0andShares

		lowerTick.TickData.Reserve1[feeIndex] = trueAmount1

	} else {

		// If feeList has been updated via a governance proposal updates future ticks to support this fee tier.
		if uint64(len(upperTick.TickData.Reserve1)) < k.GetFeeListCount(ctx) {
			upperTick.TickData.Reserve1 = append(upperTick.TickData.Reserve1, sdk.ZeroDec())
			upperTick.TickData.Reserve0AndShares = append(upperTick.TickData.Reserve0AndShares, &types.Reserve0AndSharesType{})
		}

		// If feeList has been updated via a governance proposal updates future ticks to support this fee tier.
		if uint64(len(lowerTick.TickData.Reserve1)) < k.GetFeeListCount(ctx) {
			lowerTick.TickData.Reserve1 = append(lowerTick.TickData.Reserve1, sdk.ZeroDec())
			lowerTick.TickData.Reserve0AndShares = append(lowerTick.TickData.Reserve0AndShares, &types.Reserve0AndSharesType{})
		}

		// Balance trueAmount1 to the pool ratio
		if upperTick.TickData.Reserve0AndShares[feeIndex].Reserve0.GT(sdk.ZeroDec()) {
			trueAmount1 = k.Min(amount1, lowerTick.TickData.Reserve1[feeIndex].Mul(amount0).Quo(upperTick.TickData.Reserve0AndShares[feeIndex].Reserve0))
			// trueAmount1 = min(amt1 , (reserve1 * amt0)/reserve0 )
		}

		// Balance trueAmount0 to the pool ratio
		if lowerTick.TickData.Reserve1[feeIndex].GT(sdk.ZeroDec()) {
			trueAmount0 = k.Min(amount0, upperTick.TickData.Reserve0AndShares[feeIndex].Reserve0.Mul(amount1).Quo(lowerTick.TickData.Reserve1[feeIndex]))
		}

		// if amount0 is 0 amt1/reserve1 * totalShares = sharesMinted
		// else if amt0/reserve0 * totalShares
		if trueAmount0.GT(sdk.ZeroDec()) {
			sharesMinted = (trueAmount0.Quo(upperTick.TickData.Reserve0AndShares[feeIndex].Reserve0).Mul(upperTick.TickData.Reserve0AndShares[feeIndex].TotalShares))
		} else {
			sharesMinted = (trueAmount1.Quo(lowerTick.TickData.Reserve1[feeIndex]).Mul(upperTick.TickData.Reserve0AndShares[feeIndex].TotalShares))
		}

		// Adds trueamount0 and sharesMinted to upperTick
		upperTick.TickData.Reserve0AndShares[feeIndex].Reserve0 = upperTick.TickData.Reserve0AndShares[feeIndex].Reserve0.Add(trueAmount0)
		upperTick.TickData.Reserve0AndShares[feeIndex].TotalShares = upperTick.TickData.Reserve0AndShares[feeIndex].TotalShares.Add(sharesMinted)

		// Adds trueamount1 to lowerTick
		lowerTick.TickData.Reserve1[feeIndex] = lowerTick.TickData.Reserve1[feeIndex].Add(trueAmount1)

	}

	// If a new tick has been placed that tigtens the range between currentTick0to1 and currentTick0to1 update CurrentTicks to the tighest ticks
	if trueAmount0.GT(sdk.ZeroDec()) && ((tickIndex+fee > pair.TokenPair.CurrentTick0To1) && (tickIndex+fee < pair.TokenPair.CurrentTick1To0)) {
		pair.TokenPair.CurrentTick1To0 = tickIndex + fee
	}

	if trueAmount1.GT(sdk.ZeroDec()) && ((tickIndex-fee > pair.TokenPair.CurrentTick0To1) && (tickIndex-fee < pair.TokenPair.CurrentTick1To0)) {
		pair.TokenPair.CurrentTick0To1 = tickIndex - fee
	}

	// Set pair, lower and upperTick KVStores
	k.SetPairMap(ctx, pair)
	k.SetTickMap(ctx, pairId, lowerTick)
	k.SetTickMap(ctx, pairId, upperTick)

	return trueAmount0, trueAmount1, sharesMinted, upperTick.TickData.Reserve0AndShares[feeIndex].Reserve0, lowerTick.TickData.Reserve1[feeIndex], nil

}

// FIX ME
// While this works this s a rough get around to doing float based math with sdk.Decs
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

// Returns the smaller of two sdk.Decs
func (k Keeper) Min(a, b sdk.Dec) sdk.Dec {
	if a.LT(b) {
		return a
	}
	return b
}

func (k Keeper) SingleDeposit(goCtx context.Context, msg *types.MsgDeposit, token0 string, token1 string, callerAddr sdk.AccAddress, amount0 sdk.Dec, amount1 sdk.Dec) error {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Note feeIndex is the corresponding index to a fee Value not the feeValue itself
	feeValue, _ := k.GetFeeList(ctx, uint64(msg.FeeIndex))
	fee := feeValue.Fee

	//Checks to see if given pair has been initialied, if not intializes, and returns pairId and pairMap
	pairId, err := k.DepositPairInit(goCtx, token0, token1, msg.TickIndex, fee)

	if err != nil {
		return err
	}

	pair, _ := k.GetPairMap(ctx, pairId)

	fmt.Println(amount0)
	fmt.Println(pair)
	fmt.Println(msg.TickIndex)
	// Errors if depositing amount0 at a tick less than CurrentTick1to0
	if amount0.GT(sdk.ZeroDec()) && ((msg.TickIndex + fee) < pair.TokenPair.CurrentTick0To1) {
		return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Cannot depsoit amount 0 at a tick less than the CurrentTick1to0")
	}

	// Errors if depositing amount1 at a tick less than CurrentTick0to1
	if amount1.GT(sdk.ZeroDec()) && ((msg.TickIndex - fee) > pair.TokenPair.CurrentTick1To0) {
		return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Cannot depsoit amount 0 at a tick greater than the CurrentTick0to1")
	}

	// Calls k.DepositHelper which calculates the true amounts of token0, token1, and sets the corresponding pair and tick maps
	trueAmount0, trueAmount1, sharesMinted, newReserve0, newReserve1, err := k.DepositHelper(goCtx, pairId, pair, msg.TickIndex, amount0, amount1, fee, msg.FeeIndex)

	if err != nil {
		return nil
	}

	// Calculate Shares
	shares, sharesFound := k.GetShares(ctx, msg.Receiver, pairId, msg.TickIndex, msg.FeeIndex)

	// Initializes a new tick if sharesFound does not exists
	if !sharesFound {
		shares = types.Shares{
			Address:     msg.Receiver,
			PairId:      pairId,
			TickIndex:   msg.TickIndex,
			FeeIndex:    msg.FeeIndex,
			SharesOwned: sharesMinted,
		}
	} else {
		//Updates shares.SharesOwned
		shares.SharesOwned = shares.SharesOwned.Add(sharesMinted)
	}

	// Send TrueAmount0 to Module
	/// @Dev Due to a sdk.send constraint this only sends if trueAmount0 is greater than 0
	if trueAmount0.GT(sdk.ZeroDec()) {
		coin0 := sdk.NewCoin(token0, sdk.NewIntFromBigInt(trueAmount0.BigInt()))
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin0}); err != nil {
			return err
		}
	}

	// Send TrueAmount1 to Module
	/// @Dev Due to a sdk.send constraint this only sends if trueAmount1 is greater than 0
	if trueAmount1.GT(sdk.ZeroDec()) {
		coin1 := sdk.NewCoin(token1, sdk.NewIntFromBigInt(trueAmount1.BigInt()))
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin1}); err != nil {
			return err
		}
	}

	// Update share logic to KVStore
	k.SetShares(ctx, shares)

	// Event is defined in types/Events.go
	ctx.EventManager().EmitEvent(types.CreateDepositEvent(msg.Creator, msg.Receiver,
		token0, token1, fmt.Sprint(msg.TickIndex), fmt.Sprint(msg.FeeIndex),
		newReserve0.Sub(trueAmount0).String(), newReserve1.Sub(trueAmount1).String(), newReserve0.String(), newReserve1.String(),
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

	addtick, addtickFound := k.GetTickMap(ctx, pairId, msg.TickIndex+int64(fee))
	subtick, subTickFound := k.GetTickMap(ctx, pairId, msg.TickIndex-int64(fee))

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

	shareOwner, shareOwnerFound := k.GetShares(ctx, msg.Creator, pairId, msg.TickIndex, msg.FeeIndex)

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
	if removeTick {
		pair.PairCount = pair.PairCount - 1
	}
	if removeTick && (msg.TickIndex+int64(fee) == pair.TokenPair.CurrentTick1To0) {

		tickFound := false
		c := 0
		for tickFound != true && pair.PairCount < int64(c) {
			c++
			_, tickFound = k.GetTickMap(ctx, pairId, (msg.TickIndex + fee + int64(c)))

		}

		pair.TokenPair.CurrentTick1To0 = (msg.TickIndex + fee + int64(c))
	}

	if removeTick && (msg.TickIndex-int64(fee) == pair.TokenPair.CurrentTick0To1) {

		tickFound := false
		c := 0
		for tickFound != true && pair.PairCount < int64(c) {
			c++
			_, tickFound = k.GetTickMap(ctx, pairId, (msg.TickIndex - fee - int64(c)))

		}

		pair.TokenPair.CurrentTick1To0 = (msg.TickIndex - fee - int64(c))
	}

	k.SetPairMap(ctx, pair)
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

	ctx.EventManager().EmitEvent(types.CreateWithdrawEvent(msg.Creator, msg.Receiver,
		token0, token1, fmt.Sprint(msg.TickIndex), fmt.Sprint(msg.FeeIndex), OldReserve0.String(), OldReserve1.String(),
		addtick.TickData.Reserve0AndShares[msg.FeeIndex].Reserve0.String(), subtick.TickData.Reserve1[msg.FeeIndex].String(),
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

	count := 0
	amount_left := amountIn
	amount_out := sdk.ZeroDec()

	for !amount_left.Equal(sdk.ZeroDec()) && (count < int(pair.PairCount)) {
		Current1Data, Current1Found := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick0To1)

		count++

		fmt.Println(count)
		if !Current1Found {
			continue
		}

		var i uint64
		i = 0
		for i < feeSize && !amount_left.Equal(sdk.ZeroDec()) {
			fee, _ := k.GetFeeList(ctx, i)
			fmt.Println(fee)
			feeIndex := fee.Fee
			Current0Data, Current0Found := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick0To1+2*feeIndex)
			//Current0Datam := Current0Data.TickData.Reserve1[i]

			if !Current0Found {
				i++
				continue
			}
			price, err := k.Calc_price(pair.TokenPair.CurrentTick0To1)

			if price.Mul(amount_left).Add(amount_out).LT(minOut) {
				return sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Amount Out is less than minium amount out specified: swap failed")
			}

			if err != nil {
				return err
			}

			if price.Mul(Current1Data.TickData.Reserve1[i]).LT(amount_left) {
				amount_out = amount_out.Add(Current1Data.TickData.Reserve1[i])
				amount_left = amount_left.Sub(price.Mul(Current1Data.TickData.Reserve1[i]))
				Current0Data.TickData.Reserve0AndShares[i].Reserve0 = Current0Data.TickData.Reserve0AndShares[i].Reserve0.Add(price.Mul(Current1Data.TickData.Reserve1[i]))
				Current1Data.TickData.Reserve1[i] = sdk.ZeroDec()

			} else {
				amount_out = amount_out.Add(amount_left.Mul(price))
				Current0Data.TickData.Reserve0AndShares[i].Reserve0 = Current0Data.TickData.Reserve0AndShares[i].Reserve0.Add(amount_left)
				Current1Data.TickData.Reserve1[i] = Current1Data.TickData.Reserve1[i].Sub(amount_left.Mul(price))
				amount_left = sdk.ZeroDec()
			}
			i++
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
		coinIn := sdk.NewCoin(token0, sdk.NewIntFromBigInt(amountIn.BigInt()))
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coinIn}); err != nil {
			return err
		}
	} else {
		return sdkerrors.Wrapf(types.ErrNotEnoughCoins, "AmountIn cannot be zero")
	}

	if amount_out.GT(sdk.ZeroDec()) {
		coinOut := sdk.NewCoin(token1, sdk.NewIntFromBigInt(amount_out.BigInt()))
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(msg.Receiver), sdk.Coins{coinOut}); err != nil {
			return err
		}
	}

	ctx.EventManager().EmitEvent(types.CreateSwapEvent(msg.Creator, msg.Receiver,
		token0, token1, msg.TokenIn, amountIn.String(), amount_out.String(), msg.MinOut,
	))

	return nil
}

func (k Keeper) Swap1to0(goCtx context.Context, msg *types.MsgSwap, token0 string, token1 string, callerAddr sdk.AccAddress, amountIn sdk.Dec, minOut sdk.Dec) error {

	ctx := sdk.UnwrapSDKContext(goCtx)

	pairId := k.CreatePairId(token0, token1)

	feeSize := k.GetFeeListCount(ctx)
	pair, pairFound := k.GetPairMap(ctx, pairId)

	fmt.Println("Token0: ", token0)
	if !pairFound {
		return sdkerrors.Wrapf(types.ErrValidPairNotFound, "Pair not found")
	}

	count := 0
	amount_left := amountIn
	amount_out := sdk.ZeroDec()

	for !amount_left.Equal(sdk.ZeroDec()) && (count < int(pair.PairCount)) {

		Current0Data, Current0Found := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick1To0)

		count++
		if !Current0Found {

			continue
		}
		var i uint64
		i = 0
		for i < feeSize && !amount_left.Equal(sdk.ZeroDec()) {
			fee, _ := k.GetFeeList(ctx, i)
			feeIndex := fee.Fee

			Current1Data, Current1Found := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick1To0-2*feeIndex)

			if !Current1Found {
				i++
				continue
			}
			//Current0Datam := Current0Data.TickData.Reserve1[i]

			price, err := k.Calc_price(pair.TokenPair.CurrentTick1To0)

			if err != nil {
				return err
			}

			if price.Mul(amount_left).Add(amount_out).LT(minOut) {
				return sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Amount Out is less than minium amount out specified: swap failed")
			}

			if price.Mul(Current0Data.TickData.Reserve0AndShares[i].Reserve0).LT(amount_left) {
				amount_out = amount_out.Add(Current0Data.TickData.Reserve0AndShares[i].Reserve0)
				amount_left = amount_left.Sub(price.Mul(Current0Data.TickData.Reserve0AndShares[i].Reserve0))
				Current1Data.TickData.Reserve1[i] = Current1Data.TickData.Reserve1[i].Add(price.Mul(Current0Data.TickData.Reserve0AndShares[i].Reserve0))
				Current0Data.TickData.Reserve0AndShares[i].Reserve0 = sdk.ZeroDec()

			} else {
				amount_out = amount_out.Add(amount_left.Mul(price))
				Current1Data.TickData.Reserve1[i] = Current1Data.TickData.Reserve1[i].Add(amount_left)
				Current0Data.TickData.Reserve0AndShares[i].Reserve0 = Current0Data.TickData.Reserve0AndShares[i].Reserve0.Sub(amount_left.Mul(price))
				amount_left = sdk.ZeroDec()
			}

			i++
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

		coinOut := sdk.NewCoin(token0, sdk.NewIntFromBigInt(amount_out.BigInt()))
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(msg.Receiver), sdk.Coins{coinOut}); err != nil {
			return err
		}
	}

	return nil
}
