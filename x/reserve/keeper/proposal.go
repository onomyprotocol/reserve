package keeper

import (
	"reserve/x/reserve/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// FundTreasuryProposal submits the FundTreasuryProposal.
func (k Keeper) CreateDenomProposal(ctx sdk.Context, request *types.CreateDenomProposal) error {

	_, found := k.GetDenom(ctx, request.Metadata.Base)
	if found {
		return sdkerrors.Wrapf(types.ErrDenomExists, "Denom: %s already exists", request.Metadata.Base)
	}

	k.SetDenom(
		ctx,
		request.Metadata.Base,
		types.Denom{
			Display: request.Metadata.Display,
			Rate:    request.Rate,
			Total:   sdk.ZeroUint(),
		},
	)

	k.bankKeeper.SetDenomMetaData(ctx, *request.Metadata)

	return nil
}
