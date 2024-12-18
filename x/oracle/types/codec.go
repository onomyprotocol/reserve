package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers the necessary x/oracle interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgUpdateParams{}, "oracle/MsgUpdateParams", nil)
	cdc.RegisterConcrete(&MsgRequestBandRates{}, "oracle/MsgRequestBandRates", nil)
	cdc.RegisterConcrete(&MsgUpdateBandParams{}, "oracle/MsgUpdateBandParams", nil)
	cdc.RegisterConcrete(&MsgUpdateBandOracleRequestRequest{}, "oracle/MsgUpdateBandOracleRequest", nil)
	cdc.RegisterConcrete(&MsgDeleteBandOracleRequests{}, "oracle/MsgDeleteBandOracleRequests", nil)
	cdc.RegisterConcrete(&MsgUpdateBandOracleRequestParamsRequest{}, "oracle/MsgUpdateBandOracleRequestParamsRequest", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
		&MsgRequestBandRates{},
		&MsgUpdateBandParams{},
		&MsgUpdateBandOracleRequestRequest{},
		&MsgDeleteBandOracleRequests{},
		&MsgUpdateBandOracleRequestParamsRequest{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()
	// ModuleCdc references the global x/ibc-transfer module codec. Note, the codec
	// should ONLY be used in certain instances of tests and for JSON encoding.
	//
	// The actual codec used for serialization should be provided to x/ibc transfer and
	// defined at the application level.
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}