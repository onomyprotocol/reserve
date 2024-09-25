package keeper_test

import (
	"testing"
	"time"

	"cosmossdk.io/math"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/app"
	"github.com/onomyprotocol/reserve/x/oracle/types"
	"github.com/stretchr/testify/require"
)

func TestBandPriceState(t *testing.T) {
	app := app.Setup(t, false)
	ctx := app.BaseApp.NewContextLegacy(false, tmproto.Header{Height: 1, ChainID: "3", Time: time.Unix(1618997040, 0)})

	// Band price state is nil now
	data := app.OracleKeeper.GetBandPriceState(ctx, "ATOM")
	require.Nil(t, data)

	states := app.OracleKeeper.GetAllBandPriceStates(ctx)
	require.Equal(t, 0, len(states))

	price := app.OracleKeeper.GetPrice(ctx, "ATOM", "USD")
	require.Nil(t, price)

	bandPriceState := &types.BandPriceState{
		Symbol:      "ATOM",
		Rate:        math.NewInt(10),
		ResolveTime: 1,
		Request_ID:  1,
		PriceState:  *types.NewPriceState(math.LegacyNewDec(10), 1),
	}
	// set band price state for ATOM
	err := app.OracleKeeper.SetBandPriceState(ctx, "ATOM", bandPriceState)
	require.NoError(t, err)

	data = app.OracleKeeper.GetBandPriceState(ctx, "ATOM")
	require.Equal(t, bandPriceState, data)

	price = app.OracleKeeper.GetPrice(ctx, "ATOM", "USD")
	expect := math.LegacyNewDec(10)
	require.Equal(t, &expect, price)

	states = app.OracleKeeper.GetAllBandPriceStates(ctx)
	require.Equal(t, 1, len(states))
}

func TestBandOracleRequest(t *testing.T) {
	app := app.Setup(t, false)
	ctx := app.BaseApp.NewContextLegacy(false, tmproto.Header{Height: 1, ChainID: "3", Time: time.Unix(1618997040, 0)})

	req := app.OracleKeeper.GetBandOracleRequest(ctx, 1)
	require.Nil(t, req)

	reqs := app.OracleKeeper.GetAllBandOracleRequests(ctx)
	require.Equal(t, 0, len(reqs))

	bandOracleRequest := types.BandOracleRequest{
		RequestId:      1,
		OracleScriptId: 1,
		Symbols:        []string{"INJ"},
		AskCount:       1,
		MinCount:       1,
		FeeLimit:       sdk.Coins{sdk.NewInt64Coin("INJ", 1)},
		PrepareGas:     100,
		ExecuteGas:     200,
	}
	err := app.OracleKeeper.SetBandOracleRequest(ctx, bandOracleRequest)
	require.NoError(t, err)

	req = app.OracleKeeper.GetBandOracleRequest(ctx, 1)
	require.Equal(t, &bandOracleRequest, req)
	reqs = app.OracleKeeper.GetAllBandOracleRequests(ctx)
	require.Equal(t, 1, len(reqs))

	// delete request and try again
	err = app.OracleKeeper.DeleteBandOracleRequest(ctx, 1)
	require.NoError(t, err)
	reqs = app.OracleKeeper.GetAllBandOracleRequests(ctx)
	require.Equal(t, 0, len(reqs))
}

func TestBandLatestClientId(t *testing.T) {
	app := app.Setup(t, false)
	ctx := app.BaseApp.NewContextLegacy(false, tmproto.Header{Height: 1, ChainID: "3", Time: time.Unix(1618997040, 0)})

	id := app.OracleKeeper.GetBandLatestClientID(ctx)
	require.Equal(t, uint64(0), id)

	err := app.OracleKeeper.SetBandLatestClientID(ctx, 10)
	require.NoError(t, err)

	id = app.OracleKeeper.GetBandLatestClientID(ctx)
	require.Equal(t, uint64(10), id)
}

func TestBandLatestRequestId(t *testing.T) {
	app := app.Setup(t, false)
	ctx := app.BaseApp.NewContextLegacy(false, tmproto.Header{Height: 1, ChainID: "3", Time: time.Unix(1618997040, 0)})

	id := app.OracleKeeper.GetBandLatestRequestID(ctx)
	require.Equal(t, uint64(0), id)

	err := app.OracleKeeper.SetBandLatestRequestID(ctx, 1)
	require.NoError(t, err)

	id = app.OracleKeeper.GetBandLatestRequestID(ctx)
	require.Equal(t, uint64(1), id)
}

func TestBandCallDataRecord(t *testing.T) {
	app := app.Setup(t, false)
	ctx := app.BaseApp.NewContextLegacy(false, tmproto.Header{Height: 1, ChainID: "3", Time: time.Unix(1618997040, 0)})

	record := app.OracleKeeper.GetBandCallDataRecord(ctx, 1)
	require.Nil(t, record)

	recordA := &types.CalldataRecord{
		ClientId: 1,
		Calldata: []byte("123"),
	}
	err := app.OracleKeeper.SetBandCallDataRecord(ctx, recordA)
	require.NoError(t, err)
	record = app.OracleKeeper.GetBandCallDataRecord(ctx, 1)
	require.Equal(t, recordA, record)

	err = app.OracleKeeper.DeleteBandCallDataRecord(ctx, 1)
	require.NoError(t, err)

	record = app.OracleKeeper.GetBandCallDataRecord(ctx, 1)
	require.Nil(t, record)
}

func TestGetPrice(t *testing.T) {
	app := app.Setup(t, false)
	ctx := app.BaseApp.NewContextLegacy(false, tmproto.Header{Height: 1, ChainID: "3", Time: time.Unix(1618997040, 0)})

	// Case 1: Base, quote price, does not exist, expect nil
	price := app.OracleKeeper.GetPrice(ctx, "ATOM", "USD")
	require.Nil(t, price)

	// Case 2: Set base price state for ATOM
	bandPriceStateATOM := &types.BandPriceState{
		Symbol:      "ATOM",
		Rate:        math.NewInt(10),
		ResolveTime: 1,
		Request_ID:  1,
		PriceState:  *types.NewPriceState(math.LegacyNewDec(10), 1),
	}
	err := app.OracleKeeper.SetBandPriceState(ctx, "ATOM", bandPriceStateATOM)
	require.NoError(t, err)

	// // Case 3: Set quote price state for USD
	bandPriceStateUSD := &types.BandPriceState{
		Symbol:      "USD",
		Rate:        math.NewInt(1),
		ResolveTime: 1,
		Request_ID:  1,
		PriceState:  *types.NewPriceState(math.LegacyNewDec(1), 1),
	}
	err = app.OracleKeeper.SetBandPriceState(ctx, "USD", bandPriceStateUSD)
	require.NoError(t, err)

	// Case 4: Base price is invalid (rate is zero), expect nil
	invalidPriceState := &types.BandPriceState{
		Symbol:      "ATOM",
		Rate:        math.NewInt(0), // Invalid base rate
		ResolveTime: 1,
		Request_ID:  1,
		PriceState:  *types.NewPriceState(math.LegacyNewDec(0), 1),
	}
	err = app.OracleKeeper.SetBandPriceState(ctx, "ATOM", invalidPriceState)
	require.NoError(t, err)

	price = app.OracleKeeper.GetPrice(ctx, "ATOM", "USD")
	require.Nil(t, price)

	// Case 5: Quote price is invalid (rate is zero), expect nil
	invalidQuoteState := &types.BandPriceState{
		Symbol:      "USD",
		Rate:        math.NewInt(0), // Invalid quote rate
		ResolveTime: 1,
		Request_ID:  1,
		PriceState:  *types.NewPriceState(math.LegacyNewDec(0), 1),
	}
	err = app.OracleKeeper.SetBandPriceState(ctx, "USD", invalidQuoteState)
	require.NoError(t, err)

	price = app.OracleKeeper.GetPrice(ctx, "ATOM", "USD")
	require.Nil(t, price)

	// Case 6: Reset to valid state for both ATOM and USD, expect valid price (10)
	err = app.OracleKeeper.SetBandPriceState(ctx, "ATOM", bandPriceStateATOM)
	require.NoError(t, err)
	err = app.OracleKeeper.SetBandPriceState(ctx, "USD", bandPriceStateUSD)
	require.NoError(t, err)

	price = app.OracleKeeper.GetPrice(ctx, "ATOM", "USD")
	expect := math.LegacyNewDec(10)
	require.Equal(t, &expect, price)

	// Case 7: Reverse price (USD to ATOM), expect 0.1
	price = app.OracleKeeper.GetPrice(ctx, "USD", "ATOM")
	expect = math.LegacyNewDec(1).Quo(math.LegacyNewDec(10)) // 0.1
	require.Equal(t, &expect, price)
}
