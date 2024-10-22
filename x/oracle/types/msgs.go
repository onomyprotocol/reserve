package types

import (
	"cosmossdk.io/errors"
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// oracle message types
const (
	RouterKey               = ModuleName
	TypeMsgRequestBandRates = "requestBandRates"
	TypeMsgUpdateParams     = "updateParams"
)

var (
	_ sdk.Msg = &MsgRequestBandRates{}
	_ sdk.Msg = &MsgUpdateParams{}
	_ sdk.Msg = &MsgUpdateBandParams{}
)

func (msg MsgUpdateParams) Route() string { return RouterKey }

func (msg MsgUpdateParams) Type() string { return TypeMsgUpdateParams }

func (m *MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrap(err, "invalid authority address")
	}

	if err := m.Params.Validate(); err != nil {
		return err
	}

	return nil
}

func (msg *MsgUpdateParams) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshal(msg))
}

func (msg MsgUpdateParams) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{addr}
}

// NewMsgRequestBandRates creates a new MsgRequestBandRates instance.
func NewMsgRequestBandRates(
	sender sdk.AccAddress,
	requestID uint64,
) *MsgRequestBandRates {
	return &MsgRequestBandRates{
		Sender:    sender.String(),
		RequestId: requestID,
	}
}

// Route implements the sdk.Msg interface for MsgRequestBandRates.
func (msg MsgRequestBandRates) Route() string { return RouterKey }

// Type implements the sdk.Msg interface for MsgRequestBandRates.
func (msg MsgRequestBandRates) Type() string { return TypeMsgRequestBandRates }

// ValidateBasic implements the sdk.Msg interface for MsgRequestBandRates.
func (msg MsgRequestBandRates) ValidateBasic() error {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return err
	}
	if sender.Empty() {
		return errors.Wrapf(ErrInvalidBandRequest, "MsgRequestBandRates: Sender address must not be empty.")
	}

	if msg.RequestId == 0 {
		return errors.Wrapf(ErrInvalidBandRequest, "MsgRequestBandRates: requestID should be greater than zero")
	}
	return nil
}

// GetSigners implements the sdk.Msg interface for MsgRequestBandRates.
func (msg MsgRequestBandRates) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// GetSignBytes implements the sdk.Msg interface for MsgRequestBandRates.
func (msg MsgRequestBandRates) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}
