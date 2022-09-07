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
	TokensKey      = "Tokens-value-"
	TokensCountKey = "Tokens-count-"

	// TokenMapKeyPrefix is the prefix to retrieve all TokenMap
	TokenMapKeyPrefix = "TokenMap/value/"

	// TickMapKeyPrefix is the prefix to retrieve all TickMap
	TickMapKeyPrefix = "TickMap/value/"

	// PairMapKeyPrefix is the prefix to retrieve all PairMap
	PairMapKeyPrefix = "PairMap/value/"
)

// TokenMapKey returns the store key to retrieve a TokenMap from the index fields
func TokenMapKey(address string) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}

func TickMapKey(tickIndex string) []byte {
	var key []byte

	tickIndexBytes := []byte(tickIndex)
	key = append(key, tickIndexBytes...)
	key = append(key, []byte("/")...)

	return key
}

func PairMapKey(pairId string) []byte {
	var key []byte

	pairIdBytes := []byte(pairId)
	key = append(key, pairIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
