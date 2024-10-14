package oracle_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/onomyprotocol/reserve/app"
	"github.com/onomyprotocol/reserve/x/oracle/types"
	"github.com/onomyprotocol/reserve/x/oracle"
)

func TestUpdateBandParamsProposal(t *testing.T) {
	// check default band params
	app := app.Setup(t, false)
	ctx := app.BaseApp.NewContextLegacy(false, tmproto.Header{Height: 1, ChainID: "3", Time: time.Unix(1618997040, 0)})

	bandParams := app.OracleKeeper.GetBandParams(ctx)
	require.Equal(t ,types.DefaultBandParams(), bandParams)

	handler := oracle.NewOracleProposalHandler(app.OracleKeeper)
	new_BandParams := types.BandParams{
		IbcRequestInterval: 2,
		IbcSourceChannel: "channel-1",
		IbcVersion: "bandchain-2",
		IbcPortId: "oracle",
	}
	err := handler(ctx, &types.UpdateBandParamsProposal{
		Title: "Update Band param proposal",
		Description: "Update band param proposal",
		BandParams: new_BandParams,
	})

	require.NoError(t, err)
	portID := app.OracleKeeper.GetPort(ctx)
	require.Equal(t, new_BandParams.IbcPortId, portID)

	isBound := app.OracleKeeper.IsBound(ctx, portID)
	require.False(t, isBound)

	bandParams = app.OracleKeeper.GetBandParams(ctx)
	require.Equal(t ,new_BandParams, bandParams)

}
