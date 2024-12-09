package keeper

import (
	"context"

	"cosmossdk.io/collections"
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

	allAuction := []types.Auction{}

	err := k.k.Auctions.Walk(ctx, nil, func(key uint64, value types.Auction) (stop bool, err error) {
		allAuction = append(allAuction, value)
		return false, nil
	})

	return &types.QueryAllAuctionResponse{
		Auctions: allAuction,
	}, err
}

func (k Querier) QueryAllBids(ctx context.Context, req *types.QueryAllBidsRequest) (*types.QueryAllBidsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	allBids := []types.Bid{}

	err := k.k.Bids.Walk(ctx, nil, func(key uint64, value types.BidQueue) (stop bool, err error) {
		allBids = append(allBids, value.Bids...)
		return false, nil
	})

	return &types.QueryAllBidsResponse{
		Bids: allBids,
	}, err
}

func (k Querier) QueryAllBidderBids(ctx context.Context, req *types.QueryAllBidderBidsRequest) (*types.QueryAllBidderBidsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	bidderAddr, err := k.k.authKeeper.AddressCodec().StringToBytes(req.Bidder)
	if err != nil {
		return nil, err
	}

	bidsByAddress, err := k.k.BidByAddress.Get(ctx, collections.Join(req.AuctionId, sdk.AccAddress(bidderAddr)))
	if err != nil {
		return nil, err
	}

	return &types.QueryAllBidderBidsResponse{
		Bids: bidsByAddress.Bids,
	}, err
}
