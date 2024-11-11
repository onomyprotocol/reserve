package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/onomyprotocol/reserve/app/apptesting"
	"github.com/onomyprotocol/reserve/x/auction/keeper"
	"github.com/onomyprotocol/reserve/x/auction/types"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper

	k         keeper.Keeper
	msgServer types.MsgServer
}

func (s *KeeperTestSuite) SetupTest() {
	s.Setup()

	s.k = s.App.AuctionKeeper
	s.msgServer = keeper.NewMsgServerImpl(s.k)
	// s.queryServer = keeper.NewQueryServerImpl(s.k)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (s *KeeperTestSuite) TestParams() {
	s.SetupTest()

	err := s.k.SetParams(s.Ctx, types.DefaultParams())
	s.Require().NoError(err)

	s.k.GetParams(s.Ctx)
	s.Require().NoError(err)
	// s.Require().Equal(p.LimitTotal, types.DefaultLimitTotal)
	// s.Require().Equal(p.AcceptablePriceRatio, types.DefaultAcceptablePriceRatio)
}
