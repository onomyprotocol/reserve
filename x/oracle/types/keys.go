package types

import (
	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "oracle"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_oracle"

	// Version defines the current version the IBC module supports
	// Version = "oracle-1"

	// PortID is the default port id that module binds to
	PortID = "oracle"
)

var (
	ParamsKey              = []byte("p_oracle")
	BandParamsKey          = []byte{1}
	BandCallDataRecordKey  = []byte{2}
	LatestClientIDKey      = []byte{3}
	BandOracleRequestIDKey = []byte{4}
	// BandPriceKey               = []byte{0x05}
	LatestRequestIDKey         = []byte{5}
	BandOracleRequestParamsKey = []byte{6}
	BandPriceKey               = collections.NewPrefix(7)
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("oracle-port-")
)

func GetBandCallDataRecordKey(clientID uint64) []byte {
	return append(BandCallDataRecordKey, sdk.Uint64ToBigEndian(clientID)...)
}

func GetBandOracleRequestIDKey(requestID uint64) []byte {
	return append(BandOracleRequestIDKey, sdk.Uint64ToBigEndian(requestID)...)
}

// func GetBandPriceStoreKey(symbol string) []byte {
// 	return append(BandPriceKey, []byte(symbol)...)
// }

func KeyPrefix(p string) []byte {
	return []byte(p)
}
