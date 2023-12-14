package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateVault{}, "reserve/CreateVault", nil)
	cdc.RegisterConcrete(&MsgDepositCollateral{}, "reserve/DepositCollateral", nil)
	cdc.RegisterConcrete(&MsgMintDenom{}, "reserve/MintDenom", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateVault{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDepositCollateral{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMintDenom{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
