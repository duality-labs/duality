package keeper

import (
	"context"
	//"fmt"

	"strconv"

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
	token0, token1, error := k.sortTokens(ctx, token0, token1, amount0, amount1)

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
			ShareAmount: "0",
		}
	}

	tickOld, tickFound := k.GetTicks(
		ctx,
		token0[0],
		token1[0],
	)

	var OneToZeroOld types.Pool
	var ZeroToOneOld types.Pool
	OneToZeroFound := false
	ZeroToOneFound := false

	var SharesMinted sdk.Dec
	var trueAmounts0 = amount0
	var trueAmounts1 = amount1

	if tickFound {
		OnetoZeroOld, OneToZeroFound := k.getPool(&tickOld.PoolsOneToZero, msg.Fee, msg.Price)
		ZeroToOneOld, ZeroToOneFound := k.getPool(&tickOld.PoolsZeroToOne, msg.Fee, msg.Price)

		if OneToZeroFound{
			trueAmounts0, trueAmounts1, SharesMinted, err := k.depositHelperSub(&OnetoZeroOld, amount0, amount1, msg.Fee, msg.Price)

			if err != nil {
				return nil, err
			}
		} else if ZeroToOneFound{
			trueAmounts0, trueAmounts1, SharesMinted, err := k.depositHelperAdd(&OnetoZeroOld, amount0, amount1, msg.Fee, msg.Price)

			if err != nil {
				return nil, err
			}
		}

		if !OneToZeroFound && ZeroToOneFound {

			OnetoZeroOld = types.Pool {
				Reserve0: "0",
				Reserve1: "0",
				Fee: msg.Fee,
				Price: msg.Price,
				TotalShares: "0",
				Index: 0,} 

			trueAmounts0, trueAmounts1, SharesMinted, err := k.depositHelperAdd(&OnetoZeroOld , amount0, amount1, msg.Fee, msg.Price)
		}

	} else {

		price, err := sdk.NewDecFromStr(msg.Price)
		if err != nil {
			return nil, err
		}

		SharesMinted = amount0.Add(amount1.Mul(price))
	}

	var NewPool types.Pool
	if OneToZeroFound {
		NewPool = types.Pool {
			Reserve0: One,
		}
	} else if ZeroToOneFound {

	} else {
		NewPool = types.Pool {
			
		}
	}
	
	//Token 0
	if trueAmounts0.GT(sdk.ZeroDec()) {
		coin0 := sdk.NewCoin(token0[0], sdk.NewIntFromBigInt(trueAmounts0.BigInt()) )
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin0}); err != nil {
			return nil, err
		}
	}
	
	//Token 1
	if trueAmounts1.GT(sdk.ZeroDec()) {
		coin1 := sdk.NewCoin(token1[0], sdk.NewIntFromBigInt(trueAmounts1.BigInt()) )
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin1}); err != nil {
			return nil, err
		}
	}


	tickNew := types.Tick{
		Token0:      token0[0],
		Token1:      token1[0],
		Price:       msg.Price,
		Fee:         msg.Fee,
		Reserves0:   tickOld.Reserves0 + uint64(trueAmounts0),
		Reserves1:   tickOld.Reserves1 + uint64(trueAmounts1),
		TotalShares: tickOld.TotalShares + uint64(SharesMinted),
	}

	shareNew := types.Share{
		Owner:       msg.Creator,
		Token0:      token0[0],
		Token1:      token1[0],
		Price:       msg.Price,
		Fee:         msg.Fee,
		ShareAmount: shareOld.ShareAmount + uint64(SharesMinted),
	}

	
	

	// SharesMinted, TotalAmounts0, TotalAmounts1

	//GetTicks -> ticks.ZeroToOne ticks.OneToZero (OUR PQS) actually []Pool
	//k.Update(&tickOld.ZeroToOne, OneToZeroOld (Our Pool we are revising), token0[0], token1[0], msg.Price, msg.Fee, newResrve0, newReserve1, totalShares)
	//k.Update(&ticksOld.OneToZero)

	k.SetTicks(
		ctx,
		//token0
		//token1
		//&tickOld.ZeroToOne
		//&tickOld.OneToZero
	)


	k.SetShare(
		ctx,
		shareNew,
	)



	
	var event = sdk.NewEvent(sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, "duality"),
		sdk.NewAttribute(sdk.AttributeKeyAction, types.DepositEventKey),
		sdk.NewAttribute(types.DepositEventCreator, msg.Creator),
		sdk.NewAttribute(types.DepositEventToken0, msg.Token0),
		sdk.NewAttribute(types.DepositEventToken1, msg.Token1),
		sdk.NewAttribute(types.DepositEventPrice, msg.Price),
		sdk.NewAttribute(types.DepositEventFee, msg.Fee ),
		sdk.NewAttribute(types.DepositEventOldReserves0, tickOld.Reserves0),
		sdk.NewAttribute(types.DepositEventOldReserves1, tickOld.Reserves1),
		sdk.NewAttribute(types.DepositEventNewReserves0, strconv.FormatUint(uint64(tickNew.Reserves0), 10)),
		sdk.NewAttribute(types.DepositEventNewReserves1, strconv.FormatUint(uint64(tickNew.Reserves1), 10)),
		sdk.NewAttribute(types.DepositEventReceiver, msg.Receiver),
		sdk.NewAttribute(types.DepositEventSharesMinted,SharesMinted.String()),

	)
	ctx.EventManager().EmitEvent(event)


	return &types.MsgSingleDepositResponse{ SharesMinted.String() }, nil
}
