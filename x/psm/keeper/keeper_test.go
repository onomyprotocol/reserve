package keeper_test

import (
	"context"
	"testing"

	"cosmossdk.io/math"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/app/apptesting"
	"github.com/onomyprotocol/reserve/x/psm/keeper"
	"github.com/onomyprotocol/reserve/x/psm/types"
)

var (
	usdt = "usdt"
	usdc = "usdc"

	amountStableInit = math.NewInt(1_000_000)

	limitUSDT = math.NewInt(100000000)
	limitUSDC = math.NewInt(100000000)
)

type MockOracleKeeper struct {
	prices map[string]math.LegacyDec
}

func (m MockOracleKeeper) SetPrice(ctx context.Context, denom string, price math.LegacyDec) {
	m.prices[denom] = price
}

func (s MockOracleKeeper) GetPrice(ctx context.Context, denom1 string, denom2 string) (math.LegacyDec, error) {
	price1, ok := s.prices[denom1]

	if !ok {
		panic("not found price " + denom1)
	}
	price2, ok := s.prices[denom2]
	if !ok {
		panic("not found price " + denom2)
	}
	p := price1.Quo(price2)
	return p, nil
}

func (s MockOracleKeeper) AddNewSymbolToBandOracleRequest(ctx context.Context, symbol string, oracleScriptId int64) error {
	_, ok := s.prices[symbol]

	if !ok {
		s.SetPrice(ctx, symbol, math.LegacyMustNewDecFromStr("1"))
	}
	return nil
}

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper

	k                keeper.Keeper
	msgServer        types.MsgServer
	queryServer      types.QueryServer
	mockOracleKeeper *MockOracleKeeper

	communityAddress sdk.Address
}

func (s *KeeperTestSuite) SetupTest() {
	s.Setup()

	mockOracleKeeper := MockOracleKeeper{
		prices: make(map[string]math.LegacyDec),
	}
	mockOracleKeeper.SetPrice(s.Ctx, types.ReserveStableCoinDenom, math.LegacyMustNewDecFromStr("1"))
	mockOracleKeeper.SetPrice(s.Ctx, types.SymBolUSD, math.LegacyMustNewDecFromStr("1"))

	s.App.PSMKeeper.OracleKeeper = mockOracleKeeper
	s.mockOracleKeeper = &mockOracleKeeper

	s.k = s.App.PSMKeeper
	s.msgServer = keeper.NewMsgServerImpl(s.k)
	s.queryServer = keeper.NewQueryServerImpl(s.k)
	s.communityAddress = s.App.AccountKeeper.GetModuleAddress("distribution")

	// mint coin for community pool,
	coinmints := sdk.NewCoins(sdk.NewCoin(usdc, math.NewInt(100_000_000_000)), sdk.NewCoin(usdt, math.NewInt(100_000_000_000)))
	err := s.App.BankKeeper.MintCoins(s.Ctx, types.ModuleName, coinmints)
	s.Require().NoError(err)
	err = s.k.BankKeeper.SendCoinsFromModuleToModule(s.Ctx, types.ModuleName, "distribution", coinmints)
	s.Require().NoError(err)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (s *KeeperTestSuite) TestParams() {
	s.SetupTest()

	err := s.k.SetParams(s.Ctx, types.DefaultParams())
	s.Require().NoError(err)

	p, err := s.k.GetParams(s.Ctx)
	s.Require().NoError(err)
	s.Require().Equal(p.LimitTotal, types.DefaultLimitTotal)
	s.Require().Equal(p.AcceptablePriceRatio, types.DefaultAcceptablePriceRatio)
}
