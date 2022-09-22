package keeper

import (
	"context"
	"sort"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Decide if we want to keep struct
type Route struct {
	// string of tokens
	path  []string
	price sdk.Dec
}

// TODO: Update intermediary pairs to be a KV store only upgradeable by governance

// Route and intermediary routes needed to be added to each other
func getIntermediaryPaths() []Route {
	// Hardcoded for now!!!
	return []Route{
		{
			path: []string{"DUAL"},
		}, {
			path: []string{"USDC"},
		},
		{
			path: []string{"DUAL", "USDC"},
		},
		{
			path: []string{"USDC", "DUAL"},
		},
	}
}

func getRoutes(tokenIn string, tokenOut string) []Route {
	allRoutes := []Route{}
	baseRoute := Route{
		path: []string{tokenIn, tokenOut},
	}
	allRoutes = append(allRoutes, baseRoute)
	intermediaryPaths := getIntermediaryPaths()
	for _, route := range intermediaryPaths {
		newPath := append(append([]string{tokenIn}, route.path[:]...), tokenOut)
		newRoute := Route{
			path: newPath,
		}
		allRoutes = append(allRoutes, newRoute)
	}
	return allRoutes
}

// Get valid routes from routes & calculates the price of given valid routes
func (k Keeper) getValidRoutes(goCtx context.Context, tokenIn string, tokenOut string, routes []Route) ([]Route, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// List of pair ids in order

	validRoutes := []Route{}
	for _, route := range routes {
		price := sdk.NewDec(1)

		isValidRoute := false
		for i := 0; i < len(route.path)-1; i++ {
			// Gets each pair sequentially
			token0, token1, err := k.SortTokens(route.path[i], route.path[i+1])
			if err == nil {
				pairId := k.CreatePairId(token0, token1)
				pair, pairFound := k.GetPairMap(ctx, pairId)

				if pairFound {
					if i == len(route.path)-2 {
						isValidRoute = true
					}
					// Multiply price according to tick
					if route.path[i] == token0 {
						tickPrice, err := k.Calc_price(pair.TokenPair.CurrentTick0To1)
						if err != nil {
							return nil, err
						}
						price = price.Mul(tickPrice)
					} else {
						tickPrice, err := k.Calc_price(pair.TokenPair.CurrentTick1To0)
						if err != nil {
							return nil, err
						}
						price = price.Mul(tickPrice)
					}
				} else {
					break
				}
			} else {
				break
			}

		}
		// If all pairs are valid, add to valid routes
		if isValidRoute {
			route.price = price
			validRoutes = append(validRoutes, route)
		}
	}
	return validRoutes, nil
}

// Assumes all routes are valid
func (k Keeper) updatePrices(goCtx context.Context, tokenIn string, tokenOut string, routes []Route) ([]Route, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	for _, route := range routes {
		price := sdk.NewDec(1)
		for i := 0; i < len(route.path)-1; i++ {
			// Gets each pair sequentially
			token0, token1, err := k.SortTokens(route.path[i], route.path[i+1])
			if err == nil {
				pairId := k.CreatePairId(token0, token1)
				pair, pairFound := k.GetPairMap(ctx, pairId)

				if pairFound {
					// Multiply price according to tick
					if route.path[i] == token0 {
						tickPrice, err := k.Calc_price(pair.TokenPair.CurrentTick0To1)
						if err != nil {
							return nil, err
						}
						price = price.Mul(tickPrice)
					} else {
						tickPrice, err := k.Calc_price(pair.TokenPair.CurrentTick1To0)
						if err != nil {
							return nil, err
						}
						price = price.Mul(tickPrice)
					}
				} else {
					break
				}
			} else {
				break
			}

		}
		// If all pairs are valid, add to valid routes
		route.price = price
	}
	return routes, nil
}

/* TODO: Need to figure out how to compare against updated prices of all routes
// Working theory: Not an issue because we save the price in a variable



*/
func (k Keeper) SwapDynamicRouter(goCtx context.Context, msg *types.MsgSwap, callerAddress sdk.AccAddress, tokenIn string, tokenOut string, amountIn sdk.Dec, minOut sdk.Dec) (sdk.Dec, error) {
	routes, err := k.getValidRoutes(goCtx, tokenIn, tokenOut, getRoutes(tokenIn, tokenOut))
	// No valid routes found! Cannot perform swap
	if len(routes) == 0 {
		return sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrNoValidRoutes, "No valid routes found")
	}

	// Valid routes are incorrect
	if err != nil {
		panic(err)
	}

	amountLeft := amountIn
	totalAmountOut := sdk.ZeroDec()
	// Swap while there is still amountIn
	for amountLeft.GT(sdk.ZeroDec()) {
		sort.SliceStable(routes, func(i, j int) bool {
			return routes[j].price.GT(routes[i].price)
		})
		// Get the best route & the second best price
		bestRoute := routes[0]
		secondBestPrice := routes[1].price

		// ((bestRoute.price*amountLeft)+totalAmountOut) < minOut
		// If the price is no longer good enough (at best) to reach minOut, return
		if (bestRoute.price.Mul(amountLeft).Add(totalAmountOut)).LT(minOut) {
			return sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Amount Out is less than minium amount out specified: swap failed")
		}

		// Swap in 5% chunks across the route (arbitrary)
		amountToSwap := sdk.MinDec(amountIn.QuoInt64(20), amountLeft)

		// Swap the 5% chunk and see what amountOutFromSwap is
		amountOutFromSwap, err := k.swapWhileBestPrice(goCtx, msg, callerAddress, bestRoute, secondBestPrice, amountToSwap)

		if err != nil {
			return sdk.ZeroDec(), err
		}

		// Add amountOut from swap to totalAmountOut
		totalAmountOut = totalAmountOut.Add(amountOutFromSwap)

		// Update prices according to new updates from swap
		routes, err = k.updatePrices(goCtx, tokenIn, tokenOut, routes)

		// Prices failed to Update
		if err != nil {
			return sdk.ZeroDec(), err
		}
	}

	return totalAmountOut, nil
}

// Use same msg type for swapRoute
func (k Keeper) swapWhileBestPrice(goCtx context.Context, msg *types.MsgSwap, callerAddress sdk.AccAddress, bestRoute Route, secondBestPrice sdk.Dec, amountIn sdk.Dec) (amountFromSwap sdk.Dec, err error) {
	amountToSwap := amountIn

	// Passes in the amountOut from the previous pair into the next pair until we swap to the end of route
	for i := 0; i < len(bestRoute.path)-1; i++ {
		// Gets each pair sequentially
		token0, token1, _ := k.SortTokens(bestRoute.path[i], bestRoute.path[i+1])
		if token0 == bestRoute.path[i] {
			// Use sdk.ZeroDec() for minOut as we can set a tighter bound later
			amountToSwap, err = k.Swap0to1(goCtx, msg, token0, token1, callerAddress, amountToSwap, sdk.ZeroDec())
			if err != nil {
				return sdk.ZeroDec(), err
			}

		} else {
			// Use sdk.ZeroDec() for minOut as we can set a tighter bound later
			amountToSwap, err = k.Swap1to0(goCtx, msg, token0, token1, callerAddress, amountToSwap, sdk.ZeroDec())
			if err != nil {
				return sdk.ZeroDec(), err
			}

		}
	}
	return amountToSwap, nil
}
