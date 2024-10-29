---
sidebar_position: 4
title: Governance Proposal Messages
---

# Governance Proposal Messages

## MsgUpdateBandParams Gov

Band oracle parameters can be updated through a message gov `MsgUpdateBandParams`

```protobuf
// MsgUpdateBandParams define defines a SDK message for update band parameters
message MsgUpdateBandParams {
  option (cosmos.msg.v1.signer) = "authority";
  option (amino.name) = "oracle/UpdateBandParams";
  option (gogoproto.goproto_getters) = false;
  option (gogoproto.equal)           = false;
  // authority is the address that controls the module (defaults to x/gov unless overwritten).
  string   authority     = 1 [(cosmos_proto.scalar) = "cosmos.AddressString"];
  BandParams band_params = 2 [ (gogoproto.nullable) = false ];
}
```

Example usage:

```json
{
    "messages": [{
      "@type": "/reserve.oracle.MsgUpdateBandParams",
      "authority": "onomy10d07y265gmmuvt4z0w9aw880jnsr700jqr8n8k",
      "band_params" : {
        "ibc_request_interval": 7,
        "ibc_source_channel": "channel-0",
        "ibc_version":"bandchain-2",
        "ibc_port_id": "oracle",
        "legacy_oracle_ids": [42]
      }
    }],
    "deposit": "100000000stake",
    "title": "My proposal",
    "summary": "A short summary of my proposal"
}
```

## UpdateBandOracleRequest Gov

Band oracle request can be updated through a message gov `MsgUpdateBandOracleRequestRequest`

```protobuf
// MsgUpdateBandOracleRequestRequest define defines a SDK message for update band oracle requests
message MsgUpdateBandOracleRequestRequest{
  option (cosmos.msg.v1.signer) = "authority";
  option (amino.name) = "oracle/UpdateBandOracleRequest";
  option (gogoproto.goproto_getters) = false;
  option (gogoproto.equal)           = false;
  // authority is the address that controls the module (defaults to x/gov unless overwritten).
  string   authority                      = 1 [(cosmos_proto.scalar) = "cosmos.AddressString"];

  BandOracleRequest update_oracle_request = 2;
}
```

Example usage:

```json
{
    "messages": [{
      "@type": "/reserve.oracle.MsgUpdateBandOracleRequestRequest",
      "authority": "onomy10d07y265gmmuvt4z0w9aw880jnsr700jqr8n8k",
      "update_oracle_request" : {
        "request_id": 1,
        "oracle_script_id": 42,
        "symbols":["BTC","USD","EUR"],
        "ask_count": 1,
        "min_count": 1,
        "fee_limit": [{"denom":"uband","amount":"300000"}],
        "prepare_gas": 100,
        "execute_gas": 200,
        "min_source_count": 6
      }
    }],
    "deposit": "100000000stake",
    "title": "My proposal",
    "summary": "A short summary of my proposal"
}
```

## DeleteBandOracleRequests Gov

Band oracle requests can be deleted through a message gov `MsgDeleteBandOracleRequests`

```protobuf
// MsgDeleteBandOracleRequests define defines a SDK message for delete band oracle requests
message MsgDeleteBandOracleRequests{
  option (cosmos.msg.v1.signer) = "authority";
  option (amino.name) = "oracle/UpdateBandOracleRequest";
  option (gogoproto.goproto_getters) = false;
  option (gogoproto.equal)           = false;
  // authority is the address that controls the module (defaults to x/gov unless overwritten).
  string   authority                 = 1 [(cosmos_proto.scalar) = "cosmos.AddressString"];

  repeated uint64 delete_request_ids = 2;
}
```

Example usage:

```json
{
    "messages": [{
      "@type": "/reserve.oracle.MsgDeleteBandOracleRequests",
      "authority": "onomy10d07y265gmmuvt4z0w9aw880jnsr700jqr8n8k",
      "delete_request_ids" : [1, 2]
    }],
    "deposit": "100000000stake",
    "title": "My proposal",
    "summary": "A short summary of my proposal"
}
```

## UpdateBandOracleRequestParams Gov

Band oracle request parameters can be updated through a message gov `MsgUpdateBandOracleRequestParamsRequest`

```protobuf
// MsgUpdateBandOracleRequestParamsRequest define defines a SDK message for update band oracle request parameters
message MsgUpdateBandOracleRequestParamsRequest {
  option (cosmos.msg.v1.signer) = "authority";
  option (amino.name) = "oracle/UpdateBandOracleRequestParams";
  option (gogoproto.goproto_getters) = false;
  option (gogoproto.equal)           = false;
  // authority is the address that controls the module (defaults to x/gov unless overwritten).
  string   authority                 = 1 [(cosmos_proto.scalar) = "cosmos.AddressString"];

  BandOracleRequestParams update_band_oracle_request_params = 2;
}
```

Example usage:

```json
{
    "messages": [{
      "@type": "/reserve.oracle.MsgUpdateBandOracleRequestParamsRequest",
      "authority": "onomy10d07y265gmmuvt4z0w9aw880jnsr700jqr8n8k",
      "update_band_oracle_request_params" : {
        "ask_count": 1,
        "min_count": 1,
        "fee_limit": [{"denom":"uband","amount":"300000"}],
        "prepare_gas": 100,
        "execute_gas": 200,
        "min_source_count": 6
      }
    }],
    "deposit": "100000000stake",
    "title": "My proposal",
    "summary": "A short summary of my proposal"
}
```