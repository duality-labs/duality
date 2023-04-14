package keeper

import (
	"context"
	"encoding/json"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"

	dextypes "github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/incentives/types"
)

var _ types.QueryServer = QueryServer{}

// QueryServer defines a wrapper around the incentives module keeper providing gRPC method handlers.
type QueryServer struct {
	Keeper
}

// NewQueryServer creates a new QueryServer struct.
func NewQueryServer(k Keeper) QueryServer {
	return QueryServer{Keeper: k}
}

func (q QueryServer) GetModuleStatus(goCtx context.Context, req *types.GetModuleStatusRequest) (*types.GetModuleStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.GetModuleStatusResponse{
		RewardCoins: q.Keeper.GetModuleCoinsToBeDistributed(ctx),
		StakedCoins: q.Keeper.GetModuleStakedCoins(ctx),
		Params:      q.Keeper.GetParams(ctx),
	}, nil
}

func (q QueryServer) GetGaugeByID(goCtx context.Context, req *types.GetGaugeByIDRequest) (*types.GetGaugeByIDResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	gauge, err := q.Keeper.GetGaugeByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &types.GetGaugeByIDResponse{Gauge: gauge}, nil
}

func (q QueryServer) GetGauges(goCtx context.Context, req *types.GetGaugesRequest) (*types.GetGaugesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	var gauges []*types.Gauge
	// Pagination defines pagination for the response
	var pagination *query.PageResponse

	var prefix []byte
	switch req.Filter.Status {
	case types.GetGaugesRequest_Filter_ACTIVE_UPCOMING:
		prefix = types.KeyPrefixGaugeIndex
	case types.GetGaugesRequest_Filter_ACTIVE:
		prefix = types.KeyPrefixGaugeIndexActive
	case types.GetGaugesRequest_Filter_UPCOMING:
		prefix = types.KeyPrefixGaugeIndexUpcoming
	case types.GetGaugesRequest_Filter_FINISHED:
		prefix = types.KeyPrefixGaugeIndexFinished
	default:
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid status filter value")
	}
	pagination, gauges, err := q.filterByPrefixAndDenom(ctx, prefix, req.Filter.Denom, req.Pagination)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.GetGaugesResponse{
		Gauges:     gauges,
		Pagination: pagination,
	}, nil
}

func (q QueryServer) GetStakeByID(goCtx context.Context, req *types.GetStakeByIDRequest) (*types.GetStakeByIDResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	stake, err := q.Keeper.GetStakeByID(ctx, req.StakeId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.GetStakeByIDResponse{Stake: stake}, nil
}

func (q QueryServer) GetStakes(goCtx context.Context, req *types.GetStakesRequest) (*types.GetStakesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	hasOwner := len(req.Owner) > 0
	if !hasOwner {
		// TODO: Verify this protection is necessary
		return nil, status.Error(codes.InvalidArgument, "for performance reasons will not return all stakes")
	}

	owner, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, err
	}

	stakes := q.Keeper.getStakesByAccount(ctx, owner)
	return &types.GetStakesResponse{
		Stakes: stakes,
	}, nil
}

func (q QueryServer) GetFutureRewardEstimate(goCtx context.Context, req *types.GetFutureRewardEstimateRequest) (*types.GetFutureRewardEstimateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	if len(req.Owner) == 0 && len(req.StakeIds) == 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty owner")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	diff := req.EndEpoch - q.Keeper.GetEpochInfo(ctx).CurrentEpoch
	if diff > 365 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "end epoch out of ranges")
	}

	var ownerAddress sdk.AccAddress
	if len(req.Owner) != 0 {
		owner, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			return nil, err
		}
		ownerAddress = owner
	}

	stakes := make(types.Stakes, 0, len(req.StakeIds))
	for _, stakeId := range req.StakeIds {
		stake, err := q.Keeper.GetStakeByID(ctx, stakeId)
		if err != nil {
			return nil, err
		}
		stakes = append(stakes, stake)
	}

	rewards, err := q.Keeper.GetRewardsEstimate(ctx, ownerAddress, stakes, req.EndEpoch)
	if err != nil {
		return nil, err
	}
	return &types.GetFutureRewardEstimateResponse{Coins: rewards}, nil
}

// getGaugeFromIDJsonBytes returns gauges from the json bytes of gaugeIDs.
func (q QueryServer) getGaugeFromIDJsonBytes(ctx sdk.Context, refValue []byte) (types.Gauges, error) {
	gauges := types.Gauges{}
	gaugeIDs := []uint64{}

	err := json.Unmarshal(refValue, &gaugeIDs)
	if err != nil {
		return gauges, err
	}

	for _, gaugeID := range gaugeIDs {
		gauge, err := q.Keeper.GetGaugeByID(ctx, gaugeID)
		if err != nil {
			return types.Gauges{}, err
		}

		gauges = append(gauges, gauge)
	}

	return gauges, nil
}

// filterByPrefixAndDenom filters gauges based on a given key prefix and denom
func (q QueryServer) filterByPrefixAndDenom(ctx sdk.Context, prefixType []byte, denom string, pagination *query.PageRequest) (*query.PageResponse, types.Gauges, error) {
	gauges := types.Gauges{}
	store := ctx.KVStore(q.Keeper.storeKey)
	valStore := prefix.NewStore(store, prefixType)
	depositDenom, err := dextypes.NewDepositDenomFromString(denom)
	if err != nil {
		return nil, nil, err
	}
	lowerTick := depositDenom.Tick - int64(depositDenom.Fee)
	upperTick := depositDenom.Tick + int64(depositDenom.Fee)

	pageRes, err := query.FilteredPaginate(valStore, pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		// this may return multiple gauges at once if two gauges start at the same time.
		// for now this is treated as an edge case that is not of importance
		newGauges, err := q.getGaugeFromIDJsonBytes(ctx, value)
		if err != nil {
			return false, err
		}
		if accumulate {
			if denom != "" {
				for _, gauge := range newGauges {
					if *gauge.DistributeTo.PairID != *depositDenom.PairID {
						return false, nil
					}
					lowerTickInRange := gauge.DistributeTo.StartTick <= lowerTick && lowerTick <= gauge.DistributeTo.EndTick
					upperTickInRange := gauge.DistributeTo.StartTick <= upperTick && upperTick <= gauge.DistributeTo.EndTick
					if !lowerTickInRange || !upperTickInRange {
						return false, nil
					}
					gauges = append(gauges, gauge)
				}
			} else {
				gauges = append(gauges, newGauges...)
			}
		}
		return true, nil
	})
	return pageRes, gauges, err
}
