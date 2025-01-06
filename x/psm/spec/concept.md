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

