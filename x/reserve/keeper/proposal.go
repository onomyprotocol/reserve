package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/onomyprotocol/reserve/x/reserve/types"
)

// FundTreasuryProposal submits the FundTreasuryProposal.
func (k Keeper) CreateDenomProposal(ctx sdk.Context, request *types.CreateDenomProposal) error {
	
	return k.bankKeeper.SendCoinsFromAccountToModule(ctx, senderAddr, types.ModuleName, amountToSend)
}
