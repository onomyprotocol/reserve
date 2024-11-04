# Vaults Module

The Vaults module provides functionality for managing user collateralized vaults, allowing users to deposit collateral, mint stable tokens, repay debt, and withdraw collateral. This module supports operations that ensure the safety and proper management of collateralized debt positions.

## Contents
- [Concept](#concept)
- [State](#state)
- [Messages](#messages)
- [Events](#events)
- [Params](#params)
- [Keepers](#keepers)
- [Hooks](#hooks)
- [Future Improvements](#future-improvements)

## Concept

The Vaults module enables users to create and manage vaults secured by collateral assets. Users can deposit assets as collateral and mint a specified amount of tokens based on the collateral ratio. The module ensures that the minted tokens and collateral adhere to a minimum collateralization ratio, thus maintaining the stability and solvency of the system.

### Key Features
- **Vault Creation**: Users can lock supported assets as collateral in their vaults and mint stable asset
- **Collateral Management**: User can deposit more collateral asset into vault or withdraw their collateral.
- **Debt Management**: User can burn stable tokens to reduce their debt or mint more tokens as long as the collateral ratio of the vault meets the necessary threshold.
- **Liquidation**: Vaults that fall below the required collateral ratio can be subject to liquidation to maintain the stability and security of the system. The liquidation process involves seizing and selling the collateral to repay the outstanding debt and ensure that the system remains solvent.

## State

### Params

Params contains:

- `MinInitialDebt`: Min debt for vault creation.
- `MintDenom`: Stable denom supported.
- `ChargingPeriod`: The compounding interest period is based on the debt amount of the vault.

```protobuf
message Params {
  option (amino.name) = "reserve/x/vaults/Params";
  option (gogoproto.equal) = true;

  string min_initial_debt = 1 [
    (cosmos_proto.scalar) = "cosmos.Int",
    (gogoproto.customtype) = "cosmossdk.io/math.Int",
    (amino.dont_omitempty) = true,
    (gogoproto.nullable) = false
  ];

  string mint_denom = 2;

  google.protobuf.Duration charging_period = 3 [
    (gogoproto.stdduration) = true,
    (gogoproto.nullable) = false,
    (amino.dont_omitempty) = true
  ];
}
```

### Vault Manager

Each vault manager will contains:

#### Vault Manager Params

Params config for each vault type:

- `MinCollateralRatio`: The minimum ratio of collateral value to debt.
- `LiquidationRatio`: ratio that vault will be liquidate.
- `MaxDebt`: How much stable tokens can be minted for each vault type.
- `StabilityFee`: interest fee in annual.
- `LiquidationPenalty`: The penalty applied during the liquidation process.
- `MintingFee`: The fee applied when minting new tokens.

Beside of that, vault manager keep track of:

- `Denom`: The collateral denomination accepted for vault.
- `MintAvailable`: The remaining capacity for minting tokens.

```protobuf
message VaultMamager {
  VaultMamagerParams params = 1
      [ (amino.dont_omitempty) = true, (gogoproto.nullable) = false ];

  string denom = 2;

  string mint_available = 3 [
    (cosmos_proto.scalar) = "cosmos.Int",
    (gogoproto.customtype) = "cosmossdk.io/math.Int",
    (amino.dont_omitempty) = true,
    (gogoproto.nullable) = false
  ];
}
```

### Vault Structure
Each vault contains:
- `Id`: A unique identifier for the vault.
- `Owner`: The address of the user who owns the vault.
- `CollateralLocked`: The amount of collateral locked in the vault.
- `Debt`: The total amount of debt issued from the vault.
- `Status`: Indicates status of vault (active, close, liquidating, liquidated).
- `Address`: The on-chain address associated with the vault.

```protobuf
message Vault {
  uint64 id = 1;
  string owner = 2 [ (cosmos_proto.scalar) = "cosmos.AddressString" ];

  cosmos.base.v1beta1.Coin debt = 3
      [ (gogoproto.nullable) = false, (amino.dont_omitempty) = true ];

  cosmos.base.v1beta1.Coin collateral_locked = 4
      [ (gogoproto.nullable) = false, (amino.dont_omitempty) = true ];

  VaultStatus status = 5;

  string liquidation_price = 6 [
    (cosmos_proto.scalar) = "cosmos.Dec",
    (gogoproto.customtype) = "cosmossdk.io/math.LegacyDec",
    (amino.dont_omitempty) = true,
    (gogoproto.nullable) = false
  ];

  string address = 7 [ (cosmos_proto.scalar) = "cosmos.AddressString" ];
}
```

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

## Events

The Vaults module emits events for various operations:
- **CreateVault**: Emitted when a new vault is created.
- **Deposit**: Emitted when collateral is deposited into a vault.
- **Withdraw**: Emitted when collateral is withdrawn.
- **Mint**: Emitted when minting token from vault.
- **Repay**: Emitted when debt is repaid.
- **CloseVault**: Emitted when a vault is closed.

## ABCI

### Debt recalculate
Each `ChargingPeriod`, all vaults debt is recalculated in `BeginBlock`. Using vault manager `StabilityFee`

### Liquidate
`auction` module will take all collateral locked in vault to auction. The auction result will return to `vaults` module to handle liquidate logic. The 2 flows can happen after auction:

**Flow 1: Auction raises enough IST to cover debt**

The following steps occur in this order

1. IST raised by the auction is burned to reduce debt in 1:1 ratio.

- Definitionally, this should result in zero debt. Since the auction should stop
  when it raises enough IST, it should result in zero IST remaining as well.
  However, if some excess IST exists, it should be transferred to the reserve.

2. From any remaining collateral, the liquidation penalty is taken and
   transferred to the reserve.

- Liquidation penalty is calculated as debt / current oracle price at auction
  start * liquidation penalty

3. Excess collateral - if any - is returned to vault holders

- Vault holders receive collateral back sequenced by highest collateralization
  ratio at liquidation time first.
- The max amount of collateral a vault should be able to receive back is:
  original collateral - collateral covering their share of debt (using average
  liquidation price) - collateral covering their share of the penalty (their
  debt / total debt \* total penalty)

**Flow 2: Auction does not raise enough to cover IST debt**

This flow further bifurcates based on whether the auction has sold all its
collateral asset and still has not covered the debt or has collateral
remaining (which simply did not receive bidders)

**Flow 2a: all collateral sold and debt is not covered**

1. IST raised by the auction is burned to reduce debt in 1:1 ratio.

- Definitionally, this should result in zero IST remaining and some debt
  remaining.

2. Remaining debt is recorded in the reserve as a shortfall
   
   *sequence ends; no penalty is taken and vaults receive nothing back*

**Flow 2b: collateral remains but debt is still not covered by IST raised by
auction end**

1. IST raised by the auction is burned to reduce debt in 1:1 ratio.

- Definitionally, this should result in zero IST remaining and some debt remaining.

2. From any remaining collateral, the liquidation penalty is taken and
   transferred to the reserve.

- Liquidation penalty is calculated as debt / current oracle price at auction
  start \* liquidation penalty.

  _Note: there now should be debt remaining and possibly collateral remaining_

3. The vault manager now iterates through the vaults that were liquidated and
   attempts to reconstitute them (minus collateral from the liquidation penalty)
   starting from the best CR to worst.

- Reconstitution means full prior debt AND full prior collateral minus
  collateral used from penalty.
- Collateral used for penalty = vault debt / total debt \* total liquidation penalty.
- Debt that is given back to a vault should be subtracted from the vault
  manager's view of remaining liquidation debt (i.e., it shouldn't be
  double counted)
- Reconstituted vaults are set to OPEN status (i.e., they are live again
  and able to be interacted with)

4. When the vault manager reaches a vault it cannot fully reconstitute
   (both full debt and collateral as described above), it marks that vault as
   liquidated. It then marks all other lower CR vaults as liquidated.
5. Any remaining collateral is transferred to the reserve.
6. Any remaining debt (subtracting debt that was given back to reconstituted
   vaults, as described above) is transferred to the reserve as shortfall.