package oracle_test

import (
	"cosmossdk.io/math"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v8/modules/core/24-host"
	ibctesting "github.com/cosmos/ibc-go/v8/testing"

	oracletypes "github.com/onomyprotocol/reserve/x/oracle/types"
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
