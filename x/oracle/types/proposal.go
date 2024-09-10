package types

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	errorsmod "cosmossdk.io/errors"
)

// constants
const (
	ProposalUpdateBandParamsProposal string = "ProposalTypeEnableBandIBC"
)

func init() {
	govtypes.RegisterProposalType(ProposalUpdateBandParamsProposal)
}

// Implements Proposal Interface
var _ govtypes.Content = &UpdateBandParamsProposal{}

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
	return ProposalUpdateBandParamsProposal
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
