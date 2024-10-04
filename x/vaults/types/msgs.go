package types

import (
	sdkerrors "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ProposalTypeActiveCollateralProposal string = "ActiveCollateralProposal"
)

var (
	Query_serviceDesc = _Query_serviceDesc
	Msg_serviceDesc   = _Msg_serviceDesc
)

func NewMsgCreateVault(owner string, collateral, minted sdk.Coin) MsgCreateVault {
	return MsgCreateVault{
		Owner:      owner,
		Collateral: collateral,
		Minted:     minted,
	}
}

func NewMsgDeposit(vaultId uint64, sender string, amount sdk.Coin) MsgDeposit {
	return MsgDeposit{
		VaultId: vaultId,
		Sender:  sender,
		Amount:  amount,
	}
}

func NewMsgWithdraw(vaultId uint64, sender string, amount sdk.Coin) MsgWithdraw {
	return MsgWithdraw{
		VaultId: vaultId,
		Sender:  sender,
		Amount:  amount,
	}
}

func NewMsgMint(vaultId uint64, sender string, amount sdk.Coin) MsgMint {
	return MsgMint{
		VaultId: vaultId,
		Sender:  sender,
		Amount:  amount,
	}
}

func NewMsgRepay(vaultId uint64, sender string, amount sdk.Coin) MsgRepay {
	return MsgRepay{
		VaultId: vaultId,
		Sender:  sender,
		Amount:  amount,
	}
}

func NewMsgClose(vaultId uint64, sender string) MsgClose {
	return MsgClose{
		VaultId: vaultId,
		Sender:  sender,
	}
}

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
