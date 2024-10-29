---
sidebar_position: 5
title: Events
---
# Events

The oracle module emits the following events:

## Band

```protobuf
message EventBandAckSuccess {
  string ack_result = 1;
  int64 client_id = 2;
}
  
message EventBandAckError {
  string ack_error = 1;
  int64 client_id = 2;
}

message EventBandResponseTimeout { 
  int64 client_id = 1; 
}

message SetBandPriceEvent {
  string relayer = 1;
  repeated string symbols = 2;
  repeated string prices = 3 [
    (gogoproto.customtype) = "cosmossdk.io/math.LegacyDec",
    (gogoproto.nullable) = false
  ];
  uint64 resolve_time = 4;
  uint64 request_id = 5;
  int64 client_id = 6;
}
```