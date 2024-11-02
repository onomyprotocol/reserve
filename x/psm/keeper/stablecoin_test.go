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

	err := s.k.Stablecoins.Set(s.Ctx, s1.Denom, s1)
	s.Require().NoError(err)
	err = s.k.Stablecoins.Set(s.Ctx, s2.Denom, s2)
	s.Require().NoError(err)

	stablecoin1, err := s.k.Stablecoins.Get(s.Ctx, usdt)
	s.Require().NoError(err)
	s.Require().Equal(stablecoin1.Denom, usdt)
	s.Require().Equal(stablecoin1.LimitTotal, limitUSDT)

	stablecoin2, err := s.k.Stablecoins.Get(s.Ctx, usdc)
	s.Require().NoError(err)
	s.Require().Equal(stablecoin2.Denom, usdc)
	s.Require().Equal(stablecoin2.LimitTotal, limitUSDC)

	count := 0
	err = s.k.Stablecoins.Walk(s.Ctx, nil, func(key string, value types.Stablecoin) (stop bool, err error) {
		count += 1
		return false, nil
	})

	s.Require().NoError(err)
	s.Require().Equal(count, 2)
}
