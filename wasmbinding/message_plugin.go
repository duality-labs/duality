package wasmbinding

import (
	"encoding/json"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/duality-labs/duality/wasmbinding/bindings"

	dexkeeper "github.com/duality-labs/duality/x/dex/keeper"
	dextypes "github.com/duality-labs/duality/x/dex/types"
)

func CustomMessageDecorator(
	dexKeeper *dexkeeper.Keeper,
) func(messenger wasmkeeper.Messenger) wasmkeeper.Messenger {
	return func(old wasmkeeper.Messenger) wasmkeeper.Messenger {
		return &CustomMessenger{
			Wrapped:      old,
			DexMsgServer: dexkeeper.NewMsgServerImpl(*dexKeeper),
		}
	}
}

type CustomMessenger struct {
	Wrapped      wasmkeeper.Messenger
	DexMsgServer dextypes.MsgServer
}

var _ wasmkeeper.Messenger = (*CustomMessenger)(nil)

func (m *CustomMessenger) DispatchMsg(
	ctx sdk.Context,
	contractAddr sdk.AccAddress,
	contractIBCPortID string,
	msg wasmvmtypes.CosmosMsg,
) ([]sdk.Event, [][]byte, error) {
	if msg.Custom != nil {
		var contractMsg bindings.DualityMsg
		if err := json.Unmarshal(msg.Custom, &contractMsg); err != nil {
			ctx.Logger().Debug("json.Unmarshal: failed to decode incoming custom cosmos message",
				"from_address", contractAddr.String(),
				"message", string(msg.Custom),
				"error", err,
			)
			return nil, nil, sdkerrors.Wrap(err, "failed to decode incoming custom cosmos message")
		}

		if contractMsg.MultiHopSwap != nil {
			return m.multiHopSwap(ctx, contractAddr, contractMsg.MultiHopSwap)
		}
	}

	return m.Wrapped.DispatchMsg(ctx, contractAddr, contractIBCPortID, msg)
}

func (m *CustomMessenger) multiHopSwap(
	ctx sdk.Context,
	contractAddr sdk.AccAddress,
	multiHopSwap *bindings.MultiHopSwap,
) ([]sdk.Event, [][]byte, error) {
	response, err := m.performMultiHopSwap(ctx, contractAddr, multiHopSwap)
	if err != nil {
		ctx.Logger().Debug("performSubmitTx: failed to submit dex transaction",
			"creator", contractAddr.String(),
			"receiver", multiHopSwap.Receiver,
			"routes", multiHopSwap.Routes,
			"amount", multiHopSwap.AmountIn,
			"exit_limit_price", multiHopSwap.ExitLimitPrice,
			"pick_best_route", multiHopSwap.PickBestRoute,
			"error", err,
		)
		return nil, nil, sdkerrors.Wrap(err, "failed to submit dex transaction")
	}

	data, err := json.Marshal(response)
	if err != nil {
		ctx.Logger().Error("json.Marshal: failed to marshal submitTx response to JSON",
			"creator", contractAddr.String(),
			"receiver", multiHopSwap.Receiver,
			"routes", multiHopSwap.Routes,
			"amount", multiHopSwap.AmountIn,
			"exit_limit_price", multiHopSwap.ExitLimitPrice,
			"pick_best_route", multiHopSwap.PickBestRoute,
			"error", err,
		)
		return nil, nil, sdkerrors.Wrap(err, "marshal json failed")
	}

	ctx.Logger().Debug("dex transaction submitted",
		"creator", contractAddr.String(),
		"receiver", multiHopSwap.Receiver,
		"routes", multiHopSwap.Routes,
		"amount", multiHopSwap.AmountIn,
		"exit_limit_price", multiHopSwap.ExitLimitPrice,
		"pick_best_route", multiHopSwap.PickBestRoute,
	)
	return nil, [][]byte{data}, nil
}

func (m *CustomMessenger) performMultiHopSwap(
	ctx sdk.Context,
	contractAddr sdk.AccAddress,
	multiHopSwap *bindings.MultiHopSwap,
) (*bindings.MultiHopSwapResponse, error) {
	routes := make([]*dextypes.MultiHopRoute, len(multiHopSwap.Routes))
	for i, route := range multiHopSwap.Routes {
		routes[i] = &dextypes.MultiHopRoute{
			Hops: route.Hops,
		}
	}
	tx := &dextypes.MsgMultiHopSwap{
		Creator:        contractAddr.String(),
		Receiver:       multiHopSwap.Receiver,
		Routes:         routes,
		AmountIn:       multiHopSwap.AmountIn,
		ExitLimitPrice: multiHopSwap.ExitLimitPrice,
		PickBestRoute:  multiHopSwap.PickBestRoute,
	}

	if err := tx.ValidateBasic(); err != nil {
		return nil, sdkerrors.Wrap(err, "failed to validate incoming MultiHopSwap message")
	}

	grpcResponse, err := m.DexMsgServer.MultiHopSwap(sdk.WrapSDKContext(ctx), tx)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to submit dex transaction")
	}

	bindingsResponse := &bindings.MultiHopSwapResponse{
		CoinOut: grpcResponse.CoinOut,
	}

	return bindingsResponse, nil
}
