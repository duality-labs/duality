package types

const (
	// ModuleName defines the module name
	ModuleName = "dex"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_dex"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	NodesKey      = "Nodes-value-"
	NodesCountKey = "Nodes-count-"
)

const (
	VirtualPriceTickQueueKey      = "VirtualPriceTickQueue-value-"
	VirtualPriceTickQueueCountKey = "VirtualPriceTickQueue-count-"
)

const (
	BitArrKey      = "BitArr-value-"
	BitArrCountKey = "BitArr-count-"
)
