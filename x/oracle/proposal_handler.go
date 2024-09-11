package keeper

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/onomyprotocol/reserve/x/oracle/keeper"
	"github.com/onomyprotocol/reserve/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "cosmossdk.io/errors"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewOracleProposalHandler creates a governance handler to manage new oracles
func NewOracleProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.UpdateBandParamsProposal:
			return handleUpdateBandParamsProposal(ctx, k, c)
		case *types.AuthorizeBandOracleRequestProposal:
			return handleAuthorizeBandOracleRequestProposal(ctx, k, c)
		default:
			return errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized param proposal content type: %T", c)
		}
	}
}

func handleUpdateBandParamsProposal(ctx sdk.Context, k keeper.Keeper, p *types.UpdateBandParamsProposal) error {
	if err := p.ValidateBasic(); err != nil {
		return err
	}

	k.SetPort(ctx, p.BandParams.IbcPortId)
	// Only try to bind to port if it is not already bound, since we may already own port capability
	if !k.IsBound(ctx, p.BandParams.IbcPortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, p.BandParams.IbcPortId)
		if err != nil {
			return errorsmod.Wrap(types.ErrBadIBCPortBind, err.Error())
		}
	}

	k.SetBandParams(ctx, p.BandParams)
	return nil
}

func handleAuthorizeBandOracleRequestProposal(ctx sdk.Context, k keeper.Keeper, p *types.AuthorizeBandOracleRequestProposal) error {
	if err := p.ValidateBasic(); err != nil {
		return err
	}

	requestID := k.GetBandLatestRequestID(ctx) + 1
	p.Request.RequestId = requestID

	k.SetBandOracleRequest(ctx, p.Request)

	k.SetBandLatestRequestID(ctx, requestID)
	return nil
}
