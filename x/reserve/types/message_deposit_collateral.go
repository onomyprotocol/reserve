package types

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDepositCollateral = "deposit_collateral"

var _ sdk.Msg = &MsgDepositCollateral{}

func NewMsgDepositCollateral(creator string, uid uint64, collateral string) *MsgDepositCollateral {
	return &MsgDepositCollateral{
		Creator:    creator,
		Uid:        uid,
		Collateral: collateral,
	}
}

func (msg *MsgDepositCollateral) Route() string {
	return RouterKey
}

func (msg *MsgDepositCollateral) Type() string {
	return TypeMsgDepositCollateral
}

func (msg *MsgDepositCollateral) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDepositCollateral) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDepositCollateral) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, err = strconv.ParseUint(strconv.FormatUint(msg.Uid, 10), 10, 64)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "drop uid is not an integer or is negative")
	}
	collateral, _ := sdk.ParseCoinNormalized(msg.Collateral)
	if !collateral.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "collateral is not a valid Coin object")
	}

	return nil
}
