package cli

import (
	"context"

	"github.com/cosmos/gogoproto/proto"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

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
		CmdGetPrice(),
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
			queryClient := types.NewQueryClient(clientCtx)

			var res proto.Message
			req := &types.QueryBandPriceStatesRequest{}
			res, err = queryClient.BandPriceStates(context.Background(), req)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdGetPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-price [denom]",
		Short: "shows info price denom",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			msg := types.NewGetPrice(args[0])
			res, err := queryClient.QueryPrice(context.Background(), &msg)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetAllPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-all-price",
		Short: "shows info all price denom",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			msg := types.NewAllGetPrice()
			res, err := queryClient.QueryAllPrice(context.Background(), &msg)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
