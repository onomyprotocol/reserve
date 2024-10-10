package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
		&MsgActiveCollateral{},
		&MsgUpdatesCollateral{},
		&MsgCreateVault{},
		&MsgDeposit{},
		&MsgWithdraw{},
		&MsgMint{},
		&MsgRepay{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)

}
