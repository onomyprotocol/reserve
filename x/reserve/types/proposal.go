package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
func NewCreateDenomProposal(sender sdk.AccAddress, title, description string, amount sdk.Coins) *CreateDenomProposal {
	return &CreateDenomProposal{sender.String(), title, description, amount}
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

	if !m.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, m.Amount.String())
	}

	if !m.Amount.IsAllPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, m.Amount.String())
	}

	return nil
}

// GetProposer returns the proposer from the proposal struct.
func (m *CreateDenomProposal) GetProposer() string { return m.Sender }

// String implements the Stringer interface.
func (m CreateDenomProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`Fund treasury proposal:
  Sender: %s
  Title: %s
  Description: %s
  Amount: %s
`, m.Sender, m.Title, m.Description, m.Amount))
	return b.String()
}
