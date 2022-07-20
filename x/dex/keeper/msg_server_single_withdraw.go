package keeper

import (
	"context"
	//"fmt"
	//"github.com/holiman/uint256"
	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SingleWithdraw(goCtx context.Context, msg *types.MsgSingleWithdraw) (*types.MsgSingleWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	receiverAddr, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}

	token0 := []string{msg.Token0}
	token1 := []string{msg.Token1}
	token0, token1, error := k.sortTokens(ctx, token0, token1)

	if error != nil {
		return nil, error
	}

	tickOld, tickFound := k.GetTicks(
		ctx,
		token0[0],
		token1[0],
	)

	if !tickFound {
		return nil, sdkerrors.Wrapf(types.ErrValidTickNotFound, "Valid tick not found")
	}
	price, err := sdk.NewDecFromStr(msg.Price) 

	if err != nil {
		return nil, err
	}

	fee, err := sdk.NewDecFromStr(msg.Fee) 

	if err != nil {
		return nil, err
	}

	OneToZeroOld, OneToZeroFound := k.getPool(&tickOld.PoolsOneToZero, fee, price)
	ZeroToOneOld, ZeroToOneFound := k.getPool(&tickOld.PoolsZeroToOne, fee, price)

	if !OneToZeroFound && !ZeroToOneFound {
		return nil, sdkerrors.Wrapf(types.ErrValidTickNotFound, "Valid Pool not found")
	}

	var reserve0 sdk.Dec
	var reserve1 sdk.Dec
	var totalShares sdk.Dec

	if OneToZeroFound {
		reserve0 = OneToZeroOld.Reserve0
		reserve1 = OneToZeroOld.Reserve1
		totalShares = OneToZeroOld.TotalShares
	} else {
		reserve0 = ZeroToOneOld.Reserve0
		reserve1 = ZeroToOneOld.Reserve1
		totalShares = ZeroToOneOld.TotalShares
	}

	shareOld, shareFound := k.GetShare(
		ctx,
		msg.Receiver,
		token0[0],
		token1[0],
		msg.Price,
		msg.Fee,
	)

	if !shareFound {
		return nil, sdkerrors.Wrapf(types.ErrValidShareNotFound, "Valid share not found")
	}


	sharesRemoving, err := sdk.NewDecFromStr(msg.SharesRemoving)
	if err != nil {
		return nil, err
	}

	if sharesRemoving.GT(shareOld.ShareAmount)  {
		return nil, sdkerrors.Wrapf(types.ErrNotEnoughShares, " Not enough shares are owned by:  %s", msg.Receiver)
	}


	//uint128 amount0Withdraw = uint128(_input.sharesRemoving * tick.reserves0 / tick.totalShares);
    //uint128 amount1Withdraw = uint128(_input.sharesRemoving * tick.reserves1 / tick.totalShares);

	amount0Withdraw := (sharesRemoving.Mul(reserve0)).Quo(totalShares)
	amount1Withdraw := (sharesRemoving.Mul(reserve1)).Quo(totalShares)

	NewPool := types.Pool {
		Price: price,
		Fee: fee,
		Reserve0: reserve0.Sub(amount0Withdraw),
		Reserve1: reserve1.Sub(amount1Withdraw),
		TotalShares: totalShares.Sub(sharesRemoving),
		Index: 0,
	}

	if NewPool.Reserve0 == sdk.ZeroDec() && ZeroToOneFound {
		k.Remove0to1(&tickOld.PoolsZeroToOne, ZeroToOneOld.Index)

	} else if NewPool.Reserve0 != sdk.ZeroDec() && ZeroToOneFound {
		k.Update0to1(&tickOld.PoolsZeroToOne, &ZeroToOneOld, NewPool.Reserve0, NewPool.Reserve1, NewPool.Fee, NewPool.TotalShares, NewPool.Price)
	}

	if NewPool.Reserve1 == sdk.ZeroDec() && OneToZeroFound {
		k.Remove1to0(&tickOld.PoolsOneToZero, OneToZeroOld.Index)
	}  else if NewPool.Reserve0 != sdk.ZeroDec() && ZeroToOneFound {
		k.Update0to1(&tickOld.PoolsOneToZero, &OneToZeroOld, NewPool.Reserve0, NewPool.Reserve1, NewPool.Fee, NewPool.TotalShares, NewPool.Price)
	}

	shareNew := types.Share{
		Owner:       msg.Creator,
		Token0:      token0[0],
		Token1:      token1[0],
		Price:       msg.Price,
		Fee:         msg.Fee,
		ShareAmount: totalShares.Sub(sharesRemoving),
	}

	tickNew := types.Ticks{
		Token0: token0[0],
		Token1: token1[0],
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

	//Token 0
	if amount0Withdraw.GT(sdk.ZeroDec()) {
		coin0 := sdk.NewCoin(token0[0], sdk.NewIntFromBigInt(amount0Withdraw.BigInt()) )
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coin0}); err != nil {
			return nil, err
		}
	}

	//Token 1
	if amount1Withdraw.GT(sdk.ZeroDec()) {
		coin1 := sdk.NewCoin(token1[0], sdk.NewIntFromBigInt(amount1Withdraw.BigInt()))
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coin1}); err != nil {
			return nil, err
		}
	}

	var event = sdk.NewEvent(sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, "duality"),
		sdk.NewAttribute(sdk.AttributeKeyAction, types.WithdrawEventKey),
		sdk.NewAttribute(types.WithdrawEventCreator, msg.Creator),
		sdk.NewAttribute(types.WithdrawEventToken0, msg.Token0),
		sdk.NewAttribute(types.WithdrawEventToken1, msg.Token1),
		sdk.NewAttribute(types.WithdrawEventPrice, msg.Price),
		sdk.NewAttribute(types.WithdrawEventFee, msg.Fee ),
		sdk.NewAttribute(types.WithdrawEventOldReserves0, reserve0.String()),
		sdk.NewAttribute(types.WithdrawEventOldReserves1, reserve1.String()),
		sdk.NewAttribute(types.WithdrawEventNewReserves0, NewPool.Reserve0.String()),
		sdk.NewAttribute(types.WithdrawEventNewReserves1, NewPool.Reserve1.String()),
		sdk.NewAttribute(types.WithdrawEventReceiver, msg.Receiver),
		sdk.NewAttribute(types.WithdrawEventAmounts0, amount0Withdraw.String()),
		sdk.NewAttribute(types.WithdrawEventAmounts0, amount1Withdraw.String()),

	)

	ctx.EventManager().EmitEvent(event)

	return &types.MsgSingleWithdrawResponse{amount0Withdraw.String(), amount1Withdraw.String()}, nil
}