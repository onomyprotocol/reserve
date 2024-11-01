package types

// this line is used by starport scaffolding # 1

func NewEmptyLiquidation(debtDenom, mintDenom string) *Liquidation {
	return &Liquidation{
		DebtDenom:              debtDenom,
		MintDenom:              mintDenom,
		LiquidatingVaults:      []*Vault{},
		VaultLiquidationStatus: make(map[uint64]*VaultLiquidationStatus),
	}
}
