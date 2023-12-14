package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateVault = "create_vault"

var _ sdk.Msg = &MsgCreateVault{}

func NewMsgCreateVault(creator string, collateral string) *MsgCreateVault {
	return &MsgCreateVault{
		Creator:    creator,
		Collateral: collateral,
	}
}

func (msg *MsgCreateVault) Route() string {
	return RouterKey
}

func (msg *MsgCreateVault) Type() string {
	return TypeMsgCreateVault
}

func (msg *MsgCreateVault) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateVault) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateVault) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	collateral, _ := sdk.ParseCoinNormalized(msg.Collateral)
	if !collateral.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "collateral is not a valid Coin object")
	}
	return nil
}
