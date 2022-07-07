package keeper

import (
	"context"
	//"fmt"
	"strconv"
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

	tickOld, tickFound := k.GetTick(
		ctx,
		token0[0],
		token1[0],
		msg.Price,
		msg.Fee,
	)

	if !tickFound {
		return nil, sdkerrors.Wrapf(types.ErrValidTickNotFound, "Valid tick not found")
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

	sharesRemoving, err := strconv.ParseUint(msg.SharesRemoving,10, 64)
	if err != nil {
		return nil, err
	}

	if shareOld.ShareAmount <= sharesRemoving {
		return nil, sdkerrors.Wrapf(types.ErrNotEnoughShares, " Not enough shares are owned by:  %s", msg.Receiver)
	}


	//uint128 amount0Withdraw = uint128(_input.sharesRemoving * tick.reserves0 / tick.totalShares);
    //uint128 amount1Withdraw = uint128(_input.sharesRemoving * tick.reserves1 / tick.totalShares);

	amount0Withdraw := (sharesRemoving * tickOld.Reserves0) / tickOld.TotalShares 
	amount1Withdraw := (sharesRemoving * tickOld.Reserves1) / tickOld.TotalShares

	tickNew := types.Tick{
		Token0:      token0[0],
		Token1:      token1[0],
		Price:       msg.Price,
		Fee:         msg.Fee,
		Reserves0:   tickOld.Reserves0 - amount0Withdraw,
		Reserves1:   tickOld.Reserves1 - amount1Withdraw,
		TotalShares: tickOld.TotalShares - sharesRemoving,
	}


	shareNew := types.Share{
		Owner:       msg.Creator,
		Token0:      token0[0],
		Token1:      token1[0],
		Price:       msg.Price,
		Fee:         msg.Fee,
		ShareAmount: shareOld.ShareAmount - sharesRemoving,
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
	if amount0Withdraw > 0 {
		coin0 := sdk.NewInt64Coin(token0[0], int64(amount0Withdraw))
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coin0}); err != nil {
			return nil, err
		}
	}

	//Token 1
	if amount1Withdraw > 0 {
		coin1 := sdk.NewInt64Coin(token1[0], int64(amount1Withdraw))
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.Coins{coin1}); err != nil {
			return nil, err
		}
	}

	return &types.MsgSingleWithdrawResponse{}, nil
}
