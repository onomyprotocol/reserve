package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"cosmossdk.io/math"
	"github.com/onomyprotocol/reserve/app/apptesting"
	"github.com/onomyprotocol/reserve/x/vaults/keeper"
	"github.com/onomyprotocol/reserve/x/vaults/keeper/mock"
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

	mockOK := mock.NewMockOracleKeeper()
	mockOK.SetPrice("atom", math.LegacyMustNewDecFromStr("8.0"))
	mockOK.SetPrice(types.DefaultMintDenoms[0], math.LegacyMustNewDecFromStr("1"))
	mockOK.SetPrice("USD", math.LegacyMustNewDecFromStr("1"))
	s.App.VaultsKeeper.OracleKeeper = mockOK
	s.k = s.App.VaultsKeeper
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

	p := s.k.GetParams(s.Ctx)
	s.Require().Equal(p.MinInitialDebt, types.DefaultMinInitialDebt)
}
