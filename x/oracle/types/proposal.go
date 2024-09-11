package types

import (
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/onomyprotocol/reserve/x/oracle/utils"
)

// constants
const (
	ProposalUpdateBandParams           string = "ProposalUpdateBandParams"
	ProposalAuthorizeBandOracleRequest string = "ProposalTypeAuthorizeBandOracleRequest"
	ProposalUpdateBandOracleRequest    string = "ProposalUpdateBandOracleRequest"
	ProposalDeleteBandOracleRequest    string = "ProposalDeleteBandOracleRequest"
)

func init() {
	govtypes.RegisterProposalType(ProposalUpdateBandParams)
	govtypes.RegisterProposalType(ProposalAuthorizeBandOracleRequest)
	govtypes.RegisterProposalType(ProposalUpdateBandOracleRequest)
	govtypes.RegisterProposalType(ProposalDeleteBandOracleRequest)
}

// Implements Proposal Interface
var _ govtypes.Content = &UpdateBandParamsProposal{}
var _ govtypes.Content = &AuthorizeBandOracleRequestProposal{}
var _ govtypes.Content = &UpdateBandOracleRequestProposal{}
var _ govtypes.Content = &DeleteBandOracleRequestProposal{}

// GetTitle returns the title of this proposal.
func (p *UpdateBandParamsProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *UpdateBandParamsProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *UpdateBandParamsProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *UpdateBandParamsProposal) ProposalType() string {
	return ProposalUpdateBandParams
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *UpdateBandParamsProposal) ValidateBasic() error {

	if p.BandParams.IbcRequestInterval == 0 {
		return ErrBadRequestInterval
	}

	if p.BandParams.IbcSourceChannel == "" {
		return errorsmod.Wrap(ErrInvalidSourceChannel, "UpdateBandParamsProposal: IBC Source Channel must not be empty.")
	}
	if p.BandParams.IbcVersion == "" {
		return errorsmod.Wrap(ErrInvalidVersion, "UpdateBandParamsProposal: IBC Version must not be empty.")
	}

	return govtypes.ValidateAbstract(p)
}

// GetTitle returns the title of this proposal.
func (p *AuthorizeBandOracleRequestProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *AuthorizeBandOracleRequestProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *AuthorizeBandOracleRequestProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *AuthorizeBandOracleRequestProposal) ProposalType() string {
	return ProposalAuthorizeBandOracleRequest
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *AuthorizeBandOracleRequestProposal) ValidateBasic() error {
	if p.Request.OracleScriptId <= 0 {
		return errorsmod.Wrapf(ErrInvalidBandRequest, "AuthorizeBandOracleRequestProposal: Oracle script id (%d) must be positive.", p.Request.OracleScriptId)
	}

	if len(p.Request.Symbols) == 0 {
		return errorsmod.Wrap(ErrBadSymbolsCount, "AuthorizeBandOracleRequestProposal")
	}

	callData, err := utils.Encode(SymbolInput{
		Symbols:            p.Request.Symbols,
		MinimumSourceCount: uint8(p.Request.MinCount),
	})
	if err != nil {
		return err
	}

	if len(callData) > MaxDataSize {
		return errorsmod.Wrapf(ErrTooLargeCalldata, "got: %d, maximum: %d", len(callData), MaxDataSize)
	}

	if p.Request.MinCount <= 0 {
		return errorsmod.Wrapf(ErrInvalidMinCount, "AuthorizeBandOracleRequestProposal: Minimum validator count (%d) must be positive.", p.Request.MinCount)
	}

	if p.Request.AskCount <= 0 {
		return errorsmod.Wrapf(ErrInvalidAskCount, "AuthorizeBandOracleRequestProposal: Request validator count (%d) must be positive.", p.Request.AskCount)
	}

	if p.Request.AskCount < p.Request.MinCount {
		return errorsmod.Wrapf(ErrInvalidAskCount, "AuthorizeBandOracleRequestProposal: Request validator count (%d) must not be less than sufficient validator count (%d).", p.Request.AskCount, p.Request.MinCount)
	}

	if !p.Request.FeeLimit.IsValid() {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, "AuthorizeBandOracleRequestProposal: Invalid Fee Limit (%s)", p.Request.GetFeeLimit().String())
	}

	if p.Request.PrepareGas <= 0 {
		return errorsmod.Wrapf(ErrInvalidOwasmGas, "AuthorizeBandOracleRequestProposal: Invalid Prepare Gas (%d)", p.Request.GetPrepareGas())
	}

	if p.Request.ExecuteGas <= 0 {
		return errorsmod.Wrapf(ErrInvalidOwasmGas, "AuthorizeBandOracleRequestProposal: Invalid Execute Gas (%d)", p.Request.ExecuteGas)
	}

	return govtypes.ValidateAbstract(p)
}

// GetTitle returns the title of this proposal.
func (p *UpdateBandOracleRequestProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *UpdateBandOracleRequestProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *UpdateBandOracleRequestProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *UpdateBandOracleRequestProposal) ProposalType() string {
	return ProposalUpdateBandOracleRequest
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *UpdateBandOracleRequestProposal) ValidateBasic() error {
	if p.UpdateOracleRequest == nil {
		return ErrInvalidBandUpdateRequest
	}

	if p.UpdateOracleRequest != nil && len(p.UpdateOracleRequest.Symbols) > 0 {
		callData, err := utils.Encode(SymbolInput{
			Symbols:            p.UpdateOracleRequest.Symbols,
			MinimumSourceCount: uint8(p.UpdateOracleRequest.MinCount),
		})

		if err != nil {
			return err
		}

		if len(callData) > MaxDataSize {
			return errorsmod.Wrapf(ErrTooLargeCalldata, "got: %d, maximum: %d", len(callData), MaxDataSize)
		}
	}

	if p.UpdateOracleRequest != nil && p.UpdateOracleRequest.AskCount > 0 && p.UpdateOracleRequest.MinCount > 0 && p.UpdateOracleRequest.AskCount < p.UpdateOracleRequest.MinCount {
		return errorsmod.Wrapf(ErrInvalidAskCount, "UpdateBandOracleRequestProposal: Request validator count (%d) must not be less than sufficient validator count (%d).", p.UpdateOracleRequest.AskCount, p.UpdateOracleRequest.MinCount)
	}

	if p.UpdateOracleRequest != nil && p.UpdateOracleRequest.FeeLimit != nil && !p.UpdateOracleRequest.FeeLimit.IsValid() {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, "UpdateBandOracleRequestProposal: Invalid Fee Limit (%s)", p.UpdateOracleRequest.GetFeeLimit().String())
	}

	if p.UpdateOracleRequest != nil && p.UpdateOracleRequest.PrepareGas <= 0 && p.UpdateOracleRequest.ExecuteGas > 0 {
		return errorsmod.Wrapf(ErrInvalidOwasmGas, "UpdateBandOracleRequestProposal: Invalid Prepare Gas (%d)", p.UpdateOracleRequest.PrepareGas)
	}

	if p.UpdateOracleRequest != nil && p.UpdateOracleRequest.ExecuteGas <= 0 {
		return errorsmod.Wrapf(ErrInvalidOwasmGas, "UpdateBandOracleRequestProposal: Invalid Execute Gas (%d)", p.UpdateOracleRequest.ExecuteGas)
	}

	return govtypes.ValidateAbstract(p)
}

// GetTitle returns the title of this proposal.
func (p *DeleteBandOracleRequestProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *DeleteBandOracleRequestProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *DeleteBandOracleRequestProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *DeleteBandOracleRequestProposal) ProposalType() string {
	return ProposalDeleteBandOracleRequest
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *DeleteBandOracleRequestProposal) ValidateBasic() error {
	if len(p.DeleteRequestIds) == 0 {
		return ErrInvalidBandDeleteRequest
	}

	return govtypes.ValidateAbstract(p)
}
