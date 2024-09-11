package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name
	ModuleName = "auction"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_auction"
)

var (
	ParamsKey = []byte("p_auction")
)

var (
	AuctionsPrefix     = collections.NewPrefix(1)
	BidsPrefix         = collections.NewPrefix(2)
	BidByAddressPrefix = collections.NewPrefix(3)
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
