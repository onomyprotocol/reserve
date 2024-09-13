package keeper_test

import (
	"cosmossdk.io/math"
	"github.com/onomyprotocol/reserve/x/psm/types"
)

func (s *KeeperTestSuite) TestQueryParams() {
	s.SetupTest()

	rp, err := s.queryServer.Params(s.Ctx, &types.QueryParamsRequest{})
	s.Require().NoError(err)
	s.Require().Equal(rp.Params.LimitTotal, types.DefaultLimitTotal)
}

func (s *KeeperTestSuite) TestStablecoin() {
	s.SetupTest()

	sc := types.Stablecoin{
		Denom:      usdt,
		LimitTotal: limitUSDT,
		Price:      math.LegacyMustNewDecFromStr("1"),
		FeeIn:      math.LegacyMustNewDecFromStr("0.001"),
		FeeOut:     math.LegacyMustNewDecFromStr("0.001"),
	}
	err := s.k.SetStablecoin(s.Ctx, sc)
	s.Require().NoError(err)

	rp, err := s.queryServer.Stablecoin(s.Ctx, &types.QueryStablecoinRequest{Denom: usdt})
	s.Require().NoError(err)
	s.Require().Equal(rp.Stablecoin, sc)
	s.Require().Equal(rp.CurrentTotal, math.ZeroInt())
	s.Require().Equal(rp.SwapableQuantity, limitUSDT)
}
