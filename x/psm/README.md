# PSM Module

The PSM module is a stabilization mechanism for the fxUSD stablecoin, allowing users to swap approved stablecoins for fxUSD at a 1-1 ratio.

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

PSM issues fxUSD in exchange for approved stablecoins, with a maximum issuance limit set by governance. It also allows users to swap fxUSD for stablecoins at a 1-1 ratio.fxUSD is issued from PSM which is backed by stablecoins within it. Parameters such as limit per stablecoin, swap-in fees and swap-out fees are governed and can change based on economic conditions. PSM is an effective support module for Vaults module in case users need fxUSD to repay off debts and close vault (redeem collateral).

### Key Features

Key Features

- **Swap to fxUSD**: Users can convert from stablecoin to fxUSD provided that the stablecoin has been added to the list of accepted stablecoins.
- **Swap to Stablecoin**: Users can convert from fxUSD to stablecoin provided that the stablecoin has been added to the list of accepted stablecoins.

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

### MsgStableSwap (swap to fxUSD)
Allows users to swap accepted stablecoins for fxUSD. `Address` is the sender's address, `OfferCoin` is the amount of stablecoin the user sent. `ExpectedDenom` is the type of denom expected to be received.

```{.go}
type MsgStableSwap struct {
	Address       string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	ExpectedDenom string     `protobuf:"bytes,2,opt,name=expected_denom,json=expectedDenom,proto3" json:"expected_denom,omitempty"`
	OfferCoin     types.Coin `protobuf:"bytes,3,opt,name=offer_coin,json=offerCoin,proto3" json:"offer_coin"`
}
```

**State Modifications:**

- Safety Checks that are being done before running swap logic:
  - Check stablecoin is suport
  - Check limit swap
  - Check balance user and calculate amount of coins received and fee in

- Transfer stablecoin from user to psm module.
- Mint fxUSD and send for user

### MsgStableSwap (swap to stablecoin)
Allows users to swap accepted fxUSD for stablecoins. `Address` is the sender's address, `ExpectedDenom` is the stablecoin name to receive ,`OfferCoin` is the amount of fxUSD the user sent

```{.go}
type MsgStableSwap struct {
	Address       string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	ExpectedDenom string     `protobuf:"bytes,2,opt,name=expected_denom,json=expectedDenom,proto3" json:"expected_denom,omitempty"`
	OfferCoin     types.Coin `protobuf:"bytes,3,opt,name=offer_coin,json=offerCoin,proto3" json:"offer_coin"`
}
```

**State Modifications:**

- Safety Checks that are being done before running swap logic:
  - Check stablecoin is suport
  - Check total stablecoin lock enough to swap
  - Check balance user and calculate amount of coins received and fee out

- Transfer fxUSD from user to psm module and burn.
- Transfer stablecoin to user

## Events

The PSM module emits events for various operations:
- **AddStablecoin**: Emitted when a new stablecoin is added.
- **UpdateStablecoin**: Emitted when a stablecoin is updates
- **Swap**: Emitted when exchanging fxUSD for stablecoin. Emitted when exchanging stablecoin for fxUSD


## ABCI

### Fee recalculate
To maintain the peg of 1 fxUSD = 1 USD, the swap fees need to be adjusted whenever the price of stablecoins fluctuates. This adjustment ensures that deviations from the target price of 1 USD per fxUSD are counterbalanced by the fees.

Adjustment Logic:
- If the stablecoin price is above 1, fee_out (the fee for converting from fxUSD to the stablecoin) will be higher, and fee_in (the fee for converting from the stablecoin to fxUSD) will be lower. This setup discourages swaps that would increase the stablecoin holdings when its value is above 1, helping to bring the price back down.
- If the stablecoin price is below 1, fee_out will be lower, and fee_in will be higher. This makes it cheaper to convert fxUSD to the stablecoin and more costly to convert the stablecoin to fxUSD, which encourages activity that pushes the price back up toward the target.


#### How to calculate fee:
The fee adjustments are scaled using the `AdjustmentFee` parameter (k), which controls the responsiveness of the fee to price deviations.

Suppose:
- `newPrice`: new market price of stablecoin relative to fxUSD.
- `feeIn`: inbound fee (to exchange stablecoin to fxUSD).
- `feeOut`: outbound fee (to exchange fxUSD to stablecoin).
- `maxFee`: maximum total fee for both directions (usually feeIn + feeOut).
- `k`: adjustment factor (`AdjustmentFee`) that controls the sensitivity of the fee to price changes.

Calculate new price ratio:`rate`= 1/`newPrice`
​
Fee adjustment:
    - `rate` < 1: 
        `newFeeOut` = `feeOut`/ `rate`^`k`
        `newFeeIn`=`maxFee`−`newFeeOut`
        `newFeeOut` will not exceed `maxFee`
    - `rate` > 1: 
        `newFeeIn` = `feeIn`*`rate`^`k`
        `newFeeOut`=`maxFee`−`newFeeIn`
        `newFeeIn` will not exceed `maxFee`
