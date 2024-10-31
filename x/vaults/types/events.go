package types

const (
	TypeEvtCreateVault = "create_vault"
	TypeEvtDeposit     = "deposit"
	TypeEvtWithdraw    = "withdraw"
	TypeEvtMint        = "mint"
	TypeEvtRepay       = "repay"
	TypeEvtLiquidate   = "liquidate"

	AttributeKeyVaultId         = "vault_id"
	AttributeKeyMintAmount      = "mint_amount"
	AttributeKeyBurnAmount      = "burn_amount"
	AttributeKeyRepayAmount     = "repay_amount"
	AttributeKeyCollateral      = "collateral"
	AttributeKeyDebt            = "debt"
	AttributeKeyOwner           = "owner"
	AttributeKeyVaultAddress    = "vault_address"
	AttributeKeyShortfallAmount = "shortfall_amount"
	AttributeKeyReserve         = "reserve"
	AttributeKeyLiquidateVaults = "liquidate_vaults"
)
