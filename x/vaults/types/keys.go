package types

import "cosmossdk.io/collections"

const (
	ModuleName = "vaults"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	ReserveModuleName = "reserve"
	RouterKey         = ModuleName
)

var (
	ParamsKey              = collections.NewPrefix(1)
	VaultKeyPrefix         = collections.NewPrefix(2)
	VaultManagerKeyPrefix  = collections.NewPrefix(3)
	VaultSequenceKeyPrefix = collections.NewPrefix(4)
	LastUpdateKeyPrefix    = collections.NewPrefix(5)
	ShortfallKeyPrefix     = collections.NewPrefix(6)
)
