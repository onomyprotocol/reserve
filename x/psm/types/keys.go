package types

const (
	// ModuleName defines the module name.
	ModuleName = "psm"

	// StoreKey defines the primary module store key.
	StoreKey = ModuleName

	// RouterKey is the message route for slashing.
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key.
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key.
	MemStoreKey = "mem_psm"

	// Inter Stable Token.
	DenomStable = "nomUSD"
)

var (
	KeyStableCoin       = []byte{0x01}
	KeyLockStableCoin   = []byte{0x02}
	KeyUnlockStableCoin = []byte{0x03}
	ParamsKey           = []byte{0x4}
)

func GetKeyStableCoin(denom string) []byte {
	return append(KeyStableCoin, []byte(denom)...)
}

func GetKeyLockCoin(denom string) []byte {
	return append(KeyLockStableCoin, []byte(denom)...)
}
