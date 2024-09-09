package cli

import (
	"fmt"

	"cosmossdk.io/math"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
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
	cmd.AddCommand(NewSwapToISTCmd())

	return cmd
}

func NewCmdSubmitAddStableCoinProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-stable-coin [title] [description] [denom] [limit-total] [price] [fee_in] [fee_out] [proposer] [deposit]",
		Args:  cobra.ExactArgs(9),
		Short: "Submit an add stable coin proposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			limitTotal, ok := math.NewIntFromString(args[3])
			if !ok {
				return fmt.Errorf("value %s cannot constructs Int from string", args[3])
			}

			price, err := math.LegacyNewDecFromStr(args[4])
			if err != nil {
				return err
			}
			feeIn, err := math.LegacyNewDecFromStr(args[5])
			if err != nil {
				return err
			}
			feeOut, err := math.LegacyNewDecFromStr(args[6])
			if err != nil {
				return err
			}
			from := sdk.MustAccAddressFromBech32(args[7])

			content := types.NewAddStableCoinProposal(
				args[0], args[1], args[2], limitTotal, price, feeIn, feeOut,
			)

			deposit, err := sdk.ParseCoinsNormalized(args[8])
			if err != nil {
				return err
			}

			msg, err := govtypes.NewMsgSubmitProposal(&content, deposit, from)
			if err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	return cmd
}

func NewCmdUpdatesStableCoinProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-limit-total-stable-coin [title] [description] [denom] [limit-total-update] [price] [fee_in] [fee_out] [deposit]",
		Args:  cobra.ExactArgs(8),
		Short: "Submit update limit total stable coin proposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			limitTotalUpdate, ok := math.NewIntFromString(args[3])
			if !ok {
				return fmt.Errorf("value %s cannot constructs Int from string", args[3])
			}
			price, err := math.LegacyNewDecFromStr(args[4])
			feeIn, err := math.LegacyNewDecFromStr(args[5])
			feeOut, err := math.LegacyNewDecFromStr(args[6])
			from := clientCtx.GetFromAddress()
			content := types.NewUpdatesStableCoinProposal(
				args[0], args[1], args[2], limitTotalUpdate, price, feeIn, feeOut,
			)

			deposit, err := sdk.ParseCoinsNormalized(args[7])
			if err != nil {
				return err
			}

			msg, err := govtypes.NewMsgSubmitProposal(&content, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	return cmd
}

func NewSwapToISTCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap-to-ist [stablecoin]",
		Args:  cobra.ExactArgs(1),
		Short: "swap stablecoin to $ist ",
		Long: `swap stablecoin to $ist.

			Example:
			$ onomyd tx psm swap-to-ist 1000usdt --from mykey
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
			msg := types.NewMsgSwapToIST(addr.String(), &stablecoin)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewSwapToStablecoinCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap-to-stablecoin [stable-coin-type] [amount-ist]",
		Args:  cobra.ExactArgs(2),
		Short: "swap $ist to stablecoin ",
		Long: `swap $ist to stablecoin.

			Example:
			$ onomyd tx psm swap-to-stablecoin usdt 10000ist --from mykey
	`,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			ISTcoin, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			addr := clientCtx.GetFromAddress()
			msg := types.NewMsgSwapToStablecoin(addr.String(), args[0], ISTcoin.Amount)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
