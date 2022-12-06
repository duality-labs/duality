package testing_scripts

import (
	"testing"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


func TestSingleLimitOrderFill(t *testing.T) {
	amount_placed := sdk.NewInt(10)
	price_filled_at := sdk.NewDec(10)
	amount_to_swap := sdk.NewInt(40)

	amount_in, amount_out := SingleLimitOrderFill(amount_placed, price_filled_at, amount_to_swap)

	// amount_out = min(amount_placed, amount_to_swap * price_filled_at)
	amount_out_expected := sdk.NewInt(10)
	// amount_in = min(amount_placed / price_filled_at, amount_to_swap)
	amount_in_expected := sdk.NewInt(1)

	if !amount_in_expected.ToDec().Equal(amount_in) {
		t.Errorf("amount_in: %d; want %d", amount_in, amount_in_expected)
	}

	if !amount_out_expected.ToDec().Equal(amount_out) {
		t.Errorf("amount_out: %d; want %d", amount_out, amount_out_expected)
	}
}
