package keeper

import (
	"context"
	"fmt"
	"slices"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/oracle/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) BandPriceStates(c context.Context, _ *types.QueryBandPriceStatesRequest) (*types.QueryBandPriceStatesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	res := &types.QueryBandPriceStatesResponse{
		PriceStates: k.GetAllBandPriceStates(ctx),
	}

	return res, nil
}

func (k Keeper) Price(c context.Context, q *types.QueryPriceRequest) (*types.QueryPriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	allowedPriceDelay := k.GetParams(ctx).AllowedPriceDelay
	// query ref by using GetBandPriceState
	basePriceState := k.GetBandPriceState(ctx, q.BaseDenom)
	if basePriceState == nil || basePriceState.Rate.IsZero() {
		return nil, fmt.Errorf("can not get price state of base denom %s: price state is nil or rate is zero", q.BaseDenom)
	}
	if ctx.BlockTime().Sub(time.Unix(basePriceState.ResolveTime, 0)) > allowedPriceDelay {
		return nil, fmt.Errorf("symbol %s old price state", q.BaseDenom)
	}
	if q.QuoteDenom == types.QuoteUSD {
		return &types.QueryPriceResponse{
			Price: basePriceState.PriceState.Price.String(),
		}, nil
	}

	quotePriceState := k.GetBandPriceState(ctx, q.QuoteDenom)
	if quotePriceState == nil || quotePriceState.Rate.IsZero() {
		return nil, fmt.Errorf("can not get price state of base denom %s: price state is nil or rate is zero", q.QuoteDenom)
	}
	if ctx.BlockTime().Sub(time.Unix(quotePriceState.ResolveTime, 0)) > allowedPriceDelay {
		return nil, fmt.Errorf("symbol %s old price state", q.QuoteDenom)
	}

	baseRate := basePriceState.Rate.ToLegacyDec()
	quoteRate := quotePriceState.Rate.ToLegacyDec()

	if baseRate.IsNil() || quoteRate.IsNil() || !baseRate.IsPositive() || !quoteRate.IsPositive() {
		return nil, fmt.Errorf("get price error validate for baseRate %s(%s) or quoteRate %s(%s)", q.BaseDenom, baseRate.String(), q.QuoteDenom, quoteRate.String())
	}

	price := baseRate.Quo(quoteRate)

	return &types.QueryPriceResponse{
		Price: price.String(),
	}, nil
}

func (k Keeper) BandParams(c context.Context, q *types.QueryBandParamsRequest) (*types.QueryBandParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	bandParams := k.GetBandParams(ctx)
	res := &types.QueryBandParamsResponse{
		BandParams: &bandParams,
	}
	return res, nil
}

func (k Keeper) BandOracleRequestParams(c context.Context, q *types.QueryBandOracleRequestParamsRequest) (*types.QueryBandOracleRequestParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	bandOracleRequestParams := k.GetBandOracleRequestParams(ctx)
	res := &types.QueryBandOracleRequestParamsResponse{
		BandOracleRequestParams: &bandOracleRequestParams,
	}
	return res, nil
}

func (k Keeper) BandOracleRequest(c context.Context, q *types.QueryBandOracleRequestRequest) (*types.QueryBandOracleRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	requestID, err := strconv.ParseUint(q.RequestId, 10, 64)
	if err != nil {
		return nil, err
	}

	bandOracleRequest := k.GetBandOracleRequest(ctx, requestID)
	res := &types.QueryBandOracleRequestResponse{
		BandOracleRequest: bandOracleRequest,
	}
	return res, nil
}

func (k Keeper) QueryOracleScriptIdByDenom(c context.Context, q *types.QueryOracleScriptIdByDenomRequest) (*types.QueryOracleScriptIdByDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	allIds := []int64{}

	err := k.IteratorOracleRequests(ctx, func(bandOracleRequest types.BandOracleRequest) bool {
		if slices.Contains(bandOracleRequest.Symbols, q.Denom) {
			allIds = append(allIds, bandOracleRequest.OracleScriptId)
		}
		return false
	})
	if err != nil {
		return nil, err
	}
	res := &types.QueryOracleScriptIdByDenomResponse{

		OracleScriptIds: allIds,
	}
	return res, nil
}
