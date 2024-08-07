package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateVault{}, "reserve/CreateVault", nil)
	cdc.RegisterConcrete(&MsgDeposit{}, "reserve/Deposit", nil)
	cdc.RegisterConcrete(&MsgWithdraw{}, "reserve/Withdraw", nil)
	cdc.RegisterConcrete(&MsgLiquidate{}, "reserve/Liquidate", nil)
	cdc.RegisterConcrete(&MsgBond{}, "reserve/Bond", nil)
	cdc.RegisterConcrete(&MsgUnbond{}, "reserve/Unbond", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateVault{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeposit{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWithdraw{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgLiquidate{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBond{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnbond{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
