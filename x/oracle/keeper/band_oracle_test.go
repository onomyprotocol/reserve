package keeper_test

import (
	"testing"
	"time"

	"cosmossdk.io/math"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/app"
	"github.com/onomyprotocol/reserve/x/oracle/types"
	"github.com/onomyprotocol/reserve/x/oracle/utils"
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

func TestProcessBandOraclePrices(t *testing.T) {
	// Set up the application and context
	app := app.Setup(t, false)
	ctx := app.BaseApp.NewContextLegacy(false, tmproto.Header{Height: 1, ChainID: "3", Time: time.Unix(1618997040, 0)})

	// Define table-driven test cases
	tests := []struct {
		name          string
		clientID      string
		calldata      *types.CalldataRecord
		oracleOutput  interface{}
		expectedError bool
		expectedRate  int64
	}{
		{
			name:          "Fail when no CallDataRecord found",
			clientID:      "1",
			calldata:      nil,
			oracleOutput:  nil,
			expectedError: false,
		},
		{
			name:     "Fail when decoding OracleInput",
			clientID: "1",
			calldata: &types.CalldataRecord{
				ClientId: 1,
				Calldata: []byte{0xFF, 0xFF},
			},
			oracleOutput:  nil,
			expectedError: true,
		},
		{
			name:     "Fail when decoding OracleOutput",
			clientID: "1",
			calldata: &types.CalldataRecord{
				ClientId: 1,
				Calldata: utils.MustEncode(types.Input{
					Symbols:    []string{"ATOM", "BTC"},
					Multiplier: types.BandPriceMultiplier,
				}),
			},
			oracleOutput:  []byte{0xFF, 0xFF},
			expectedError: true,
		},
		{
			name:     "Success with valid OracleResponsePacketData",
			clientID: "1",
			calldata: &types.CalldataRecord{
				ClientId: 1,
				Calldata: utils.MustEncode(types.Input{
					Symbols:    []string{"ATOM", "BTC"},
					Multiplier: types.BandPriceMultiplier,
				}),
			},
			oracleOutput: types.BandOutput{
				Responses: []types.Response{
					{Symbol: "ATOM", ResponseCode: 0, Rate: 100},
					{Symbol: "BTC", ResponseCode: 0, Rate: 50000},
				},
			},
			expectedError: false,
			expectedRate:  100,
		},
		{
			name:          "Fail when ClientID is not a valid integer",
			clientID:      "invalid-id", 
			calldata:      nil,
			oracleOutput:  nil,
			expectedError: true,
		},
	}

	// Iterate over each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.calldata != nil {
				err := app.OracleKeeper.SetBandCallDataRecord(ctx, tt.calldata)
				require.NoError(t, err)
			}

			var result []byte
			if tt.oracleOutput != nil {
				switch v := tt.oracleOutput.(type) {
				case types.BandOutput:
					result = utils.MustEncode(v)
				case []byte:
					result = v
				}
			}

			oraclePacketData := types.OracleResponsePacketData{
				ClientID:    tt.clientID,
				RequestID:   1,
				Result:      result,
				ResolveTime: time.Now().Unix(),
			}

			relayer := sdk.AccAddress("mock-relayer-address")

			err := app.OracleKeeper.ProcessBandOraclePrices(ctx, relayer, oraclePacketData)

			if tt.expectedError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)

				record := app.OracleKeeper.GetBandCallDataRecord(ctx, 1)
				require.Nil(t, record)

				if tt.expectedRate != 0 {
					priceState := app.OracleKeeper.GetBandPriceState(ctx, "ATOM")
					require.NotNil(t, priceState)
					require.Equal(t, tt.expectedRate, priceState.Rate.Int64())
				}
			}
		})
	}
}
