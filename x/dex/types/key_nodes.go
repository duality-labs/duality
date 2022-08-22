package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NodesKeyPrefix is the prefix to retrieve all Nodes
	NodesKeyPrefix = "Nodes/value/"
)

// NodesKey returns the store key to retrieve a Nodes from the index fields
func NodesKey(
	node string,
) []byte {
	var key []byte

	nodeBytes := []byte(node)
	key = append(key, nodeBytes...)
	key = append(key, []byte("/")...)

	return key
}
