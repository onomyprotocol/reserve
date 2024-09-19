package cli

import (
	"fmt"

	"cosmossdk.io/math"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/spf13/cobra"

	"github.com/onomyprotocol/reserve/x/vaults/types"
)

func NewCmdSubmitActiveCollateralProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "active-collateral [title] [description] [denom] [min-collateral-ratio] [liquidation-ratio] [max-debt] [deposit]",
		Args:  cobra.ExactArgs(7),
		Short: "Active collateral proposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			minCollateralRatio, err := math.LegacyNewDecFromStr(args[3])
			if err != nil {
				return err
			}

			liquidationRatio, err := math.LegacyNewDecFromStr(args[4])
			if err != nil {
				return err
			}
			maxDebt, ok := math.NewIntFromString(args[5])
			if !ok {
				return fmt.Errorf("value %s cannot constructs Int from string", args[5])
			}

			content := types.NewActiveCollateralProposal(
				args[0], args[1], args[2], minCollateralRatio, liquidationRatio, maxDebt,
			)

			deposit, err := sdk.ParseCoinsNormalized(args[6])
			if err != nil {
				return err
			}

			msg, err := govtypes.NewMsgSubmitProposal(&content, deposit, clientCtx.GetFromAddress())
			if err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	return cmd
}
