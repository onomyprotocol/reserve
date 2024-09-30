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
	ProposalTypeActiveCollateral string = "MsgActiveCollateral"
)

func (m *MsgActiveCollateral) GetDescription() string {
	return " "
}

func (m *MsgActiveCollateral) GetTitle() string {
	return " "
}

func (m *MsgActiveCollateral) ProposalRoute() string {
	return RouterKey
}

func (m *MsgActiveCollateral) ProposalType() string {
	return ProposalTypeActiveCollateral
}

func (a *MsgActiveCollateral) ValidateBasic() error {
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
