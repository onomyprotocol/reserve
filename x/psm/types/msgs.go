package types

import (
	sdkerrors "cosmossdk.io/errors"
	"cosmossdk.io/math"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

var (
	_ sdk.Msg              = &MsgSwapTonomUSD{}
	_ sdk.Msg              = &MsgSwapToStablecoin{}
	_ sdk.Msg              = &MsgAddStableCoin{}
	_ sdk.Msg              = &MsgUpdatesStableCoin{}
	_ getStablecoinFromMsg = &MsgAddStableCoin{}
	_ getStablecoinFromMsg = &MsgUpdatesStableCoin{}

	_ govtypes.Content = &MsgAddStableCoin{}
	_ govtypes.Content = &MsgUpdatesStableCoin{}
)

const (
	ProposalTypeAddStableCoinProposal     string = "MsgAddStableCoin"
	ProposalTypeUpdatesStableCoinProposal string = "MsgUpdatesStableCoin"
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeAddStableCoinProposal)
	govtypes.RegisterProposalType(ProposalTypeUpdatesStableCoinProposal)

}

func NewMsgSwapTonomUSD(addr string, coin sdk.Coin) *MsgSwapTonomUSD {
	return &MsgSwapTonomUSD{
		Address: addr,
		Coin:    coin,
	}
}

func (msg MsgSwapTonomUSD) ValidateBasic() error {
	if msg.Address == "" {
		return fmt.Errorf("empty address")
	}

	return msg.Coin.Validate()
}

func (msg MsgSwapTonomUSD) GetSigners() []sdk.AccAddress {
	acc, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{acc}
}

// Route implements the sdk.Msg interface.
func (msg MsgSwapTonomUSD) Route() string { return RouterKey }

// ///////////
func NewMsgSwapToStablecoin(addr, toDenom string, amount sdk.Coin) *MsgSwapToStablecoin {
	return &MsgSwapToStablecoin{
		Address: addr,
		ToDenom: toDenom,
		Coin:    amount,
	}
}

func (msg MsgSwapToStablecoin) ValidateBasic() error {
	if msg.Address == "" {
		return fmt.Errorf("empty address")
	}
	if msg.ToDenom == "" {
		return fmt.Errorf("empty denom")
	}

	return msg.Coin.Validate()
}

func (msg MsgSwapToStablecoin) GetSigners() []sdk.AccAddress {
	acc, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{acc}
}

// Route implements the sdk.Msg interface.
func (msg MsgSwapToStablecoin) Route() string { return RouterKey }

var (
	Query_serviceDesc = _Query_serviceDesc
	Msg_serviceDesc   = _Msg_serviceDesc
)

// func (msg MsgAddStableCoin) GetPrice() math.LegacyDec {
// 	return msg.Price
// }

func (msg MsgAddStableCoin) GetLimitTotal() math.Int {
	return msg.LimitTotal
}

func (msg MsgAddStableCoin) GetFeeIn() math.LegacyDec {
	return msg.FeeIn
}
func (msg MsgAddStableCoin) GetFeeOut() math.LegacyDec {
	return msg.FeeOut
}

func (a *MsgAddStableCoin) ProposalRoute() string { return RouterKey }

func (a *MsgAddStableCoin) ProposalType() string {
	return ProposalTypeAddStableCoinProposal
}

func (a *MsgAddStableCoin) GetDescription() string { return RouterKey }
func (a *MsgAddStableCoin) GetTitle() string       { return RouterKey }

func (msg MsgAddStableCoin) ValidateBasic() error {
	if msg.Denom == "" {
		return sdkerrors.Wrap(ErrInvalidAddStableCoinProposal, "empty denom")
	}

	if msg.LimitTotal.LT(math.ZeroInt()) {
		return sdkerrors.Wrap(ErrInvalidAddStableCoinProposal, "limittotal less than zero")
	}

	if msg.NomType == "" {
		return sdkerrors.Wrap(ErrInvalidAddStableCoinProposal, "empty nom type")
	}

	if msg.FeeIn.LT(math.LegacyZeroDec()) {
		return sdkerrors.Wrap(ErrInvalidAddStableCoinProposal, "feein less than zero")
	}

	if msg.FeeOut.LT(math.LegacyZeroDec()) {
		return sdkerrors.Wrap(ErrInvalidAddStableCoinProposal, "feeout less than zero")
	}

	return nil
}

// func (msg MsgUpdatesStableCoin) GetPrice() math.LegacyDec {
// 	return msg.Price
// }

func (msg MsgUpdatesStableCoin) GetLimitTotal() math.Int {
	return msg.LimitTotal
}

func (msg MsgUpdatesStableCoin) GetFeeIn() math.LegacyDec {
	return msg.FeeIn
}
func (msg MsgUpdatesStableCoin) GetFeeOut() math.LegacyDec {
	return msg.FeeOut
}

func (msg MsgUpdatesStableCoin) ValidateBasic() error {
	if msg.Denom == "" {
		return sdkerrors.Wrap(ErrInvalidAddStableCoinProposal, "empty denom")
	}

	if msg.LimitTotal.LT(math.ZeroInt()) {
		return sdkerrors.Wrap(ErrInvalidAddStableCoinProposal, "limittotal less than zero")
	}

	if msg.NomType == "" {
		return sdkerrors.Wrap(ErrInvalidAddStableCoinProposal, "empty nom type")
	}

	if msg.FeeIn.LT(math.LegacyZeroDec()) {
		return sdkerrors.Wrap(ErrInvalidAddStableCoinProposal, "feein less than zero")
	}

	if msg.FeeOut.LT(math.LegacyZeroDec()) {
		return sdkerrors.Wrap(ErrInvalidAddStableCoinProposal, "feeout less than zero")
	}

	return nil
}

func (a *MsgUpdatesStableCoin) ProposalRoute() string { return RouterKey }

func (a *MsgUpdatesStableCoin) ProposalType() string {
	return ProposalTypeUpdatesStableCoinProposal
}

func (a *MsgUpdatesStableCoin) GetDescription() string { return RouterKey }
func (a *MsgUpdatesStableCoin) GetTitle() string       { return RouterKey }
