package keeper_test

import (
	"cosmossdk.io/math"

	"github.com/onomyprotocol/reserve/x/psm/types"
)

func (s *KeeperTestSuite) TestStoreStablecoin() {
	s.SetupTest()

	s1 := types.Stablecoin{
		Denom:      usdt,
		LimitTotal: limitUSDT,
		FeeIn:      math.LegacyMustNewDecFromStr("0.001"),
		FeeOut:     math.LegacyMustNewDecFromStr("0.001"),
	}
	s2 := types.Stablecoin{
		Denom:      usdc,
		LimitTotal: limitUSDC,
		FeeIn:      math.LegacyMustNewDecFromStr("0.001"),
		FeeOut:     math.LegacyMustNewDecFromStr("0.001"),
	}

	err := s.k.SetStablecoin(s.Ctx, s1)
	s.Require().NoError(err)
	err = s.k.SetStablecoin(s.Ctx, s2)
	s.Require().NoError(err)

	stablecoin1, found := s.k.GetStablecoin(s.Ctx, usdt)
	s.Require().True(found)
	s.Require().Equal(stablecoin1.Denom, usdt)
	s.Require().Equal(stablecoin1.LimitTotal, limitUSDT)

	stablecoin2, found := s.k.GetStablecoin(s.Ctx, usdc)
	s.Require().True(found)
	s.Require().Equal(stablecoin2.Denom, usdc)
	s.Require().Equal(stablecoin2.LimitTotal, limitUSDC)

	count := 0
	err = s.k.IterateStablecoin(s.Ctx, func(red types.Stablecoin) (stop bool) {
		count += 1
		return false
	})
	s.Require().NoError(err)
	s.Require().Equal(count, 2)
}
