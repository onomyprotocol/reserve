package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DenomKeyPrefix is the prefix to retrieve all Denoms
	DenomKeyPrefix = "Denom/"
)

// DenomKey returns the store key to retrieve a Denom from the index fields
func DenomKey(
	base string,
) []byte {
	var key []byte

	denomBytes := []byte(base)
	key = append(key, denomBytes...)
	key = append(key, []byte("/")...)

	return key
}
