package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/dex/utils"
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

	if req.Pagination == nil {
		req.Pagination = &query.PageRequest{}
	}

	offset := req.Pagination.Offset
	key := req.Pagination.Key
	limit := req.Pagination.Limit
	countTotal := req.Pagination.CountTotal

	if req.Pagination.Reverse {
		return nil, fmt.Errorf("invalid request, reverse pagination is not enabled")
	}

	if offset > 0 && key != nil {
		return nil, fmt.Errorf("invalid request, either offset or key is expected, got both")
	}

	if limit == 0 {
		limit = query.DefaultLimit

		// count total results when the limit is zero/not supplied
		countTotal = true
	}

	var depositArr []*types.DepositRecord

	// paginate with key
	if len(key) != 0 {
		var (
			startAccum false
			numHits    uint64
			nextKey    []byte
		)
		k.bankKeeper.IterateAccountBalances(ctx, addr, func(coin sdk.Coin) bool {
			if coin.Denom == string(key) {
				startAccum = true
			}
			if numHits == limit {
				nextKey == []byte(coin.Denom)
				return true
			}
			if startAccum {
				depositDenom, err := types.NewDepositDenomFromString(sharesMaybe.Denom)
				if err != nil {
					return false
				}

				numHits++
				depositRecord := &types.DepositRecord{
					PairID:          depositDenom.PairID,
					SharesOwned:     sharesMaybe.Amount,
					CenterTickIndex: depositDenom.Tick,
					LowerTickIndex:  depositDenom.Tick - utils.MustSafeUint64(depositDenom.Fee),
					UpperTickIndex:  depositDenom.Tick + utils.MustSafeUint64(depositDenom.Fee),
					Fee:             depositDenom.Fee,
				}
				depositArr = append(depositArr, depositRecord)

				return false
			}
		})
	} else {
		iterator := getIterator(prefixStore, nil, reverse)
		defer iterator.Close()

		end := offset + limit

		var (
			numHits uint64
			nextKey []byte
		)

		for ; iterator.Valid(); iterator.Next() {
			if iterator.Error() != nil {
				return nil, iterator.Error()
			}

			accumulate := numHits >= offset && numHits < end
			hit, err := onResult(iterator.Key(), iterator.Value(), accumulate)
			if err != nil {
				return nil, err
			}

			if hit {
				numHits++
			}

			if numHits == end+1 {
				if nextKey == nil {
					nextKey = iterator.Key()
				}

				if !countTotal {
					break
				}
			}
		}

		res := &PageResponse{NextKey: nextKey}
		if countTotal {
			res.Total = numHits
		}

		return res, nil

	}

	return &types.QueryAllUserDepositsResponse{
		Deposits: k.GetAllDepositsForAddress(ctx, addr),
	}, nil
}
