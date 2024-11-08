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
			sc := types.GetMsgStablecoin(&types.MsgAddStableCoin{
				Denom:        usdt,
				LimitTotal:   limitUSDT,
				FeeIn:        t.feeIn,
				FeeOut:       t.feeOut,
				OracleScript: 44,
			})
			s.mockOracleKeeper.SetPrice(s.Ctx, sc.Denom, t.priceCurrent)
			err := s.k.StablecoinInfos.Set(s.Ctx, sc.Denom, sc)
			s.Require().NoError(err)

			s.mockOracleKeeper.SetPrice(s.Ctx, usdt, t.priceUpdate)

			err = s.k.UpdatesStablecoinEpoch(s.Ctx)
			s.Require().NoError(err)

			scUpdate, err := s.k.StablecoinInfos.Get(s.Ctx, usdt)
			s.Require().NoError(err)
			// s.Require().Equal(t.priceUpdate, scUpdate.Price)
			s.Require().Equal(t.expectFeeIn.String(), scUpdate.FeeIn.String())
			s.Require().Equal(t.expectFeeOut.String(), scUpdate.FeeOut.String())
		})
	}

}
