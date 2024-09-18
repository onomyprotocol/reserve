package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/onomyprotocol/reserve/app/apptesting"
	"github.com/onomyprotocol/reserve/x/vaults/keeper"
	"github.com/onomyprotocol/reserve/x/vaults/types"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper

	k         keeper.Keeper
	msgServer types.MsgServer
	// queryServer types.QueryServer
}

func (s *KeeperTestSuite) SetupTest() {
	s.Setup()

	s.k = s.App.VaultsKeeper
	s.msgServer = keeper.NewMsgServerImpl(s.k)
	// s.queryServer = keeper.NewQueryServerImpl(s.k)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (s *KeeperTestSuite) TestParams() {
	s.SetupTest()

	s.k.SetParams(s.Ctx, types.DefaultParams())

	p := s.k.GetParams(s.Ctx)
	s.Require().Equal(p.MintingFee, types.DefaultMintingFee)
	s.Require().Equal(p.StabilityFee, types.DefaultStabilityFee)
	s.Require().Equal(p.LiquidationPenalty, types.DefaultLiquidationPenalty)
	s.Require().Equal(p.MinInitialDebt, types.DefaultMinInitialDebt)
	s.Require().Equal(p.RecalculateDebtPeriod, types.DefaultRecalculateDebtPeriod)
	s.Require().Equal(p.LiquidatePeriod, types.DefaultLiquidatePeriod)
}
