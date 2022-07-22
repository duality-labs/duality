package keeper

import (
	"context"
	//"math/big"
	//"fmt"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SingleDeposit(goCtx context.Context, msg *types.MsgSingleDeposit) (*types.MsgSingleDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	callerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	receiverAddr, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}

	_ = receiverAddr

	AccountsToken0Balance := sdk.NewDecFromInt(k.bankKeeper.GetBalance(ctx, callerAddr, msg.Token0).Amount)

	amount0, err := sdk.NewDecFromStr(msg.Amounts0)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Not a valid decimal type: %s", err)
	}

	amount1, err := sdk.NewDecFromStr(msg.Amounts1)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Not a valid decimal type: %s", err)
	}

	if AccountsToken0Balance.LT(amount0) {
		return nil, sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Address %s  does not have enough of token 0", callerAddr)
	}

	AccountsToken1Balance := sdk.NewDecFromInt(k.bankKeeper.GetBalance(ctx, callerAddr, msg.Token1).Amount)
	if AccountsToken1Balance.LT(amount1) {
		return nil, sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Address %s does not have enough  of token 1", callerAddr)
	}

	token0 := []string{msg.Token0}
	token1 := []string{msg.Token1}
	amounts0 := []sdk.Dec{amount0}
	amounts1 := []sdk.Dec{amount1}

	token0, token1, amounts0, amounts1, error := k.SortTokensDeposit(ctx, token0, token1, amounts0, amounts1)
	amount0 = amounts0[0]
	amount1 = amounts1[0]

	if error != nil {
		return nil, error
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
		shareOld = types.Share{
			Owner:       msg.Receiver,
			Token0:      token0[0],
			Token1:      token1[0],
			Price:       msg.Price,
			Fee:         msg.Fee,
			ShareAmount: sdk.ZeroDec(),
		}
	}

	//fmt.Println("All ticks contracts:", k.GetAllTicks(ctx))
	tickOld, tickFound := k.GetTicks(
		ctx,
		token0[0],
		token1[0],
	)

	price, err := sdk.NewDecFromStr(msg.Price)

	if err != nil {
		return nil, err
	}

	fee, err := sdk.NewDecFromStr(msg.Fee)

	if err != nil {
		return nil, err
	}

	var OneToZeroOld types.Pool
	var ZeroToOneOld types.Pool
	OneToZeroFound := false
	ZeroToOneFound := false

	var SharesMinted sdk.Dec
	var trueAmounts0 = amount0
	var trueAmounts1 = amount1

	if tickFound {

		OneToZeroOld, OneToZeroFound = k.getPool(&tickOld.PoolsOneToZero, fee, price)
		ZeroToOneOld, ZeroToOneFound = k.getPool(&tickOld.PoolsZeroToOne, fee, price)

		if OneToZeroFound {
			trueAmounts0, trueAmounts1, SharesMinted, err = k.depositHelperAdd(&OneToZeroOld, amount0, amount1)

			if err != nil {
				return nil, err
			}
		} else if ZeroToOneFound {
			trueAmounts0, trueAmounts1, SharesMinted, err = k.depositHelperAdd(&ZeroToOneOld, amount0, amount1)

			if err != nil {
				return nil, err
			}

		} else if !OneToZeroFound && !ZeroToOneFound {

			SharesMinted = amount0.Add(amount1.Mul(price))
			// OneToZeroOld = types.Pool {
			// 	Reserve0: sdk.ZeroDec(),
			// 	Reserve1: sdk.ZeroDec(),
			// 	Fee: fee,
			// 	Price: price,
			// 	TotalShares: sdk.ZeroDec(),
			// 	Index: 0,}

			// trueAmounts0, trueAmounts1, SharesMinted, err = k.depositHelperAdd(&OneToZeroOld , amount0, amount1)

			// if err != nil {
			// 	return nil, err
			// }
		}

	} else {
		SharesMinted = amount0.Add(amount1.Mul(price))
	}

	var NewPool types.Pool

	if OneToZeroFound {
		NewPool = types.Pool{
			Reserve0:    OneToZeroOld.Reserve0.Add(trueAmounts0),
			Reserve1:    OneToZeroOld.Reserve1.Add(trueAmounts1),
			Fee:         fee,
			Price:       price,
			TotalShares: OneToZeroOld.TotalShares.Add(SharesMinted),
			Index:       0,
		}
	} else if ZeroToOneFound {
		NewPool = types.Pool{
			Reserve0:    ZeroToOneOld.Reserve0.Add(trueAmounts0),
			Reserve1:    ZeroToOneOld.Reserve1.Add(trueAmounts1),
			Fee:         fee,
			Price:       price,
			TotalShares: ZeroToOneOld.TotalShares.Add(SharesMinted),
			Index:       0,
		}

	} else {
		NewPool = types.Pool{
			Reserve0:    trueAmounts0,
			Reserve1:    trueAmounts1,
			Fee:         fee,
			Price:       price,
			TotalShares: SharesMinted,
			Index:       0,
		}
	}

	//Token 0
	if trueAmounts0.GT(sdk.ZeroDec()) {
		coin0 := sdk.NewCoin(token0[0], sdk.NewIntFromBigInt(trueAmounts0.BigInt()))
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin0}); err != nil {
			return nil, err
		}
	}

	//Token 1
	if trueAmounts1.GT(sdk.ZeroDec()) {
		coin1 := sdk.NewCoin(token1[0], sdk.NewIntFromBigInt(trueAmounts1.BigInt()))
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin1}); err != nil {
			return nil, err
		}
	}

	if ZeroToOneFound {
		k.Update0to1(&tickOld.PoolsZeroToOne, &ZeroToOneOld, NewPool.Reserve0, NewPool.Reserve1, NewPool.Fee, NewPool.TotalShares, NewPool.Price)

	} else if NewPool.Reserve0.GT(sdk.ZeroDec()) && !ZeroToOneFound {
		k.Push0to1(&tickOld.PoolsZeroToOne, &NewPool)
	}

	if OneToZeroFound {
		k.Update1to0(&tickOld.PoolsOneToZero, &OneToZeroOld, NewPool.Reserve0, NewPool.Reserve1, NewPool.Fee, NewPool.TotalShares, NewPool.Price)
	} else if NewPool.Reserve1.GT(sdk.ZeroDec()) && !OneToZeroFound {
		k.Push1to0(&tickOld.PoolsOneToZero, &NewPool)
	}

	tickNew := types.Ticks{
		Token0:         token0[0],
		Token1:         token1[0],
		PoolsZeroToOne: tickOld.PoolsZeroToOne,
		PoolsOneToZero: tickOld.PoolsOneToZero,
	}

	shareNew := types.Share{
		Owner:       msg.Creator,
		Token0:      token0[0],
		Token1:      token1[0],
		Price:       msg.Price,
		Fee:         msg.Fee,
		ShareAmount: shareOld.ShareAmount.Add(SharesMinted),
	}

	k.SetTicks(
		ctx,
		tickNew,
	)

	k.SetShare(
		ctx,
		shareNew,
	)

	var event = sdk.NewEvent(sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, "duality"),
		sdk.NewAttribute(sdk.AttributeKeyAction, types.DepositEventKey),
		sdk.NewAttribute(types.DepositEventCreator, msg.Creator),
		sdk.NewAttribute(types.DepositEventToken0, token0[0]),
		sdk.NewAttribute(types.DepositEventToken1, token1[0]),
		sdk.NewAttribute(types.DepositEventPrice, msg.Price),
		sdk.NewAttribute(types.DepositEventFee, msg.Fee),
		sdk.NewAttribute(types.DepositEventNewReserves0, NewPool.Reserve0.String()),
		sdk.NewAttribute(types.DepositEventNewReserves1, NewPool.Reserve1.String()),
		sdk.NewAttribute(types.DepositEventReceiver, msg.Receiver),
		sdk.NewAttribute(types.DepositEventSharesMinted, SharesMinted.String()),
	)
	ctx.EventManager().EmitEvent(event)

	return &types.MsgSingleDepositResponse{SharesMinted.String()}, nil
}
