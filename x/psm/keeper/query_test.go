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

	sc := types.StablecoinInfo{
		Denom:               usdt,
		LimitTotal:          limitUSDT,
		FeeIn:               math.LegacyMustNewDecFromStr("0.001"),
		FeeOut:              math.LegacyMustNewDecFromStr("0.001"),
		TotalStablecoinLock: math.ZeroInt(),
		FeeMaxStablecoin:    math.LegacyZeroDec(),
	}
	err := s.k.StablecoinInfos.Set(s.Ctx, sc.Denom, sc)
	s.Require().NoError(err)

	rp, err := s.queryServer.Stablecoin(s.Ctx, &types.QueryStablecoinRequest{Denom: usdt})
	s.Require().NoError(err)
	s.Require().Equal(rp.Stablecoin.Denom, sc.Denom)
	s.Require().Equal(rp.Stablecoin.LimitTotal, sc.LimitTotal)
	s.Require().Equal(rp.Stablecoin.TotalStablecoinLock, sc.TotalStablecoinLock)
	s.Require().Equal(rp.SwapableQuantity, limitUSDT)
}
