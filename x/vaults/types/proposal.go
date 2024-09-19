package types

import (
	sdkerrors "cosmossdk.io/errors"
	"cosmossdk.io/math"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	ProposalTypeActiveCollateralProposal string = "ActiveCollateralProposal"
)

var (
	_ govtypes.Content = &ActiveCollateralProposal{}
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeActiveCollateralProposal)

}

func NewActiveCollateralProposal(title, description, denom string, minCollateralRatio, liquidationRatio math.LegacyDec, maxDebt math.Int) ActiveCollateralProposal {
	return ActiveCollateralProposal{
		Title:              title,
		Description:        description,
		Denom:              denom,
		MinCollateralRatio: minCollateralRatio,
		LiquidationRatio:   liquidationRatio,
		MaxDebt:            maxDebt,
	}
}

func (a *ActiveCollateralProposal) ProposalRoute() string { return RouterKey }

func (a *ActiveCollateralProposal) ProposalType() string {
	return ProposalTypeActiveCollateralProposal
}

func (a *ActiveCollateralProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(a)
	if err != nil {
		return err
	}

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
