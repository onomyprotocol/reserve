package types

import "cosmossdk.io/collections"

const (
	ModuleName = "vaults"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	RouterKey         = ModuleName
	ReserveModuleName = "reserve"
)

var (
	ParamsKey        = collections.NewPrefix(1)
	VaultKey         = collections.NewPrefix(2)
	VaultManagerKey  = collections.NewPrefix(3)
	VaultSequenceKey = collections.NewPrefix(4)
)
