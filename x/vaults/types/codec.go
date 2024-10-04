package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

// RegisterLegacyAminoCodec registers all necessary param module types with a given LegacyAmino codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&ActiveCollateralProposal{}, "reserve/ActiveCollateralProposal", nil)
}

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

	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&ActiveCollateralProposal{},
	)

}
