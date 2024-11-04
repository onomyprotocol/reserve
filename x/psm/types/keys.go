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
)

var (
	KeyStableCoin          = []byte{0x01}
	KeyLockStableCoin      = []byte{0x02}
	KeyUnlockStableCoin    = []byte{0x03}
	ParamsKey              = []byte{0x4}
	KeyTotalStablecoinLock = []byte{0x5}
	KeyFeeMax              = []byte{0x6}
	KeyNoms                = []byte{0x07}
)
