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

	cmd.AddCommand(NewSwapToNomCmd())

	return cmd
}

func NewSwapToNomCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap [offer_stable_coin] [expected_denom]",
		Args:  cobra.ExactArgs(2),
		Short: "stable swap  ",
		Long: `swap between stable coins.

			Example:
			$ onomyd tx psm swap 100000000000000000000000nomUSD ibc/xxxxx --from validator1 --keyring-backend test --home ~/.reserved/validator1 --chain-id testing-1 -y --fees 20stake

	`,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			offerCoin, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			addr := clientCtx.GetFromAddress()
			msg := types.NewMsgStableSwap(addr.String(), offerCoin, args[1])

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
