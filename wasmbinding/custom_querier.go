package wasmbinding

import (
	"encoding/json"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/duality-labs/duality/wasmbinding/bindings"
)

// CustomQuerier returns a function that is an implementation of custom querier mechanism for specific messages
func CustomQuerier(qp *QueryPlugin) func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
	return func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
		var contractQuery bindings.DualityQuery
		if err := json.Unmarshal(request, &contractQuery); err != nil {
			return nil, sdkerrors.Wrapf(err, "failed to unmarshal duality query: %v", err)
		}

		switch {
		case contractQuery.EstimateMultiHopSwapResult != nil:
			response, err := qp.GetEstimateMultiHopSwapResult(
				ctx,
				contractQuery.EstimateMultiHopSwapResult.Creator,
				contractQuery.EstimateMultiHopSwapResult.Receiver,
				contractQuery.EstimateMultiHopSwapResult.Routes,
				contractQuery.EstimateMultiHopSwapResult.AmountIn,
				contractQuery.EstimateMultiHopSwapResult.ExitLimitPrice,
				contractQuery.EstimateMultiHopSwapResult.PickBestRoute,
			)
			if err != nil {
				return nil, sdkerrors.Wrapf(err, "failed to get interchain query result: %v", err)
			}

			bz, err := json.Marshal(response)
			if err != nil {
				return nil, sdkerrors.Wrapf(
					err,
					"failed to marshal interchain query result: %v",
					err,
				)
			}

			return bz, nil
		default:
			return nil, wasmvmtypes.UnsupportedRequest{Kind: "unknown duality query type"}
		}
	}
}
