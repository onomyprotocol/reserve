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

func (s *MockOracleKeeper) GetPrice(ctx context.Context, denom, quote string) *math.LegacyDec {
	price := s.prices[denom]
	return &price
}

func (s *MockOracleKeeper) SetPrice(denom string, price math.LegacyDec) {
	s.prices[denom] = price
}

func (s *MockOracleKeeper) AddNewSymbolToBandOracleRequest(ctx context.Context, symbol string, oracleScriptId int64) error {
	return nil
}
