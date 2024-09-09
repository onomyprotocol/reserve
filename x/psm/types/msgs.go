package types

import (
	"cosmossdk.io/math"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &MsgSwapToIST{}
	_ sdk.Msg = &MsgSwapToStablecoin{}
)

func NewMsgSwapToIST(addr string, coin *sdk.Coin) *MsgSwapToIST {
	return &MsgSwapToIST{
		Address: addr,
		Coin:    coin,
	}
}

func (msg MsgSwapToIST) ValidateBasic() error {
	if msg.Address == "" {
		return fmt.Errorf("empty address")
	}

	return msg.Coin.Validate()
}

func (msg MsgSwapToIST) GetSigners() []sdk.AccAddress {
	acc, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{acc}
}

// Route implements the sdk.Msg interface.
func (msg MsgSwapToIST) Route() string { return RouterKey }

// ///////////
func NewMsgSwapToStablecoin(addr, toDenom string, amount math.Int) *MsgSwapToStablecoin {
	return &MsgSwapToStablecoin{
		Address: addr,
		ToDenom: toDenom,
		Amount:  amount,
	}
}

func (msg MsgSwapToStablecoin) ValidateBasic() error {
	if msg.Address == "" {
		return fmt.Errorf("empty address")
	}
	if msg.ToDenom == "" {
		return fmt.Errorf("empty denom")
	}
	if msg.Amount.LT(math.ZeroInt()) {
		return fmt.Errorf("total limit less than zero")
	}
	return nil
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
