package keeper

import (
	"context"
	"fmt"
	"math"
	"strconv"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Handles core logic for MsgDeposit, checking and initializing data structures (tick, pair), calculating shares based on amount deposited, and sending funds to moduleAddress
func (k Keeper) DepositCore(
	goCtx context.Context,
	msg *types.MsgDeposit,
	token0 string,
	token1 string,
	callerAddr sdk.AccAddress,
	amounts0 []sdk.Dec,
	amounts1 []sdk.Dec,
) (amounts0Deposit []sdk.Dec, amounts1Deposit []sdk.Dec, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pair := k.GetOrInitPair(
		goCtx,
		token0,
		token1,
	)
	pairId := pair.PairId
	totalAmountReserve0 := sdk.ZeroDec()
	totalAmountReserve1 := sdk.ZeroDec()
	passedDeposit := 0
	amounts0Deposited := make([]sdk.Dec, len(amounts0))
	amounts1Deposited := make([]sdk.Dec, len(amounts1))
	for i := 0; i < len(amounts0); i++ {
		amounts0Deposited[i] = sdk.ZeroDec()
		amounts1Deposited[i] = sdk.ZeroDec()
	}

	feelist := k.GetAllFeeList(ctx)

	for i, amount0 := range amounts0 {
		amount1 := amounts1[i]
		tickIndex := msg.TickIndexes[i]
		price1To0 := CalcPrice1To0(tickIndex)
		feeIndex := msg.FeeIndexes[i]
		fee := feelist[feeIndex].Fee
		curTick0to1 := pair.TokenPair.CurrentTick0To1
		curTick1to0 := pair.TokenPair.CurrentTick1To0
		lowerTickIndex := tickIndex - fee
		upperTickIndex := tickIndex + fee

		// TODO: Allow user to deposit "behind enemy lines"
		if amount0.GT(sdk.ZeroDec()) && curTick0to1 <= lowerTickIndex {
			return nil, nil, types.ErrDepositBehindPairLiquidity
		}

		// TODO: Allow user to deposit "behind enemy lines"
		if amount1.GT(sdk.ZeroDec()) && upperTickIndex <= curTick1to0 {
			return nil, nil, types.ErrDepositBehindPairLiquidity
		}

		lowerTick := k.GetOrInitTick(goCtx, pairId, lowerTickIndex)
		upperTick := k.GetOrInitTick(goCtx, pairId, upperTickIndex)

		lowerReserve0 := &lowerTick.TickData.Reserve0AndShares[feeIndex].Reserve0
		lowerTotalShares := &lowerTick.TickData.Reserve0AndShares[feeIndex].TotalShares
		upperReserve1 := &upperTick.TickData.Reserve1[feeIndex]

		trueAmount0, trueAmount1, sharesMinted := CalcTrueAmounts(
			price1To0,
			*lowerReserve0,
			*upperReserve1,
			amount0,
			amount1,
			*lowerTotalShares,
		)

		if trueAmount0.Equal(sdk.ZeroDec()) && trueAmount1.Equal(sdk.ZeroDec()) {
			ctx.EventManager().EmitEvent(types.CreateDepositFailedEvent(
				msg.Creator,
				msg.Receiver,
				token0,
				token1,
				fmt.Sprint(tickIndex),
				fmt.Sprint(tickIndex),
				lowerReserve0.String(),
				upperReserve1.String(),
				amount0.String(),
				amount1.String(),
			))
			continue
		}

		*lowerReserve0 = lowerReserve0.Add(trueAmount0)
		*lowerTotalShares = lowerTotalShares.Add(sharesMinted)
		*upperReserve1 = upperReserve1.Add(trueAmount1)
		k.SetPairMap(ctx, pair)
		k.SetTickMap(ctx, pairId, lowerTick)
		k.SetTickMap(ctx, pairId, upperTick)

		k.Logger(ctx).Error("AFTER SET PAIR AND TICK MAP")

		k.UpdateTickPointersPostAddToken0(goCtx, &pair, &lowerTick)

		k.Logger(ctx).Error("AFTER UPDATE TICK POINTERS TOKEN0")

		k.UpdateTickPointersPostAddToken1(goCtx, &pair, &upperTick)

		k.Logger(ctx).Error("AFTER UPDATE TICK POINTERS")

		amounts0Deposited[i] = trueAmount0
		amounts1Deposited[i] = trueAmount1

		passedDeposit++

		shares, sharesFound := k.GetShares(ctx, msg.Receiver, pairId, tickIndex, feeIndex)
		if !sharesFound {
			shares = types.Shares{
				Address:     msg.Receiver,
				PairId:      pairId,
				TickIndex:   tickIndex,
				FeeIndex:    feeIndex,
				SharesOwned: sharesMinted,
			}
		} else {
			shares.SharesOwned = shares.SharesOwned.Add(sharesMinted)
		}

		k.SetShares(ctx, shares)

		totalAmountReserve0 = totalAmountReserve0.Add(trueAmount0)
		totalAmountReserve1 = totalAmountReserve1.Add(trueAmount1)

		ctx.EventManager().EmitEvent(types.CreateDepositEvent(
			msg.Creator,
			msg.Receiver,
			token0,
			token1,
			fmt.Sprint(msg.TickIndexes[i]),
			fmt.Sprint(msg.FeeIndexes[i]),
			lowerReserve0.Sub(trueAmount0).String(),
			upperReserve1.Sub(trueAmount1).String(),
			lowerReserve0.String(),
			upperReserve1.String(),
			sharesMinted.String(),
		),
		)
	}

	if passedDeposit == 0 {
		return nil, nil, sdkerrors.Wrapf(types.ErrAllDepositsFailed, "All deposits failed")
	}

	if totalAmountReserve0.GT(sdk.ZeroDec()) {
		coin0 := sdk.NewCoin(token0, sdk.NewIntFromBigInt(totalAmountReserve0.BigInt()))
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin0}); err != nil {
			return nil, nil, err
		}
	}

	if totalAmountReserve1.GT(sdk.ZeroDec()) {
		coin1 := sdk.NewCoin(token1, sdk.NewIntFromBigInt(totalAmountReserve1.BigInt()))
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin1}); err != nil {
			return nil, nil, err
		}
	}
	_ = goCtx
	return amounts0Deposited, amounts1Deposited, nil
}

// Handles core logic for MsgWithdrawl; calculating and withdrawing reserve0,reserve1 from a specified tick given a specfied number of shares to remove.
// Calculates the amount of reserve0, reserve1 to withdraw based on the percetange of the desired number of shares to remove compared to the total number of shares at the given tick
func (k Keeper) WithdrawCore(goCtx context.Context, msg *types.MsgWithdrawl, token0 string, token1 string, callerAddr sdk.AccAddress, receiverAddr sdk.AccAddress) error {

	ctx := sdk.UnwrapSDKContext(goCtx)
	pairId := k.CreatePairId(token0, token1)
	pair, found := k.GetPairMap(ctx, pairId)
	if !found {
		return types.ErrValidPairNotFound
	}
	totalReserve0ToRemove := sdk.ZeroDec()
	totalReserve1ToRemove := sdk.ZeroDec()

	for i, feeIndex := range msg.FeeIndexes {
		sharesToRemove := msg.SharesToRemove[i]
		tickIndex := msg.TickIndexes[i]

		shareOwner, found := k.GetShares(
			ctx,
			msg.Creator,
			pairId,
			tickIndex,
			feeIndex,
		)
		if !found {
			return types.ErrValidShareNotFound
		}
		userSharesOwned := &shareOwner.SharesOwned

		feeValue, found := k.GetFeeList(ctx, feeIndex)
		if !found {
			return types.ErrValidFeeIndexNotFound
		}
		fee := feeValue.Fee
		lowerTickIndex := tickIndex - fee
		upperTickIndex := tickIndex + fee
		lowerTick, lowerTickFound := k.GetTickMap(ctx, pairId, lowerTickIndex)
		upperTick, upperTickFound := k.GetTickMap(ctx, pairId, upperTickIndex)
		if !lowerTickFound || !upperTickFound {
			return types.ErrValidTickNotFound
		}

		lowerTickFeeTotalShares := &lowerTick.TickData.Reserve0AndShares[feeIndex].TotalShares
		lowerTickFeeReserve0 := &lowerTick.TickData.Reserve0AndShares[feeIndex].Reserve0
		upperTickFeeReserve1 := &upperTick.TickData.Reserve1[feeIndex]
		if lowerTickFeeTotalShares.Equal(sdk.ZeroDec()) {
			return types.ErrNotEnoughShares
		}

		sharesToRemove = MinDec(sharesToRemove, *userSharesOwned)
		ownershipRatio := sharesToRemove.Quo(*lowerTickFeeTotalShares)
		reserve1ToRemove := ownershipRatio.Mul(*upperTickFeeReserve1)
		reserve0ToRemove := ownershipRatio.Mul(*lowerTickFeeReserve0)

		*lowerTickFeeReserve0 = lowerTickFeeReserve0.Sub(reserve0ToRemove)
		*upperTickFeeReserve1 = upperTickFeeReserve1.Sub(reserve1ToRemove)
		*lowerTickFeeTotalShares = lowerTickFeeTotalShares.Sub(sharesToRemove)
		*userSharesOwned = userSharesOwned.Sub(sharesToRemove)

		totalReserve0ToRemove = totalReserve0ToRemove.Add(reserve0ToRemove)
		totalReserve1ToRemove = totalReserve1ToRemove.Add(reserve1ToRemove)

		k.SetShares(ctx, shareOwner)
		k.SetTickMap(ctx, pairId, upperTick)
		k.SetTickMap(ctx, pairId, lowerTick)

		if totalReserve0ToRemove.GT(sdk.ZeroDec()) {
			k.UpdateTickPointersPostRemoveToken0(goCtx, &pair, &lowerTick)
		}

		if totalReserve1ToRemove.GT(sdk.ZeroDec()) {
			k.UpdateTickPointersPostRemoveToken1(goCtx, &pair, &upperTick)
		}

		ctx.EventManager().EmitEvent(types.CreateWithdrawEvent(
			msg.Creator,
			msg.Receiver,
			token0,
			token1,
			fmt.Sprint(msg.TickIndexes[i]),
			fmt.Sprint(msg.FeeIndexes[i]),
			lowerTickFeeReserve0.Add(reserve0ToRemove).String(),
			upperTickFeeReserve1.Add(reserve1ToRemove).String(),
			lowerTickFeeReserve0.String(),
			upperTickFeeReserve1.String(),
			sharesToRemove.String(),
		))
	}
	k.SetPairMap(ctx, pair)
	if totalReserve0ToRemove.GT(sdk.ZeroDec()) {
		coin0 := sdk.NewCoin(token0, sdk.NewIntFromBigInt(totalReserve0ToRemove.BigInt()))
		err := k.bankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			types.ModuleName,
			receiverAddr,
			sdk.Coins{coin0},
		)
		if err != nil {
			return err
		}
	}

	// sends totalReserve1ToRemove to msg.Receiver
	if totalReserve1ToRemove.GT(sdk.ZeroDec()) {
		coin1 := sdk.NewCoin(token1, sdk.NewIntFromBigInt(totalReserve1ToRemove.BigInt()))
		err := k.bankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			types.ModuleName,
			receiverAddr,
			sdk.Coins{coin1},
		)
		if err != nil {
			return err
		}
	}

	return nil
}

// Handles core logic for the asset 0 to asset1 direction of MsgSwap; faciliates swapping amount0 for some amount of amount1, given a specified pair (token0, token1)
func (k Keeper) Swap0to1(goCtx context.Context, msg *types.MsgSwap, token0 string, token1 string, callerAddr sdk.AccAddress) (sdk.Dec, sdk.Dec, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pairId := k.CreatePairId(token0, token1)
	feeSize := k.GetFeeListCount(ctx)
	feelist := k.GetAllFeeList(ctx)
	pair, pairFound := k.GetPairMap(ctx, pairId)
	if !pairFound {
		return sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrValidPairNotFound, "Pair not found")
	}
	if pair.TokenPair.CurrentTick0To1 == math.MaxInt64 {
		return sdk.ZeroDec(), sdk.ZeroDec(), types.ErrNotEnoughLiquidity
	}

	amount_left := msg.AmountIn
	amount_out := sdk.ZeroDec()

	// verify that amount left is not zero and that there are additional valid ticks to check
	for !amount_left.Equal(sdk.ZeroDec()) && pair.TokenPair.CurrentTick0To1 <= pair.MaxTick {
		Current1Data, Current1Found := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick0To1)
		if !Current1Found {
			pair.TokenPair.CurrentTick0To1++
			continue
		}

		var i uint64 = 0

		for i < feeSize && !amount_left.Equal(sdk.ZeroDec()) {
			fee := feelist[i].Fee
			Current0Data, found := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick0To1-2*fee)
			if !found {
				i++
				continue
			}

			price_0to1 := CalcPrice0To1(pair.TokenPair.CurrentTick0To1)

			if price_0to1.Mul(amount_left).Add(amount_out).LT(msg.MinOut) {
				return sdk.ZeroDec(), sdk.ZeroDec(), types.ErrNotEnoughLiquidity
			}

			if Current1Data.TickData.Reserve1[i].LT(amount_left.Mul(price_0to1)) {
				amount_out = amount_out.Add(Current1Data.TickData.Reserve1[i])
				amountInTemp := Current1Data.TickData.Reserve1[i].Quo(price_0to1)
				amount_left = amount_left.Sub(amountInTemp)
				Current0Data.TickData.Reserve0AndShares[i].Reserve0 = Current0Data.TickData.Reserve0AndShares[i].Reserve0.Add(amountInTemp)
				Current1Data.TickData.Reserve1[i] = sdk.ZeroDec()

			} else {
				amountOutTemp := amount_left.Mul(price_0to1)
				amount_out = amount_out.Add(amountOutTemp)
				Current0Data.TickData.Reserve0AndShares[i].Reserve0 = Current0Data.TickData.Reserve0AndShares[i].Reserve0.Add(amount_left)
				Current1Data.TickData.Reserve1[i] = Current1Data.TickData.Reserve1[i].Sub(amountOutTemp)
				amount_left = sdk.ZeroDec()
			}

			i++

			//Make updates to tickMap containing reserve0/1 data to the KVStore
			k.SetTickMap(ctx, pairId, Current0Data)
			// TODO: Return to this, maybe is receiving the wrong tick
			k.UpdateTickPointersPostAddToken0(goCtx, &pair, &Current0Data)
		}

		k.SetTickMap(ctx, pairId, Current1Data)
		if i == feeSize && amount_left.GT(sdk.ZeroDec()) {
			var err error
			amount_left, amount_out, err = k.SwapLimitOrder0to1(goCtx, pairId, token1, amount_out, amount_left, pair.TokenPair.CurrentTick0To1)
			if err != nil {
				return sdk.ZeroDec(), sdk.ZeroDec(), err
			}
		}
		k.UpdateTickPointersPostRemoveToken1(goCtx, &pair, &Current1Data)
	}

	k.SetPairMap(ctx, pair)

	// Check to see if amount_out meets the threshold of minOut
	if amount_out.LT(msg.MinOut) {
		return sdk.ZeroDec(), sdk.ZeroDec(), types.ErrNotEnoughLiquidity
	}

	ctx.EventManager().EmitEvent(types.CreateSwapEvent(msg.Creator, msg.Receiver,
		token0, token1, msg.TokenIn, msg.AmountIn.String(), amount_out.String(), msg.MinOut.String(),
	))

	// Returns amount_out to keeper/msg.server: Swap
	// @Dev token transfers happen in keeper/msg.server: Swap
	return amount_out, amount_left, nil
}

// Handles core logic for the asset 1 to asset 0 direction of MsgSwap; faciliates swapping amount1 for some amount of amount0, given a specified pair (token0, token1)
func (k Keeper) Swap1to0(goCtx context.Context, msg *types.MsgSwap, token0 string, token1 string, callerAddr sdk.AccAddress) (sdk.Dec, sdk.Dec, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pairId := k.CreatePairId(token0, token1)
	feeSize := k.GetFeeListCount(ctx)
	feelist := k.GetAllFeeList(ctx)
	pair, found := k.GetPairMap(ctx, pairId)
	if !found {
		return sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrValidPairNotFound, "Pair not found")
	}
	if pair.TokenPair.CurrentTick1To0 == math.MinInt64 {
		return sdk.ZeroDec(), sdk.ZeroDec(), types.ErrNotEnoughLiquidity
	}

	amount_left := msg.AmountIn
	amount_out := sdk.ZeroDec()
	for !amount_left.Equal(sdk.ZeroDec()) && pair.TokenPair.CurrentTick1To0 >= pair.MinTick {

		Current0Data, Current0Found := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick1To0)
		if !Current0Found {
			pair.TokenPair.CurrentTick1To0 = pair.TokenPair.CurrentTick1To0 - 1
			continue
		}

		var i uint64 = 0
		for i < feeSize && !amount_left.Equal(sdk.ZeroDec()) {
			fee := feelist[i].Fee

			Current1Data, found := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick1To0+2*fee)
			if !found {
				i++
				continue
			}

			price_1to0 := CalcPrice1To0(pair.TokenPair.CurrentTick1To0)
			if price_1to0.Mul(amount_left).Add(amount_out).LT(msg.MinOut) {
				return sdk.ZeroDec(), sdk.ZeroDec(), types.ErrNotEnoughLiquidity
			}

			// If there is not enough to complete the trade
			if Current0Data.TickData.Reserve0AndShares[i].Reserve0.LT(amount_left.Mul(price_1to0)) {
				amount_out = amount_out.Add(Current0Data.TickData.Reserve0AndShares[i].Reserve0)
				amountInTemp := Current0Data.TickData.Reserve0AndShares[i].Reserve0.Quo(price_1to0)
				amount_left = amount_left.Sub(amountInTemp)
				Current1Data.TickData.Reserve1[i] = Current1Data.TickData.Reserve1[i].Add(amountInTemp)
				Current0Data.TickData.Reserve0AndShares[i].Reserve0 = sdk.ZeroDec()
			} else {
				amountOutTemp := amount_left.Mul(price_1to0)
				amount_out = amount_out.Add(amountOutTemp)
				Current0Data.TickData.Reserve0AndShares[i].Reserve0 = Current0Data.TickData.Reserve0AndShares[i].Reserve0.Sub(amountOutTemp)
				Current1Data.TickData.Reserve1[i] = Current1Data.TickData.Reserve1[i].Add(amount_left)
				amount_left = sdk.ZeroDec()
			}

			i++

			k.SetTickMap(ctx, pairId, Current1Data)
			k.UpdateTickPointersPostAddToken1(goCtx, &pair, &Current1Data)
		}

		k.SetTickMap(ctx, pairId, Current0Data)

		if i == feeSize && amount_left.GT(sdk.ZeroDec()) {
			var err error
			amount_left, amount_out, err = k.SwapLimitOrder1to0(goCtx, pairId, token0, amount_out, amount_left, pair.TokenPair.CurrentTick1To0)

			if err != nil {
				return sdk.ZeroDec(), sdk.ZeroDec(), err
			}
		}
		k.UpdateTickPointersPostRemoveToken0(goCtx, &pair, &Current0Data)
	}

	k.SetPairMap(ctx, pair)

	if amount_out.LT(msg.MinOut) {
		return sdk.ZeroDec(), sdk.ZeroDec(), types.ErrNotEnoughLiquidity
	}

	ctx.EventManager().EmitEvent(types.CreateSwapEvent(msg.Creator, msg.Receiver,
		token0, token1, msg.TokenIn, msg.AmountIn.String(), amount_out.String(), msg.MinOut.String(),
	))

	return amount_out, amount_left, nil
}

// Handles swapping asset 0 for asset 1 through any active limit orders at a specified tick
// Returns amount_out, amount_left, error
func (k Keeper) SwapLimitOrder0to1(goCtx context.Context, pairId string, tokenOut string, amountOut sdk.Dec, amountRemainingTokenIn sdk.Dec, tickIndex int64) (sdk.Dec, sdk.Dec, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	priceInToOut := CalcPrice0To1(tickIndex)
	// TODO: check this out
	priceOutToIn := sdk.OneDec().Quo(priceInToOut)

	tick, tickFound := k.GetTickMap(ctx, pairId, tickIndex)
	if !tickFound {
		return amountRemainingTokenIn, amountOut, nil
	}

	fillTranche := &tick.LimitOrderPool1To0.CurrentLimitOrderKey
	placeTranche := &tick.LimitOrderPool1To0.Count

	reserveData, found := k.GetLimitOrderPoolReserveMap(ctx, pairId, tickIndex, tokenOut, *fillTranche)
	if !found || reserveData.Reserves.Equal(sdk.ZeroDec()) {
		return amountRemainingTokenIn, amountOut, nil
	}
	reservesTokenOut := &reserveData.Reserves

	fillData, found := k.GetLimitOrderPoolFillMap(
		ctx,
		pairId,
		tickIndex,
		tokenOut,
		*fillTranche,
	)
	if !found {
		return sdk.ZeroDec(), sdk.ZeroDec(), fmt.Errorf("should not happen1")
	}
	fillTokenIn := &fillData.FilledReserves

	if reservesTokenOut.LTE(amountRemainingTokenIn.Mul(priceInToOut)) {
		amountOut = amountOut.Add(*reservesTokenOut)
		amountFilledTokenIn := reservesTokenOut.Mul(priceOutToIn)
		amountRemainingTokenIn = amountRemainingTokenIn.Sub(amountFilledTokenIn)
		*reservesTokenOut = sdk.ZeroDec()
		*fillTokenIn = fillTokenIn.Add(amountFilledTokenIn)
		k.SetLimitOrderPoolReserveMap(ctx, reserveData)
		k.SetLimitOrderPoolFillMap(ctx, fillData)

		if *fillTranche != *placeTranche {
			*fillTranche++
			k.SetTickMap(ctx, pairId, tick)

			reserveData, found = k.GetLimitOrderPoolReserveMap(ctx, pairId, tickIndex, tokenOut, *fillTranche)
			fillData, _ = k.GetLimitOrderPoolFillMap(ctx, pairId, tickIndex, tokenOut, *fillTranche)
			if !found {
				return sdk.ZeroDec(), sdk.ZeroDec(), fmt.Errorf("should not happen2")
			}
			reservesTokenOut = &reserveData.Reserves
			fillTokenIn = &fillData.FilledReserves

			if reservesTokenOut.LTE(amountRemainingTokenIn.Mul(priceInToOut)) {
				amountOut = amountOut.Add(*reservesTokenOut)
				amountFilledTokenIn := reservesTokenOut.Mul(priceOutToIn)
				amountRemainingTokenIn = amountRemainingTokenIn.Sub(amountFilledTokenIn)
				*reservesTokenOut = sdk.ZeroDec()
				*fillTokenIn = fillTokenIn.Add(amountFilledTokenIn)
			} else {
				amountFilledTokenOut := amountRemainingTokenIn.Mul(priceInToOut)
				amountOut = amountOut.Add(amountFilledTokenOut)
				*fillTokenIn = fillTokenIn.Add(amountRemainingTokenIn)
				*reservesTokenOut = reservesTokenOut.Sub(amountFilledTokenOut)
				amountRemainingTokenIn = sdk.ZeroDec()
			}
			k.SetLimitOrderPoolReserveMap(ctx, reserveData)
			k.SetLimitOrderPoolFillMap(ctx, fillData)
		}
	} else {
		amountFilledTokenOut := amountRemainingTokenIn.Mul(priceInToOut)
		amountOut = amountOut.Add(amountFilledTokenOut)
		*fillTokenIn = fillTokenIn.Add(amountRemainingTokenIn)
		*reservesTokenOut = reservesTokenOut.Sub(amountFilledTokenOut)
		amountRemainingTokenIn = sdk.ZeroDec()
		k.SetLimitOrderPoolReserveMap(ctx, reserveData)
		k.SetLimitOrderPoolFillMap(ctx, fillData)
	}

	pair, _ := k.GetPairMap(ctx, pairId)
	k.UpdateTickPointersPostRemoveToken1(goCtx, &pair, &tick)

	return amountRemainingTokenIn, amountOut, nil
}

// Handles swapping asset 1 for asset 0 through any active limit orders at a specified tick
// Returns amount_out, amount_left, error
func (k Keeper) SwapLimitOrder1to0(
	goCtx context.Context,
	pairId string,
	tokenOut string,
	amountOut sdk.Dec,
	amountRemainingTokenIn sdk.Dec,
	tickIndex int64,
) (sdk.Dec, sdk.Dec, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	priceInToOut := CalcPrice1To0(tickIndex)
	priceOutToIn := sdk.OneDec().Quo(priceInToOut)

	tick, tickFound := k.GetTickMap(ctx, pairId, tickIndex)
	if !tickFound {
		return amountRemainingTokenIn, amountOut, nil
	}

	fillTranche := &tick.LimitOrderPool0To1.CurrentLimitOrderKey
	placeTranche := &tick.LimitOrderPool0To1.Count

	reserveData, found := k.GetLimitOrderPoolReserveMap(ctx, pairId, tickIndex, tokenOut, *fillTranche)
	if !found || reserveData.Reserves.Equal(sdk.ZeroDec()) {
		return amountRemainingTokenIn, amountOut, nil
	}
	reservesTokenOut := &reserveData.Reserves

	fillData, found := k.GetLimitOrderPoolFillMap(
		ctx,
		pairId,
		tickIndex,
		tokenOut,
		*fillTranche,
	)
	if !found {
		return sdk.ZeroDec(), sdk.ZeroDec(), fmt.Errorf("should not happen3")
	}
	fillTokenIn := &fillData.FilledReserves

	if reservesTokenOut.LTE(amountRemainingTokenIn.Mul(priceInToOut)) {
		amountOut = amountOut.Add(*reservesTokenOut)
		amountFilledTokenIn := reservesTokenOut.Mul(priceOutToIn)
		amountRemainingTokenIn = amountRemainingTokenIn.Sub(amountFilledTokenIn)
		*reservesTokenOut = sdk.ZeroDec()
		*fillTokenIn = fillTokenIn.Add(amountFilledTokenIn)
		k.SetLimitOrderPoolReserveMap(ctx, reserveData)
		k.SetLimitOrderPoolFillMap(ctx, fillData)

		if *fillTranche != *placeTranche {
			*fillTranche++
			k.SetTickMap(ctx, pairId, tick)

			reserveData, found = k.GetLimitOrderPoolReserveMap(ctx, pairId, tickIndex, tokenOut, *fillTranche)
			fillData, _ = k.GetLimitOrderPoolFillMap(ctx, pairId, tickIndex, tokenOut, *fillTranche)
			if !found {
				return sdk.ZeroDec(), sdk.ZeroDec(), fmt.Errorf("should not happen4")
			}
			reservesTokenOut = &reserveData.Reserves
			fillTokenIn = &fillData.FilledReserves

			if reservesTokenOut.LTE(amountRemainingTokenIn.Mul(priceInToOut)) {
				amountOut = amountOut.Add(*reservesTokenOut)
				amountFilledTokenIn := reservesTokenOut.Mul(priceOutToIn)
				amountRemainingTokenIn = amountRemainingTokenIn.Sub(amountFilledTokenIn)
				*reservesTokenOut = sdk.ZeroDec()
				*fillTokenIn = fillTokenIn.Add(amountFilledTokenIn)
			} else {
				amountFilledTokenOut := amountRemainingTokenIn.Mul(priceInToOut)
				amountOut = amountOut.Add(amountFilledTokenOut)
				*fillTokenIn = fillTokenIn.Add(amountRemainingTokenIn)
				*reservesTokenOut = reservesTokenOut.Sub(amountFilledTokenOut)
				amountRemainingTokenIn = sdk.ZeroDec()
			}
			k.SetLimitOrderPoolReserveMap(ctx, reserveData)
			k.SetLimitOrderPoolFillMap(ctx, fillData)
		}
	} else {
		amountFilledTokenOut := amountRemainingTokenIn.Mul(priceInToOut)
		amountOut = amountOut.Add(amountFilledTokenOut)
		fillData.FilledReserves = fillData.FilledReserves.Add(amountRemainingTokenIn)
		reserveData.Reserves = reserveData.Reserves.Sub(amountFilledTokenOut)
		amountRemainingTokenIn = sdk.ZeroDec()
		k.SetLimitOrderPoolReserveMap(ctx, reserveData)
		k.SetLimitOrderPoolFillMap(ctx, fillData)
	}

	pair, _ := k.GetPairMap(ctx, pairId)
	k.UpdateTickPointersPostRemoveToken0(goCtx, &pair, &tick)

	return amountRemainingTokenIn, amountOut, nil
}

// Handles MsgPlaceLimitOrder, initializing (tick, pair) data structures if needed, calculating and storing information for a new limit order at a specific tick
func (k Keeper) PlaceLimitOrderCore(goCtx context.Context, msg *types.MsgPlaceLimitOrder, token0 string, token1 string, callerAddr sdk.AccAddress) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pair := k.GetOrInitPair(goCtx, token0, token1)
	pairId := pair.PairId
	tick := k.GetOrInitTick(goCtx, pair.PairId, msg.TickIndex)

	var fillTrancheIndex *uint64
	var placeTrancheIndex *uint64

	if msg.TokenIn == token0 {
		if msg.TickIndex > pair.TokenPair.CurrentTick0To1 {
			return types.ErrPlaceLimitOrderBehindPairLiquidity
		}
		fillTrancheIndex = &tick.LimitOrderPool0To1.CurrentLimitOrderKey
		placeTrancheIndex = &tick.LimitOrderPool0To1.Count
	} else {
		if msg.TickIndex < pair.TokenPair.CurrentTick1To0 {
			return types.ErrPlaceLimitOrderBehindPairLiquidity
		}
		fillTrancheIndex = &tick.LimitOrderPool1To0.CurrentLimitOrderKey
		placeTrancheIndex = &tick.LimitOrderPool1To0.Count
	}

	FillData := k.GetOrInitTickTrancheFillMap(goCtx, pairId, msg.TickIndex, *placeTrancheIndex, msg.TokenIn)
	ReserveData, UserShareData, TotalSharesData := k.GetOrInitLimitOrderMaps(goCtx, pairId, msg.TickIndex, msg.TokenIn, *fillTrancheIndex, msg.Receiver)
	if FillData.FilledReserves.GT(sdk.ZeroDec()) {
		*placeTrancheIndex++
		if FillData.FilledReserves.Equal(TotalSharesData.TotalShares) {
			*fillTrancheIndex++
		}
		k.SetTickMap(ctx, pairId, tick)
		FillData = k.GetOrInitTickTrancheFillMap(goCtx, pairId, msg.TickIndex, *placeTrancheIndex, msg.TokenIn)
		ReserveData, UserShareData, TotalSharesData = k.GetOrInitLimitOrderMaps(goCtx, pairId, msg.TickIndex, msg.TokenIn, *fillTrancheIndex, msg.Receiver)
	}
	ReserveData.Reserves = ReserveData.Reserves.Add(msg.AmountIn)
	UserShareData.SharesOwned = UserShareData.SharesOwned.Add(msg.AmountIn)
	TotalSharesData.TotalShares = TotalSharesData.TotalShares.Add(msg.AmountIn)

	k.SetLimitOrderPoolFillMap(ctx, FillData)
	k.SetLimitOrderPoolReserveMap(ctx, ReserveData)
	k.SetLimitOrderPoolUserShareMap(ctx, UserShareData)
	k.SetLimitOrderPoolTotalSharesMap(ctx, TotalSharesData)
	k.SetPairMap(ctx, pair)

	if msg.TokenIn == token0 {
		k.UpdateTickPointersPostAddToken0(goCtx, &pair, &tick)
	} else if msg.TokenIn == token1 {
		k.UpdateTickPointersPostAddToken1(goCtx, &pair, &tick)
	}

	if msg.AmountIn.GT(sdk.ZeroDec()) {
		coin0 := sdk.NewCoin(msg.TokenIn, sdk.NewIntFromBigInt(msg.AmountIn.BigInt()))
		err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin0})
		if err != nil {
			return err
		}
	}

	ctx.EventManager().EmitEvent(types.CreatePlaceLimitOrderEvent(msg.Creator, msg.Receiver,
		token0, token1, msg.TokenIn, msg.AmountIn.String(), msg.AmountIn.String(), strconv.Itoa(int(*fillTrancheIndex)),
	))

	return nil
}

// Handles MsgCancelLimitOrder, removing a specifed number of shares from a limit order and returning the respective amount in terms of the reserve to the user
func (k Keeper) CancelLimitOrderCore(goCtx context.Context, msg *types.MsgCancelLimitOrder, token0 string, token1 string, callerAddr sdk.AccAddress, receiverAddr sdk.AccAddress) error {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pairId := k.CreatePairId(token0, token1)
	tick, tickFound := k.GetTickMap(ctx, pairId, msg.TickIndex)
	if !tickFound {
		return sdkerrors.Wrapf(types.ErrValidTickNotFound, "Valid tick not found ")
	}

	UserSharesData, UserSharesDataFound := k.GetLimitOrderPoolUserShareMap(ctx, pairId, msg.TickIndex, msg.KeyToken, msg.Key, msg.Creator)
	ReserveData, ReserveDataFound := k.GetLimitOrderPoolReserveMap(ctx, pairId, msg.TickIndex, msg.KeyToken, msg.Key)
	TotalSharesData, TotalShareDataFound := k.GetLimitOrderPoolTotalSharesMap(ctx, pairId, msg.TickIndex, msg.KeyToken, msg.Key)

	if !UserSharesDataFound || !ReserveDataFound || !TotalShareDataFound {
		return sdkerrors.Wrapf(types.ErrValidLimitOrderMapsNotFound, "UserShareMap not found")
	}

	if msg.SharesOut.GT(UserSharesData.SharesOwned) {
		return sdkerrors.Wrapf(types.ErrCannotWithdrawLimitOrder, "sharesOut is larger than shares Owned at the specified tick")
	}

	amountOut := msg.SharesOut.Mul(ReserveData.Reserves).Quo(TotalSharesData.TotalShares)

	UserSharesData.SharesOwned = UserSharesData.SharesOwned.Sub(msg.SharesOut)
	ReserveData.Reserves = ReserveData.Reserves.Sub(amountOut)
	TotalSharesData.TotalShares = TotalSharesData.TotalShares.Sub(msg.SharesOut)

	k.SetLimitOrderPoolUserShareMap(ctx, UserSharesData)
	k.SetLimitOrderPoolReserveMap(ctx, ReserveData)
	k.SetLimitOrderPoolTotalSharesMap(ctx, TotalSharesData)

	if amountOut.GT(sdk.ZeroDec()) {
		coinOut := sdk.NewCoin(msg.KeyToken, sdk.NewIntFromBigInt(amountOut.BigInt()))
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coinOut}); err != nil {
			return err
		}
	} else {
		return sdkerrors.Wrapf(types.ErrCannotWithdrawLimitOrder, "Cannot cancel liqudity from this limit order at this time")
	}

	ctx.EventManager().EmitEvent(types.CancelLimitOrderEvent(msg.Creator, msg.Receiver,
		token0, token1, msg.KeyToken, strconv.Itoa(int(msg.Key)), amountOut.String(),
	))

	pair, _ := k.GetPairMap(ctx, pairId)
	if msg.KeyToken == token0 {
		k.UpdateTickPointersPostRemoveToken0(goCtx, &pair, &tick)
	} else {
		k.UpdateTickPointersPostRemoveToken1(goCtx, &pair, &tick)
	}

	return nil
}

// Handles MsgWithdrawFilledLimitOrder, calculates and sends filled liqudity from module to user for a limit order based on amount wished to receive.
func (k Keeper) WithdrawFilledLimitOrderCore(
	goCtx context.Context,
	msg *types.MsgWithdrawFilledLimitOrder,
	token0 string,
	token1 string,
	callerAddr sdk.AccAddress,
	receiverAddr sdk.AccAddress,
) error {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pairId := k.CreatePairId(token0, token1)
	tick, found := k.GetTickMap(ctx, pairId, msg.TickIndex)
	if !found {
		return sdkerrors.Wrapf(types.ErrValidTickNotFound, "Valid tick not found ")
	}

	orderTokenIn := msg.KeyToken
	var orderTokenOut string
	if msg.KeyToken == token0 {
		orderTokenOut = token1
	} else {
		orderTokenOut = token0
	}
	trancheIndex := msg.Key
	tickIndex := msg.TickIndex

	FillData, FillDataFound := k.GetLimitOrderPoolFillMap(ctx, pairId, tickIndex, orderTokenIn, trancheIndex)
	UserShareData, UserShareDataFound := k.GetLimitOrderPoolUserShareMap(ctx, pairId, tickIndex, orderTokenIn, trancheIndex, msg.Creator)
	TotalSharesData, TotalSharesDataFound := k.GetLimitOrderPoolTotalSharesMap(ctx, pairId, tickIndex, orderTokenIn, trancheIndex)
	if !FillDataFound || !UserShareDataFound || !TotalSharesDataFound {
		return sdkerrors.Wrapf(types.ErrValidLimitOrderMapsNotFound, "Valid mappings for limit order withdraw not found")
	}
	amountFilledLimitTokenOut := FillData.FilledReserves
	userSharesLimitTokenIn := &UserShareData.SharesOwned
	totalSharesLimitTokenIn := TotalSharesData.TotalShares

	UserSharesWithdrawnData, UserSharesWithdrawnDataFound := k.GetLimitOrderPoolUserSharesWithdrawn(
		ctx,
		pairId,
		tickIndex,
		orderTokenIn,
		trancheIndex,
		msg.Creator,
	)
	if !UserSharesWithdrawnDataFound {
		UserSharesWithdrawnData = types.LimitOrderPoolUserSharesWithdrawn{
			PairId:          pairId,
			TickIndex:       msg.TickIndex,
			Token:           msg.KeyToken,
			Count:           msg.Key,
			Address:         msg.Creator,
			SharesWithdrawn: sdk.ZeroDec(),
		}
	}
	userSharesWithdrawnLimitTokenIn := &UserSharesWithdrawnData.SharesWithdrawn

	var priceLimitInToOut sdk.Dec
	if orderTokenIn == token0 {
		priceLimitInToOut = CalcPrice0To1(tick.TickIndex)
	} else {
		priceLimitInToOut = CalcPrice1To0(tick.TickIndex)
	}
	priceLimitOutToIn := sdk.OneDec().Quo(priceLimitInToOut)
	amountFilledLimitTokenIn := amountFilledLimitTokenOut.Mul(priceLimitOutToIn)
	ratioFilledTokenIn := amountFilledLimitTokenIn.Quo(totalSharesLimitTokenIn)
	sharesFilledTokenIn := userSharesLimitTokenIn.Mul(ratioFilledTokenIn)
	sharesToRemoveTokenIn := sharesFilledTokenIn.Sub(*userSharesWithdrawnLimitTokenIn)
	amountOut := sharesToRemoveTokenIn.Mul(priceLimitInToOut)

	*userSharesWithdrawnLimitTokenIn = userSharesWithdrawnLimitTokenIn.Add(sharesToRemoveTokenIn)
	k.SetLimitOrderPoolUserSharesWithdrawn(ctx, UserSharesWithdrawnData)

	if amountOut.GT(sdk.ZeroDec()) {
		coinOut := sdk.NewCoin(orderTokenOut, sdk.NewIntFromBigInt(amountOut.BigInt()))
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coinOut}); err != nil {
			return err
		}
	} else {
		return sdkerrors.Wrapf(types.ErrCannotWithdrawLimitOrder, "Cannot withdraw additional liqudity from this limit order at this time")
	}

	ctx.EventManager().EmitEvent(types.WithdrawFilledLimitOrderEvent(msg.Creator, msg.Receiver,
		token0, token1, msg.KeyToken, strconv.Itoa(int(msg.Key)), amountOut.String(),
	))

	return nil
}
