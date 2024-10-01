package types

import (
	sdkerrors "cosmossdk.io/errors"
	"cosmossdk.io/math"
)

var (
	Query_serviceDesc = _Query_serviceDesc
	Msg_serviceDesc   = _Msg_serviceDesc
)

const (
	ProposalTypeActiveCollateralProposal string = "ActiveCollateralProposal"
)

func (m *ActiveCollateralProposal) GetDescription() string {
	return " "
}

func (m *ActiveCollateralProposal) GetTitle() string {
	return " "
}

func (m *ActiveCollateralProposal) ProposalRoute() string {
	return RouterKey
}

func (m *ActiveCollateralProposal) ProposalType() string {
	return ProposalTypeActiveCollateralProposal
}

func (m *ActiveCollateralProposal) ValidateBasic() error {
	a := m.ActiveCollateral
	if a.Denom == "" {
		return sdkerrors.Wrap(ErrInvalidActiveCollateralProposal, "empty denom")
	}

	if a.MinCollateralRatio.LT(math.LegacyZeroDec()) {
		return sdkerrors.Wrap(ErrInvalidActiveCollateralProposal, "less than zero")
	}

	if a.LiquidationRatio.LT(math.LegacyZeroDec()) {
		return sdkerrors.Wrap(ErrInvalidActiveCollateralProposal, "less than zero")
	}

	if a.MaxDebt.LT(math.ZeroInt()) {
		return sdkerrors.Wrap(ErrInvalidActiveCollateralProposal, "less than zero")
	}

	return nil
}
