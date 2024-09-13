package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/onomyprotocol/reserve/x/psm/types"
)

var _ types.QueryServer = queryServer{}

// NewQueryServerImpl returns an implementation of the QueryServer interface
// for the provided Keeper.
func NewQueryServerImpl(k Keeper) types.QueryServer {
	return queryServer{k}
}

type queryServer struct {
	keeper Keeper
}

func (q queryServer) Params(ctx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	params, err := q.keeper.GetParams(ctx)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryParamsResponse{Params: params}, nil
}

func (q queryServer) Stablecoin(c context.Context, req *types.QueryStablecoinRequest) (*types.QueryStablecoinResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	stablecoin, found := q.keeper.GetStablecoin(ctx, req.Denom)
	if !found {
		return nil, status.Errorf(codes.NotFound, "not found stablecoin %s", req.Denom)
	}

	totalStablecoinLock, err := q.keeper.TotalStablecoinLock(ctx, req.Denom)
	if err != nil {
		return nil, err
	}

	return &types.QueryStablecoinResponse{
		Stablecoin:       stablecoin,
		CurrentTotal:     totalStablecoinLock,
		SwapableQuantity: stablecoin.LimitTotal.Sub(totalStablecoinLock),
	}, nil
}
