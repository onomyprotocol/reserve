package keeper

import (
	"context"

	"reserve/x/reserve/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetVaultByUid(goCtx context.Context, req *types.QueryGetVaultByUidRequest) (*types.QueryGetVaultByUidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	vault, found := k.GetVault(ctx, req.Uid)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "vault not found")
	}

	return &types.QueryGetVaultByUidResponse{
		Vault: &vault,
	}, nil
}

func (k Keeper) GetAllVaults(goCtx context.Context, req *types.QueryGetAllVaultsRequest) (*types.QueryGetAllVaultsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	vaults := k.GetVaults(ctx)

	return &types.QueryGetAllVaultsResponse{
		Vaults: vaults,
	}, nil
}

func (k Keeper) GetAllVaultsByOwner(goCtx context.Context, req *types.QueryGetAllVaultsByOwnerRequest) (*types.QueryGetAllVaultsByOwnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	vaults := k.GetVaultsByOwner(ctx, req.Address)

	return &types.QueryGetAllVaultsByOwnerResponse{
		Vaults: vaults,
	}, nil
}

func (k Keeper) GetAllVaultsInDefault(goCtx context.Context, req *types.QueryGetAllVaultsInDefaultRequest) (*types.QueryGetAllVaultsInDefaultResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	vaults := k.GetVaultsInDefault(ctx)

	return &types.QueryGetAllVaultsInDefaultResponse{
		Vaults: vaults,
	}, nil
}
