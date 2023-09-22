//nolint:revive,stylecheck  // if we change the names of var-naming things here, we harm some kind of mapping.
package bindings

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	math_utils "github.com/duality-labs/duality/utils/math"
)

// DualityMsg is used like a sum type to hold one of custom Duality messages.
// Follow https://github.com/duality-labs/duality-contracts/tree/main/packages/bindings/src/msg.rs
// for more information.
type DualityMsg struct {
	MultiHopSwap *MultiHopSwap `json:"multi_hop_swap,omitempty"`
}

// SubmitTx submits interchain transaction on a remote chain.
type MultiHopSwap struct {
	Receiver       string           `json:"receiver"`
	Routes         []*MultiHopRoute `json:"routes"`
	AmountIn       sdk.Int          `json:"amount_in"`
	ExitLimitPrice math_utils.PrecDec `json:"exit_limit_price"`
	// If pickBestRoute == true then all routes are run and the route with the best price is chosen
	// otherwise, the first succesful route is used.
	PickBestRoute bool `json:"pick_best_route"`
}

type MultiHopRoute struct {
	Hops []string `json:"hops"`
}

type MultiHopSwapResponse struct {
	CoinOut sdk.Coin `json:"coin_out"`
}
