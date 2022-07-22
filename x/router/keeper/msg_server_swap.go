package keeper

import (
	"context"
	//"sort"
	"fmt"

	
	"github.com/NicholasDotSol/duality/x/router/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Swap(goCtx context.Context, msg *types.MsgSwap) (*types.MsgSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//k.dexKeeper.GetAllTicks(ctx)

	token0 := []string{msg.TokenIn}
	token1 := []string{msg.TokenOut}
	token0, token1, err := k.dexKeeper.SortTokens(ctx, token0, token1)

	if err != nil {
		return nil, err
	}

	oldTick, tickFound := k.dexKeeper.GetTicks(ctx, token0[0], token1[0])

	if !tickFound {
		return nil, err
	}

	if token0[0] == msg.TokenIn {
		if len(oldTick.PoolsZeroToOne) != 0 {
			fmt.Println(oldTick.PoolsZeroToOne)
		} else {
			fmt.Println("fail")
		}

	} else {
		if len(oldTick.PoolsOneToZero) != 0 {
			fmt.Println(oldTick.PoolsOneToZero)
		} else {
			fmt.Println("fail")
		}
	}

	_ = ctx

	return &types.MsgSwapResponse{}, nil
}
