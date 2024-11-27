package mock

import (
	"context"

	"cosmossdk.io/math"
)

type MockOracleKeeper struct {
	prices map[string]math.LegacyDec
}

func NewMockOracleKeeper() *MockOracleKeeper {
	return &MockOracleKeeper{
		prices: make(map[string]math.LegacyDec),
	}
}

func (s *MockOracleKeeper) GetPrice(ctx context.Context, denom1 string, denom2 string) (math.LegacyDec, error) {
	price1, ok := s.prices[denom1]
	if !ok {
		panic("not found price" + denom1)
	}
	price2, ok := s.prices[denom2]
	if !ok {
		panic("not found price" + denom2)
	}
	p := price1.Quo(price2)
	return p, nil
}
func (s *MockOracleKeeper) SetPrice(denom string, price math.LegacyDec) {
	s.prices[denom] = price
}

func (s *MockOracleKeeper) AddNewSymbolToBandOracleRequest(ctx context.Context, symbol string, oracleScriptId int64) error {
	return nil
}
