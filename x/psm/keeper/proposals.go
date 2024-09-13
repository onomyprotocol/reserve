package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/onomyprotocol/reserve/x/psm/types"
)

func (k Keeper) AddStableCoinProposal(ctx sdk.Context, c *types.AddStableCoinProposal) error {
	if err := c.ValidateBasic(); err != nil {
		return err
	}

	_, found := k.GetStablecoin(ctx, c.Denom)
	if found {
		return fmt.Errorf("%s has existed", c.Denom)
	}

	err := k.SetStablecoin(ctx, types.ConvertAddStableCoinToStablecoin(*c))
	if err != nil {
		return err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventAddStablecoin,
			sdk.NewAttribute(types.AttributeStablecoinName, c.Denom),
		),
	)
	return nil
}

func (k Keeper) UpdatesStableCoinProposal(ctx sdk.Context, c *types.UpdatesStableCoinProposal) error {
	if err := c.ValidateBasic(); err != nil {
		return err
	}

	_, found := k.GetStablecoin(ctx, c.Denom)
	if !found {
		return fmt.Errorf("%s not existed", c.Denom)
	}

	err := k.SetStablecoin(ctx, types.ConvertUpdateStableCoinToStablecoin(*c))
	if err != nil {
		return err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventAddStablecoin,
			sdk.NewAttribute(types.AttributeStablecoinName, c.Denom),
		),
	)
	return nil
}
