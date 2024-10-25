---
sidebar_position: 1
title: State
---

# State

## Params
The oracle module parameters. 
```protobuf
message Params {
  option (amino.name) = "reserve/x/oracle/Params";
  option (gogoproto.equal) = true;
}
```

## Band Oracle

This section describes all the state management to maintain the price by connecting to Band chain via IBC.

- LatestClientID is maintained to manage unique clientID for band IBC packets. It is increased by 1 when sending price request packet into bandchain.

* LatestClientID: `0x03 -> Formated(LatestClientID)`

- LatestRequestID is maintained to manage unique `BandOracleRequests`. Incremented by 1 when creating a new `BandOracleRequest`.

* LatestRequestID: `0x06 -> Formated(LatestRequestID)`

- Band price data for a given symbol is stored as follows:

* BandPriceState: `0x05 | []byte(symbol) -> ProtocolBuffer(BandPriceState)`

```protobuf
message BandPriceState {
  string symbol = 1;
  string rate = 2 [
    (gogoproto.customtype) = "cosmossdk.io/math.Int",
    (gogoproto.nullable) = false
  ];
  uint64 resolve_time = 3;
  uint64 request_ID = 4;
  PriceState price_state = 5 [ (gogoproto.nullable) = false ];
}
```

- BandCallDataRecord is stored as follows when sending price request packet into bandchain:

* CalldataRecord: `0x02 | []byte(ClientId) -> ProtocolBuffer(CalldataRecord)`

```protobuf
message CalldataRecord {
  uint64 client_id = 1;
  bytes calldata = 2;
}
```

- BandOracleRequest is stored as follows when the governance configure oracle requests to send:

* BandOracleRequest: `0x04 | []byte(RequestId) -> ProtocolBuffer(BandOracleRequest)`

```protobuf
message BandOracleRequest {
  // Unique Identifier for band ibc oracle request
  uint64 request_id = 1;

  // OracleScriptID is the unique identifier of the oracle script to be executed.
  int64 oracle_script_id = 2;

  // Symbols is the list of symbols to prepare in the calldata
  repeated string symbols = 3;

  // AskCount is the number of validators that are requested to respond to this
  // oracle request. Higher value means more security, at a higher gas cost.
  uint64 ask_count = 4;

  // MinCount is the minimum number of validators necessary for the request to
  // proceed to the execution phase. Higher value means more security, at the
  // cost of liveness.
  uint64 min_count = 5;

  // FeeLimit is the maximum tokens that will be paid to all data source providers.
  repeated cosmos.base.v1beta1.Coin fee_limit = 6 [(gogoproto.nullable) = false, (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coins"];

  // PrepareGas is amount of gas to pay to prepare raw requests
  uint64 prepare_gas = 7;
  // ExecuteGas is amount of gas to reserve for executing
  uint64 execute_gas = 8;
  // MinSourceCount is the minimum number of data sources that must be used by each validator
  uint64 min_source_count = 9;
}
```

- BandParams is stored as follows and configured by governance:

* BandParams: `0x01 -> ProtocolBuffer(BandParams)`

`BandParams` contains the information for IBC connection with band chain.

```protobuf
message BandParams {
  // block request interval to send Band IBC prices
  int64 ibc_request_interval = 1;
  // band IBC source channel
  string ibc_source_channel = 2;
  // band IBC version
  string ibc_version = 3;
  // band IBC portID
  string ibc_port_id = 4;
  //  legacy oracle scheme ids
  repeated int64 legacy_oracle_ids = 5;
}
```

- BandOracleRequestParams is stored as follows and configured by governance:

* BandOracleRequestParams: `0x07 -> ProtocolBuffer(BandOracleRequestParams)`

`BandOracleRequestParams` contains the information for Band Oracle request.

```protobuf
message BandOracleRequestParams {
  // AskCount is the number of validators that are requested to respond to this
  // oracle request. Higher value means more security, at a higher gas cost.
  uint64 ask_count = 1;

  // MinCount is the minimum number of validators necessary for the request to
  // proceed to the execution phase. Higher value means more security, at the
  // cost of liveness.
  uint64 min_count = 2;

  // FeeLimit is the maximum tokens that will be paid to all data source
  // providers.
  repeated cosmos.base.v1beta1.Coin fee_limit = 3 [
    (gogoproto.nullable) = false,
    (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coins"
  ];

  // PrepareGas is amount of gas to pay to prepare raw requests
  uint64 prepare_gas = 4;
  // ExecuteGas is amount of gas to reserve for executing
  uint64 execute_gas = 5;
  // MinSourceCount is the minimum number of data sources that must be used by
  // each validator
  uint64 min_source_count = 6;
}
```

Note:

1. `IbcSourceChannel`, `IbcVersion`, `IbcPortId` are common parameters required for IBC connection.
2. `IbcRequestInterval` describes the automatic price fetch request interval that is automatically triggered on onomy chain on beginblocker.
