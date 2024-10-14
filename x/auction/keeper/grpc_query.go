package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/onomyprotocol/reserve/x/auction/types"
)

var _ types.QueryServer = Querier{}

type Querier struct {
	k Keeper
}

func NewQueryServerImpl(k Keeper) types.QueryServer {
	return Querier{k: k}
}

func (k Querier) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryParamsResponse{Params: k.k.GetParams(ctx)}, nil
}

func (k Querier) QueryAllAuction(ctx context.Context, req *types.QueryAllAuctionRequest) (*types.QueryAllAuctionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	t, _ := k.k.Auctions.Has(ctx, 0)
	fmt.Println()
	fmt.Println(t)

	allAuction := []types.Auction{}

	k.k.Auctions.Walk(ctx, nil, func(key uint64, value types.Auction) (stop bool, err error) {
		fmt.Println("----1----")
		allAuction = append(allAuction, value)
		return false, nil
	})

	return &types.QueryAllAuctionResponse{
		Auctions: allAuction,
	}, nil
}
