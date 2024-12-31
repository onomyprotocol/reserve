package cli

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/onomyprotocol/reserve/x/auction/types"
)

// GetTxCmd returns the transaction commands for this module.
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(NewBidCmd())
	cmd.AddCommand(NewCancelBidCmd())

	return cmd
}

func NewBidCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bid [auction-id] [amount] [recive-price]",
		Args:  cobra.ExactArgs(3),
		Short: "create vaults ",
		Long: `create vaults.

			Example:
			$ onomyd tx bid 0 1000nomUSD 0.8 --from mykey
	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			addr := clientCtx.GetFromAddress()

			auctionID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}
			msg := types.NewMsgBid(addr.String(), auctionID, amount, args[2])

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCancelBidCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel-bid [bid-id] [auction_id]",
		Args:  cobra.ExactArgs(2),
		Short: "create vaults ",
		Long: `create vaults.

			Example:
			$ onomyd tx cancel-bid 1 0 --from mykey
	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			addr := clientCtx.GetFromAddress()

			auctionID, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			bidID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgCancelBid(addr.String(), bidID, auctionID)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
