package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/oracle/types"
)

var _ types.QueryServer = &Keeper{}

func (k *Keeper) BandPriceStates(c context.Context, _ *types.QueryBandPriceStatesRequest) (*types.QueryBandPriceStatesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	println("Check query band prices state .......")
	result := k.GetAllBandPriceStates(ctx)
	res := &types.QueryBandPriceStatesResponse{
		PriceStates: result,
	}

	return res, nil
}