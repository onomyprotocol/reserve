package mock

import (
	"context"

	"cosmossdk.io/math"
)

type MockOracleKeeper struct {
	prices map[string]math.LegacyDec
}

func NewMockOracleKeeper() MockOracleKeeper {
	return MockOracleKeeper{
		prices: make(map[string]math.LegacyDec),
	}
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
