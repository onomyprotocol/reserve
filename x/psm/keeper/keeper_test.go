package keeper_test

import (
	"cosmossdk.io/math"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/onomyprotocol/reserve/app/apptesting"
	"github.com/onomyprotocol/reserve/x/psm/keeper"
	"github.com/onomyprotocol/reserve/x/psm/types"
)

var (
	usdt = "usdt"
	usdc = "usdc"

	limitUSDT = math.NewInt(1000000)
	limitUSDC = math.NewInt(1000000)
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper

	k           keeper.Keeper
	msgServer   types.MsgServer
	queryServer types.QueryServer
}

func (s *KeeperTestSuite) SetupTest() {
	s.Setup()

	s.k = s.App.PSMKeeper
	s.msgServer = keeper.NewMsgServerImpl(s.k)
	s.queryServer = keeper.NewQueryServerImpl(s.k)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (s *KeeperTestSuite) TestParams() {
	s.SetupTest()

	s.k.SetParams(s.Ctx, types.DefaultParams())

	p, err := s.k.GetParams(s.Ctx)
	s.Require().NoError(err)
	s.Require().Equal(p.LimitTotal, types.DefaultLimitTotal)
	s.Require().Equal(p.AcceptablePriceRatio, types.DefaultAcceptablePriceRatio)
}
