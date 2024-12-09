package oracle_test

import (
	"encoding/hex"
	"time"

	"cosmossdk.io/math"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types" // nolint:all
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v8/modules/core/24-host"
	ibctesting "github.com/cosmos/ibc-go/v8/testing"
	reserveapp "github.com/onomyprotocol/reserve/app"
	oracletypes "github.com/onomyprotocol/reserve/x/oracle/types"
	utils "github.com/onomyprotocol/reserve/x/oracle/utils"
)

func (suite *PriceRelayTestSuite) TestOnChanOpenInit() {
	var (
		channel *channeltypes.Channel
		path    *ibctesting.Path
		chanCap *capabilitytypes.Capability
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{

		{
			"success", func() {}, true,
		},
		{
			"invalid port ID", func() {
				path.EndpointA.ChannelConfig.PortID = ibctesting.MockPort
			}, false,
		},
		{
			"invalid version", func() {
				channel.Version = "version"
			}, false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.TearDownTest()
			suite.SetupTest() // reset
			path = NewPriceRelayPath(suite.chainO, suite.chainB)
			suite.coordinator.SetupConnections(path)
			path.EndpointA.ChannelID = ibctesting.FirstChannelID

			counterparty := channeltypes.NewCounterparty(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)
			channel = &channeltypes.Channel{
				State:          channeltypes.INIT,
				Ordering:       channeltypes.UNORDERED,
				Counterparty:   counterparty,
				ConnectionHops: []string{path.EndpointA.ConnectionID},
				Version:        oracletypes.DefaultTestBandIbcParams().IbcVersion,
			}
			module, _, err := suite.chainO.App.GetIBCKeeper().PortKeeper.LookupModuleByPort(suite.chainO.GetContext(), oracletypes.DefaultTestBandIbcParams().IbcPortId)
			suite.Require().NoError(err)

			chanCap, err = suite.chainO.App.GetScopedIBCKeeper().NewCapability(suite.chainO.GetContext(), host.ChannelCapabilityPath(oracletypes.DefaultTestBandIbcParams().IbcPortId, path.EndpointA.ChannelID))
			suite.Require().NoError(err)

			cbs, ok := suite.chainO.App.GetIBCKeeper().Router.GetRoute(module)
			suite.Require().True(ok)

			tc.malleate() // explicitly change fields in channel and testChannel

			_, err = cbs.OnChanOpenInit(suite.chainO.GetContext(), channel.Ordering, channel.GetConnectionHops(),
				path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, chanCap, channel.Counterparty, channel.GetVersion(),
			)

			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}

		})
	}
}

func (suite *PriceRelayTestSuite) TestOnChanOpenTry() {
	var (
		channel             *channeltypes.Channel
		chanCap             *capabilitytypes.Capability
		path                *ibctesting.Path
		counterpartyVersion string
	)

	testCases := []struct {
		name          string
		malleate      func()
		expPass       bool
		expAppVersion string
	}{

		{
			"success", func() {}, true, oracletypes.DefaultTestBandIbcParams().IbcVersion,
		},
		{
			"invalid port ID", func() {
				path.EndpointA.ChannelConfig.PortID = ibctesting.MockPort
			}, false, "",
		},
		{
			"invalid channel version", func() {
				channel.Version = "version"
			}, true, oracletypes.DefaultTestBandIbcParams().IbcVersion,
		},
		{
			"invalid counterparty version", func() {
				counterpartyVersion = "version"
			}, false, "",
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.TearDownTest()
			suite.SetupTest() // reset

			path = NewPriceRelayPath(suite.chainO, suite.chainB)
			suite.coordinator.SetupConnections(path)
			path.EndpointA.ChannelID = ibctesting.FirstChannelID

			counterparty := channeltypes.NewCounterparty(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)
			channel = &channeltypes.Channel{
				State:          channeltypes.TRYOPEN,
				Ordering:       channeltypes.UNORDERED,
				Counterparty:   counterparty,
				ConnectionHops: []string{path.EndpointA.ConnectionID},
				Version:        oracletypes.DefaultTestBandIbcParams().IbcVersion,
			}
			counterpartyVersion = oracletypes.DefaultTestBandIbcParams().IbcVersion

			module, _, err := suite.chainO.App.GetIBCKeeper().PortKeeper.LookupModuleByPort(suite.chainO.GetContext(), oracletypes.DefaultTestBandIbcParams().IbcPortId)
			suite.Require().NoError(err)

			chanCap, err = suite.chainO.App.GetScopedIBCKeeper().NewCapability(suite.chainO.GetContext(), host.ChannelCapabilityPath(oracletypes.DefaultTestBandIbcParams().IbcPortId, path.EndpointA.ChannelID))
			suite.Require().NoError(err)

			cbs, ok := suite.chainO.App.GetIBCKeeper().Router.GetRoute(module)
			suite.Require().True(ok)

			tc.malleate() // explicitly change fields in channel and testChannel

			appVersion, err := cbs.OnChanOpenTry(suite.chainO.GetContext(), channel.Ordering, channel.GetConnectionHops(),
				path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, chanCap, channel.Counterparty, counterpartyVersion,
			)

			suite.Assert().Equal(tc.expAppVersion, appVersion)

			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *PriceRelayTestSuite) TestOnChanOpenAck() {
	var counterpartyVersion string

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{

		{
			"success", func() {}, true,
		},
		{
			"invalid counterparty version", func() {
				counterpartyVersion = "version"
			}, false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.TearDownTest()
			suite.SetupTest() // reset

			path := NewPriceRelayPath(suite.chainO, suite.chainB)
			suite.coordinator.SetupConnections(path)
			path.EndpointA.ChannelID = ibctesting.FirstChannelID
			counterpartyVersion = oracletypes.DefaultTestBandIbcParams().IbcVersion

			module, _, err := suite.chainO.App.GetIBCKeeper().PortKeeper.LookupModuleByPort(suite.chainO.GetContext(), oracletypes.DefaultTestBandIbcParams().IbcPortId)
			suite.Require().NoError(err)

			cbs, ok := suite.chainO.App.GetIBCKeeper().Router.GetRoute(module)
			suite.Require().True(ok)

			tc.malleate() // explicitly change fields in channel and testChannel

			err = cbs.OnChanOpenAck(suite.chainO.GetContext(), path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointA.Counterparty.ChannelID, counterpartyVersion)

			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *PriceRelayTestSuite) TestOnRecvPacket() {
	// TODO: Add more test case to cover all branch
	var packetData []byte
	var msg oracletypes.OracleResponsePacketData
	var symbolsInput = oracletypes.SymbolInput{
		Symbols:            []string{"ATOM", "BNB", "BTC", "ETH", "INJ", "USDT", "OSMO", "STX", "SOL"},
		MinimumSourceCount: 1,
	}
	data := utils.MustEncode(symbolsInput)
	testCases := []struct {
		name          string
		malleate      func()
		expAckSuccess bool
	}{
		{
			"success", func() {}, true,
		},
		{
			"fails - cannot unmarshal packet data", func() {
				packetData = []byte("invalid data")
			}, false,
		},
		{
			"fails - request is not resolved successfully", func() {
				msg.ResolveStatus = oracletypes.RESOLVE_STATUS_FAILURE
				packetData = msg.GetBytes()
			}, false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.TearDownTest()
			suite.SetupTest() // reset
			path := NewPriceRelayPath(suite.chainO, suite.chainB)
			suite.coordinator.SetupConnections(path)

			result, _ := hex.DecodeString(
				"000000090000000441544f4d00000000028510582500000003424e42000000004b269758800000000342544300000019cde9ff0a340000000345544800000001b055202e8c00000003494e4a0000000001b991af3c000000045553445400000000003ba159af000000044f534d4f00000000002b0682d70000000353545800000000002f7a459e00000003534f4c0000000004fa3a37e8",
			)
			// prepare packet
			msg = oracletypes.OracleResponsePacketData{
				ClientID:      "1",
				RequestID:     1,
				AnsCount:      1,
				RequestTime:   1000000000,
				ResolveTime:   time.Now().Unix(),
				ResolveStatus: oracletypes.RESOLVE_STATUS_SUCCESS,
				Result:        result,
			}
			packetData = msg.GetBytes()

			// modify test data
			tc.malleate()

			packet := channeltypes.NewPacket(
				packetData,
				uint64(1),
				path.EndpointA.ChannelConfig.PortID,
				path.EndpointA.ChannelID,
				path.EndpointB.ChannelConfig.PortID,
				path.EndpointB.ChannelID,
				clienttypes.NewHeight(0, 100),
				0,
			)

			// prepare expected ack
			expectedAck := channeltypes.NewResultAcknowledgement([]byte{byte(1)})
			// get module
			module, _, err := suite.chainO.App.GetIBCKeeper().PortKeeper.LookupModuleByPort(
				suite.chainO.GetContext(),
				path.EndpointA.ChannelConfig.PortID,
			)
			suite.Require().NoError(err)

			// get routeq
			cbs, ok := suite.chainO.App.GetIBCKeeper().Router.GetRoute(module)
			suite.Require().True(ok)

			onomyApp := suite.chainO.App.(*reserveapp.App)
			err = onomyApp.OracleKeeper.SetBandCallDataRecord(suite.chainO.GetContext(), &oracletypes.CalldataRecord{
				ClientId: 1,
				Calldata: data,
			})
			suite.Require().NoError(err)

			// call recv packet
			ack := cbs.OnRecvPacket(suite.chainO.GetContext(), packet, nil)

			// state check blabla
			// check result
			if tc.expAckSuccess {
				suite.Require().True(ack.Success())
				suite.Require().Equal(expectedAck, ack)

				price, err := onomyApp.OracleKeeper.GetPrice(suite.chainO.GetContext(), "ATOM", "USD")
				suite.Require().NoError(err)
				suite.Require().Equal("10.822375461000000000", price.String())
			} else {
				suite.Require().False(ack.Success())
			}
		})
	}
}

func (suite *PriceRelayTestSuite) TestPriceFeedThreshold() {

	currentBTCPrice, _ := math.LegacyNewDecFromStr("48495.410")
	withinThresholdBTCPrice, _ := math.LegacyNewDecFromStr("49523.620")
	minThresholdBTCPrice, _ := math.LegacyNewDecFromStr("484.9540")
	maxThresholdBTCPrice, _ := math.LegacyNewDecFromStr("4952362.012")

	testCases := []struct {
		name         string
		lastPrice    math.LegacyDec
		newPrice     math.LegacyDec
		expThreshold bool
	}{
		{
			"Within Threshold", math.LegacyNewDec(100), math.LegacyNewDec(120), false,
		},
		{
			"Min Threshold", math.LegacyNewDec(101), math.LegacyNewDec(1), true,
		},
		{
			"Max Threshold", math.LegacyNewDec(2), math.LegacyNewDec(201), true,
		},
		{
			"Within Threshold BTC", currentBTCPrice, withinThresholdBTCPrice, false,
		},
		{
			"Min Threshold BTC", currentBTCPrice, minThresholdBTCPrice, true,
		},
		{
			"Max Threshold BTC", currentBTCPrice, maxThresholdBTCPrice, true,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			isThresholdExceeded := oracletypes.CheckPriceFeedThreshold(tc.lastPrice, tc.newPrice)
			suite.Assert().Equal(tc.expThreshold, isThresholdExceeded)
		})
	}
}
