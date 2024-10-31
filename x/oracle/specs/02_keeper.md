---
sidebar_position: 2
title: Keepers
---

# Keepers

The oracle module currently provides the price state information of tokens which can be used by other modules
which need to read price feeds.

# Band oracle

The band oracle in oracle module provides the ability to create/modify/read/delete BandParams, BandOracleRequestParams, BandOracleRequest, BandPriceState, BandLatestClientID, BandOracleRequest and BandCallDataRecord.

```go
    GetBandParams(ctx context.Context) types.BandParams
    SetBandParams(ctx context.Context, bandParams types.BandParams) error

    GetBandOracleRequestParams(ctx context.Context) types.BandOracleRequestParams
    SetBandOracleRequestParams(ctx context.Context, bandOracleRequestParams types.BandOracleRequestParams) error

    GetBandOracleRequest(ctx context.Context, requestID uint64) *types.BandOracleRequest
    SetBandOracleRequest(ctx context.Context, req types.BandOracleRequest) error
    DeleteBandOracleRequest(ctx context.Context, requestID uint64) error
    GetAllBandOracleRequests(ctx context.Context) []*types.BandOracleRequest

    GetBandPriceState(ctx context.Context, symbol string) *types.BandPriceState
    SetBandPriceState(ctx context.Context, symbol string, priceState *types.BandPriceState) error
    GetAllBandPriceStates(ctx context.Context) []*types.BandPriceState

    GetBandLatestClientID(ctx context.Context) uint64 
    SetBandLatestClientID(ctx context.Context, clientID uint64) error

    GetBandCallDataRecord(ctx context.Context, clientID uint64) *types.CalldataRecord
    SetBandCallDataRecord(ctx context.Context, record *types.CalldataRecord) error
    GetAllBandCalldataRecords(ctx context.Context) []*types.CalldataRecord
    DeleteBandCallDataRecord(ctx context.Context, clientID uint64) error

    GetBandOracleRequest(ctx context.Context, requestID uint64) *types.BandOracleRequest
    SetBandOracleRequest(ctx context.Context, req types.BandOracleRequest) error
    GetAllBandOracleRequests(ctx context.Context) []*types.BandOracleRequest
    DeleteBandOracleRequest(ctx context.Context, requestID uint64) error
```