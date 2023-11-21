package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/onomyprotocol/reserve/x/reserve/types"
)

// FundTreasuryProposal submits the FundTreasuryProposal.
func (k Keeper) CreateDenomProposal(ctx sdk.Context, request *types.CreateDenomProposal) error {
	senderAddr, err := sdk.AccAddressFromBech32(request.Sender)
	if err != nil {
		return err
	}

	senderBalance := k.bankKeeper.GetAllBalances(ctx, senderAddr)
	amountToSend := request.Amount
	if _, isNegative := senderBalance.SafeSub(amountToSend); isNegative {
		return sdkerrors.Wrapf(types.ErrInsufficientBalance, "sender balance is less than amount to send")
	}

	return k.bankKeeper.SendCoinsFromAccountToModule(ctx, senderAddr, types.ModuleName, amountToSend)
}
