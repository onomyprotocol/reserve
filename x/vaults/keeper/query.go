package keeper

import (
	"context"
	// "errors"

	// "cosmossdk.io/collections"

	"cosmossdk.io/math"
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
	allCollateral := []*types.VaultManager{}

	err := q.keeper.VaultsManager.Walk(ctx, nil, func(key string, value types.VaultManager) (stop bool, err error) {
		allCollateral = append(allCollateral, &value)
		return false, nil
	})

	return &types.QueryAllCollateralResponse{
		AllVaultManager: allCollateral,
	}, err
}

func (q queryServer) QueryCollateralsByDenom(ctx context.Context, req *types.QueryCollateralsByDenomRequest) (*types.QueryCollateralsByDenomResponse, error) {
	allCollateral := []*types.VaultManager{}

	err := q.keeper.VaultsManager.Walk(ctx, nil, func(key string, value types.VaultManager) (stop bool, err error) {
		if value.Denom == req.Denom {
			allCollateral = append(allCollateral, &value)
		}

		return false, nil
	})

	return &types.QueryCollateralsByDenomResponse{
		AllVaultManagerByDenom: allCollateral,
	}, err
}

func (q queryServer) QueryCollateralsByMintDenom(ctx context.Context, req *types.QueryCollateralsByMintDenomRequest) (*types.QueryCollateralsByMintDenomResponse, error) {
	allCollateral := []*types.VaultManager{}

	err := q.keeper.VaultsManager.Walk(ctx, nil, func(key string, value types.VaultManager) (stop bool, err error) {
		if value.Params.MintDenom == req.MintDenom {
			allCollateral = append(allCollateral, &value)
		}

		return false, nil
	})

	return &types.QueryCollateralsByMintDenomResponse{
		AllVaultManagerByMintDenom: allCollateral,
	}, err
}

func (q queryServer) QueryCollateralsByDenomMintDenom(ctx context.Context, req *types.QueryCollateralsByDenomMintDenomRequest) (*types.QueryCollateralsByDenomMintDenomResponse, error) {
	var collateral *types.VaultManager

	err := q.keeper.VaultsManager.Walk(ctx, nil, func(key string, value types.VaultManager) (stop bool, err error) {
		if value.Denom == req.Denom && value.Params.MintDenom == req.MintDenom {
			collateral = &value
			return true, nil
		}

		return false, nil
	})

	return &types.QueryCollateralsByDenomMintDenomResponse{
		VaultManager: collateral,
	}, err
}

func (q queryServer) QueryAllVaults(ctx context.Context, req *types.QueryAllVaultsRequest) (*types.QueryAllVaultsResponse, error) {
	allVaults := []*types.Vault{}

	err := q.keeper.Vaults.Walk(ctx, nil, func(key uint64, value types.Vault) (stop bool, err error) {
		allVaults = append(allVaults, &value)
		return false, nil
	})

	return &types.QueryAllVaultsResponse{
		AllVault: allVaults,
	}, err
}

func (q queryServer) QueryVaultsByID(ctx context.Context, req *types.QueryVaultIdRequest) (*types.QueryVaultIdResponse, error) {
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

func (q queryServer) QueryVaultByOwner(ctx context.Context, req *types.QueryVaultByOwnerRequest) (*types.QueryVaultByOwnerResponse, error) {
	allVaults := []*types.Vault{}

	err := q.keeper.Vaults.Walk(ctx, nil, func(key uint64, value types.Vault) (stop bool, err error) {
		if req.Address == value.Owner {
			allVaults = append(allVaults, &value)
		}
		return false, nil
	})

	return &types.QueryVaultByOwnerResponse{
		Vaults: allVaults,
	}, err
}

func (q queryServer) QueryTotalCollateralLockedByDenom(ctx context.Context, req *types.QueryTotalCollateralLockedByDenomRequest) (*types.QueryTotalCollateralLockedByDenomResponse, error) {
	total := math.ZeroInt()

	err := q.keeper.Vaults.Walk(ctx, nil, func(key uint64, value types.Vault) (stop bool, err error) {
		if req.Denom == value.CollateralLocked.Denom && value.Status == types.ACTIVE {
			total = total.Add(value.CollateralLocked.Amount)
		}
		return false, nil
	})

	return &types.QueryTotalCollateralLockedByDenomResponse{
		Total: total,
	}, err
}
