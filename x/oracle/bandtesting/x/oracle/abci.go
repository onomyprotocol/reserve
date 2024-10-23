package oracle

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/onomyprotocol/reserve/x/oracle/bandtesting/x/oracle/keeper"
)

// handleEndBlock cleans up the state during end block. See comment in the implementation!
func handleEndBlock(ctx sdk.Context, k keeper.Keeper) {
	// Loops through all requests to resolve all of them!
	requests := k.GetAllRequests(ctx)
	for i := range requests {
		k.ProcessRequest(ctx, requests[i])
	}
	k.DeleteAllRequests(ctx)
}
