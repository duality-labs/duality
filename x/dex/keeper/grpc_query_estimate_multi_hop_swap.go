package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duality-labs/duality/x/dex/types"
	"github.com/duality-labs/duality/x/dex/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EstimateMultiHopSwap(
	goCtx context.Context,
	req *types.QueryEstimateMultiHopSwapRequest,
) (*types.QueryEstimateMultiHopSwapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	amountIn := req.AmountIn
	routes := req.Routes
	exitLimitPrice := req.ExitLimitPrice
	pickBestRoute := req.PickBestRoute

	ctx := sdk.UnwrapSDKContext(goCtx)
	var routeErrors []error
	initialInCoin := sdk.NewCoin(routes[0].Hops[0], amountIn)
	stepCache := make(map[multihopCacheKey]StepResult)
	var bestRoute struct {
		coinOut sdk.Coin
		route   []string
	}
	bestRoute.coinOut = sdk.Coin{Amount: sdk.ZeroInt()}

	for _, route := range routes {
		routeCoinOut, _, err := k.RunMultihopRoute(
			ctx,
			*route,
			initialInCoin,
			exitLimitPrice,
			stepCache,
		)
		if err != nil {
			routeErrors = append(routeErrors, err)
			continue
		}

		if !pickBestRoute || bestRoute.coinOut.Amount.LT(routeCoinOut.Amount) {
			bestRoute.coinOut = routeCoinOut
			bestRoute.route = route.Hops
		}
		if !pickBestRoute {
			break
		}
	}

	if len(routeErrors) == len(routes) {
		// All routes have failed
		allErr := utils.JoinErrors(types.ErrAllMultiHopRoutesFailed, routeErrors...)

		return nil, allErr
	}

	// NB: Critically, we do not write the best route's buffered state context since this is only an estimate.

	return &types.QueryEstimateMultiHopSwapResponse{CoinOut: bestRoute.coinOut}, nil
}
