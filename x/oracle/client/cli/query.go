package cli

import (
	"fmt"
	"context"
	"strconv"

	"github.com/cosmos/gogoproto/proto"
	errors "github.com/pkg/errors"

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
		GetPrice(),
		GetBandParams(),
		GetBandOracleRequestParams(),
		GetBandOracleRequest(),
	)
	return cmd
}

// GetBandPriceStates queries the state for all band price states
func GetBandPriceStates() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "band-price-states",
		Short: "Get Band price states",
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
				BaseDenom:  baseDenom,
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

// GetBandParams queries the band parameters
func GetBandParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "band-params",
		Short: "Get band parameters",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			var res proto.Message
			req := &types.QueryBandParamsRequest{}
			res, err = queryClient.BandParams(context.Background(), req)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetBandOracleRequestParams queries the band oracle request parameters
func GetBandOracleRequestParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "band-oracle-request-params",
		Short: "Get band oracle request parameters",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			var res proto.Message
			req := &types.QueryBandOracleRequestParamsRequest{}
			res, err = queryClient.BandOracleRequestParams(context.Background(), req)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetBandOracleRequest queries the band oracle request parameters
func GetBandOracleRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "band-oracle-request",
		Short: "Get band oracle request",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			requestID, err := strconv.Atoi(args[0])
			if err != nil {
				return errors.New("requestID should be a positive number")
			} else if requestID <= 0 {
				return errors.New("requestID should be a positive number")
			}
			queryClient := types.NewQueryClient(clientCtx)

			var res proto.Message
			req := &types.QueryBandOracleRequestRequest{
				RequestId: fmt.Sprint(requestID),
			}
			res, err = queryClient.BandOracleRequest(context.Background(), req)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
