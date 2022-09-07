package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
)

func (k Keeper) SingleDeposit(goCtx context.Context, msg *types.MsgDeposit) error {

	_ = goCtx
	return nil
}

func (k Keeper) MultiDeposit(goCtx context.Context, msg *types.MsgDeposit) error {

	_ = goCtx
	return nil
}
