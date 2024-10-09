package cli

import (
	"fmt"
	"context"

	"github.com/cosmos/gogoproto/proto"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/onomyprotocol/reserve/x/oracle/types"
)

// GetQueryCmd returns the parent command for all modules/oracle CLi query commands.
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the oracle module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetBandPriceStates(),
		GetPrice(),
	)
	return cmd
}

// GetBandPriceStates queries the state for all band price states
func GetBandPriceStates() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "band-price-states",
		Short: "Gets Band price states",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			ctx := cmd.Context()
			sdkContext := sdk.UnwrapSDKContext(ctx)
			queryClient := types.NewQueryClient(clientCtx)
			fmt.Println("check cmd context:", sdkContext)
			var res proto.Message
			req := &types.QueryBandPriceStatesRequest{}
			res, err = queryClient.BandPriceStates(ctx, req)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetPrice queries the price based on rate of base/quote
func GetPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "price [base] [quote]",
		Short: "Get price based on rate of base/quote",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			baseDenom := args[0]
			quoteDenom := args[1]

			queryClient := types.NewQueryClient(clientCtx)

			var res proto.Message
			req := &types.QueryPriceRequest{
				BaseDenom: baseDenom,
				QuoteDenom: quoteDenom,
			}
			res, err = queryClient.Price(context.Background(), req)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
