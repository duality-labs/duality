package keeper

import (
	"context"
	//"fmt"
	//"github.com/holiman/uint256"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Specifications of types.Msg.SingleWithdraw can be found in ../proto/dex/tx.proto
func (k msgServer) SingleWithdraw(goCtx context.Context, msg *types.MsgSingleWithdraw) (*types.MsgSingleWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Converts receiver address (string) to sdk.AccAddress
	receiverAddr, err := sdk.AccAddressFromBech32(msg.Receiver)
	// Error checking for the valid receiver address
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}

	// Sorts token0, token1 for uniformity for use in internal mappings
	token0, token1, error := k.SortTokens(ctx, msg.Token0, msg.Token1)

	// Error handling for sort
	if error != nil {
		return nil, error
	}

	// Checks if a given any given pools of the given token pair exists
	tickOld, tickFound := k.GetTicks(
		ctx,
		token0,
		token1,
	)

	// Error handling if no pair exists
	if !tickFound {
		return nil, sdkerrors.Wrapf(types.ErrValidTickNotFound, "Valid tick not found")
	}

	// Converts msg.Price (string) to sdk.Dec
	price, err := sdk.NewDecFromStr(msg.Price)

	// Error handling for sdk.Dec
	if err != nil {
		return nil, err
	}

	// Converts msg.Fee (string) to sdk.Dec
	fee, err := sdk.NewDecFromStr(msg.Fee)

	// Error handling for sdk.Dec
	if err != nil {
		return nil, err
	}

	// Checks to see whether the given fee/priced pool exists in either priority queue

	OneToZeroOld, OneToZeroFound := k.GetPool(&tickOld.PoolsOneToZero, fee, price)
	ZeroToOneOld, ZeroToOneFound := k.GetPool(&tickOld.PoolsZeroToOne, fee, price)

	// Error handling for when the token pair has been initialized and then emptied, and thus no liquidity shares to be withdrawn at this time.
	if !OneToZeroFound && !ZeroToOneFound {
		return nil, sdkerrors.Wrapf(types.ErrValidTickNotFound, "Valid Pool not found")
	}

	// NewPool variable initializations
	var reserve0 sdk.Dec
	var reserve1 sdk.Dec
	var totalShares sdk.Dec

	/* Checks to see which priority queue has the given price/fee pool and updates the NewPool with the respective state information.
	Functionality for OneToZero and ZeroToOne is the same, however we do not which pool exists upon creation of the message and thus
	must check for each case
	*/
	if OneToZeroFound {
		reserve0 = OneToZeroOld.Reserve0
		reserve1 = OneToZeroOld.Reserve1
		totalShares = OneToZeroOld.TotalShares
	} else {
		reserve0 = ZeroToOneOld.Reserve0
		reserve1 = ZeroToOneOld.Reserve1
		totalShares = ZeroToOneOld.TotalShares
	}

	// Gets number of shares owned for a specified pool of a given address
	shareOld, shareFound := k.GetShare(
		ctx,
		msg.Receiver,
		token0,
		token1,
		msg.Price,
		msg.Fee,
	)

	// If no shares were found, user cannot withdrawn funds from the specified pool
	if !shareFound {
		return nil, sdkerrors.Wrapf(types.ErrValidShareNotFound, "Valid share not found")
	}

	// Converts msg.SharesRemoving (string) to sdk.Dec
	sharesRemoving, err := sdk.NewDecFromStr(msg.SharesRemoving)
	// Error handling for sdk.Dec
	if err != nil {
		return nil, err
	}

	// Error handling to determine user has less than or equal to the number of shares they wish to remove
	if sharesRemoving.GT(shareOld.ShareAmount) {
		return nil, sdkerrors.Wrapf(types.ErrNotEnoughShares, " Not enough shares are owned by:  %s", msg.Receiver)
	}


	// Calculates amount0, amount1 to withdraw
	amount0Withdraw := (sharesRemoving.Mul(reserve0)).Quo(totalShares)
	amount1Withdraw := (sharesRemoving.Mul(reserve1)).Quo(totalShares)

	// Creates an updated pool given tokens/shares removed
	NewPool := types.Pool{
		Price:       price,
		Fee:         fee,
		Reserve0:    reserve0.Sub(amount0Withdraw),
		Reserve1:    reserve1.Sub(amount1Withdraw),
		TotalShares: totalShares.Sub(sharesRemoving),
		Index:       0,
	}

	// If the Pool previously existed and now no longer has any tokens of Reserve1, remove the pool from the priority queue 
	if NewPool.Reserve1.Equal(sdk.ZeroDec()) && ZeroToOneFound {
		k.Remove0to1(&tickOld.PoolsZeroToOne, ZeroToOneOld.Index)

	// If the Pool previously existed and still has some number of tokens of reserve0, update the pool in the priority queue
	} else if !(NewPool.Reserve0.Equal(sdk.ZeroDec() )) && ZeroToOneFound {
		k.Update0to1(&tickOld.PoolsZeroToOne, &ZeroToOneOld, NewPool.Reserve0, NewPool.Reserve1, NewPool.Fee, NewPool.TotalShares, NewPool.Price)
	}

	// If the Pool previously existed and now no longer has any tokens of Reserve0, remove the pool from the priority queue 
	if NewPool.Reserve0.Equal(sdk.ZeroDec()) && OneToZeroFound {
		k.Remove1to0(&tickOld.PoolsOneToZero, OneToZeroOld.Index)

	// If the Pool previously existed and still has some number of tokens of reserve0, update the pool in the priority queue
	} else if !(NewPool.Reserve1.Equal(sdk.ZeroDec())) && OneToZeroFound {
		k.Update1to0(&tickOld.PoolsOneToZero, &OneToZeroOld, NewPool.Reserve0, NewPool.Reserve1, NewPool.Fee, NewPool.TotalShares, NewPool.Price)
	}

	// Initialize a new Share object
	shareNew := types.Share{
		Owner:       msg.Creator,
		Token0:      token0,
		Token1:      token1,
		Price:       msg.Price,
		Fee:         msg.Fee,
		ShareAmount: shareOld.ShareAmount.Sub(sharesRemoving),
	}

	//Intialize a new  tick object (note: changes to PoolsZeroToOne/ PoolsOneToZero are done above by reference)
	tickNew := types.Ticks{
		Token0:         token0,
		Token1:         token1,
		PoolsZeroToOne: tickOld.PoolsZeroToOne,
		PoolsOneToZero: tickOld.PoolsOneToZero,
	}

	k.SetShare(
		ctx,
		shareNew,
	)

	k.SetTicks(
		ctx,
		tickNew,
	)

	//Send amount0 to withdrawl from module to receiver address
	if amount0Withdraw.GT(sdk.ZeroDec()) {
		coin0 := sdk.NewCoin(token0, sdk.NewIntFromBigInt(amount0Withdraw.BigInt()))
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coin0}); err != nil {
			return nil, err
		}
	}

	//Send amount1 to withdrawl from module to receiver address
	if amount1Withdraw.GT(sdk.ZeroDec()) {
		coin1 := sdk.NewCoin(token1, sdk.NewIntFromBigInt(amount1Withdraw.BigInt()))
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coin1}); err != nil {
			return nil, err
		}
	}


	newReserve0 := NewPool.Reserve0
	newReserve1 := NewPool.Reserve1
	ctx.EventManager().EmitEvent(types.CreateWithdrawEvent(msg.Creator,
		token0, token1, price.String(), fee.String(),
		
		newReserve0.Add(amount0Withdraw).String(), newReserve1.Add(amount1Withdraw).String(),
		newReserve0.String(), newReserve1.String(),
		sdk.NewAttribute(types.WithdrawEventAmounts0, amount0Withdraw.String()),
		sdk.NewAttribute(types.WithdrawEventAmounts1, amount1Withdraw.String()),

	))

	return &types.MsgSingleWithdrawResponse{amount0Withdraw.String(), amount1Withdraw.String()}, nil
}
