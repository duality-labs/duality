package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreatePair(goCtx context.Context, msg *types.MsgCreatePair) (*types.MsgCreatePairResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	msgAddLP := &types.MsgAddLiquidity{
		Creator:        msg.Creator,
		TokenA:         msg.TokenA,
		TokenB:         msg.TokenB,
		TokenDirection: msg.TokenDirection,
		Index:          msg.Index,
		Amount:         msg.Amount,
		Price:          msg.Price,
		Fee:            msg.Fee,
		OrderType:      msg.OrderType,
		Receiver:       msg.Receiver,
	}

	token0, token1, callerAdr, receiverAdr, amounts, price, err := k.AddLiquidityVerification(goCtx, msgAddLP)
	if err != nil {
		return nil, err
	}

	pair, pairFound := k.GetPairs(ctx, token0, token1)

	if pairFound {
		sdkerrors.Wrapf(types.ErrValidPairNotFound, "Pair has already been initialized")
	}

	k.SetPairs(ctx, types.Pairs{
		Token0: token0,
		Token1: token1,
		//FIX Me
		TickSpacing:  0,
		CurrentIndex: msg.Index,
		Tickmap:      nil,
		IndexMap:     nil,
	})

	_ = pair

	err = k.SingleDeposit(goCtx, token0, token1, amounts, price, msgAddLP, callerAdr, receiverAdr)

	if err != nil {
		return nil, err
	}
	_ = ctx

	return &types.MsgCreatePairResponse{}, nil
}
