package keeper

import (
	"context"
	"fmt"
	"math"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Route struct
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

// Temporary setting of intermediary paths
// Uses intermediary paths to create potential routes
func getRoutes(tokenIn string, tokenOut string) []Route {
	allRoutes := []Route{}
	baseRoute := Route{
		path: []string{tokenIn, tokenOut},
	}
	allRoutes = append(allRoutes, baseRoute)
	// TODO: Source from KV store for active pairs
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

// ORDERED LIST OF ROUTES, AMOUNT YOU WANT TO SWAP THROUGH THEM
// Core function, can use any arbitrary SwapDynamicRouter
// i.e. Optimal Routing from Bain, BF, etc.
func (k Keeper) swapAcrossRoute(goCtx context.Context, callerAddress sdk.AccAddress, receiverAddress sdk.AccAddress, bestRoute Route, amountIn sdk.Dec, minOut sdk.Dec) (amountFromSwap sdk.Dec, err error) {
	amountToSwap := amountIn
	amountLeft := sdk.ZeroDec()
	// Passes in the amountOut from the previous pair into the next pair until we swap to the end of route
	var amountOut sdk.Dec
	for i := 0; i < len(bestRoute.path)-1; i++ {
		// Gets each pair sequentially
		token0, token1, _ := k.SortTokens(bestRoute.path[i], bestRoute.path[i+1])
		// fmt.Println("Token0: ", token0, "Token1: ", token1)

		msgSwap := &types.MsgSwap{
			Creator:  string(callerAddress),
			Receiver: string(receiverAddress),
			TokenA:   token0,
			TokenB:   token1,
			AmountIn: amountToSwap,
			TokenIn:  bestRoute.path[i],
			MinOut:   minOut,
		}
		if token0 == bestRoute.path[i] {
			// TODO: Slippage check for the route
			// minAmountOut
			// Use sdk.ZeroDec() for minOut as we can set a tighter bound later
			// amountToSwap is the nextAmountOut that we want to use

			amountOut, err = k.Swap0to1(goCtx, msgSwap, token0, token1, callerAddress)
			if err != nil {
				return sdk.ZeroDec(), err
			}

		} else {
			// TODO: Slippage check for the route
			// minAmountOut

			// Use sdk.ZeroDec() for minOut as we can set a tighter bound later
			amountOut, err = k.Swap1to0(goCtx, msgSwap, token0, token1, callerAddress)
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

func (k Keeper) DynamicRouteSwap(goCtx context.Context, callerAddress sdk.AccAddress, receiverAddress sdk.AccAddress, tokenIn string, tokenOut string, amountIn sdk.Dec, minOut sdk.Dec, numChunks int64) (sdk.Dec, error) {
	//TODO: Add checks that all arguments are passed in!

	// Unnecessary for now!
	// ctx := sdk.UnwrapSDKContext(goCtx)

	// fmt.Println("In Dynamic Route Swap")
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
	// fmt.Println("Price Matrix: ", priceMatrix)
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
		var routeAmountIn sdk.Dec
		if i == numChunks-1 {
			routeAmountIn = amountLeft
		} else {
			routeAmountIn = amountIn.QuoInt64(numChunks)
		}

		// TODO: Temporarily setting minOut to sdk.ZeroDec()
		routeAmountOut, err := k.swapAcrossRoute(goCtx, callerAddress, receiverAddress, chunkToRoute[i], amountLeft, sdk.ZeroDec())
		if err != nil {
			return sdk.ZeroDec(), err
		}
		// Subtract amount routed into this chunk from amountLeft
		amountLeft = amountLeft.Sub(routeAmountIn)
		// Add amount received from this chunk to amountOut
		amountOut = amountOut.Add(routeAmountOut)
	}

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
	// fmt.Println("In Calculate Chunk Prices")
	// fmt.Println("CALCULATING PRICES FOR: ", route.path)
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
		// print("Chunk Prices:", len(chunkPrices))
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
				fmt.Println("Inside Loop: ", j, chunkPrices[j])

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

	fmt.Println("0to1")
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
	for !amount_left.Equal(sdk.ZeroDec()) {

		// fmt.Println("Num Chunks So Far: ", numChunksSoFar)
		// Tick data for tick that holds information about reserve1
		Current1Data, Current1Found := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick0To1)

		// TODO: Exit after certain # of ticks
		if !Current1Found {
			// iterate count
			count += 1
			if count > int(pair.TotalTickCount) {
				break
			}
			pair.TokenPair.CurrentTick0To1 = pair.TokenPair.CurrentTick0To1 - 1
			continue
		}

		count = 0

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
				// fmt.Println("Calculate price failed")
				return chunkPrices, total_amount_out, err
			}
			// fmt.Println("Fee Tick ", i)
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
			// fmt.Println("Reached end of loop ")

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

	fmt.Println("1to0")
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
	for !amount_left.Equal(sdk.ZeroDec()) {

		Current0Data, Current0Found := k.GetTickMap(ctx, pairId, pair.TokenPair.CurrentTick1To0)
		//Current0Datam := Current0Data.TickData.Reserve1[i]

		// If tick/feeIndex pair is not found continue

		// TODO: Exit after certain # of ticks
		if !Current0Found {
			// iterate count
			count += 1
			if count > int(pair.TotalTickCount) {
				break
			}
			pair.TokenPair.CurrentTick0To1 = pair.TokenPair.CurrentTick1To0 + 1
			continue
		}

		// iterate count
		count = 0

		var i uint64

		// iterator for feeList
		i = 0
		for i < feeSize && !amount_left.Equal(sdk.ZeroDec()) {
			// fmt.Println("iteration", numChunksSoFar)
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
				// fmt.Println("Calculate price failed")
				return chunkPrices, total_amount_out, err
			}

			// price * r1 < amount_left
			if price.Mul(Current0Data.TickData.Reserve0AndShares[i].Reserve0).LT(amount_left) {
				// fmt.Println("not enough in reserves")
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
				if numChunksSoFar < numChunks {
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
