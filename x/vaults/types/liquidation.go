package types

import (
	"cosmossdk.io/math"
	// this line is used by starport scaffolding # 1
)

func NewEmptyLiquidation(denom string, price math.LegacyDec) *Liquidation {
	return &Liquidation{
		Denom:                  denom,
		MarkPrice:              price,
		LiquidatingVaults:      []*Vault{},
		VaultLiquidationStatus: make(map[uint64]*VaultLiquidationStatus),
	}
}
