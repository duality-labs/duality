package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/dex/utils"
)

// NOTE: Currently we are using TruncateInt in multiple places for converting Decs back into sdk.Ints.
// This may create some accounting anomalies but seems preferable to other alternatives.
// See full ADR here: https://www.notion.so/dualityxyz/A-Modest-Proposal-For-Truncating-696a919d59254876a617f82fb9567895

// Handles core logic for MsgDeposit, checking and initializing data structures (tick, pair), calculating
// shares based on amount deposited, and sending funds to moduleAddress.
func (k Keeper) DepositCore(
	goCtx context.Context,
	token0 string,
	token1 string,
	callerAddr sdk.AccAddress,
	receiverAddr sdk.AccAddress,
	amounts0 []sdk.Int,
	amounts1 []sdk.Int,
	tickIndices []int64,
	fees []uint64,
	options []*types.DepositOptions,
) (amounts0Deposit []sdk.Int, amounts1Deposit []sdk.Int, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pairId := CreatePairId(token0, token1)
	totalAmountReserve0 := sdk.ZeroInt()
	totalAmountReserve1 := sdk.ZeroInt()
	amounts0Deposited := make([]sdk.Int, len(amounts0))
	amounts1Deposited := make([]sdk.Int, len(amounts1))
	for i := 0; i < len(amounts0); i++ {
		amounts0Deposited[i] = sdk.ZeroInt()
		amounts1Deposited[i] = sdk.ZeroInt()
	}

	for i, amount0 := range amounts0 {
		amount1 := amounts1[i]
		tickIndex := tickIndices[i]
		fee := fees[i]
		autoswap := options[i].Autoswap

		feeUInt := utils.MustSafeUint64(fee)
		lowerTickIndex := tickIndex - feeUInt
		upperTickIndex := tickIndex + feeUInt

		// behind enemy lines checks
		// TODO: Allow user to deposit "behind enemy lines"
		if amount0.IsPositive() && k.IsBehindEnemyLines(ctx, pairId, pairId.Token0, lowerTickIndex) {
			return nil, nil, types.ErrDepositBehindPairLiquidity
		}
		// TODO: Allow user to deposit "behind enemy lines"
		if amount1.IsPositive() && k.IsBehindEnemyLines(ctx, pairId, pairId.Token1, upperTickIndex) {
			return nil, nil, types.ErrDepositBehindPairLiquidity
		}

		sharesId := CreateSharesId(token0, token1, tickIndex, fee)
		existingShares := k.bankKeeper.GetSupply(ctx, sharesId).Amount

		pool, err := k.GetOrInitPool(
			ctx,
			pairId,
			tickIndex,
			fee,
		)

		if err != nil {
			return nil, nil, err
		}

		inAmount0, inAmount1, outShares := pool.Deposit(amount0, amount1, existingShares, autoswap)

		k.SavePool(ctx, pool)

		if inAmount0.IsZero() && inAmount1.IsZero() {
			return nil, nil, types.ErrZeroTrueDeposit
		}

		if outShares.IsZero() {
			return nil, nil, types.ErrDepositShareUnderflow
		}

		if err := k.MintShares(ctx, receiverAddr, outShares, sharesId); err != nil {
			return nil, nil, err
		}

		amounts0Deposited[i] = inAmount0
		amounts1Deposited[i] = inAmount1
		totalAmountReserve0 = totalAmountReserve0.Add(inAmount0)
		totalAmountReserve1 = totalAmountReserve1.Add(inAmount1)

		ctx.EventManager().EmitEvent(types.CreateDepositEvent(
			callerAddr.String(),
			receiverAddr.String(),
			token0,
			token1,
			fmt.Sprint(tickIndices[i]),
			fmt.Sprint(fees[i]),
			pool.GetLowerReserve0().Sub(inAmount0).String(),
			pool.GetUpperReserve1().Sub(inAmount1).String(),
			pool.GetLowerReserve0().String(),
			pool.GetUpperReserve1().String(),
			outShares.String(),
		))
	}

	if totalAmountReserve0.IsPositive() {
		coin0 := sdk.NewCoin(token0, totalAmountReserve0)
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin0}); err != nil {
			return nil, nil, err
		}
	}

	if totalAmountReserve1.IsPositive() {
		coin1 := sdk.NewCoin(token1, totalAmountReserve1)
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin1}); err != nil {
			return nil, nil, err
		}
	}

	return amounts0Deposited, amounts1Deposited, nil
}

// Handles core logic for MsgWithdrawl; calculating and withdrawing reserve0,reserve1 from a specified tick given a specfied number of shares to remove.
// Calculates the amount of reserve0, reserve1 to withdraw based on the percetange of the desired number of shares to remove compared to the total number of shares at the given tick
func (k Keeper) WithdrawCore(
	goCtx context.Context,
	token0 string,
	token1 string,
	callerAddr sdk.AccAddress,
	receiverAddr sdk.AccAddress,
	sharesToRemoveList []sdk.Int,
	tickIndices []int64,
	fees []uint64,
) error {

	ctx := sdk.UnwrapSDKContext(goCtx)
	pairId := CreatePairId(token0, token1)
	totalReserve0ToRemove := sdk.ZeroInt()
	totalReserve1ToRemove := sdk.ZeroInt()

	for i, fee := range fees {
		sharesToRemove := sharesToRemoveList[i]
		tickIndex := tickIndices[i]

		pool, err := k.GetOrInitPool(ctx, pairId, tickIndex, fee)
		if err != nil {
			return err
		}

		sharesId := CreateSharesId(token0, token1, tickIndex, fee)
		totalShares := k.bankKeeper.GetSupply(ctx, sharesId).Amount

		if totalShares.LT(sharesToRemove) {
			return sdkerrors.Wrapf(types.ErrInsufficientShares, "%s does not have %s shares of type %s", callerAddr, sharesToRemove, sharesId)
		}

		outAmount0, outAmount1 := pool.Withdraw(sharesToRemove, totalShares)
		k.SavePool(ctx, pool)
		if sharesToRemove.IsPositive() {
			if err := k.BurnShares(ctx, callerAddr, sharesToRemove, sharesId); err != nil {
				return err
			}
		}

		totalReserve0ToRemove = totalReserve0ToRemove.Add(outAmount0)
		totalReserve1ToRemove = totalReserve1ToRemove.Add(outAmount1)

		ctx.EventManager().EmitEvent(types.CreateWithdrawEvent(
			callerAddr.String(),
			receiverAddr.String(),
			token0,
			token1,
			fmt.Sprint(tickIndices[i]),
			fmt.Sprint(fees[i]),
			pool.LowerTick0.Reserves.Add(outAmount0).String(),
			pool.UpperTick1.Reserves.Add(outAmount1).String(),
			pool.LowerTick0.Reserves.String(),
			pool.UpperTick1.Reserves.String(),
			sharesToRemove.String(),
		))
	}

	if totalReserve0ToRemove.IsPositive() {
		coin0 := sdk.NewCoin(token0, totalReserve0ToRemove)
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

	// sends totalReserve1ToRemove to receiverAddr
	if totalReserve1ToRemove.IsPositive() {
		coin1 := sdk.NewCoin(token1, totalReserve1ToRemove)
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
func (k Keeper) SwapCore(goCtx context.Context,
	tokenIn string,
	tokenOut string,
	amountIn sdk.Int,
	callerAddr sdk.AccAddress,
	receiverAddr sdk.AccAddress,
) (coinOut sdk.Coin, error error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	pairId, err := CreatePairIdFromUnsorted(tokenIn, tokenOut)
	if err != nil {
		return sdk.Coin{}, err
	}

	amountIn, amountOut, err := k.Swap(ctx, pairId, tokenIn, tokenOut, amountIn, nil)
	if err != nil {
		return sdk.Coin{}, err
	}

	if amountOut.IsZero() {
		return sdk.Coin{}, types.ErrInsufficientLiquidity
	}

	coinIn := sdk.NewCoin(tokenIn, amountIn)
	coinOut = sdk.NewCoin(tokenOut, amountOut)

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coinIn}); err != nil {
		return sdk.Coin{}, err
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coinOut}); err != nil {
		return sdk.Coin{}, err
	}

	ctx.EventManager().EmitEvent(types.CreateSwapEvent(callerAddr.String(), receiverAddr.String(),
		tokenIn, tokenOut, amountIn.String(), amountOut.String()))

	return coinOut, nil
}

// Handles MsgPlaceLimitOrder, initializing (tick, pair) data structures if needed, calculating and storing information for a new limit order at a specific tick
func (k Keeper) PlaceLimitOrderCore(
	goCtx context.Context,
	tokenIn string,
	tokenOut string,
	amountIn sdk.Int,
	tickIndex int64,
	orderType types.LimitOrderType,
	goodTil *time.Time,
	callerAddr sdk.AccAddress,
	receiverAddr sdk.AccAddress,
) (trancheKeyP *string, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pairId, err := CreatePairIdFromUnsorted(tokenIn, tokenOut)
	if err != nil {
		return nil, err
	}

	placeTranche, err := k.GetOrInitPlaceTranche(ctx, pairId, tokenIn, tickIndex, goodTil, orderType)
	trancheKey := placeTranche.TrancheKey
	trancheUser := k.GetOrInitLimitOrderTrancheUser(ctx, pairId, tickIndex, tokenIn, trancheKey, orderType, receiverAddr.String())

	amountLeft, totalIn := amountIn, sdk.ZeroInt()
	// For everything except just-in-time (JIT) orders try to execute as a swap first
	if !orderType.IsJIT() {
		var amountInSwap, amountOutSwap sdk.Int
		limitPrice := placeTranche.PriceMakerToTaker().ToDec()
		amountInSwap, amountOutSwap, err = k.Swap(
			ctx,
			pairId,
			tokenIn,
			tokenOut,
			amountIn,
			&limitPrice,
		)
		if err != nil {
			return nil, err
		}
		totalIn = amountInSwap
		trancheUser.TakerReserves = amountOutSwap
		amountLeft = amountLeft.Sub(amountInSwap)
	}

	if amountLeft.IsPositive() && orderType.IsFoK() {
		return nil, types.ErrFoKLimitOrderNotFilled
	}

	sharesIssued := sdk.ZeroInt()
	// FOR GTC, JIT & GoodTil try to place a maker limitOrder with remaining Amount
	if amountLeft.IsPositive() && (orderType.IsGTC() || orderType.IsJIT() || orderType.IsGoodTil()) {
		placeTranche.PlaceMakerLimitOrder(ctx, amountLeft)
		trancheUser.SharesOwned = trancheUser.SharesOwned.Add(amountLeft)

		if orderType.HasExpiration() {
			goodTilRecord := NewLimitOrderExpiration(placeTranche)
			k.SetLimitOrderExpiration(ctx, goodTilRecord)
			ctx.GasMeter().ConsumeGas(types.ExpiringLimitOrderGas, "Expiring LimitOrder Fee")
		}
		k.SaveTranche(ctx, placeTranche)
		totalIn = totalIn.Add(amountLeft)
		sharesIssued = amountLeft
	}

	k.SaveTrancheUser(ctx, trancheUser)

	if totalIn.IsPositive() {
		coin0 := sdk.NewCoin(tokenIn, totalIn)
		err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin0})
		if err != nil {
			return nil, err
		}
	}
	ctx.EventManager().EmitEvent(types.CreatePlaceLimitOrderEvent(
		callerAddr.String(),
		receiverAddr.String(),
		tokenIn,
		tokenOut,
		totalIn.String(),
		sharesIssued.String(),
		trancheKey,
	))

	return &trancheKey, nil
}

// Handles MsgCancelLimitOrder, removing a specifed number of shares from a limit order and returning the respective amount in terms of the reserve to the user
func (k Keeper) CancelLimitOrderCore(
	goCtx context.Context,
	trancheKey string,
	callerAddr sdk.AccAddress,
) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	trancheUser, foundTrancheUser := k.GetLimitOrderTrancheUser(ctx, callerAddr.String(), trancheKey)
	if !foundTrancheUser {
		return types.ErrActiveLimitOrderNotFound
	}

	pairId, tickIndex, tokenIn := trancheUser.PairId, trancheUser.TickIndex, trancheUser.Token
	tranche, foundTranche := k.GetLimitOrderTranche(ctx, pairId, tokenIn, tickIndex, trancheKey)
	if !foundTranche {
		return types.ErrActiveLimitOrderNotFound
	}

	amountToCancel := tranche.RemoveTokenIn(trancheUser)
	trancheUser.SharesCancelled = trancheUser.SharesCancelled.Add(amountToCancel)

	if amountToCancel.IsPositive() {
		coinOut := sdk.NewCoin(tokenIn, amountToCancel)
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, callerAddr, sdk.Coins{coinOut}); err != nil {
			return err
		}
		k.SaveTrancheUser(ctx, trancheUser)
		k.SaveTranche(ctx, *tranche)
		if trancheUser.OrderType.HasExpiration() {
			k.RemoveLimitOrderExpiration(ctx, *tranche.ExpirationTime, tranche.Ref())
		}

	} else {
		return sdkerrors.Wrapf(types.ErrCancelEmptyLimitOrder, "%s", tranche.TrancheKey)
	}

	ctx.EventManager().EmitEvent(types.CancelLimitOrderEvent(
		callerAddr.String(),
		tokenIn,
		pairId.MustOppositeToken(tokenIn),
		trancheKey,
		amountToCancel.String(),
	))

	return nil
}

// Handles MsgWithdrawFilledLimitOrder, calculates and sends filled liqudity from module to user for a limit order based on amount wished to receive.
func (k Keeper) WithdrawFilledLimitOrderCore(
	goCtx context.Context,
	trancheKey string,
	callerAddr sdk.AccAddress,
) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	trancheUser, foundTrancheUser := k.GetLimitOrderTrancheUser(
		ctx,
		callerAddr.String(),
		trancheKey,
	)
	if !foundTrancheUser {
		return sdkerrors.Wrapf(types.ErrValidLimitOrderTrancheNotFound, "%s", trancheKey)
	}
	pairId, tickIndex, tokenIn := trancheUser.PairId, trancheUser.TickIndex, trancheUser.Token

	tranche, wasFilled, foundTranche := k.FindLimitOrderTranche(ctx, pairId, tickIndex, tokenIn, trancheKey)

	amountOutTokenOut := sdk.ZeroDec()
	remainingTokenIn := sdk.ZeroInt()
	// It's possible that a TrancheUser exists but tranche does not if LO was filled entirely through a swap
	if foundTranche {
		var amountOutTokenIn sdk.Int
		amountOutTokenIn, amountOutTokenOut = tranche.Withdraw(trancheUser)

		if wasFilled {
			//This is only relevant for inactive JIT and GoodTil limit orders
			remainingTokenIn = tranche.RemoveTokenIn(trancheUser)
			k.SaveInactiveTranche(ctx, tranche)
		} else {
			k.SetLimitOrderTranche(ctx, tranche)
		}

		trancheUser.SharesWithdrawn = trancheUser.SharesWithdrawn.Add(amountOutTokenIn)
	}

	takerReserves := trancheUser.WithdrawTakerReserves()
	amountOutTokenOut = amountOutTokenOut.Add(takerReserves.ToDec())

	k.SaveTrancheUser(ctx, trancheUser)

	tokenOut := pairId.MustOppositeToken(tokenIn)
	if amountOutTokenOut.IsPositive() || remainingTokenIn.IsPositive() {
		coinOut := sdk.NewCoin(tokenOut, amountOutTokenOut.TruncateInt())
		coinInRefund := sdk.NewCoin(tokenIn, remainingTokenIn)
		coins := sdk.NewCoins(coinOut, coinInRefund)
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, callerAddr, coins); err != nil {
			return err
		}
	} else {
		return types.ErrWithdrawEmptyLimitOrder
	}

	ctx.EventManager().EmitEvent(types.WithdrawFilledLimitOrderEvent(
		callerAddr.String(),
		tokenIn,
		tokenOut,
		trancheKey,
		amountOutTokenOut.String(),
	))

	return nil
}
