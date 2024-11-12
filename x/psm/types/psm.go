package types

import (
	"cosmossdk.io/math"
)

func GetMsgStablecoin(msg getStablecoinFromMsg) StablecoinInfo {
	return StablecoinInfo{
		Denom:               msg.GetDenom(),
		LimitTotal:          msg.GetLimitTotal(),
		FeeIn:               msg.GetFeeIn(),
		FeeOut:              msg.GetFeeOut(),
		TotalStablecoinLock: math.ZeroInt(),
		FeeMaxStablecoin:    msg.GetFeeIn().Add(msg.GetFeeOut()),
		Symbol:              msg.GetSymbol(),
	}
}

type getStablecoinFromMsg interface {
	GetDenom() string
	GetLimitTotal() math.Int
	GetFeeIn() math.LegacyDec
	GetFeeOut() math.LegacyDec
	GetSymbol() string
}
