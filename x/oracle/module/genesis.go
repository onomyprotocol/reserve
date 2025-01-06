package oracle

import (
	"context"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/onomyprotocol/reserve/x/oracle/keeper"
	"github.com/onomyprotocol/reserve/x/oracle/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx context.Context, k keeper.Keeper, genState types.GenesisState) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}

	for _, bandPriceState := range genState.BandPriceStates {
		err := k.SetBandPriceState(ctx, bandPriceState.Symbol, bandPriceState)
		if err != nil {
			sdkCtx.Logger().Info("can not set band price state for symbol %s", bandPriceState.Symbol)
		}
	}

	for _, bandOracleRequest := range genState.BandOracleRequests {
		err := k.SetBandOracleRequest(ctx, *bandOracleRequest)
		if err != nil {
			sdkCtx.Logger().Info("can not set band oracle request for request id %v", bandOracleRequest.RequestId)
		}
	}

	err := k.SetBandParams(ctx, genState.BandParams)
	if err != nil {
		sdkCtx.Logger().Info("can not set band params")
		// should we panic here?
		panic(err)
	}

	if genState.BandParams.IbcPortId != "" {
		k.SetPort(sdkCtx, genState.BandParams.IbcPortId)
		// Only try to bind to port if it is not already bound, since we may already own port capability
		if !k.IsBound(sdkCtx, genState.BandParams.IbcPortId) {
			// module binds to the port on InitChain
			// and claims the returned capability
			err := k.BindPort(sdkCtx, genState.BandParams.IbcPortId)
			if err != nil {
				panic(types.ErrBandPortBind.Error() + err.Error())
			}
		}
	}

	if genState.BandLatestClientId != 0 {
		err := k.SetBandLatestClientID(ctx, genState.BandLatestClientId)
		if err != nil {
			sdkCtx.Logger().Info("can not set band latest client id %v", genState.BandLatestClientId)
		}
	}

	for _, record := range genState.CalldataRecords {
		err := k.SetBandCallDataRecord(ctx, record)
		if err != nil {
			sdkCtx.Logger().Info("can not set band call data record with client id %v", record.ClientId)
		}
	}

	if genState.BandLatestRequestId != 0 {
		err := k.SetBandLatestRequestID(ctx, genState.BandLatestRequestId)
		if err != nil {
			sdkCtx.Logger().Info("can not set band latest request id %v", genState.BandLatestRequestId)
		}
	}

	err = k.SetBandOracleRequestParams(ctx, genState.BandOracleRequestParams)
	if err != nil {
		sdkCtx.Logger().Info("can not set band oracle request params")
		// should we set panic here
	}
	for _, pair := range genState.PairDecimalsRates {
		err = k.PairDecimalsRate.Set(ctx, collections.Join(pair.Base, pair.Quote), pair)
		if err != nil {
			sdkCtx.Logger().Info("can not set pair decimals rate")
		}
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx context.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Params:                  k.GetParams(ctx),
		BandParams:              k.GetBandParams(ctx),
		BandPriceStates:         k.GetAllBandPriceStates(ctx),
		BandOracleRequests:      k.GetAllBandOracleRequests(ctx),
		BandLatestClientId:      k.GetBandLatestClientID(ctx),
		CalldataRecords:         k.GetAllBandCalldataRecords(ctx),
		BandLatestRequestId:     k.GetBandLatestRequestID(ctx),
		BandOracleRequestParams: k.GetBandOracleRequestParams(ctx),
		PairDecimalsRates:       k.GetAllPairDecimalsRate(ctx),
	}
}
