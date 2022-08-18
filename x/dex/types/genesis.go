package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		NodesList:             []Nodes{},
		TicksList:             []Ticks{},
		BitArrList:            []BitArr{},
		PairsList:             []Pairs{},
		VirtualPriceQueueList: []VirtualPriceQueue{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in nodes
	nodesIdMap := make(map[uint64]bool)
	nodesCount := gs.GetNodesCount()
	for _, elem := range gs.NodesList {
		if _, ok := nodesIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for nodes")
		}
		if elem.Id >= nodesCount {
			return fmt.Errorf("nodes id should be lower or equal than the last id")
		}
		nodesIdMap[elem.Id] = true
	}

	// Check for duplicated index in ticks
	ticksIndexMap := make(map[string]struct{})

	for _, elem := range gs.TicksList {
		index := string(TicksKey(elem.Price, elem.Fee, elem.Direction, elem.OrderType))
		if _, ok := ticksIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for ticks")
		}
		ticksIndexMap[index] = struct{}{}
	}

	// Check for duplicated ID in bitArr
	bitArrIdMap := make(map[uint64]bool)
	bitArrCount := gs.GetBitArrCount()
	for _, elem := range gs.BitArrList {
		if _, ok := bitArrIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for bitArr")
		}
		if elem.Id >= bitArrCount {
			return fmt.Errorf("bitArr id should be lower or equal than the last id")
		}
		bitArrIdMap[elem.Id] = true
	}
	// Check for duplicated index in pairs
	pairsIndexMap := make(map[string]struct{})

	for _, elem := range gs.PairsList {
		index := string(PairsKey(elem.Token0, elem.Token1))
		if _, ok := pairsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for pairs")
		}
		pairsIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in virtualPriceQueue
	virtualPriceQueueIndexMap := make(map[string]struct{})

	for _, elem := range gs.VirtualPriceQueueList {
		index := string(VirtualPriceQueueKey(elem.VPrice, elem.Direction, elem.OrderType))
		if _, ok := virtualPriceQueueIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for virtualPriceQueue")
		}
		virtualPriceQueueIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
