package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBond = "bond"

var _ sdk.Msg = &MsgBond{}

func NewMsgBond(creator string, denom string) *MsgBond {
	return &MsgBond{
		Creator: creator,
		Denom:   denom,
	}
}

func (msg *MsgBond) Route() string {
	return RouterKey
}

func (msg *MsgBond) Type() string {
	return TypeMsgBond
}

func (msg *MsgBond) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBond) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBond) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
