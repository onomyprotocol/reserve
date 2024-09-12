package keeper_test

import (
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/psm/types"
)

func (s *KeeperTestSuite) TestStoreStablecoin() {
	s.SetupTest()

	s1 := types.Stablecoin{
		Denom:      usdt,
		LimitTotal: limitUSDT,
		Price:      math.LegacyMustNewDecFromStr("1"),
		FeeIn:      math.LegacyMustNewDecFromStr("0.001"),
		FeeOut:     math.LegacyMustNewDecFromStr("0.001"),
	}
	s2 := types.Stablecoin{
		Denom:      usdc,
		LimitTotal: limitUSDC,
		Price:      math.LegacyMustNewDecFromStr("1"),
		FeeIn:      math.LegacyMustNewDecFromStr("0.001"),
		FeeOut:     math.LegacyMustNewDecFromStr("0.001"),
	}

	s.k.SetStablecoin(s.Ctx, s1)
	s.k.SetStablecoin(s.Ctx, s2)

	stablecoin1, found := s.k.GetStablecoin(s.Ctx, usdt)
	s.Require().True(found)
	s.Require().Equal(stablecoin1.Denom, usdt)
	s.Require().Equal(stablecoin1.LimitTotal, limitUSDT)

	stablecoin2, found := s.k.GetStablecoin(s.Ctx, usdc)
	s.Require().True(found)
	s.Require().Equal(stablecoin2.Denom, usdc)
	s.Require().Equal(stablecoin2.LimitTotal, limitUSDC)

	count := 0
	s.k.IterateStablecoin(s.Ctx, func(red types.Stablecoin) (stop bool) {
		count += 1
		return false
	})
	s.Require().Equal(count, 2)
}

func (s *KeeperTestSuite) TestStoreLockcoin() {
	s.SetupTest()

	coinLock1 := sdk.NewCoin(usdt, math.NewInt(1000))
	coinLock2 := sdk.NewCoin(usdc, math.NewInt(1000))

	l1 := types.LockCoin{
		Address: s.TestAccs[0].String(),
		Coin:    &coinLock1,
		Time:    time.Now().Unix(),
	}
	l2 := types.LockCoin{
		Address: s.TestAccs[1].String(),
		Coin:    &coinLock2,
		Time:    time.Now().Unix(),
	}

	s.k.SetLockCoin(s.Ctx, l1)
	s.k.SetLockCoin(s.Ctx, l2)

	lockCoin1, found := s.k.GetLockCoin(s.Ctx, s.TestAccs[0].String())
	s.Require().True(found)
	s.Require().Equal(coinLock1, *lockCoin1.Coin)

	lockCoin2, found := s.k.GetLockCoin(s.Ctx, s.TestAccs[1].String())
	s.Require().True(found)
	s.Require().Equal(coinLock2, *lockCoin2.Coin)

	count := 0
	s.k.IterateLockCoin(s.Ctx, func(red types.LockCoin) (stop bool) {
		count += 1
		return false
	})
	s.Require().Equal(count, 2)

	l3 := types.LockCoin{
		Address: s.TestAccs[1].String(),
		Coin:    &coinLock1,
		Time:    time.Now().Unix(),
	}
	s.k.SetLockCoin(s.Ctx, l3)

	totalLock := s.k.TotalStablecoinLock(s.Ctx, usdt)
	s.Require().Equal(l1.Coin.Add(*l3.Coin).Amount.String(), totalLock.String())
}
