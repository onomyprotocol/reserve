package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// VaultKeyPrefix is the prefix to retrieve all Drop
	VaultKeyPrefix = "Vault/"
	// VaultsKeyPrefix is the prefix to retrieve all Vaults of Owner
	VaultMapKeyPrefix = "Vault/Owner/Name/"
)

// DropKey returns the store key to retrieve a Drop from the index fields
func VaultKey(
	uid uint64,
) []byte {
	var key []byte

	uidBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(uidBytes, uid)
	key = append(key, uidBytes...)
	key = append(key, []byte("/")...)

	return key
}
