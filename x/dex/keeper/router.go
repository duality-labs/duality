package keeper

import (
	"context"
	"fmt"
	"math"
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
// Intermediary paths need to be stored in both directions
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

// Uses intermediary paths to create potential routes
func getRoutes(tokenIn string, tokenOut string) []Route {
	allRoutes := []Route{}
	baseRoute := Route{
		path: []string{tokenIn, tokenOut},
	}
	allRoutes = append(allRoutes, baseRoute)
	intermediaryPaths := getIntermediaryPaths()
	for _, route := range intermediaryPaths {
		// Create path: tokenIn -> intermediaryPath -> tokenOut
		newPath := append(append([]string{tokenIn}, route.path[:]...), tokenOut)
		newRoute := Route{
			path: newPath,
		}
		allRoutes = append(allRoutes, newRoute)
	}
	return allRoutes
}

// Validate that routes exist & calculate the price of valid routes
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
						tickPrice, err := k.Calc_price(pair.TokenPair.CurrentTick0To1, false)
						if err != nil {
							return nil, err
						}
						price = price.Mul(tickPrice)
					} else {
						tickPrice, err := k.Calc_price(pair.TokenPair.CurrentTick1To0, true)
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

// Updates the prices across all routes passed in
// Assumes all routes passed in are valid
func (k Keeper) updatePrices(goCtx context.Context, tokenIn string, tokenOut string, routes []Route) ([]Route, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	newRoutes := []Route{}
	for _, route := range routes {
		// No err from route price
		price, _ := k.updateRoutePrice(ctx, route)
		if price.GT(sdk.ZeroDec()) {
			newRoutes = append(newRoutes, Route{route.path, price})
		}
	}
	return newRoutes, nil
}

// Updates price for a specific route
// TODO: Check that route has liquidity!
// Return 0 if route has no liquidity
func (k Keeper) updateRoutePrice(ctx sdk.Context, route Route) (sdk.Dec, error) {
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
					// Checks if there are active ticks
					// If no liquidity at tick, then nothing exists
					// if k.GetTotalReservesAtTick(pairId, pair.TokenPair.CurrentTick0To1, true) > 0 {
					tickPrice, err := k.Calc_price(pair.TokenPair.CurrentTick0To1, false)
					if err != nil {
						return sdk.ZeroDec(), err
					}
					price = price.Mul(tickPrice)

				} else {
					// Checks if there are active ticks
					// TODO: THIS DOES NOT WORK B/C IT DOESN"T CHECK IF THERE ARE NO TICKS ON ONE SIDE
					// if pair.PairCount > 0 {
					tickPrice, err := k.Calc_price(pair.TokenPair.CurrentTick1To0, true)
					if err != nil {
						return sdk.ZeroDec(), err
					}
					price = price.Mul(tickPrice)
					// }
				}
			}
		}
	}
	return price, nil
}

/* TODO: Need to figure out how to compare against updated prices of all routes
// Working theory: Not an issue because we save the price in a variable
DUMMY ALGO
*/
func (k Keeper) SwapDynamicRouter(goCtx context.Context, callerAddress sdk.AccAddress, tokenIn string, tokenOut string, amountIn sdk.Dec, minOut sdk.Dec) (sdk.Dec, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	routes, err := k.getValidRoutes(goCtx, tokenIn, tokenOut, getRoutes(tokenIn, tokenOut))
	// No valid routes found! Cannot perform swap
	if len(routes) == 0 {
		return sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrNoValidRoutes, "No valid routes found")
	}

	// Valid routes failure
	if err != nil {
		return sdk.ZeroDec(), err
	}

	amountLeft := amountIn
	totalAmountOut := sdk.ZeroDec()
	// Swap while there is still amountIn
	for amountLeft.GT(sdk.ZeroDec()) {
		// TODO: Check this works, sort routes by price
		sort.SliceStable(routes, func(i, j int) bool {
			return routes[j].price.GT(routes[i].price)
		})
		fmt.Println("Sorted routes: ", routes)
		// Get the best route & the second best price
		bestRoute := routes[0]
		secondBestPrice := sdk.ZeroDec()
		if len(routes) > 1 {
			secondBestPrice = routes[1].price
		}

		// ((bestRoute.price*amountLeft)+totalAmountOut) < minOut
		// If the price is no longer good enough (at best) to reach minOut, return
		if (bestRoute.price.Mul(amountLeft).Add(totalAmountOut)).LT(minOut) {
			return sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Amount Out is less than minium amount out specified: swap failed")
		}

		amountToSwap := sdk.MinDec(amountIn.QuoInt64(20), amountLeft)
		minOutToSwitchRoutes := secondBestPrice.Mul(amountLeft).Add(totalAmountOut)

		fmt.Println("BestRoute:  ", bestRoute, "AmountToSwap: ", amountToSwap)
		// Swap the 5% chunk and see what amountOutFromSwap is
		amountOutFromSwap, err := k.swapAcrossRoute(goCtx, callerAddress, bestRoute, amountToSwap, minOutToSwitchRoutes)

		if err != nil {
			return sdk.ZeroDec(), err
		}
		// Subtract amountToSwap from amountLeft
		amountLeft = amountLeft.Sub(amountToSwap)

		// Add amountOutFromSwap to totalAmountOut
		totalAmountOut = totalAmountOut.Add(amountOutFromSwap)
		fmt.Println("totalAmountOut: ", totalAmountOut)

		// Update the route price for the best route
		updatedPrice, err := k.updateRoutePrice(ctx, bestRoute)
		if err != nil {
			return sdk.ZeroDec(), err
		}
		bestRoute.price = updatedPrice

		routes, err = k.updatePrices(goCtx, tokenIn, tokenOut, routes)

		// Prices failed to Update
		if err != nil {
			return sdk.ZeroDec(), err
		}
	}

	return totalAmountOut, nil
}

// ORDERED LIST OF ROUTES, AMOUNT YOU WANT TO SWAP THROUGH THEM
// Core function, can use any arbitrary SwapDynamicRouter
// i.e. Optimal Routing from Bain, BF, etc.
func (k Keeper) swapAcrossRoute(goCtx context.Context, callerAddress sdk.AccAddress, bestRoute Route, amountIn sdk.Dec, minOut sdk.Dec) (amountFromSwap sdk.Dec, err error) {
	amountToSwap := amountIn
	amountLeft := sdk.ZeroDec()
	// Passes in the amountOut from the previous pair into the next pair until we swap to the end of route
	var amountOut sdk.Dec
	for i := 0; i < len(bestRoute.path)-1; i++ {
		// Gets each pair sequentially
		token0, token1, _ := k.SortTokens(bestRoute.path[i], bestRoute.path[i+1])
		fmt.Println("Token0: ", token0, "Token1: ", token1)
		if token0 == bestRoute.path[i] {
			// TODO: Slippage check for the route
			// minAmountOut

			// Use sdk.ZeroDec() for minOut as we can set a tighter bound later
			// amountToSwap is the nextAmountOut that we want to use
			amountLeft, amountOut, err = k.Swap0to1(goCtx, token0, token1, callerAddress, amountToSwap, sdk.ZeroDec())
			if err != nil {
				return sdk.ZeroDec(), err
			}

		} else {
			// TODO: Slippage check for the route
			// minAmountOut

			// Use sdk.ZeroDec() for minOut as we can set a tighter bound later
			amountLeft, amountOut, err = k.Swap1to0(goCtx, token0, token1, callerAddress, amountToSwap, sdk.ZeroDec())
			if err != nil {
				return sdk.ZeroDec(), err
			}

		}
		amountToSwap = amountOut
	}
	_ = amountLeft
	// Return amountOut over the route
	return amountOut, nil
}

/*

SIMULATE SWAP ALGORITHM

*/

func (k Keeper) SimulateSwap(goCtx context.Context, callerAddress sdk.AccAddress, tokenIn string, tokenOut string, amountIn sdk.Dec, minOut sdk.Dec, numChunks int64) (sdk.Dec, error) {
	// Unnecessary for now!
	// ctx := sdk.UnwrapSDKContext(goCtx)

	// Confirm that numChunks > 0
	if numChunks < 1 {
		return sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrInvalidNumChunks, "Number of chunks must be greater than 0")
	}

	chunkPrices := make([]sdk.Dec, numChunks)
	_ = chunkPrices
	routes, _ := k.getValidRoutes(goCtx, tokenIn, tokenOut, getRoutes(tokenIn, tokenOut))

	// CREATING PRICE MATRIX
	priceMatrix := make([][]sdk.Dec, len(routes))
	for i, route := range routes {
		chunkPrices, _ := k.CalculateChunkPrices(goCtx, callerAddress, route, amountIn, minOut, numChunks)
		// Set ith row of priceMatrix to chunkPrices
		priceMatrix[i] = chunkPrices
	}

	chunkToRoute := make([]Route, numChunks)

	// SELECTING THE BEST COMBINATION OF CHUNKS
	counters := make([]int64, len(routes))
	// Init counter as int64
	var i int64
	for i = 0; i < numChunks; i++ {
		bestPrice := sdk.NewDec(math.MinInt64)
		bestRouteIdx := 0
		for j := 0; j < len(routes); j++ {
			if priceMatrix[j][counters[j]].GT(bestPrice) {
				bestPrice = priceMatrix[j][counters[j]]
				bestRouteIdx = j
			}
		}
		// Set chunkToRoute[i] to be equal to routes[j]
		chunkToRoute[i] = routes[bestRouteIdx]
	}
	amountOut := sdk.ZeroDec()
	amountLeft := amountIn
	for i = 0; i < numChunks; i++ {
		routeAmountIn := sdk.ZeroDec()
		if i == numChunks-1 {
			routeAmountIn = amountLeft
		} else {
			routeAmountIn = amountIn.QuoInt64(numChunks)
		}

		// TODO: Temporarily setting minOut to sdk.ZeroDec()
		routeAmountOut, err := k.swapAcrossRoute(goCtx, callerAddress, chunkToRoute[i], amountLeft, sdk.ZeroDec())
		if err != nil {
			return sdk.ZeroDec(), err
		}
		// Subtract amount routed into this chunk from amountLeft
		amountLeft = amountLeft.Sub(routeAmountIn)
		// Add amount received from this chunk to amountOut
		amountOut = amountOut.Add(routeAmountOut)
	}

	// Confirm amountOut routed through is greater than minOut
	if amountOut.LT(minOut) {
		return sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Amount Out is less than minium amount out specified: swap failed")
	}

	return amountOut, nil

}

/*

Calculates the price of each chunk for a given route

*/
func (k Keeper) CalculateChunkPrices(goCtx context.Context, callerAddress sdk.AccAddress, route Route, amountIn sdk.Dec, minOut sdk.Dec, numChunks int64) ([]sdk.Dec, error) {
	amountToSwap := amountIn
	amountLeft := sdk.ZeroDec()

	// Tracks chunk prices (needed for non-direct routes)
	currentChunkPrices := make([]sdk.Dec, numChunks)
	for i := 0; i < len(route.path)-1; i++ {
		// Gets each pair sequentially
		token0, token1, _ := k.SortTokens(route.path[i], route.path[i+1])
		// fmt.Println("Token0: ", token0, "Token1: ", token1)
		var chunkPrices []sdk.Dec
		var amountOut sdk.Dec
		var err error
		if token0 == route.path[i] {
			chunkPrices, amountOut, err = k.SimulateSwap0to1(goCtx, token0, token1, callerAddress, amountToSwap, numChunks)
		} else {
			chunkPrices, amountOut, err = k.SimulateSwap1to0(goCtx, token0, token1, callerAddress, amountToSwap, numChunks)
		}
		if err != nil {
			return nil, err
		}
		// If we are on the first pair, init currentChunkPrices
		if i == 0 {
			currentChunkPrices = chunkPrices
		} else {
			// Element-wise multiplication of each chunk's prices, to reflect accurate prices for each chunk
			// Multiplying prices of each pair in the route
			for j := 0; j < len(currentChunkPrices); j++ {
				currentChunkPrices[j] = currentChunkPrices[j].Mul(chunkPrices[j])
			}
		}
		amountToSwap = amountOut

	}
	_ = amountLeft
	return currentChunkPrices, nil
}

/*

GHOST FUNCTIONS
- Test Swap without Changing State

*/

// Swap0to1 with No State Changes (Simulation Fn.)
// Outputs an array benchmarking every 1/Xth of the amountIn
// Note: We don't care about minOut, just gives array
func (k Keeper) SimulateSwap0to1(goCtx context.Context, token0 string, token1 string, callerAddr sdk.AccAddress, amountIn sdk.Dec, numChunks int64) ([]sdk.Dec, sdk.Dec, error) {

	// Store price of each chunk
	chunkPrices := make([]sdk.Dec, numChunks)

	ctx := sdk.UnwrapSDKContext(goCtx)

	// pair idea: "token0/token1"
	pairId := k.CreatePairId(token0, token1)

	// size of the feeList
	feeSize := k.GetFeeListCount(ctx)
	feelist := k.GetAllFeeList(ctx)
	// geets the PairMap from the KVstore given pairId
	pair, pairFound := k.GetPairMap(ctx, pairId)

	// If tokenPair does not exists a swap cannot be made through it, error
	if !pairFound {
		return nil, sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrValidPairNotFound, "Pair not found")
	}

	// Counts how many ticks we have iterated through, compare to initialized ticks in the pair
	// @Note Heuristic to remove unnecessary looping
	count := 0

	numChunksSoFar := int64(0)
	//amount_left is the amount left to deposit
	amount_left := amountIn.QuoInt64(numChunks)
	// amount to return to receiver
	amount_out := sdk.ZeroDec()
	// amount to return to receiver (in case we have a multi-hop)
	total_amount_out := sdk.ZeroDec()

	// verify that amount left is not zero and that there are additional valid ticks to check
	for !amount_left.Equal(sdk.ZeroDec()) && (count < int(pair.PairCount)) {

		// Tick data for tick that holds information about reserve1
		Current1Data, Current1Found := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick0To1)

		if !Current1Found {
			pair.TokenPair.CurrentTick0To1 = pair.TokenPair.CurrentTick0To1 - 1
			continue
		}

		// iterate count
		count++

		var i uint64

		// iterator for feeList
		i = 0
		for i < feeSize && !amount_left.Equal(sdk.ZeroDec()) {
			// gets fee for given feeIndex
			fee := feelist[i].Fee
			Current0Data, Current0Found := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick0To1+2*fee)
			//Current0Datam := Current0Data.TickData.Reserve1[i]

			// If tick/feeIndex pair is not found continue
			if !Current0Found {
				i++
				continue
			}
			// calculate currentPrice
			price, err := k.Calc_price(pair.TokenPair.CurrentTick0To1, false)

			if err != nil {
				return chunkPrices, total_amount_out, err
			}

			// price * r1 < amount_left
			if price.Mul(Current1Data.TickData.Reserve1[i]).LT(amount_left) {
				// amount_out += r1 (adds as all of reserve1 to amount_out)
				amount_out = amount_out.Add(Current1Data.TickData.Reserve1[i])
				// decrement amount_left by price * r1
				amount_left = amount_left.Sub(price.Mul(Current1Data.TickData.Reserve1[i]))
				//updates reserve0 with the new amountIn
				Current0Data.TickData.Reserve0AndShares[i].Reserve0 = Current0Data.TickData.Reserve0AndShares[i].Reserve0.Add(price.Mul(Current1Data.TickData.Reserve1[i]))
				// sets reserve1 to 0
				Current1Data.TickData.Reserve1[i] = sdk.ZeroDec()
				//updates feeIndex
				i++

			} else {
				if numChunksSoFar < numChunks {
					// amountOut += amount_left * price
					amount_out = amount_out.Add(amount_left.Mul(price))
					// increment reserve0 with amountLeft
					Current0Data.TickData.Reserve0AndShares[i].Reserve0 = Current0Data.TickData.Reserve0AndShares[i].Reserve0.Add(amount_left)
					// decrement reserve1 with amount_left * price
					Current1Data.TickData.Reserve1[i] = Current1Data.TickData.Reserve1[i].Sub(amount_left.Mul(price))

					// Store price of each chunk
					chunkPrices[numChunksSoFar] = (amount_out.Quo(amountIn.QuoInt64(numChunks)))

					// reset amountLeft to size of a chunk
					amount_left = amountIn.QuoInt64(numChunks)

					// add to total_amount_out
					total_amount_out = total_amount_out.Add(amount_out)
					// reset amountOut for chunk to 0
					amount_out = sdk.ZeroDec()

				} else {
					amount_left = sdk.ZeroDec()
				}
				numChunksSoFar++
			}

			//Make updates to tickMap containing reserve0/1 data to the KVStore

			// DO NOT UPDATE STATE WHEN SIMULATING
			// // Changes inside of the loop
			// k.SetTickMap(ctx, pairId, Current0Data)
		}
		// DO NOT UPDATE STATE WHEN SIMULATING
		// // Current1Data updates here
		// k.SetTickMap(ctx, pairId, Current1Data)

		// if feeIndex is equal to the largest index in feeList
		if i == feeSize {
			// iterates CurrentTick0to1
			pair.TokenPair.CurrentTick0To1 = pair.TokenPair.CurrentTick0To1 - 1
		}
	}

	// DO NOT UPDATE PAIR MAP IN THIS FUNCTION
	// k.SetPairMap(ctx, pair)

	// Returns amount_out to keeper/msg.server: Swap
	// @Dev token transfers happen in keeper/msg.server: Swap
	return chunkPrices, total_amount_out, nil
}

// Swap0to1 with No State Changes (Simulation Fn.)
// Outputs an array benchmarking every 1/Xth of the amountIn
// Note: We don't care about minOut, just gives array
func (k Keeper) SimulateSwap1to0(goCtx context.Context, token0 string, token1 string, callerAddr sdk.AccAddress, amountIn sdk.Dec, numChunks int64) ([]sdk.Dec, sdk.Dec, error) {

	// Store price of each chunk
	chunkPrices := make([]sdk.Dec, numChunks)

	ctx := sdk.UnwrapSDKContext(goCtx)

	// pair idea: "token0/token1"
	pairId := k.CreatePairId(token0, token1)

	// size of the feeList
	feeSize := k.GetFeeListCount(ctx)
	feelist := k.GetAllFeeList(ctx)
	// geets the PairMap from the KVstore given pairId
	pair, pairFound := k.GetPairMap(ctx, pairId)

	if !pairFound {
		return nil, sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrValidPairNotFound, "Pair not found")
	}

	// Counts how many ticks we have iterated through, compare to initialized ticks in the pair
	// @Note Heuristic to remove unecessary looping
	count := 0

	// How many chunks of the entire amount we've fulfilled
	numChunksSoFar := int64(0)

	//amount_left is the size of the chunk (will be reset every time we fill a chunk)
	amount_left := amountIn.QuoInt64(numChunks)

	//amount_out is the output of a chunk (will be reset every time we fill a chunk)
	amount_out := sdk.ZeroDec()
	// amount to return to receiver (in case we have a multi-hop)
	total_amount_out := sdk.ZeroDec()

	// verify that amount left is not zero and that there are additional valid ticks to check
	for !amount_left.Equal(sdk.ZeroDec()) && (count < int(pair.PairCount)) {

		Current0Data, Current0Found := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick1To0)
		//Current0Datam := Current0Data.TickData.Reserve1[i]

		// If tick/feeIndex pair is not found continue

		if !Current0Found {
			pair.TokenPair.CurrentTick0To1 = pair.TokenPair.CurrentTick1To0 + 1
			continue
		}

		// iterate count
		count++

		var i uint64

		// iterator for feeList
		i = 0
		for i < feeSize && !amount_left.Equal(sdk.ZeroDec()) {
			// gets fee for given feeIndex
			fee := feelist[i].Fee

			Current1Data, Current1Found := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick1To0-2*fee)

			if !Current1Found {
				i++
				continue
			}
			//Current0Datam := Current0Data.TickData.Reserve1[i]

			// calculate currentPrice and inverts
			price, err := k.Calc_price(pair.TokenPair.CurrentTick1To0, true)

			if err != nil {
				return chunkPrices, total_amount_out, err
			}

			// price * r1 < amount_left
			if price.Mul(Current0Data.TickData.Reserve0AndShares[i].Reserve0).LT(amount_left) {
				// amountOut += amount_left * price
				amount_out = amount_out.Add(Current0Data.TickData.Reserve0AndShares[i].Reserve0)
				// decrement amount_left by price * reserve0
				amount_left = amount_left.Sub(price.Mul(Current0Data.TickData.Reserve0AndShares[i].Reserve0))
				//updates reserve1 with the new amountIn
				Current1Data.TickData.Reserve1[i] = Current1Data.TickData.Reserve1[i].Add(price.Mul(Current0Data.TickData.Reserve0AndShares[i].Reserve0))
				// sets reserve0 to 0
				Current0Data.TickData.Reserve0AndShares[i].Reserve0 = sdk.ZeroDec()
				//updates feeIndex
				i++

			} else {
				if numChunksSoFar < numChunks-1 {
					// amountOut += amount_left * price
					amount_out = amount_out.Add(amount_left.Mul(price))
					// increment reserve1 with amountLeft
					Current1Data.TickData.Reserve1[i] = Current1Data.TickData.Reserve1[i].Add(amount_left)
					// decrement reserve0 with amount_left * price
					Current0Data.TickData.Reserve0AndShares[i].Reserve0 = Current0Data.TickData.Reserve0AndShares[i].Reserve0.Sub(amount_left.Mul(price))

					// Store price of each chunk
					chunkPrices[numChunksSoFar] = (amount_out.Quo(amountIn.QuoInt64(numChunks)))

					// reset amountLeft to size of a chunk
					amount_left = amountIn.QuoInt64(numChunks)

					// add to total_amount_out
					total_amount_out = total_amount_out.Add(amount_out)
					// reset amountOut for chunk to 0
					amount_out = sdk.ZeroDec()

				} else {
					amount_left = sdk.ZeroDec()
				}
				numChunksSoFar++

			}

			//Make updates to tickMap containing reserve0/1 data to the KVStore
			// DO NOT UPDATE STATE WHEN SIMULATING
			// // Changes inside of the loop
			// k.SetTickMap(ctx, pairId, Current1Data)
		}
		// DO NOT UPDATE STATE WHEN SIMULATING
		// // Current1Data updates here
		// k.SetTickMap(ctx, pairId, Current0Data)

		// if feeIndex is equal to the largest index in feeList
		if i == feeSize {

			pair.TokenPair.CurrentTick0To1 = pair.TokenPair.CurrentTick1To0 + 1
		}
	}

	// Check to see if amount_out meets the threshold of minOut
	// DO NOT UPDATE PAIR MAP IN THIS FUNCTION

	// k.SetPairMap(ctx, pair)

	// Returns amount_out to keeper/msg.server: Swap
	// @Dev token transfers happen in keeper/msg.server: Swap
	return chunkPrices, total_amount_out, nil
}
