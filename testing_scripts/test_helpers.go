package testing_scripts

import (
	"github.com/NicholasDotSol/duality/x/dex/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SingleLimitOrderFill() simulates the fill of a single limit order and returns the amount
// swapped into it, filling some of it (amount_in) and the amount swapped out (amount_out). It
// takes as input the amount that was placed for the limit order (amount_placed), the price the
// trader pays when filling it (price_filled_at) and the amount that they are swapping (amount_to_swap).
// The format of the return statement is (amount_in, amount_out).
func SingleLimitOrderFill(amount_placed sdk.Dec,
	price_filled_at sdk.Dec,
	amount_to_swap sdk.Dec) (sdk.Dec, sdk.Dec) {
	amount_out, amount_in := sdk.ZeroDec(), sdk.ZeroDec()
	// Checks if the swap will deplete the entire limit order and simulates the trade accordingly
	if amount_to_swap.GT(amount_placed.Quo(price_filled_at)) {
		amount_out = amount_placed
		amount_in = amount_placed.Quo(price_filled_at)
	} else {
		amount_in = amount_to_swap
		amount_out = amount_in.Mul(price_filled_at)
	}

	return amount_in, amount_out
}

// Calls SingleLimitOrderFill() and updates the filled and unfilled reserves.
// Returns the unfilled reserves (unfilled_reserves), filled reserves (filled_reserves) and the amount left to swap
// (amount_to_swap_remaining)
func SingleLimitOrderFillAndUpdate(amount_placed sdk.Dec,
	price_filled_at sdk.Dec,
	amount_to_swap sdk.Dec,
	unfilled_reserves sdk.Dec) (sdk.Dec, sdk.Dec, sdk.Dec) {
	amount_in, amount_out := SingleLimitOrderFill(amount_placed, price_filled_at, amount_to_swap)
	unfilled_reserves = unfilled_reserves.Sub(amount_out)
	filled_reserves := amount_placed.Add(amount_in)
	amount_to_swap_remaining := amount_to_swap.Sub(amount_in)
	return unfilled_reserves, filled_reserves, amount_to_swap_remaining
}

// MultipleLimitOrderFills() simulates the fill of multiple consecutive limit orders and returns the
// total amount filled. It takes as input the amounts that were placed for the limit
// order (amount_placed), the pricesthe trader pays when filling the orders (price_filled_at)
// and the amount that they are swapping (amount_to_swap).
func MultipleLimitOrderFills(amounts_placed []sdk.Dec, prices []sdk.Dec, amount_to_swap sdk.Dec) sdk.Dec {
	total_out, amount_remaining := sdk.ZeroDec(), amount_to_swap

	// Loops through all of the limit orders that need to be filled
	for i := 0; i < len(amounts_placed); i++ {
		amount_in, amount_out := SingleLimitOrderFill(amounts_placed[i], prices[i], amount_remaining)

		amount_remaining = amount_remaining.Sub(amount_in)
		total_out = total_out.Add(amount_out)
	}
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
	return amount_in, amount_out
}

// SinglePoolSwapAndUpdate() simulates swapping through a single liquidity pool and updates that pool's
// liquidity. Takes in all of the same inputs as SinglePoolSwap(): amount_liquidity, price_swapped_at,
// and amount_to_swap; but has additional inputs, reservesOfInToken, reservesOfOutToken. It returns the
// updated amounts for the reservesOfInToken and the reservesOfOutToken, in the format of
// (resulting_reserves_in_token, resulting_reserves_out_token, amount_in, amount_out)
func SinglePoolSwapAndUpdate(amount_liquidity sdk.Dec,
	price_swapped_at sdk.Dec,
	amount_to_swap sdk.Dec,
	reservesOfInToken sdk.Dec,
	reservesOfOutToken sdk.Dec) (sdk.Dec, sdk.Dec, sdk.Dec, sdk.Dec) {
	amount_in, amount_out := SinglePoolSwap(amount_liquidity, price_swapped_at, amount_to_swap)
	resulting_reserves_in_token := reservesOfInToken.Add(amount_in)
	resulting_reserves_out_token := reservesOfOutToken.Add(amount_out)
	return resulting_reserves_in_token, resulting_reserves_out_token, amount_in, amount_out
}

// SinglePoolSwapAndUpdateDirection() simulates swapping through a single liquidity pool and updates that pool's
// liquidity and specifies whether the in and out tokens are 0 or 1. Takes in all of the same inputs as
// SinglePoolSwapAndUpdate(): amount_liquidity, price_swapped_at, amount_to_swap, reservesOfToken0 sdk.Dec,
// reservesOfToken1 but has an additional input inToken which is a bool indicating whether 0 or 1 is swapped into
// the pool. It returns the updated amounts for the reservesOfInToken and the reservesOfOutToken, in the format
// of (reservesOfInToken,reservesOfOutToken).
func SinglePoolSwapAndUpdateDirectional(amount_liquidity sdk.Dec,
	price_swapped_at sdk.Dec,
	amount_to_swap sdk.Dec,
	reservesOfToken0 sdk.Dec,
	reservesOfToken1 sdk.Dec,
	inToken bool) (sdk.Dec, sdk.Dec) {
	resultingReservesOfToken0, resultingReservesOfToken1 := sdk.ZeroDec(), sdk.ZeroDec()
	if inToken {
		resultingReservesOfToken1, resultingReservesOfToken0, _, _ = SinglePoolSwapAndUpdate(amount_liquidity,
			price_swapped_at,
			amount_to_swap,
			reservesOfToken1,
			reservesOfToken0)
	} else {
		resultingReservesOfToken0, resultingReservesOfToken1, _, _ = SinglePoolSwapAndUpdate(amount_liquidity,
			price_swapped_at,
			amount_to_swap,
			reservesOfToken0,
			reservesOfToken1)
	}
	return resultingReservesOfToken0, resultingReservesOfToken1
}

// MultiplePoolSwapAndUpdate() simulates swapping through multiple liquidity pools and updates that pool's
// liquidity. Takes in similar inputs to SinglePoolSwapAndUpdate(): amount_liquidity, price_swapped_at,
// and amount_to_swap, reservesOfInToken, reservesOfOutToken; But they are held in arrays the size of how many
// pools are being swapped through. It returns the updated amounts for the reservesOfInToken and the
// reservesOfOutToken, in the format of (reservesOfInToken,reservesOfOutToken)
func MultiplePoolSwapAndUpdate(amounts_liquidity []sdk.Dec,
	prices_swapped_at []sdk.Dec,
	amount_to_swap sdk.Dec,
	reserves_in_token_array []sdk.Dec,
	reserves_out_token_array []sdk.Dec) ([]sdk.Dec, []sdk.Dec, sdk.Dec, sdk.Dec) {
	num_pools := len(amounts_liquidity)
	amount_remaining := amount_to_swap
	amount_out_total, amount_out_temp, amount_in := sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec()
	resulting_reserves_in_token := make([]sdk.Dec, num_pools, num_pools)
	resulting_reserves_out_token := make([]sdk.Dec, num_pools, num_pools)
	for i := 0; i < num_pools; i++ {
		resulting_reserves_in_token[i], resulting_reserves_out_token[i], amount_in, amount_out_temp = SinglePoolSwapAndUpdate(amounts_liquidity[i],
			prices_swapped_at[i],
			amount_to_swap,
			reserves_in_token_array[i],
			reserves_out_token_array[i])
		amount_out_total = amount_out_total.Add(amount_out_temp)
		amount_remaining = amount_remaining.Sub(amount_in)
		i++
	}

	return resulting_reserves_in_token, resulting_reserves_out_token, amount_remaining, amount_out_total
}

func SharesOnDeposit(existing_shares sdk.Dec, existing_amount0 sdk.Dec, existing_amount1 sdk.Dec, new_amount0 sdk.Dec, new_amount1 sdk.Dec, tickIndex int64) (shares_minted sdk.Dec) {
	price1To0 := keeper.CalcPrice1To0(tickIndex)

	new_value := new_amount0.Add(price1To0.Mul(new_amount1))

	if existing_amount0.Add(existing_amount1).GT(sdk.ZeroDec()) {
		existing_value := existing_amount0.Add(price1To0.Mul(existing_amount1))
		shares_minted = shares_minted.Mul(new_value.Quo(existing_value))
	} else {
		shares_minted = new_value
	}

	return shares_minted
}
