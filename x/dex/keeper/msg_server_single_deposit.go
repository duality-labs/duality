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

	AccountsToken0Balance := k.bankKeeper.GetBalance(ctx, callerAddr, msg.Token0)

	if AccountsToken0Balance.Amount.LT(sdk.NewIntFromUint64(msg.Amounts0)) {
		return nil, sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Address %s  does not have enough of token 0", callerAddr)
	}

	AccountsToken1Balance := k.bankKeeper.GetBalance(ctx, callerAddr, msg.Token1)
	if AccountsToken1Balance.Amount.LT(sdk.NewIntFromUint64(msg.Amounts1)) {
		return nil, sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Address %s does not have enough  of token 1", callerAddr)
	}

	token0 := []string{msg.Token0}
	token1 := []string{msg.Token1}
	token0, token1, error := k.sortTokens(ctx, token0, token1)

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
			ShareAmount: 0,
		}
	}

	tickOld, tickFound := k.GetTick(
		ctx,
		token0[0],
		token1[0],
		msg.Price,
		msg.Fee,
	)

	if !tickFound {
		tickOld = types.Tick{
			Token0:      token0[0],
			Token1:      token1[0],
			Price:       msg.Price,
			Fee:         msg.Fee,
			Reserves0:   0,
			Reserves1:   0,
			TotalShares: 0,
		}
	}

	var SharesMinted uint
	var trueAmounts0 uint = uint(msg.Amounts0)
	var trueAmounts1 uint = uint(msg.Amounts1)

	price, err := strconv.ParseFloat(msg.Price, 64)
	if err != nil {
		return nil, err
	}

	if tickOld.Reserves0 > 0 {
		trueAmounts0 = k.min(uint(msg.Amounts1), uint((uint(tickOld.Reserves1)*uint(msg.Amounts0))/uint(tickOld.Reserves0)))
	}

	if tickOld.Reserves1 > 0 {
		trueAmounts0 = k.min(uint(msg.Amounts0), uint((uint(tickOld.Reserves0)*uint(msg.Amounts1))/uint(tickOld.Reserves1)))
	}

	if trueAmounts0 == uint(msg.Amounts0) && trueAmounts1 != uint(msg.Amounts1) {
		trueAmounts1 = uint(msg.Amounts1) + (((uint(msg.Amounts1) - trueAmounts1) * uint(msg.Fee)) / uint(10000-msg.Fee))
	} else if trueAmounts1 == uint(msg.Amounts1) && trueAmounts0 != uint(msg.Amounts0) {
		trueAmounts0 = uint(msg.Amounts0) + (((uint(msg.Amounts0) - trueAmounts0) * uint(msg.Fee)) / uint(10000-msg.Fee))
	}

	if tickOld.TotalShares == 0 {
		SharesMinted = uint(float64(msg.Amounts0) + float64(msg.Amounts1)*price)
	} else {
		SharesMinted =
			uint(float64(tickOld.TotalShares) * ((float64(msg.Amounts0) + float64(msg.Amounts1)*price) / (float64(tickOld.Reserves0) + float64(tickOld.Reserves1)*price)))
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

	k.SetShare(
		ctx,
		shareNew,
	)

	k.SetTick(
		ctx,
		tickNew,
	)

	//Token 0
	if trueAmounts0 > 0 {
		coin0 := sdk.NewInt64Coin(token0[0], int64(trueAmounts0))
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin0}); err != nil {
			return nil, err
		}
	}
	
	//Token 1
	if trueAmounts1 > 0 {
		coin1 := sdk.NewInt64Coin(token1[0], int64(trueAmounts1))
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, callerAddr, types.ModuleName, sdk.Coins{coin1}); err != nil {
			return nil, err
		}
	}
	

	return &types.MsgSingleDepositResponse{}, nil
}
