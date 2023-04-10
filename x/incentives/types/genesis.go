package types

import (
	"encoding/json"
	"errors"

	"github.com/cosmos/cosmos-sdk/codec"
)

// DefaultIndex is the default incentive module's global index.
const DefaultIndex uint64 = 1

// DefaultGenesis returns the incentive module's default genesis state.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: Params{
			DistrEpochIdentifier: "week",
		},
		Gauges: Gauges{},
	}
}

// GetGenesisStateFromAppState returns an incentive module GenesisState given a raw application genesis state.
func GetGenesisStateFromAppState(cdc codec.JSONCodec, appState map[string]json.RawMessage) *GenesisState {
	var genesisState GenesisState

	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return &genesisState
}

// Validate performs basic genesis state validation, returning an error upon any failure.
func (gs GenesisState) Validate() error {
	if gs.Params.DistrEpochIdentifier == "" {
		return errors.New("epoch identifier should NOT be empty")
	}
	return nil
}
