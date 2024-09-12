package vaults

import (
	"context"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"cosmossdk.io/core/appmodule"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/onomyprotocol/reserve/x/vaults/keeper"
	"github.com/onomyprotocol/reserve/x/vaults/types"
)

const consensusVersion uint64 = 1

var (
	_ module.AppModule          = AppModule{}
	_ appmodule.HasBeginBlocker = (*AppModule)(nil)
)

type AppModuleBasic struct {
	keeper keeper.Keeper
}

func (a AppModuleBasic) Name() string {
	return types.ModuleName
}

// DefaultGenesis is an empty object
func (AppModuleBasic) DefaultGenesis(_ codec.JSONCodec) json.RawMessage {
	return []byte("{}")
}

func (AppModuleBasic) ValidateGenesis(_ codec.JSONCodec, config client.TxEncodingConfig, _ json.RawMessage) error {
	return nil
}

func (a AppModule) ExportGenesis(_ sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	return a.DefaultGenesis(cdc)
}

func (a AppModule) InitGenesis(ctx sdk.Context, marshaler codec.JSONCodec, message json.RawMessage) {
	var genesisState types.GenesisState
	marshaler.MustUnmarshalJSON(message, &genesisState)
	a.keeper.InitGenesis(ctx, genesisState)
}

func (AppModule) ConsensusVersion() uint64 { return consensusVersion }

func (a AppModuleBasic) RegisterRESTRoutes(_ client.Context, _ *mux.Router) {
}

func (a AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
}

func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	return nil
}

func (a AppModuleBasic) GetQueryCmd() *cobra.Command {
	return nil
}

func (a AppModuleBasic) RegisterLegacyAminoCodec(_ *codec.LegacyAmino) {
}

type AppModule struct {
	AppModuleBasic
}

func NewAppModule() *AppModule {
	return &AppModule{}
}

func (a AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {
}

func (a AppModule) RegisterServices(_ module.Configurator) {
}

func (a AppModule) BeginBlock(_ context.Context) error {
	return nil
}

func (a AppModule) EndBlock(_ sdk.Context) []abci.ValidatorUpdate {
	return nil
}

func (AppModule) IsOnePerModuleType() {}

func (AppModule) IsAppModule() {}

// RegisterInterfaces registers a module's interface types and their concrete implementations as proto.Message.
func (a AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(reg)
}
