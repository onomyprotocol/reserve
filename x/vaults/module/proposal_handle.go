package vaults

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	"github.com/onomyprotocol/reserve/x/vaults/keeper"
	"github.com/onomyprotocol/reserve/x/vaults/types"
)

func NewVaultsProposalHandler(k *keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.ActiveCollateralProposal:
			return k.ActiveCollateralAsset(ctx, c.ActiveCollateral.Denom, c.ActiveCollateral.MinCollateralRatio, c.ActiveCollateral.LiquidationRatio, c.ActiveCollateral.MaxDebt)
		case *types.UpdatesCollateralProposal:
			return k.UpdatesCollateralAsset(ctx, c.UpdatesCollateral.Denom, c.UpdatesCollateral.MinCollateralRatio, c.UpdatesCollateral.LiquidationRatio, c.UpdatesCollateral.MaxDebt)
		default:
			return errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s proposal content type: %T", types.ModuleName, c)
		}
	}
}
