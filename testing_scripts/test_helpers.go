package testing_scripts

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SingleLimitOrderFill() simulates the fill of a single limit order and returns the amount
// swapped into it, filling some of it (amount_in) and the amount swapped out (amount_out). It
// takes as input the amount that was placed for the limit order (amount_placed), the price the
// trader pays when filling it (price_filled_at) and the amount that they are swapping (amount_to_swap).
// The format of the return statement is (amount_in, amount_out).
func SingleLimitOrderFill(amount_placed sdk.Dec, price_filled_at sdk.Dec, amount_to_swap sdk.Dec) (sdk.Dec, sdk.Dec) {

	amount_out, amount_in := sdk.ZeroDec(), sdk.ZeroDec()

	// Checks if the swap will deplete the entire limit order and simulates the trade accordingly
	if amount_to_swap.GT(amount_placed.Quo(price_filled_at)) {
		amount_out = amount_placed
		amount_in = amount_placed.Quo(price_filled_at)
	} else {
		amount_in = amount_to_swap
		amount_out = amount_in.Mul(price_filled_at)
	}

	fmt.Println("Amount Out: ", amount_out)
	fmt.Println("Amount In: ", amount_in)
	return amount_in, amount_out
}

// LimitOrderFills() simulates the fill of multiple consecutive limit orders and returns the
// total amount filled. It takes as input the amounts that were placed for the limit
// order (amount_placed), the pricesthe trader pays when filling the orders (price_filled_at)
// and the amount that they are swapping (amount_to_swap).
func LimitOrderFills(amounts_placed []sdk.Dec, prices []sdk.Dec, amount_to_swap sdk.Dec) sdk.Dec {

	total_out, amount_remaining := sdk.ZeroDec(), amount_to_swap

	// Loops through all of the limit orders that need to be filled
	for i := 0; i < len(amounts_placed); i++ {
		fmt.Println("\nRound: ", i)
		amount_in, amount_out := SingleLimitOrderFill(amounts_placed[i], prices[i], amount_remaining)

		amount_remaining = amount_remaining.Sub(amount_in)
		total_out = total_out.Add(amount_out)
	}
	fmt.Println("Total Out: ", total_out)
	return total_out
}

// SinglePoolSwap() simulates swapping through a single liquidity pool and returns the amount
// swapped into it (amount_in) and the amount swapped out, received by the swapper (amount_out). It
// takes as input the amount of liquidity in the pool (amount_liquidity), the price the
// trader pays when swapping through it (price_swapped_at) and the amount that they are
// swapping (amount_to_swap). The format of the return statement is (amount_in, amount_out).
// Same thing as SingleLimitOrderFill() except in naming.
func SinglePoolSwap(amount_liquidity sdk.Dec, price_swapped_at sdk.Dec, amount_to_swap sdk.Dec) (sdk.Dec, sdk.Dec) {

	amount_out, amount_in := sdk.ZeroDec(), sdk.ZeroDec()

	// Checks if the swap will deplete the entire limit order and simulates the trade accordingly
	if amount_to_swap.GT(amount_liquidity.Quo(price_swapped_at)) {
		amount_out = amount_liquidity
		amount_in = amount_liquidity.Quo(price_swapped_at)
	} else {
		amount_in = amount_to_swap
		amount_out = amount_in.Mul(price_swapped_at)
	}

	fmt.Println("Amount Out: ", amount_out)
	fmt.Println("Amount In: ", amount_in)
	return amount_in, amount_out
}

// SinglePoolSwapAndUpdate() simulates swapping through a single liquidity pool and updates that pool's
// liquidity. Takes in all of the same inputs as SinglePoolSwap(): amount_liquidity, price_swapped_at,
// and amount_to_swap; but has additional inputs, reservesOfInToken, reservesOfOutToken. It returns the
// updated amounts for the reservesOfInToken and the reservesOfOutToken, in the format of
// (reservesOfInToken,reservesOfOutToken)
func SinglePoolSwapAndUpdate(amount_liquidity sdk.Dec,
	price_swapped_at sdk.Dec,
	amount_to_swap sdk.Dec,
	reservesOfInToken sdk.Dec,
	reservesOfOutToken sdk.Dec) (sdk.Dec, sdk.Dec) {

	amountIn, amountOut := SinglePoolSwap(amount_liquidity, price_swapped_at, amount_to_swap)

	resultingReservesOfInToken := reservesOfInToken.Add(amountIn)
	resultingReservesOfOutToken := reservesOfOutToken.Add(amountOut)

	return resultingReservesOfInToken, resultingReservesOfOutToken
}

// TODO: UPDATE
// SinglePoolSwapAndUpdateDirection() simulates swapping through a single liquidity pool and updates that pool's
// liquidity. Takes in all of the same inputs as SinglePoolSwap(): amount_liquidity, price_swapped_at,
// and amount_to_swap; but has additional inputs, reservesOfInToken, reservesOfOutToken. It returns the
// updated amounts for the reservesOfInToken and the reservesOfOutToken, in the format of
// (reservesOfInToken,reservesOfOutToken)
/*func SinglePoolSwapAndUpdateDirectional(amount_liquidity sdk.Dec,
										price_swapped_at sdk.Dec,
										amount_to_swap sdk.Dec,
										reservesOfInToken sdk.Dec,
										reservesOfOutToken sdk.Dec,
										inToken) (sdk.Dec, sdk.Dec) {

amountIn, amountOut := SinglePoolSwap(amount_liquidity, price_swapped_at, amount_to_swap)

resultingReservesOfInToken := reservesOfInToken.Add(amountIn)
resultingReservesOfOutToken := reservesOfOutToken.Add(amountOut)

return resultingReservesOfInToken, resultingReservesOfOutToken
}*/
