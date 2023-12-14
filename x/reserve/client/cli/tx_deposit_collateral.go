package cli

import (
	"strconv"

	"reserve/x/reserve/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdDepositCollateral() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit-collateral [uid] [amount]",
		Short: "Broadcast message deposit-collateral",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argUid, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argAmount := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDepositCollateral(
				clientCtx.GetFromAddress().String(),
				argUid,
				argAmount,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
