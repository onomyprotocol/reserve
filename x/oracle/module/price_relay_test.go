package oracle_test

import (
	"encoding/json"
	"testing"

	"cosmossdk.io/log"
	dbm "github.com/cosmos/cosmos-db"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	ibctesting "github.com/cosmos/ibc-go/v8/testing"
	testifysuite "github.com/stretchr/testify/suite"

	//"github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/onomyprotocol/reserve/app"
	reserveapp "github.com/onomyprotocol/reserve/app"
	simapp "github.com/onomyprotocol/reserve/app"
	bandapp "github.com/onomyprotocol/reserve/x/oracle/bandtesting/app"
	bandoracletypes "github.com/onomyprotocol/reserve/x/oracle/bandtesting/x/oracle/types"
	oracletypes "github.com/onomyprotocol/reserve/x/oracle/types"

)

type PriceRelayTestSuite struct {
	testifysuite.Suite

	coordinator *ibctesting.Coordinator

	// testing chains used for convenience and readability
	chainO *ibctesting.TestChain
	chainB *ibctesting.TestChain
}

func (suite *PriceRelayTestSuite) SetupTest() {
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), 0)

	// setup injective chain
	chainID := ibctesting.GetChainID(0)
	ibctesting.DefaultTestingAppInit = func() (ibctesting.TestingApp, map[string]json.RawMessage) {
		db := dbm.NewMemDB()
		encCdc := bandapp.MakeEncodingConfig()
		app, _ := reserveapp.New(log.NewNopLogger(), db, nil, true, simtestutil.EmptyAppOptions{})
		genesisState := app.DefaultGenesis()
		oracleGenesis := oracletypes.DefaultGenesis()
		oracleGenesis.BandParams = *oracletypes.DefaultTestBandIbcParams()
		oracleGenesisRaw := encCdc.Marshaler.MustMarshalJSON(oracleGenesis)
		genesisState[oracletypes.ModuleName] = oracleGenesisRaw
		return app, genesisState
	}
	suite.coordinator.Chains[chainID] = ibctesting.NewTestChain(suite.T(), suite.coordinator, chainID)

	// setup band chain
	chainID = ibctesting.GetChainID(1)
	ibctesting.DefaultTestingAppInit = func() (ibctesting.TestingApp, map[string]json.RawMessage) {
		db := dbm.NewMemDB()
		encCdc := bandapp.MakeEncodingConfig()
		app := bandapp.NewBandApp(log.NewNopLogger(), db, nil, true, map[int64]bool{}, bandapp.DefaultNodeHome, 5, encCdc, simtestutil.EmptyAppOptions{})
		return app, bandapp.NewDefaultGenesisState()
	}
	suite.coordinator.Chains[chainID] = ibctesting.NewTestChain(suite.T(), suite.coordinator, chainID)

	suite.chainO = suite.coordinator.GetChain(ibctesting.GetChainID(0))
	suite.chainB = suite.coordinator.GetChain(ibctesting.GetChainID(1))
}

func NewPriceRelayPath(chainI, chainB *ibctesting.TestChain) *ibctesting.Path {
	path := ibctesting.NewPath(chainI, chainB)
	path.EndpointA.ChannelConfig.Version = oracletypes.DefaultTestBandIbcParams().IbcVersion
	path.EndpointA.ChannelConfig.PortID = oracletypes.DefaultTestBandIbcParams().IbcPortId
	path.EndpointB.ChannelConfig.Version = oracletypes.DefaultTestBandIbcParams().IbcVersion
	path.EndpointB.ChannelConfig.PortID = oracletypes.ModuleName

	return path
}

// constructs a send from chainA to chainB on the established channel/connection
// and sends the same coin back from chainB to chainA.
func (suite *PriceRelayTestSuite) TestHandlePriceRelay() {
	// setup between chainA and chainB

	path := NewPriceRelayPath(suite.chainO, suite.chainB)

	suite.coordinator.Setup(path)

	timeoutHeight := clienttypes.NewHeight(1, 110)

	// relay send
	bandOracleReq := oracletypes.BandOracleRequest{
		OracleScriptId: 1,
		Symbols:        []string{"nom", "btc"},
		AskCount:       1,
		MinCount:       1,
		FeeLimit:       sdk.Coins{sdk.NewInt64Coin("nom", 1)},
		PrepareGas:     100,
		ExecuteGas:     200,
	}

	priceRelayPacket := oracletypes.NewOracleRequestPacketData("11", bandOracleReq.GetCalldata(true), &bandOracleReq)
	packet := channeltypes.NewPacket(priceRelayPacket.GetBytes(), 1, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, timeoutHeight, 0)
	_, err := path.EndpointA.SendPacket(packet.TimeoutHeight, packet.TimeoutTimestamp, packet.Data)
	suite.Require().NoError(err)

	// nolint:all
	// ack := channeltypes.NewResultAcknowledgement(types.ModuleCdc.MustMarshalJSON(bandoracletypes.NewOracleRequestPacketAcknowledgement(1)))
	err = path.RelayPacket(packet)
	suite.Require().NoError(err) // relay committed

	suite.chainB.NextBlock()

	oracleResponsePacket := bandoracletypes.NewOracleResponsePacketData("11", 1, 0, 1577923380, 1577923405, 1, []byte("beeb"))
	responsePacket := channeltypes.NewPacket(
		oracleResponsePacket.GetBytes(),
		1,
		path.EndpointB.ChannelConfig.PortID,
		path.EndpointB.ChannelID,
		path.EndpointA.ChannelConfig.PortID,
		path.EndpointA.ChannelID,
		clienttypes.ZeroHeight(),
		1577924005000000000,
	)

	expectCommitment := channeltypes.CommitPacket(suite.chainB.Codec, responsePacket)
	commitment := suite.chainB.App.GetIBCKeeper().ChannelKeeper.GetPacketCommitment(suite.chainB.GetContext(), path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, 1)
	suite.Equal(expectCommitment, commitment)

	injectiveApp := suite.chainO.App.(*reserveapp.App)
	injectiveApp.OracleKeeper.SetBandOracleRequest(suite.chainO.GetContext(), oracletypes.BandOracleRequest{
		RequestId:      1,
		OracleScriptId: 1,
		Symbols:        []string{"A"},
		AskCount:       1,
		MinCount:       1,
		FeeLimit:       sdk.Coins{},
		PrepareGas:     100,
		ExecuteGas:     200,
	})

	// send from chainI to chainB
	msg := oracletypes.NewMsgRequestBandRates(suite.chainO.SenderAccount.GetAddress(), 1)

	_, err = suite.chainO.SendMsgs(msg)
	suite.Require().NoError(err) // message committed
}

func (suite *PriceRelayTestSuite) TearDownTest() {
	for _, chain := range suite.coordinator.Chains {
		if app, ok := chain.App.(*app.App); ok {
			simapp.Cleanup(app) // cleanup old instance first
		}
	}
}

func TestPriceRelayTestSuite(t *testing.T) {
	testifysuite.Run(t, new(PriceRelayTestSuite))
}
