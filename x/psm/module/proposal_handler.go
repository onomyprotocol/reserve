package psm

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	"github.com/onomyprotocol/reserve/x/psm/keeper"
	"github.com/onomyprotocol/reserve/x/psm/types"
)

func NewPSMProposalHandler(k *keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		msgSv := keeper.NewMsgServerImpl(*k)
		switch c := content.(type) {
		case *types.MsgAddStableCoin:
			_, err := msgSv.AddStableCoinProposal(ctx, c)
			return err
		case *types.MsgUpdatesStableCoin:
			_, err := msgSv.UpdatesStableCoinProposal(ctx, c)
			return err
		default:
			return errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s proposal content type: %T", types.ModuleName, c)
		}
	}
}
