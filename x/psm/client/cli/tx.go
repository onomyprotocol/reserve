package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/onomyprotocol/reserve/x/psm/types"
)

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2, // nolint:gomnd
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(NewSwapToStablecoinCmd())
	cmd.AddCommand(NewSwapTonomUSDCmd())

	return cmd
}

func NewSwapTonomUSDCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap-to-nomUSD [stablecoin]",
		Args:  cobra.ExactArgs(1),
		Short: "swap stablecoin to $nomUSD ",
		Long: `swap stablecoin to $nomUSD.

			Example:
			$ onomyd tx psm swap-to-nomUSD 1000usdt --from mykey
	`,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			stablecoin, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			addr := clientCtx.GetFromAddress()
			msg := types.NewMsgSwapToNom(addr.String(), stablecoin)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewSwapToStablecoinCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap-to-stablecoin [stable-coin-type] [amount-nomUSD]",
		Args:  cobra.ExactArgs(2),
		Short: "swap $nomUSD to stablecoin ",
		Long: `swap $nomUSD to stablecoin.

			Example:
			$ onomyd tx psm swap-to-stablecoin usdt 10000nomUSD --from mykey
	`,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			nomUSDcoin, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			addr := clientCtx.GetFromAddress()
			msg := types.NewMsgSwapToStablecoin(addr.String(), args[0], nomUSDcoin)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
