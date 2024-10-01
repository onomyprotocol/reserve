package types

import (
	"cosmossdk.io/math"
)

func GetMsgStablecoin(msg getStablecoinFromMsg) Stablecoin {
	return Stablecoin{
		Denom:      msg.GetDenom(),
		LimitTotal: msg.GetLimitTotal(),
		Price:      msg.GetPrice(),
		FeeIn:      msg.GetFeeIn(),
		FeeOut:     msg.GetFeeOut(),
	}
}

type getStablecoinFromMsg interface {
	GetDenom() string
	GetLimitTotal() math.Int
	GetPrice() math.LegacyDec
	GetFeeIn() math.LegacyDec
	GetFeeOut() math.LegacyDec
}
