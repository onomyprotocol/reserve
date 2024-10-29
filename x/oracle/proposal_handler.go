package oracle

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/onomyprotocol/reserve/x/oracle/keeper"
	"github.com/onomyprotocol/reserve/x/oracle/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewOracleProposalHandler creates a governance handler to manage new oracles
func NewOracleProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.UpdateBandParamsProposal:
			return handleUpdateBandParamsProposal(ctx, k, c)
		case *types.UpdateBandOracleRequestProposal:
			return handleUpdateBandOracleRequestProposal(ctx, k, c)

		case *types.DeleteBandOracleRequestProposal:
			return handleDeleteBandOracleRequestProposal(ctx, k, c)
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
	if k.ShouldBound(ctx, p.BandParams.IbcPortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, p.BandParams.IbcPortId)
		if err != nil {
			return errorsmod.Wrap(types.ErrBandPortBind, err.Error())
		}
	}

	return k.SetBandParams(ctx, p.BandParams)
}

func handleUpdateBandOracleRequestProposal(ctx sdk.Context, k keeper.Keeper, p *types.UpdateBandOracleRequestProposal) error {
	if err := p.ValidateBasic(); err != nil {
		return err
	}

	request := k.GetBandOracleRequest(ctx, p.UpdateOracleRequest.RequestId)
	if request == nil {
		return errorsmod.Wrapf(types.ErrBandRequestNotFound, "cannot update requestID %T", p.UpdateOracleRequest.RequestId)
	}

	if p.UpdateOracleRequest.OracleScriptId > 0 {
		request.OracleScriptId = p.UpdateOracleRequest.OracleScriptId
	}

	if len(p.UpdateOracleRequest.Symbols) > 0 {
		request.Symbols = p.UpdateOracleRequest.Symbols
	}

	if p.UpdateOracleRequest.MinCount > 0 {
		request.MinCount = p.UpdateOracleRequest.MinCount
	}

	if p.UpdateOracleRequest.AskCount > 0 {
		request.AskCount = p.UpdateOracleRequest.AskCount
	}

	if p.UpdateOracleRequest.FeeLimit != nil {
		request.FeeLimit = p.UpdateOracleRequest.FeeLimit
	}

	if p.UpdateOracleRequest.PrepareGas > 0 {
		request.PrepareGas = p.UpdateOracleRequest.PrepareGas
	}

	if p.UpdateOracleRequest.ExecuteGas > 0 {
		request.ExecuteGas = p.UpdateOracleRequest.ExecuteGas
	}

	if p.UpdateOracleRequest.MinSourceCount > 0 {
		request.MinSourceCount = p.UpdateOracleRequest.MinSourceCount
	}

	return k.SetBandOracleRequest(ctx, *request)
}

func handleDeleteBandOracleRequestProposal(ctx sdk.Context, k keeper.Keeper, p *types.DeleteBandOracleRequestProposal) error {
	if err := p.ValidateBasic(); err != nil {
		return err
	}

	for _, requestID := range p.DeleteRequestIds {
		err := k.DeleteBandOracleRequest(ctx, requestID)
		if err != nil {
			return err
		}
	}

	return nil
}
