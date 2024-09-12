package types

import (
	"cosmossdk.io/math"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &MsgSwapTonomUSD{}
	_ sdk.Msg = &MsgSwapToStablecoin{}
)

func NewMsgSwapTonomUSD(addr string, coin *sdk.Coin) *MsgSwapTonomUSD {
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

var (
	Query_serviceDesc = _Query_serviceDesc
	Msg_serviceDesc   = _Msg_serviceDesc
)
