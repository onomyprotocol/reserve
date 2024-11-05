package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

// RegisterLegacyAminoCodec registers the necessary x/oracle interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgUpdateParams{}, "oracle/MsgUpdateParams", nil)
	cdc.RegisterConcrete(&MsgRequestBandRates{}, "oracle/MsgRequestBandRates", nil)

	cdc.RegisterConcrete(&UpdateBandParamsProposal{}, "oracle/UpdateBandParamsProposal", nil)
	cdc.RegisterConcrete(&UpdateBandOracleRequestProposal{}, "oracle/UpdateBandOracleRequestProposal", nil)
	cdc.RegisterConcrete(&DeleteBandOracleRequestProposal{}, "oracle/DeleteBandOracleRequestProposal", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
		&MsgRequestBandRates{},
		&MsgSetPrice{},
	)

	registry.RegisterImplementations((*govtypes.Content)(nil),
		&UpdateBandParamsProposal{},
		&UpdateBandOracleRequestProposal{},
		&DeleteBandOracleRequestProposal{},
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
