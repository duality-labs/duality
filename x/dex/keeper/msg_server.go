package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
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
	callerAddr := sdk.MustAccAddressFromBech32(msg.Creator)
	receiverAddr := sdk.MustAccAddressFromBech32(msg.Receiver)

	// lexographically sort token0, token1
	token0, token1, err := SortTokens(msg.TokenA, msg.TokenB)
	if err != nil {
		return nil, err
	}
	// sort amounts
	amounts0, amounts1 := SortAmounts(msg.TokenA, token0, msg.AmountsA, msg.AmountsB)

	tickIndexes := NormalizeAllTickIndexes(msg.TokenA, token0, msg.TickIndexesAToB)

	Amounts0Deposit, Amounts1Deposit, err := k.DepositCore(
		goCtx,
		token0,
		token1,
		callerAddr,
		receiverAddr,
		amounts0,
		amounts1,
		tickIndexes,
		msg.Fees,
		msg.Options,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgDepositResponse{Amounts0Deposit, Amounts1Deposit}, nil
}

func (k msgServer) Withdrawl(goCtx context.Context, msg *types.MsgWithdrawl) (*types.MsgWithdrawlResponse, error) {
	callerAddr := sdk.MustAccAddressFromBech32(msg.Creator)
	receiverAddr := sdk.MustAccAddressFromBech32(msg.Receiver)

	// lexographically sort token0, token1
	token0, token1, err := SortTokens(msg.TokenA, msg.TokenB)
	if err != nil {
		return nil, err
	}

	tickIndexes := NormalizeAllTickIndexes(msg.TokenA, token0, msg.TickIndexesAToB)

	err = k.WithdrawCore(
		goCtx,
		token0,
		token1,
		callerAddr,
		receiverAddr,
		msg.SharesToRemove,
		tickIndexes,
		msg.Fees,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgWithdrawlResponse{}, nil
}

func (k msgServer) Swap(goCtx context.Context, msg *types.MsgSwap) (*types.MsgSwapResponse, error) {
	callerAddr := sdk.MustAccAddressFromBech32(msg.Creator)
	receiverAddr := sdk.MustAccAddressFromBech32(msg.Receiver)

	coinOut, err := k.SwapCore(goCtx, msg.TokenIn, msg.TokenOut, msg.AmountIn, callerAddr, receiverAddr)
	if err != nil {
		return nil, err
	}

	// TODO: Inconsistent that this is the only response that returns coins instead of ints
	return &types.MsgSwapResponse{CoinOut: coinOut}, nil
}

func (k msgServer) PlaceLimitOrder(goCtx context.Context, msg *types.MsgPlaceLimitOrder) (*types.MsgPlaceLimitOrderResponse, error) {
	callerAddr := sdk.MustAccAddressFromBech32(msg.Creator)
	receiverAddr := sdk.MustAccAddressFromBech32(msg.Receiver)

	token0, _, err := SortTokens(msg.TokenIn, msg.TokenOut)
	if err != nil {
		return &types.MsgPlaceLimitOrderResponse{}, err
	}
	tickIndex := NormalizeTickIndex(msg.TokenIn, token0, msg.TickIndex)

	sdkCtx := sdk.UnwrapSDKContext(goCtx)
	err = msg.ValidateGoodTilExpiration(sdkCtx.BlockTime())
	if err != nil {
		return &types.MsgPlaceLimitOrderResponse{}, err
	}

	trancheKey, err := k.PlaceLimitOrderCore(
		goCtx,
		msg.TokenIn,
		msg.TokenOut,
		msg.AmountIn,
		tickIndex,
		msg.OrderType,
		msg.ExpirationTime,
		callerAddr,
		receiverAddr,
	)
	if err != nil {
		return &types.MsgPlaceLimitOrderResponse{}, err
	}

	return &types.MsgPlaceLimitOrderResponse{TrancheKey: *trancheKey}, nil
}

func (k msgServer) WithdrawFilledLimitOrder(goCtx context.Context, msg *types.MsgWithdrawFilledLimitOrder) (*types.MsgWithdrawFilledLimitOrderResponse, error) {
	callerAddr := sdk.MustAccAddressFromBech32(msg.Creator)

	err := k.WithdrawFilledLimitOrderCore(
		goCtx,
		msg.TrancheKey,
		callerAddr,
	)
	if err != nil {
		return &types.MsgWithdrawFilledLimitOrderResponse{}, err
	}

	return &types.MsgWithdrawFilledLimitOrderResponse{}, nil
}

func (k msgServer) CancelLimitOrder(goCtx context.Context, msg *types.MsgCancelLimitOrder) (*types.MsgCancelLimitOrderResponse, error) {
	callerAddr := sdk.MustAccAddressFromBech32(msg.Creator)

	err := k.CancelLimitOrderCore(
		goCtx,
		msg.TrancheKey,
		callerAddr,
	)
	if err != nil {
		return &types.MsgCancelLimitOrderResponse{}, err
	}

	return &types.MsgCancelLimitOrderResponse{}, nil
}
