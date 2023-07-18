package wasmbinding

import (
	dexkeeper "github.com/duality-labs/duality/x/dex/keeper"
)

type QueryPlugin struct {
	dexKeeper *dexkeeper.Keeper
}

// NewQueryPlugin returns a reference to a new QueryPlugin.
func NewQueryPlugin(
	dexKeeper *dexkeeper.Keeper,
) *QueryPlugin {
	return &QueryPlugin{
		dexKeeper: dexKeeper,
	}
}
