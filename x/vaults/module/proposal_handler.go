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
		msgSv := keeper.NewMsgServerImpl(*k)
		switch c := content.(type) {
		case *types.ActiveCollateralProposal:
			_, err := msgSv.ActiveCollateral(ctx, types.NewMsgActiveCollateral(c))
			return err
		case *types.UpdatesCollateralProposal:
			_, err := msgSv.UpdatesCollateral(ctx, types.NewMsgUpdatesCollateral(c))
			return err
		case *types.BurnShortfallProposal:
			_, err := msgSv.BurnShortfall(ctx, types.NewMsgBurnShortfall(c))
			return err
		default:
			return errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s proposal content type: %T", types.ModuleName, c)
		}
	}
}
