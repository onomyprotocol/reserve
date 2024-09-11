package types

import (
	errorsmod "cosmossdk.io/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/onomyprotocol/reserve/x/oracle/utils"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// constants
const (
	ProposalUpdateBandParams           string = "ProposalUpdateBandParams"
	ProposalAuthorizeBandOracleRequest string = "ProposalTypeAuthorizeBandOracleRequest"
)

func init() {
	govtypes.RegisterProposalType(ProposalUpdateBandParams)
	govtypes.RegisterProposalType(ProposalAuthorizeBandOracleRequest)
}

// Implements Proposal Interface
var _ govtypes.Content = &UpdateBandParamsProposal{}
var _ govtypes.Content = &AuthorizeBandOracleRequestProposal{}

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
