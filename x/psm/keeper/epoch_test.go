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
	mockOracleKeeper.SetPrice(s.Ctx, types.DenomStable, math.LegacyMustNewDecFromStr("1"))

	tests := []struct {
		name          string
		priceCurrent  math.LegacyDec
		feeInCurrent  math.LegacyDec
		feeOutCurrent math.LegacyDec
		priceUpdate   math.LegacyDec

		expectFeeIn  math.LegacyDec
		expectFeeOut math.LegacyDec
	}{
		{
			name:          "normal",
			priceCurrent:  math.LegacyMustNewDecFromStr("1"),
			feeInCurrent:  math.LegacyMustNewDecFromStr("0.001"),
			feeOutCurrent: math.LegacyMustNewDecFromStr("0.001"),
			priceUpdate:   math.LegacyMustNewDecFromStr("1.01"),
			expectFeeIn:   math.LegacyMustNewDecFromStr("0.0005"),
			expectFeeOut:  math.LegacyMustNewDecFromStr("0.0015"),
		},
		{
			name:          "fluctuation",
			priceCurrent:  math.LegacyMustNewDecFromStr("1.05"),
			feeInCurrent:  math.LegacyMustNewDecFromStr("0.001"),
			feeOutCurrent: math.LegacyMustNewDecFromStr("0.001"),
			priceUpdate:   math.LegacyMustNewDecFromStr("0.95"),
			expectFeeIn:   math.LegacyMustNewDecFromStr("0.006"),
			expectFeeOut:  math.LegacyMustNewDecFromStr("0.000"),
		},
	}

	for _, t := range tests {
		s.Run(t.name, func() {
			sc := types.Stablecoin{
				Denom:      usdt,
				LimitTotal: limitUSDT,
				Price:      t.priceCurrent,
				FeeIn:      t.feeInCurrent,
				FeeOut:     t.feeOutCurrent,
			}
			err := s.k.SetStablecoin(s.Ctx, sc)
			s.Require().NoError(err)
			mockOracleKeeper.SetPrice(s.Ctx, usdt, t.priceUpdate)

			err = s.k.UpdatesStablecoinEpoch(s.Ctx)
			s.Require().NoError(err)

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

// return multiper denom1 = 1denom2, nomal in module psm denom2 is nomUSD
func (m MockOracleKeeper) GetPrice(ctx context.Context, denom1 string, denom2 string) *math.LegacyDec {
	price1, ok := m.price[denom1]

	if !ok {
		panic("can not get price for " + denom1)
	}
	price2, ok := m.price[denom2]
	if !ok {
		panic("can not get price for " + denom2)
	}

	price := price1.Quo(price2)
	return &price
}
