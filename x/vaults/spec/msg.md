## Messages

### MsgCreateNewVault
Allows a user to create a new vault by depositing collateral and minting tokens.

```{.go}
type MsgCreateVault struct {
	Owner      string   
	Collateral types.Coin 
	Minted     types.Coin
}
```

**State Modifications:**

- Safety Checks that are being done before running create vault logic:
  - Check that `Collateral` denom is active
  - Check that `Minted` is params mint denom
  - Check that `Minted` is cross the params min debt
- Create a vault struct with Id, Address and Owner:
  - `Id` is auto increase each vault
  - `Address` is unique by each vault
  - `Owner` is msg sender
- Lock user collateral into `Collateral` field of vault:
- Mint to user:
  - Calculate mint fee base on `Minted`
  - Vault Debt = Minted + fee
  - Mint and transfer `Minted` to user
- Decrease vault manager mint available

### MsgDepositToVault
Allows users to deposit additional collateral to an existing vault.

```{.go}
type MsgDeposit struct {
	VaultId uint64     
	Sender  string     
	Amount  types.Coin 
}
```

**State Modifications:**

- Safety Checks that are being done before running deposit to vault logic:
  - Check that `VaultId` is exist
  - Check that vault is active
  - Check that `Sender` is vault owner
  - Check that `Amount` denom is accept for vault
- Send `Amount` from user to vault address
- Increase `CollateralLock`

### MsgWithdrawFromVault
Enables users to withdraw collateral from a vault, provided the new collateral ratio remains above the minimum.

```{.go}
type MsgWithdraw struct {
	VaultId uint64    
	Sender  string     
	Amount  types.Coin 
}
```

**State Modifications:**

- Safety Checks that are being done before running withdraw from vault logic:
  - Check that `VaultId` is exist
  - Check that vault is active
  - Check that `Sender` is vault owner
  - Check that `Amount` denom is accept for vault
- Check that new collateral ratio after withdraw exceed min collateral ratio.
- Send `Amount` from vault address to user
- Decrease `CollateralLock`

### MsgMint
Allows users to mint more token from a vault.

```{.go}
type MsgMint struct {
	VaultId uint64     
	Sender  string     
	Amount  types.Coin 
}
```

**State Modifications:**

- Safety Checks that are being done before running repay debt logic:
  - Check that `VaultId` is exist
  - Check that vault is active
  - Check that `Sender` is vault owner
  - Check that `Amount` denom is params mint denom
- Check that collateral ratio after mint is greater than min collateral ratio
- Mint `Amount` then transfer to user:
  - apply minting fee
- Increase vault debt.
- Decrease vault manager mint available.


### MsgRepayDebt
Allows users to repay a portion or all of their outstanding debt.

```{.go}
// MsgRepay defines a SDK message for repay debt.
type MsgRepay struct {
	VaultId uint64     
	Sender  string     
	Amount  types.Coin 
}
```

**State Modifications:**

- Safety Checks that are being done before running repay debt logic:
  - Check that `VaultId` is exist
  - Check that vault is active
  - Check that `Sender` is vault owner
  - Check that `Amount` denom is params mint denom
- Check that `Amount` is greater than debt.
  - If so, repay amount should be debt
- Send repay amount from user to `vaults` module then burn.
- Decrease vault debt.
- Increase vault manager mint available.

### MsgCloseVault
Allows users to repay all debt and withdraw all collateral locked.

```{.go}
type MsgClose struct {
	VaultId uint64 
	Sender  string 
}
```

**State Modifications:**

- Safety Checks that are being done before running repay debt logic:
  - Check that `VaultId` is exist
  - Check that vault is active
  - Check that `Sender` is vault owner
- Transfer all debt from user to vault module then burn.
- Transfer all collateral locked to user
- Mark vault as Close

