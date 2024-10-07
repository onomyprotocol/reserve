package keeper

import (
	"context"
	// "errors"

	// "cosmossdk.io/collections"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/onomyprotocol/reserve/x/vaults/types"
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

	params := q.keeper.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}

func (q queryServer) QueryAllCollateral(ctx context.Context, req *types.QueryAllCollateralRequest) (*types.QueryAllCollateralResponse, error) {
	allCollateral := []*types.VaultMamager{}

	q.keeper.VaultsManager.Walk(ctx, nil, func(key string, value types.VaultMamager) (stop bool, err error) {
		allCollateral = append(allCollateral, &value)
		return false, nil
	})

	return &types.QueryAllCollateralResponse{
		AllVaultMamager: allCollateral,
	}, nil
}

func (q queryServer) QueryAllVaults(ctx context.Context, req *types.QueryAllVaultsRequest) (*types.QueryAllVaultsResponse, error) {
	allVaults := []*types.Vault{}

	q.keeper.Vaults.Walk(ctx, nil, func(key uint64, value types.Vault) (stop bool, err error) {
		allVaults = append(allVaults, &value)
		return false, nil
	})

	return &types.QueryAllVaultsResponse{
		AllVault: allVaults,
	}, nil
}

func (q queryServer) QueryVaults(ctx context.Context, req *types.QueryVaultIdRequest) (*types.QueryVaultIdResponse, error) {
	if req == nil {
		return &types.QueryVaultIdResponse{}, status.Error(codes.InvalidArgument, "invalid request")
	}

	vault, err := q.keeper.GetVault(ctx, req.VaultId)
	if err != nil {
		return &types.QueryVaultIdResponse{}, err
	}
	return &types.QueryVaultIdResponse{
		Vault: &vault,
	}, nil
}
