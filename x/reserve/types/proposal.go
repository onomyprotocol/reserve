package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	// ProposalTypeCreateDenomProposal defines the type for a CreateDenomProposal.
	ProposalTypeCreateDenomProposal = "CreateDenomProposal"
)

var (
	_ govtypes.Content = &CreateDenomProposal{}
)

func init() { // nolint:gochecknoinits // cosmos sdk style
	govtypes.RegisterProposalType(ProposalTypeCreateDenomProposal)
	govtypes.RegisterProposalTypeCodec(&CreateDenomProposal{}, fmt.Sprintf("%s/%s", ModuleName, ProposalTypeCreateDenomProposal))
}

// NewCreateDenomProposal creates a new fund treasury proposal.
func NewCreateDenomProposal(sender sdk.AccAddress, title string, description string, metadata banktypes.Metadata, rate []sdk.Uint) *CreateDenomProposal {
	return &CreateDenomProposal{sender.String(), title, description, &metadata, rate}
}

// GetTitle returns the title of a fund treasury proposal.
func (m *CreateDenomProposal) GetTitle() string { return m.Title }

// GetDescription returns the description of a fund treasury proposal.
func (m *CreateDenomProposal) GetDescription() string { return m.Description }

// ProposalRoute returns the routing key of a fund treasury proposal.
func (m *CreateDenomProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of the fund treasury proposal.
func (m *CreateDenomProposal) ProposalType() string { return ProposalTypeCreateDenomProposal }

// ValidateBasic runs basic stateless validity checks.
func (m *CreateDenomProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(m)
	if err != nil {
		return err
	}
	sender, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return err
	}
	if err := sdk.VerifyAddressFormat(sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address: %s", err)
	}

	return nil
}

// GetProposer returns the proposer from the proposal struct.
func (m *CreateDenomProposal) GetProposer() string { return m.Sender }
