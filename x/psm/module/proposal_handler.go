package psm

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	"github.com/onomyprotocol/reserve/x/psm/client/cli"
	"github.com/onomyprotocol/reserve/x/psm/keeper"
	"github.com/onomyprotocol/reserve/x/psm/types"
)

var (
	AddStableCoinProposalHandler     = govclient.NewProposalHandler(cli.NewCmdSubmitAddStableCoinProposal)
	UpdatesStableCoinProposalHandler = govclient.NewProposalHandler(cli.NewCmdUpdatesStableCoinProposal)
)

func NewStablecoinProposalHandler(k *keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.AddStableCoinProposal:
			return k.AddStableCoinProposal(ctx, c)
		case *types.UpdatesStableCoinProposal:
			return k.UpdatesStableCoinProposal(ctx, c)
		default:
			return errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s proposal content type: %T", types.ModuleName, c)
		}
	}
}
