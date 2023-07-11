package wasmbinding

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	dexkeeper "github.com/duality-labs/duality/x/dex/keeper"
)

// RegisterCustomPlugins returns wasmkeeper.Option that we can use to connect handlers for implemented custom queries and messages to the App
func RegisterCustomPlugins(
	dexKeeper *dexkeeper.Keeper,
) []wasmkeeper.Option {
	wasmQueryPlugin := NewQueryPlugin(dexKeeper)

	queryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Custom: CustomQuerier(wasmQueryPlugin),
	})
	messagePluginOpt := wasmkeeper.WithMessageHandlerDecorator(
		CustomMessageDecorator(dexKeeper),
	)

	return []wasm.Option{
		queryPluginOpt,
		messagePluginOpt,
	}
}
