package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/onomyprotocol/reserve/x/reserve/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// FundTreasuryProposal submits the FundTreasuryProposal.
func (k Keeper) CreateDenomProposal(ctx sdk.Context, request *types.CreateDenomProposal) error {
	_, found := k.GetDenom(ctx, request.Denom)
	if found {
		sdkerrors.Wrapf(types.ErrDenomExists, "Denom: %s already exists", request.Denom)
	}

	k.SetDenom(ctx, types.Denom{
		display: request.Metadata
	})

	return k.bankKeeper.SendCoinsFromAccountToModule(ctx, senderAddr, types.ModuleName, amountToSend)
}
