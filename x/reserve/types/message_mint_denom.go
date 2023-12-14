package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMintDenom = "mint_denom"

var _ sdk.Msg = &MsgMintDenom{}

func NewMsgMintDenom(creator string, denom string, amount string) *MsgMintDenom {
	return &MsgMintDenom{
		Creator: creator,
		Denom:   denom,
		Amount:  amount,
	}
}

func (msg *MsgMintDenom) Route() string {
	return RouterKey
}

func (msg *MsgMintDenom) Type() string {
	return TypeMsgMintDenom
}

func (msg *MsgMintDenom) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMintDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMintDenom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
