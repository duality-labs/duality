package wasmbinding

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/duality-labs/duality/wasmbinding/bindings"
)

func (qp *QueryPlugin) GetEstimateMultiHopSwapResult(
	ctx sdk.Context,
	creator string,
	receiver string,
	routes []*bindings.MultiHopRoute,
	amountIn sdk.Int,
	exitLimitPrice sdk.Dec,
	// If pickBestRoute == true then all routes are run and the route with the best price is chosen
	// otherwise, the first succesful route is used.
	pickBestRoute bool,
) (*bindings.EstimateMultiHopSwapResultResponse, error) {
	// TODO

	// grpcResp, err := qp.dexKeeper.EstimateMultiHopSwap(ctx, queryID)
	// if err != nil {
	// 	return nil, err
	// }
	// resp := bindings.QueryResult{
	// 	KvResults: make([]*bindings.StorageValue, 0, len(grpcResp.KvResults)),
	// 	Height:    grpcResp.GetHeight(),
	// 	Revision:  grpcResp.GetRevision(),
	// }
	// for _, grpcKv := range grpcResp.GetKvResults() {
	// 	kv := bindings.StorageValue{
	// 		StoragePrefix: grpcKv.GetStoragePrefix(),
	// 		Key:           grpcKv.GetKey(),
	// 		Value:         grpcKv.GetValue(),
	// 	}
	// 	resp.KvResults = append(resp.KvResults, &kv)
	// }

	return &bindings.EstimateMultiHopSwapResultResponse{}, nil
}
