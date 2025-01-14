## State

### Params

Params contains:

- `LimitTotal`: Maximum amount of fxUSD that can be provided.
- `AcceptablePriceRatio`: The spread price between fxUSD and the stablecoins is at an acceptable level where the fees will remain the same.
- `AdjustmentFee`: "AdjustmentFee" is a parameter used to adjust the amount of change in swap fees based on the deviation of the stablecoin from the target price. AdjustmentFee determines the number of iterations that one of the fees (entry fee or exit fee) will be adjusted according to the current ratio between the target price and the actual price. The larger the AdjustmentFee, the more severe the fee adjustment will be. Specifically, it affects how the feeOut or feeIn is calculated by multiplying (or dividing) it by the price ratio.

```protobuf
message Params {
  // total $fxUSD can mint
  bytes limit_total = 1 [
    (cosmos_proto.scalar)  = "cosmos.Int",
    (gogoproto.customtype) = "cosmossdk.io/math.Int",
    (gogoproto.nullable)   = false,
    (amino.dont_omitempty) = true
  ];
  // The price cannot be exactly 1, an acceptable such as 0.9999 (AcceptablePriceRatio = 0.0001)
    bytes acceptable_price_ratio = 2 [
    (cosmos_proto.scalar)  = "cosmos.Dec",
    (gogoproto.customtype) = "cosmossdk.io/math.LegacyDec",
    (gogoproto.nullable)   = false
  ];
  // feeIn adjustment factor
  int64 adjustment_fee = 3;

}
```

### Stablecoin Manager

Each stablecoin will contains:

#### Stablecoin Manager Params

Params config for each stablecoin type:

- `Denom`: stablecoin name
- `LimitTotal`: limit total stablecoin 
- `FeeIn`: stablecoin to fxUSD exchange fee, fee_in when 1 stablecoin = 1fxUSD
- `FeeOut`: fxUSD to stablecoin exchange fee, fee_out when 1 stablecoin = 1fxUSD
- `TotalStablecoinLock`: amount of stablecoins locked in exchange for fxUSD, default start at 0
- `FeeMaxStablecoin`: maximum fee for when either fee = 0, default at fee_in+fee_out

```protobuf
message Stablecoin {
  // stablecoin name
  string denom = 1;
    // limit total stablecoin module support
  bytes limit_total = 2 [
    (cosmos_proto.scalar)  = "cosmos.Int",
    (gogoproto.customtype) = "cosmossdk.io/math.Int",
    (gogoproto.nullable)   = false,
    (amino.dont_omitempty) = true
  ];
  // stablecoin to fxUSD exchange fee, fee_in when 1 stablecoin = 1fxUSD
  bytes fee_in = 3 [
    (cosmos_proto.scalar)  = "cosmos.Dec",
    (gogoproto.customtype) = "cosmossdk.io/math.LegacyDec",
    (gogoproto.nullable)   = false
  ];
  // fxUSD to stablecoin exchange fee, fee_out when 1 stablecoin = 1fxUSD
  bytes fee_out = 4 [
    (cosmos_proto.scalar)  = "cosmos.Dec",
    (gogoproto.customtype) = "cosmossdk.io/math.LegacyDec",
    (gogoproto.nullable)   = false
  ];
  // amount of stablecoins locked in exchange for fxUSD
  bytes total_stablecoin_lock = 5 [
    (cosmos_proto.scalar)  = "cosmos.Int",
    (gogoproto.customtype) = "cosmossdk.io/math.Int",
    (gogoproto.nullable)   = false,
    (amino.dont_omitempty) = true
  ];
  // maximum fee for when either fee = 0
  bytes fee_max_stablecoin = 6 [
    (cosmos_proto.scalar)  = "cosmos.Dec",
    (gogoproto.customtype) = "cosmossdk.io/math.LegacyDec",
    (gogoproto.nullable)   = false
  ];
}

```


## Messages

### MsgSwapTofxUSD
Allows users to swap accepted stablecoins for fxUSD. `Address` is the sender's address, `Coin` is the amount of stablecoin the user sent

```{.go}
type MsgSwapTofxUSD struct {
	Address string      
	Coin    *types.Coin 
}
```

**State Modifications:**

- Safety Checks that are being done before running swap logic:
  - Check stablecoin is suport
  - Check limit swap
  - Check balance user and calculate amount of coins received and fee in

- Transfer stablecoin from user to psm module.
- Mint fxUSD and send for user

### MsgSwapToStablecoin
Allows users to swap accepted fxUSD for stablecoins. `Address` is the sender's address, `ToDenom` is the stablecoin name to receive ,`Amount` is the amount of fxUSD the user sent

```{.go}
type MsgSwapToStablecoin struct {
	Address string                
	ToDenom string               
	Amount  math.Int 
}
```

**State Modifications:**

- Safety Checks that are being done before running swap logic:
  - Check stablecoin is suport
  - Check total stablecoin lock enough to swap
  - Check balance user and calculate amount of coins received and fee out

- Transfer fxUSD from user to psm module and burn.
- Transfer stablecoin to user
