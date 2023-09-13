package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/utils"
	"github.com/duality-labs/duality/x/dex/types"
	dexutils "github.com/duality-labs/duality/x/dex/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) UserDepositsAll(
	goCtx context.Context,
	req *types.QueryAllUserDepositsRequest,
) (*types.QueryAllUserDepositsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	var depositArr []*types.DepositRecord

	pageRes, err := utils.FilteredPaginateAccountBalances(
		ctx,
		k.bankKeeper,
		addr,
		req.Pagination,
		func(poolCoinMaybe sdk.Coin, accumulate bool) bool {
			depositDenom, err := types.NewDepositDenomFromString(poolCoinMaybe.Denom)
			if err != nil {
				return false
			}

			if accumulate {
				depositRecord := &types.DepositRecord{
					PairID:          depositDenom.PairID,
					SharesOwned:     poolCoinMaybe.Amount,
					CenterTickIndex: depositDenom.Tick,
					LowerTickIndex:  depositDenom.Tick - dexutils.MustSafeUint64(depositDenom.Fee),
					UpperTickIndex:  depositDenom.Tick + dexutils.MustSafeUint64(depositDenom.Fee),
					Fee:             depositDenom.Fee,
				}

				depositArr = append(depositArr, depositRecord)
			}

			return true
		})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUserDepositsResponse{
		Deposits:   depositArr,
		Pagination: pageRes,
	}, nil
}
