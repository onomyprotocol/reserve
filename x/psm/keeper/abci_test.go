package keeper_test

import (
	"cosmossdk.io/math"
	"github.com/onomyprotocol/reserve/x/psm/types"
)

func (s *KeeperTestSuite) TestUpdatesStablecoinEpoch() {
	s.SetupTest()

	tests := []struct {
		name         string
		priceCurrent math.LegacyDec
		feeIn        math.LegacyDec
		feeOut       math.LegacyDec
		priceUpdate  math.LegacyDec

		expectFeeIn  math.LegacyDec
		expectFeeOut math.LegacyDec
	}{
		{
			name:         "normal",
			priceCurrent: math.LegacyMustNewDecFromStr("1"),
			feeIn:        math.LegacyMustNewDecFromStr("0.01"),
			feeOut:       math.LegacyMustNewDecFromStr("0.01"),
			priceUpdate:  math.LegacyMustNewDecFromStr("1.01"),
			expectFeeIn:  math.LegacyMustNewDecFromStr("0.005111362664117791"),
			expectFeeOut: math.LegacyMustNewDecFromStr("0.014888637335882209"),
		},
		{
			name:         "fluctuation",
			priceCurrent: math.LegacyMustNewDecFromStr("1"),
			feeIn:        math.LegacyMustNewDecFromStr("0.01"),
			feeOut:       math.LegacyMustNewDecFromStr("0.01"),
			priceUpdate:  math.LegacyMustNewDecFromStr("0.99"),
			expectFeeIn:  math.LegacyMustNewDecFromStr("0.014948314143157351"),
			expectFeeOut: math.LegacyMustNewDecFromStr("0.005051685856842649"),
		},
	}

	for _, t := range tests {
		s.Run(t.name, func() {
			sc := types.Stablecoin{
				Denom:      usdt,
				LimitTotal: limitUSDT,
				// Price:      t.priceCurrent,
				FeeIn:  t.feeIn,
				FeeOut: t.feeOut,
			}
			s.mockOracleKeeper.SetPrice(s.Ctx, sc.Denom, t.priceCurrent)
			err := s.k.FeeMaxStablecoin.Set(s.Ctx, usdt, t.feeIn.Add(t.feeOut).String())
			s.Require().NoError(err)
			err = s.k.SetStablecoin(s.Ctx, sc)
			s.Require().NoError(err)

			s.mockOracleKeeper.SetPrice(s.Ctx, usdt, t.priceUpdate)

			err = s.k.UpdatesStablecoinEpoch(s.Ctx)
			s.Require().NoError(err)

			scUpdate, found := s.k.GetStablecoin(s.Ctx, usdt)
			s.Require().True(found)
			// s.Require().Equal(t.priceUpdate, scUpdate.Price)
			s.Require().Equal(t.expectFeeIn.String(), scUpdate.FeeIn.String())
			s.Require().Equal(t.expectFeeOut.String(), scUpdate.FeeOut.String())
		})
	}

}
