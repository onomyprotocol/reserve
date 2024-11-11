# PSM Module

The PSM module is a stabilization mechanism for the nomUSD stablecoin, allowing users to swap approved stablecoins for nomUSD at a 1-1 ratio.

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

PSM issues nomUSD in exchange for approved stablecoins, with a maximum issuance limit set by governance. It also allows users to swap nomUSD for stablecoins at a 1-1 ratio.nomUSD is issued from PSM which is backed by stablecoins within it. Parameters such as limit per stablecoin, swap-in fees and swap-out fees are governed and can change based on economic conditions. PSM is an effective support module for Vaults module in case users need nomUSD to repay off debts and close vault (redeem collateral).

### Key Features

Key Features

- **Swap to nomUSD**: Users can convert from stablecoin to nomUSD provided that the stablecoin has been added to the list of accepted stablecoins.
- **Swap to Stablecoin**: Users can convert from nomUSD to stablecoin provided that the stablecoin has been added to the list of accepted stablecoins.

