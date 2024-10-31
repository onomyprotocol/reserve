package keeper

import (
	"context"
	"strconv"

	errors "cosmossdk.io/errors"
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

	price := k.GetPrice(ctx, q.BaseDenom, q.QuoteDenom)
	if price == nil || price.IsNil() {
		return nil, errors.Wrapf(types.ErrInvalidOracle, "base %s quote %s", q.BaseDenom, q.QuoteDenom)
	} else {
		res := &types.QueryPriceResponse{
			Price: price.String(),
		}
		return res, nil
	}
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
