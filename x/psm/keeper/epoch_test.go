package keeper_test

import (
	"context"

	"cosmossdk.io/math"
	"github.com/onomyprotocol/reserve/x/psm/types"
)

func (s *KeeperTestSuite) TestUpdatesStablecoinEpoch() {
	s.SetupTest()
	mockOracleKeeper := MockOracleKeeper{
		price: make(map[string]math.LegacyDec),
	}
	s.k.OracleKeeper = &mockOracleKeeper

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
			feeIn:        math.LegacyMustNewDecFromStr("0.001"),
			feeOut:       math.LegacyMustNewDecFromStr("0.001"),
			priceUpdate:  math.LegacyMustNewDecFromStr("1.01"),
			expectFeeIn:  math.LegacyMustNewDecFromStr("0.0005"),
			expectFeeOut: math.LegacyMustNewDecFromStr("0.0015"),
		},
		{
			name:         "fluctuation",
			priceCurrent: math.LegacyMustNewDecFromStr("1.05"),
			feeIn:        math.LegacyMustNewDecFromStr("0.001"),
			feeOut:       math.LegacyMustNewDecFromStr("0.001"),
			priceUpdate:  math.LegacyMustNewDecFromStr("0.95"),
			expectFeeIn:  math.LegacyMustNewDecFromStr("0.006"),
			expectFeeOut: math.LegacyMustNewDecFromStr("0.000"),
		},
	}

	for _, t := range tests {
		s.Run(t.name, func() {
			sc := types.Stablecoin{
				Denom:      usdt,
				LimitTotal: limitUSDT,
				Price:      t.priceCurrent,
				FeeIn:      t.feeIn,
				FeeOut:     t.feeOut,
			}
			err := s.k.SetStablecoin(s.Ctx, sc)
			s.Require().NoError(err)
			mockOracleKeeper.SetPrice(s.Ctx, usdt, t.priceUpdate)

			s.k.UpdatesStablecoinEpoch(s.Ctx)
			scUpdate, found := s.k.GetStablecoin(s.Ctx, usdt)
			s.Require().True(found)
			s.Require().Equal(t.priceUpdate, scUpdate.Price)
			s.Require().Equal(t.expectFeeIn, scUpdate.FeeIn)
			s.Require().Equal(t.expectFeeOut, scUpdate.FeeOut)
		})
	}

}

type MockOracleKeeper struct {
	price map[string]math.LegacyDec
}

func (m MockOracleKeeper) SetPrice(ctx context.Context, denom string, price math.LegacyDec) {
	m.price[denom] = price
}

func (m MockOracleKeeper) Price(ctx context.Context, denom string) (math.LegacyDec, error) {
	return m.price[denom], nil
}
