package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"

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

func (q queryServer) Stablecoin(ctx context.Context, req *types.QueryStablecoinRequest) (*types.QueryStablecoinResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	stablecoin, err := q.keeper.Stablecoins.Get(ctx, req.Denom)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "not found stablecoin %s", req.Denom)
	}

	return &types.QueryStablecoinResponse{
		Stablecoin:       stablecoin,
		SwapableQuantity: stablecoin.LimitTotal.Sub(stablecoin.TotalStablecoinLock),
	}, nil
}

func (q queryServer) AllStablecoin(ctx context.Context, req *types.QueryAllStablecoinRequest) (*types.QueryAllStablecoinResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	allStablecoins := []*types.StablecoinResponse{}

	adder := func(s types.Stablecoin) {
		newStablecoinResponse := types.StablecoinResponse{
			Stablecoin:       s,
			SwapableQuantity: s.LimitTotal.Sub(s.TotalStablecoinLock),
		}
		allStablecoins = append(allStablecoins, &newStablecoinResponse)
	}

	err := q.keeper.Stablecoins.Walk(ctx, nil, func(key string, value types.Stablecoin) (stop bool, err error) {
		adder(value)
		return false, nil
	})

	return &types.QueryAllStablecoinResponse{AllStablecoinResponse: allStablecoins}, err
}
