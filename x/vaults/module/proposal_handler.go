package vaults

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	"github.com/onomyprotocol/reserve/x/vaults/client/cli"
	"github.com/onomyprotocol/reserve/x/vaults/keeper"
	"github.com/onomyprotocol/reserve/x/vaults/types"
)

var (
	ActiveCollateralAssetProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitActiveCollateralProposal)
)

func NewVaultsProposalHandler(k *keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.ActiveCollateralProposal:
			return k.ActiveCollateralAsset(ctx, c.Denom, c.MinCollateralRatio, c.LiquidationRatio, c.MaxDebt)
		default:
			return errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s proposal content type: %T", types.ModuleName, c)
		}
	}
}
