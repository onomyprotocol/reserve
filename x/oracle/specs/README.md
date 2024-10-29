# `Oracle`

## Abstract

This specification specifies the oracle module, which is primarily used by the `vaults` modules to obtain external price data.

## Workflow

1. New asset must first approved through a msg governance proposal (from `vaults` module) which will be add to the `BandOracleRequest` later. <br/>
    **Example message governance proposal**: `ActiveCollateralProposal`
2. Once the message governance proposal is approved, the specified asset will be added to the `BandOracleRequest`, then for each begin block will make the request to band chain to get the asset data.
3. Upon receiving the ibc packet, oracle module will process the data to parse and update the price state of asset.
4. Other Cosmos-SDK modules can then fetch the latest price data by querying the oracle module.

## Band IBC integration flow

Cosmos SDK blockchains are able to interact with each other using IBC and Onomy support the feature to fetch price feed from bandchain via IBC.

1. To communicate with BandChain's oracle using IBC, Onomy Chain must first initialize a communication channel with the oracle module on the BandChain using relayers.

2. Once the connection has been established, a pair of channel identifiers is generated -- one for the Onomy Chain and one for Band. The channel identifier is used by Onomy Chain to route outgoing oracle request packets to Band. Similarly, Band's oracle module uses the channel identifier when sending back the oracle response.

3. The list of prices to be fetched via IBC should be determined by `ActiveCollateralProposal` and `UpdateBandOracleRequest`.

4. Chain periodically (`IbcRequestInterval`) sends price request IBC packets (`OracleRequestPacketData`) to bandchain and bandchain responds with the price via IBC packet (`OracleResponsePacketData`). Band chain is providing prices when there are threshold number of data providers confirm and it takes time to get the price after sending requests. To request price before the configured interval, any user can broadcast `MsgRequestBandRates` message which is instantly executed.

## Contents

1. **[State](./01_state.md)**
2. **[Keeper](./02_keeper.md)**
3. **[Messages](./03_messages.md)**
4. **[Proposals](./04_proposals.md)**
5. **[Events](./05_events.md)**
