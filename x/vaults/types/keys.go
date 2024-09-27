package types

import "cosmossdk.io/collections"

const (
	ModuleName        = "vaults"

	// StoreKey is the string store representation
	StoreKey = ModuleName
	
	ReserveModuleName = "reserve"
)

var (
	ParamsKey              = collections.NewPrefix(1)
	VaultKeyPrefix         = collections.NewPrefix(2)
	VaultManagerKeyPrefix  = collections.NewPrefix(3)
	VaultSequenceKeyPrefix = collections.NewPrefix(4)
)
