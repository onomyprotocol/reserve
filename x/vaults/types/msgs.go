package types

import (
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	ProposalTypeActiveCollateralProposal  string = "ActiveCollateralProposal"
	ProposalTypeUpdatesCollateralProposal string = "UpdatesCollateralProposal"
	ProposalTypeBurnShortfallProposal     string = "BurnShortfallProposal"
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

func (msg *MsgCreateVault) ValidateBasic() error {
	if msg.Owner == "" {
		return fmt.Errorf("owner address is empty")
	}

	err := msg.Collateral.Validate()
	if err != nil {
		return err
	}
	return msg.Minted.Validate()
}

func (msg *MsgDeposit) ValidateBasic() error {
	if msg.Sender == "" {
		return fmt.Errorf("sender address is empty")
	}

	return msg.Amount.Validate()
}

func (msg *MsgWithdraw) ValidateBasic() error {
	if msg.Sender == "" {
		return fmt.Errorf("sender address is empty")
	}

	return msg.Amount.Validate()
}

func (msg *MsgMint) ValidateBasic() error {
	if msg.Sender == "" {
		return fmt.Errorf("sender address is empty")
	}

	return msg.Amount.Validate()
}

func (msg *MsgRepay) ValidateBasic() error {
	if msg.Sender == "" {
		return fmt.Errorf("sender address is empty")
	}

	return msg.Amount.Validate()
}

func (msg *MsgClose) ValidateBasic() error {
	if msg.Sender == "" {
		return fmt.Errorf("sender address is empty")
	}

	return nil
}

func (msg *MsgActiveCollateral) ValidateBasic() error {
	if msg.CollateralDenom == "" {
		return fmt.Errorf("denom is empty")
	}

	if msg.CollateralSymbol == "" {
		return fmt.Errorf("symbol is empty")
	}

	if msg.MintDenom == "" {
		return fmt.Errorf("mintDenom is empty")
	}

	if msg.MintSymbol == "" {
		return fmt.Errorf("mintSymbol is empty")
	}

	if msg.Authority == "" {
		return fmt.Errorf("authority is empty")
	}

	if msg.CollateralOracleScript == 0 {
		return fmt.Errorf("collateral oracle script is empty")
	}

	if msg.MintOracleScript == 0 {
		return fmt.Errorf("mint oracle script is empty")
	}

	if msg.MinCollateralRatio.LT(math.LegacyZeroDec()) {
		return fmt.Errorf("minCollateralRatio cannot be less than 0")
	}

	if msg.LiquidationRatio.LT(math.LegacyZeroDec()) {
		return fmt.Errorf("LiquidationRatio cannot be less than 0")
	}

	if msg.StabilityFee.LTE(math.LegacyZeroDec()) {
		return fmt.Errorf("StabilityFee cannot be less than 0")
	}

	if msg.LiquidationPenalty.LTE(math.LegacyZeroDec()) {
		return fmt.Errorf("LiquidationPenalty cannot be less than 0")
	}

	if msg.MaxDebt.LTE(math.ZeroInt()) {
		return fmt.Errorf("MaxDebt cannot be less than 0")
	}

	if msg.MintingFee.LT(math.LegacyZeroDec()) {
		return fmt.Errorf("MintingFee cannot be less than 0")
	}
	return nil
}

func (msg *MsgUpdatesCollateral) ValidateBasic() error {
	if msg.CollateralDenom == "" {
		return fmt.Errorf("denom is empty")
	}

	if msg.Symbol == "" {
		return fmt.Errorf("symbol is empty")
	}

	if msg.MintDenom == "" {
		return fmt.Errorf("mintDenom is empty")
	}

	if msg.Authority == "" {
		return fmt.Errorf("authority is empty")
	}

	if msg.CollateralOracleScript == 0 {
		return fmt.Errorf("collateral oracle script is empty")
	}

	if msg.MintOracleScript == 0 {
		return fmt.Errorf("mint oracle script is empty")
	}

	if msg.MinCollateralRatio.LT(math.LegacyZeroDec()) {
		return fmt.Errorf("minCollateralRatio cannot be less than 0")
	}

	if msg.LiquidationRatio.LT(math.LegacyZeroDec()) {
		return fmt.Errorf("minCollateralRatio cannot be less than 0")
	}

	if msg.StabilityFee.LTE(math.LegacyZeroDec()) {
		return fmt.Errorf("StabilityFee cannot be less than 0")
	}

	if msg.LiquidationPenalty.LT(math.LegacyZeroDec()) {
		return fmt.Errorf("minCollateralRatio cannot be less than 0")
	}

	if msg.MaxDebt.LT(math.ZeroInt()) {
		return fmt.Errorf("minCollateralRatio cannot be less than 0")
	}

	if msg.MintingFee.LT(math.LegacyZeroDec()) {
		return fmt.Errorf("minCollateralRatio cannot be less than 0")
	}
	return nil
}

var _ govtypes.Content = &ActiveCollateralProposal{}
var _ govtypes.Content = &UpdatesCollateralProposal{}
var _ govtypes.Content = &BurnShortfallProposal{}

func NewMsgActiveCollateral(a *ActiveCollateralProposal) *MsgActiveCollateral {
	return &MsgActiveCollateral{
		CollateralDenom:        a.ActiveCollateral.CollateralDenom,
		MinCollateralRatio:     a.ActiveCollateral.MinCollateralRatio,
		LiquidationRatio:       a.ActiveCollateral.LiquidationRatio,
		MaxDebt:                a.ActiveCollateral.MaxDebt,
		StabilityFee:           a.ActiveCollateral.StabilityFee,
		LiquidationPenalty:     a.ActiveCollateral.LiquidationPenalty,
		MintingFee:             a.ActiveCollateral.MintingFee,
		CollateralOracleScript: a.ActiveCollateral.CollateralOracleScript,
		MintOracleScript:       a.ActiveCollateral.MintOracleScript,
		Authority:              a.ActiveCollateral.Authority,
		CollateralSymbol:       a.ActiveCollateral.CollateralSymbol,
		MintDenom:              a.ActiveCollateral.MintDenom,
	}
}

func NewMsgUpdatesCollateral(u *UpdatesCollateralProposal) *MsgUpdatesCollateral {
	return &MsgUpdatesCollateral{
		CollateralDenom:        u.UpdatesCollateral.CollateralDenom,
		MinCollateralRatio:     u.UpdatesCollateral.MinCollateralRatio,
		LiquidationRatio:       u.UpdatesCollateral.LiquidationRatio,
		MaxDebt:                u.UpdatesCollateral.MaxDebt,
		StabilityFee:           u.UpdatesCollateral.StabilityFee,
		LiquidationPenalty:     u.UpdatesCollateral.LiquidationPenalty,
		MintingFee:             u.UpdatesCollateral.MintingFee,
		CollateralOracleScript: u.UpdatesCollateral.CollateralOracleScript,
		MintOracleScript:       u.UpdatesCollateral.MintOracleScript,
		Authority:              u.UpdatesCollateral.Authority,
		Symbol:                 u.UpdatesCollateral.Symbol,
		MintDenom:              u.UpdatesCollateral.MintDenom,
	}
}

// NewMsgBurnShortfall
func NewMsgBurnShortfall(a *BurnShortfallProposal) *MsgBurnShortfall {
	return &MsgBurnShortfall{
		Authority: a.BurnShortfall.Authority,
		MintDenom: a.BurnShortfall.MintDenom,
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

	return a.ValidateBasic()
}

func (m *UpdatesCollateralProposal) GetDescription() string {
	return " "
}

func (m *UpdatesCollateralProposal) GetTitle() string {
	return " "
}

func (m *UpdatesCollateralProposal) ProposalRoute() string {
	return RouterKey
}

func (m *UpdatesCollateralProposal) ProposalType() string {
	return ProposalTypeUpdatesCollateralProposal
}

func (m *UpdatesCollateralProposal) ValidateBasic() error {
	a := m.UpdatesCollateral

	return a.ValidateBasic()
}

// ValidateBasic
func (msg *MsgBurnShortfall) ValidateBasic() error {
	if msg.Authority == "" {
		return fmt.Errorf("authority is empty")
	}
	if msg.MintDenom == "" {
		return fmt.Errorf("authority is empty")
	}
	if msg.Amount.LTE(math.ZeroInt()) {
		return fmt.Errorf("amount cannot be less than 0")
	}
	return nil
}

func (m *BurnShortfallProposal) GetDescription() string {
	return m.Description
}

func (m *BurnShortfallProposal) GetTitle() string {
	return m.Title
}

func (m *BurnShortfallProposal) ProposalRoute() string {
	return RouterKey
}

func (m *BurnShortfallProposal) ProposalType() string {
	return ProposalTypeBurnShortfallProposal
}

func (m *BurnShortfallProposal) ValidateBasic() error {
	return m.BurnShortfall.ValidateBasic()
}
