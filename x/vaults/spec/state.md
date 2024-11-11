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

