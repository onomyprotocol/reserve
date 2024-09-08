package types

func ConvertAddStableCoinToStablecoin(s AddStableCoinProposal) Stablecoin {
	return Stablecoin{
		Denom:      s.Denom,
		LimitTotal: s.LimitTotal,
		Price:      s.Price,
		FeeIn:      s.FeeIn,
		FeeOut:     s.FeeOut,
	}
}

func ConvertUpdateStableCoinToStablecoin(s UpdatesStableCoinProposal) Stablecoin {
	return Stablecoin{
		Denom:      s.Denom,
		LimitTotal: s.UpdatesLimitTotal,
		Price:      s.Price,
		FeeIn:      s.FeeIn,
		FeeOut:     s.FeeOut,
	}
}
