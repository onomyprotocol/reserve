package cli

// import (
// 	"fmt"

// 	"cosmossdk.io/math"

// 	"github.com/cosmos/cosmos-sdk/client"
// 	"github.com/cosmos/cosmos-sdk/client/tx"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
// 	"github.com/spf13/cobra"

// 	"github.com/onomyprotocol/reserve/x/psm/types"
// )

// func NewCmdSubmitAddStableCoinProposal() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "add-stable-coin [title] [description] [denom] [limit-total] [price] [fee_in] [fee_out] [deposit]",
// 		Args:  cobra.ExactArgs(8),
// 		Short: "Submit an add stable coin proposal",
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			clientCtx, err := client.GetClientTxContext(cmd)
// 			if err != nil {
// 				return err
// 			}

// 			limitTotal, ok := math.NewIntFromString(args[3])
// 			if !ok {
// 				return fmt.Errorf("value %s cannot constructs Int from string", args[3])
// 			}

// 			price, err := math.LegacyNewDecFromStr(args[4])
// 			if err != nil {
// 				return err
// 			}
// 			feeIn, err := math.LegacyNewDecFromStr(args[5])
// 			if err != nil {
// 				return err
// 			}
// 			feeOut, err := math.LegacyNewDecFromStr(args[6])
// 			if err != nil {
// 				return err
// 			}

// 			content := types.NewAddStableCoinProposal(
// 				args[0], args[1], args[2], limitTotal, price, feeIn, feeOut,
// 			)

// 			deposit, err := sdk.ParseCoinsNormalized(args[7])
// 			if err != nil {
// 				return err
// 			}

// 			msg, err := govtypes.NewMsgSubmitProposal(&content, deposit, clientCtx.GetFromAddress())
// 			if err != nil {
// 				return err
// 			}
// 			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
// 		},
// 	}
// 	return cmd
// }

// func NewCmdUpdatesStableCoinProposal() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "update-limit-total-stable-coin [title] [description] [denom] [limit-total-update] [price] [fee_in] [fee_out] [deposit]",
// 		Args:  cobra.ExactArgs(8),
// 		Short: "Submit update limit total stable coin proposal",
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			clientCtx, err := client.GetClientTxContext(cmd)
// 			if err != nil {
// 				return err
// 			}

// 			limitTotalUpdate, ok := math.NewIntFromString(args[3])
// 			if !ok {
// 				return fmt.Errorf("value %s cannot constructs Int from string", args[3])
// 			}
// 			price, err := math.LegacyNewDecFromStr(args[4])
// 			if err != nil {
// 				return err
// 			}
// 			feeIn, err := math.LegacyNewDecFromStr(args[5])
// 			if err != nil {
// 				return err
// 			}
// 			feeOut, err := math.LegacyNewDecFromStr(args[6])
// 			if err != nil {
// 				return err
// 			}

// 			content := types.NewUpdatesStableCoinProposal(
// 				args[0], args[1], args[2], limitTotalUpdate, price, feeIn, feeOut,
// 			)

// 			deposit, err := sdk.ParseCoinsNormalized(args[7])
// 			if err != nil {
// 				return err
// 			}

// 			msg, err := govtypes.NewMsgSubmitProposal(&content, deposit, clientCtx.GetFromAddress())
// 			if err != nil {
// 				return err
// 			}

// 			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
// 		},
// 	}

// 	return cmd
// }
