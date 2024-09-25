package types

// this line is used by starport scaffolding # 1

func NewEmptyLiquidation(denom string) *Liquidation {
	return &Liquidation{
		Denom:                  denom,
		LiquidatingVaults:      []*Vault{},
		VaultLiquidationStatus: make(map[uint64]*VaultLiquidationStatus),
	}
}
