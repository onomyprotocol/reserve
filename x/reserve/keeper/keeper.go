package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/onomyprotocol/reserve/x/reserve/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace

		bankKeeper    types.BankKeeper
		accountKeeper types.AccountKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,

	bankKeeper types.BankKeeper,
	accountKeeper types.AccountKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) validateSenderBalance(ctx sdk.Context, senderAddress sdk.AccAddress, coins sdk.Coins) error {
	for _, coin := range coins {
		balance := k.bankKeeper.GetBalance(ctx, senderAddress, coin.Denom)
		if balance.IsLT(coin) {
			return sdkerrors.Wrapf(
				types.ErrInsufficientBalance, "%s is smaller than %s", balance, coin)
		}
	}

	return nil
}

func addUid(s []uint64, r uint64) ([]uint64, bool) {
	for _, v := range s {
		if v == r {
			return s, false
		}
	}

	return append(s, r), true
}

func removeUid(s []uint64, r uint64) ([]uint64, bool) {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...), true
		}
	}
	return s, false
}

func addPair(s []string, r string) ([]string, bool) {
	for _, v := range s {
		if v == r {
			return s, false
		}
	}

	return append(s, r), true
}

func removePair(s []string, r string) ([]string, bool) {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...), true
		}
	}
	return s, false
}
