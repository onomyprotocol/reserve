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
	cmd.AddCommand(NewSwapToNomCmd())

	return cmd
}

func NewSwapToNomCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap-to-nom [stablecoin]",
		Args:  cobra.ExactArgs(1),
		Short: "swap stablecoin to $nomUSD ",
		Long: `swap stablecoin to $nomUSD.

			Example:
			$ onomyd tx psm swap-to-nom 100000000000000000000000usdt --from validator1 --keyring-backend test --home ~/.reserved/validator1 --chain-id testing-1 -y --fees 20stake

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
		Use:   "swap-to-stablecoin [stable-coin-type] [amount-nomX]",
		Args:  cobra.ExactArgs(2),
		Short: "swap $nomX to stablecoin.  ",
		Long: `swap $nomX to stablecoin.

			Example:
			$ onomyd tx psm swap-to-stablecoin usdc 10000nomUSD --from validator1 --keyring-backend test --home ~/.reserved/validator1 --chain-id testing-1 -y --fees 20stake
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
