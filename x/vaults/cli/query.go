package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/onomyprotocol/reserve/x/vaults/types"
)

// GetQueryCmd returns the cli query commands for this module.
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2, // nolint:gomnd
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryAllCollateral())
	cmd.AddCommand(CmdQueryAllVaults())
	cmd.AddCommand(CmdQueryVaultByID())
	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdQueryVaultFromAddress())
	cmd.AddCommand(CmdQueryCollateralsByDenom())
	cmd.AddCommand(CmdQueryTotalCollateralLockedByDenom())
	return cmd
}

func CmdQueryAllCollateral() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-collateral",
		Short: "show all collateral",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.QueryAllCollateral(context.Background(), &types.QueryAllCollateralRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdQueryCollateralsByDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "collateral-by-denom [denom]",
		Short: "show all collateral by denom",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.QueryCollateralsByDenom(context.Background(), &types.QueryCollateralsByDenomRequest{Denom: args[0]})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdQueryCollateralsByMintDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "collateral-by-mint-denom [mint-denom]",
		Short: "show all collateral by mint denom",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.QueryCollateralsByMintDenom(context.Background(), &types.QueryCollateralsByMintDenomRequest{MintDenom: args[0]})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
func CmdQueryCollateralsByDenomMintDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "collateral-by-denom [denom] [mint-denom]",
		Short: "show collateral by denom and mint denom",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.QueryCollateralsByDenomMintDenom(context.Background(), &types.QueryCollateralsByDenomMintDenomRequest{Denom: args[0], MintDenom: args[1]})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdQueryAllVaults() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-vaults",
		Short: "show all vaults",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.QueryAllVaults(context.Background(), &types.QueryAllVaultsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdQueryVaultByID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vault [vault-id]",
		Short: "show vault from id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			vaultID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.QueryVaultIdRequest{
				VaultId: vaultID,
			}

			res, err := queryClient.QueryVaultsByID(context.Background(), &msg)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "show params module vaults",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdQueryVaultFromAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vault-by-address [owner-address]",
		Short: "show vaults from owner address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			msg := types.QueryVaultByOwnerRequest{
				Address: args[0],
			}

			res, err := queryClient.QueryVaultByOwner(context.Background(), &msg)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdQueryTotalCollateralLockedByDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "total-collateral-locked-by-denom [denom]",
		Short: "show total collateral locked by denom",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			msg := types.QueryTotalCollateralLockedByDenomRequest{
				Denom: args[0],
			}

			res, err := queryClient.QueryTotalCollateralLockedByDenom(context.Background(), &msg)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
