package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

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

	// TODO: Do want any other verification checks for creating pair?
	token0, token1, callerAdr, receiverAdr, amounts, price, err := k.AddLiquidityVerification(goCtx, msgAddLP)
	if err != nil {
		return nil, err
	}

	pair, pairFound := k.GetPairs(ctx, token0, token1)

	_ = pair
	if pairFound {
		return nil, sdkerrors.Wrapf(types.ErrValidPairNotFound, "Pair has already been initialized")
	}

	k.SetPairs(ctx, types.Pairs{
		Token0:       token0,
		Token1:       token1,
		TickSpacing:  1,
		CurrentIndex: msg.Index,
		Tickmap:      nil,
		IndexMap:     nil,
	})

	// Get and Set Node for Token0
	NodeToken0, NodeToken0Found := k.GetNodes(ctx, token0)

	if NodeToken0Found {
		NodeToken0.OutgoingEdges = append(NodeToken0.OutgoingEdges, token1)
	} else {
		NodeToken0.OutgoingEdges = []string{token1}
	}

	k.SetNodes(ctx, types.Nodes{token0, NodeToken0.OutgoingEdges})

	// Get and Set Node for Token1
	NodeToken1, NodeToken1Found := k.GetNodes(ctx, token1)

	if NodeToken1Found {
		NodeToken1.OutgoingEdges = append(NodeToken1.OutgoingEdges, token0)
	} else {
		NodeToken1.OutgoingEdges = []string{token0}
	}

	k.SetNodes(ctx, types.Nodes{token1, NodeToken1.OutgoingEdges})

	err = k.SingleDeposit(goCtx, token0, token1, amounts, price, msgAddLP, callerAdr, receiverAdr)

	if err != nil {
		return nil, err
	}
	_ = ctx

	return &types.MsgCreatePairResponse{}, nil
}

func (k msgServer) AddLiquidity(goCtx context.Context, msg *types.MsgAddLiquidity) (*types.MsgAddLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, callerAdr, receiverAdr, amounts, price, err := k.AddLiquidityVerification(goCtx, msg)

	if err != nil {
		return nil, err
	}

	err = k.SingleDeposit(goCtx, token0, token1, amounts, price, msg, callerAdr, receiverAdr)

	if err != nil {
		return nil, err
	}

	_, _, _, _, _, _ = token0, token1, callerAdr, receiverAdr, amounts, ctx

	return &types.MsgAddLiquidityResponse{}, nil
}

func (k msgServer) RemoveLiquidity(goCtx context.Context, msg *types.MsgRemoveLiquidity) (*types.MsgRemoveLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, callerAdr, receiverAdr, shares, price, err := k.RemoveLiquidityVerification(goCtx, msg)

	if err != nil {
		return nil, err
	}

	err = k.SingleWithdraw(goCtx, token0, token1, shares, price, msg, callerAdr, receiverAdr)

	if err != nil {
		return nil, err
	}

	_, _, _, _, _, _ = token0, token1, callerAdr, receiverAdr, shares, ctx

	return &types.MsgRemoveLiquidityResponse{}, nil
}

func (k msgServer) Swap(goCtx context.Context, msg *types.MsgSwap) (*types.MsgSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSwapResponse{}, nil
}
