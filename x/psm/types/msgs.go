package types

import (
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

var (
	_ sdk.Msg              = &MsgStableSwap{}
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

func NewMsgStableSwap(addr string, coin sdk.Coin, expectedDenom string) *MsgStableSwap {
	return &MsgStableSwap{
		Address:       addr,
		OfferCoin:     coin,
		ExpectedDenom: expectedDenom,
	}
}

func (msg MsgStableSwap) ValidateBasic() error {
	if msg.Address == "" {
		return fmt.Errorf("empty address")
	}

	return msg.OfferCoin.Validate()
}

func (msg MsgStableSwap) GetSigners() []sdk.AccAddress {
	acc, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{acc}
}

// Route implements the sdk.Msg interface.
func (msg MsgStableSwap) Route() string { return RouterKey }

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

	if msg.OracleScript <= 0 {
		return sdkerrors.Wrap(ErrInvalidAddStableCoinProposal, "empty oracle script")
	}

	if msg.LimitTotal.LT(math.ZeroInt()) {
		return sdkerrors.Wrap(ErrInvalidAddStableCoinProposal, "limittotal less than zero")
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

	if msg.OracleScript <= 0 {
		return sdkerrors.Wrap(ErrInvalidAddStableCoinProposal, "empty oracle script")
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
