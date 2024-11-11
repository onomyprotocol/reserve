
## Events

The PSM module emits events for various operations:
- **AddStablecoin**: Emitted when a new stablecoin is added.
- **UpdateStablecoin**: Emitted when a stablecoin is updates
- **SwapToStablecoin**: Emitted when exchanging nomUSD for stablecoin
- **SwapTonomUSD**: Emitted when exchanging stablecoin for nomUSD


## ABCI

### Fee recalculate
To maintain the peg of 1 nomUSD = 1 USD, the swap fees need to be adjusted whenever the price of stablecoins fluctuates. This adjustment ensures that deviations from the target price of 1 USD per nomUSD are counterbalanced by the fees.

Adjustment Logic:
- If the stablecoin price is above 1, fee_out (the fee for converting from nomUSD to the stablecoin) will be higher, and fee_in (the fee for converting from the stablecoin to nomUSD) will be lower. This setup discourages swaps that would increase the stablecoin holdings when its value is above 1, helping to bring the price back down.
- If the stablecoin price is below 1, fee_out will be lower, and fee_in will be higher. This makes it cheaper to convert nomUSD to the stablecoin and more costly to convert the stablecoin to nomUSD, which encourages activity that pushes the price back up toward the target.


#### How to calculate fee:
The fee adjustments are scaled using the `AdjustmentFee` parameter (k), which controls the responsiveness of the fee to price deviations.

Suppose:
- `newPrice`: new market price of stablecoin relative to nomUSD.
- `feeIn`: inbound fee (to exchange stablecoin to nomUSD).
- `feeOut`: outbound fee (to exchange nomUSD to stablecoin).
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
