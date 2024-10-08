package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/oracle/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) BandPriceStates(c context.Context, _ *types.QueryBandPriceStatesRequest) (*types.QueryBandPriceStatesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	println("go to query BandPriceStates")
	println("================================================")
	for _, data := range k.GetAllBandPriceStates(ctx) {
		println("Check current band price state: ", data.String())
	}
	res := &types.QueryBandPriceStatesResponse{
		PriceStates: k.GetAllBandPriceStates(ctx),
	}

	return res, nil
}

func (k Keeper) Price(c context.Context, q *types.QueryPriceRequest) (*types.QueryPriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	
	price := k.GetPrice(ctx, q.BaseDenom, q.QuoteDenom)

	res := &types.QueryPriceResponse{
		Price: price.String(),
	}

	return res, nil
}