package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

func (k msgServer) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, createrAddr, amount0, amount1, err := k.depositVerification(goCtx, *msg)

	if err != nil {
		return nil, err
	}

	//TODO add cases for multiDeposit when tickIndex != 1

	//TODO remove msg if not needed
	err = k.DepositCore(goCtx, msg, token0, token1, createrAddr, amount0, amount1)

	if err != nil {
		return nil, err
	}

	_ = ctx

	return &types.MsgDepositResponse{}, nil
}

func (k msgServer) Withdrawl(goCtx context.Context, msg *types.MsgWithdrawl) (*types.MsgWithdrawlResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, createrAddr, err := k.withdrawlVerification(goCtx, *msg)

	if err != nil {
		return nil, err
	}

	err = k.WithdrawCore(goCtx, msg, token0, token1, createrAddr)
	_ = ctx

	return &types.MsgWithdrawlResponse{}, nil
}

func (k msgServer) Swap(goCtx context.Context, msg *types.MsgSwap) (*types.MsgSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, createrAddr, amountIn, minOut, err := k.swapVerification(goCtx, *msg)

	if err != nil {
		return nil, err
	}

	if msg.TokenIn == token0 {
		err = k.Swap0to1(goCtx, msg, token0, token1, createrAddr, amountIn, minOut)
		if err != nil {
			return nil, err
		}
	} else {
		err = k.Swap1to0(goCtx, msg, token0, token1, createrAddr, amountIn, minOut)
		if err != nil {
			return nil, err
		}
	}

	_ = ctx

	return &types.MsgSwapResponse{}, nil
}
