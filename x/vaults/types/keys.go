package types

import "cosmossdk.io/collections"

const (
	ModuleName        = "vaults"
	ReserveModuleName = "reserve"
)

var (
	ParamsKey              = collections.NewPrefix(1)
	VaultKeyPrefix         = collections.NewPrefix(2)
	VaultManagerKeyPrefix  = collections.NewPrefix(3)
	VaultSequenceKeyPrefix = collections.NewPrefix(4)
)
