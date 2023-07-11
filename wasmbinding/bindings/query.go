package bindings

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DualityQuery contains duality custom queries.
type DualityQuery struct {
	EstimateMultiHopSwapResult *EstimateMultiHopSwapResultRequest `json:"estimate_multi_hop_swap_result,omitempty"`
}

type EstimateMultiHopSwapResultRequest struct {
	Creator        string           `json:"creator"`
	Receiver       string           `json:"receiver"`
	Routes         []*MultiHopRoute `json:"routes"`
	AmountIn       sdk.Int          `json:"amount_in"`
	ExitLimitPrice sdk.Dec          `json:"exit_limit_price"`
	// If pickBestRoute == true then all routes are run and the route with the best price is chosen
	// otherwise, the first succesful route is used.
	PickBestRoute bool `json:"pick_best_route"`
}

type EstimateMultiHopSwapResultResponse struct {
	// TODO
}
