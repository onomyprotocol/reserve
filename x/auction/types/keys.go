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
	AuctionIdSeqPrefix = collections.NewPrefix(1)
	BidIdSeqPrefix     = collections.NewPrefix(2)
	AuctionsPrefix     = collections.NewPrefix(3)
	BidsPrefix         = collections.NewPrefix(4)
	BidByAddressPrefix = collections.NewPrefix(5)
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
