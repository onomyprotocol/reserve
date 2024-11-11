package keeper_test

import (
	"context"
	"testing"

	"cosmossdk.io/math"

	"github.com/stretchr/testify/suite"

	"github.com/onomyprotocol/reserve/app/apptesting"
	"github.com/onomyprotocol/reserve/x/auction/keeper"
	"github.com/onomyprotocol/reserve/x/auction/types"
	vaultstypes "github.com/onomyprotocol/reserve/x/vaults/types"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper

	k                keeper.Keeper
	msgServer        types.MsgServer
	mockOracleKeeper *MockOracleKeeper
}

func (s *KeeperTestSuite) SetupTest() {
	s.Setup()

	mockOracleKeeper := MockOracleKeeper{
		prices: make(map[string]math.LegacyDec),
	}
	mockOracleKeeper.SetPrice(s.Ctx, vaultstypes.DefaultMintDenoms[0], math.LegacyMustNewDecFromStr("1"))

	s.App.AuctionKeeper.OracleKeeper = mockOracleKeeper
	s.mockOracleKeeper = &mockOracleKeeper

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

type MockOracleKeeper struct {
	prices map[string]math.LegacyDec
}

func (m MockOracleKeeper) SetPrice(ctx context.Context, denom string, price math.LegacyDec) {
	m.prices[denom] = price
}

func (s MockOracleKeeper) GetPrice(ctx context.Context, denom1 string, denom2 string) *math.LegacyDec {
	price1, ok := s.prices[denom1]

	if !ok {
		panic("not found price " + denom1)
	}
	price2, ok := s.prices[denom2]
	if !ok {
		panic("not found price " + denom2)
	}
	p := price1.Quo(price2)
	return &p
}

func (s MockOracleKeeper) AddNewSymbolToBandOracleRequest(ctx context.Context, symbol string, oracleScriptId int64) error {
	_, ok := s.prices[symbol]

	if !ok {
		s.SetPrice(ctx, symbol, math.LegacyMustNewDecFromStr("1"))
	}
	return nil
}
