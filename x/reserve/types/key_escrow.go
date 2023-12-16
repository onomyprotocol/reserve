package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DenomKeyPrefix is the prefix to retrieve all Denoms
	EscrowKeyPrefix = "Escrow/"
)

// DenomKey returns the store key to retrieve a Denom from the index fields
func EscrowKey(
	base string,
) []byte {
	var key []byte

	escrowBytes := []byte(base)
	key = append(key, escrowBytes...)
	key = append(key, []byte("/")...)

	return key
}
