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
	err = k.SingleDeposit(goCtx, msg, token0, token1, createrAddr, amount0, amount1)

	if err != nil {
		return nil, err
	}

	_ = ctx

	return &types.MsgDepositResponse{}, nil
}

func (k msgServer) Withdrawal(goCtx context.Context, msg *types.MsgWithdrawal) (*types.MsgWithdrawalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, createrAddr, sharesToRemove, err := k.WithdrawalVerification(goCtx, *msg)

	if err != nil {
		return nil, err
	}

	err = k.SingleWithdrawal(goCtx, msg, token0, token1, createrAddr, sharesToRemove)
	_ = ctx

	return &types.MsgWithdrawalResponse{}, nil
}

func (k msgServer) Swap(goCtx context.Context, msg *types.MsgSwap) (*types.MsgSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, createrAddr, amountIn, minOut, err := k.swapVerification(goCtx, *msg)

	if err != nil {
		return nil, err
	}

	if msg.TokenIn == token0 {
		_, err = k.Swap0to1(goCtx, msg, token0, token1, createrAddr, amountIn, minOut)
		if err != nil {
			return nil, err
		}
	} else {
		_, err = k.Swap1to0(goCtx, msg, token0, token1, createrAddr, amountIn, minOut)
		if err != nil {
			return nil, err
		}
	}

	_ = ctx

	return &types.MsgSwapResponse{}, nil
}

// TODO: Add functionality for SwapRoute msg
func (k msgServer) SwapRoute(goCtx context.Context, msg *types.MsgSwap) (*types.MsgSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, createrAddr, amountIn, minOut, err := k.swapVerification(goCtx, *msg)

	if err != nil {
		return nil, err
	}

	_, amountOut := k.SwapDynamicRouter(goCtx, msg, createrAddr, token0, token1, amountIn, minOut)

	print(amountOut)

	_ = ctx

	// TODO: Change to use a new response type (dynamic router)
	return &types.MsgSwapResponse{}, nil
}
