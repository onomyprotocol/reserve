package cli

import (
	// "context"

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
			res, err = queryClient.BandPriceStates(cmd.Context(), req)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}