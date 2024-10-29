---
sidebar_position: 3
title: Messages
---

# Messages

## RequestBandRates

`MsgRequestBandIBCRates` is a message to instantly broadcast a request to bandchain.

```protobuf
// MsgRequestBandRates defines a SDK message for requesting data from
// BandChain using IBC.
message MsgRequestBandRates {
  option (amino.name) = "oracle/MsgRequestBandRates";
  option (gogoproto.equal) = false;
  option (gogoproto.goproto_getters) = false;
  option (cosmos.msg.v1.signer) = "sender";

  string sender = 1;
  uint64 request_id = 2;
}
```

Anyone can broadcast this message and no specific authorization is needed.