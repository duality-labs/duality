package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SingleDeposit(goCtx context.Context, token0 string, token1 string, amount sdk.Dec, msg *types.MsgAddLiquidity, callerAdr sdk.AccAddress, receiver sdk.AccAddress) error {

	ctx := sdk.UnwrapSDKContext(goCtx)

	_ = ctx
	return nil
}
