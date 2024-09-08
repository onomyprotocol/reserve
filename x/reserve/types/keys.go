package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name
	ModuleName = "reserve"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_reserve"
)

var (
	ParamsKey = collections.NewPrefix("p_reserve")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
