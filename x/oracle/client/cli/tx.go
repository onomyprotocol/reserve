package cli

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	errors "github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/onomyprotocol/reserve/x/oracle/types"
)

var DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())

const listSeparator = ","

// GetTxCmd returns the transaction commands for this module.
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		NewRequestBandRatesTxCmd(),
	)

	return cmd
}

// NewRequestBandRatesTxCmd implements the request command handler.
func NewRequestBandRatesTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-band-rates [request-id]",
		Short: "Make a new data request via an existing oracle script",
		Args:  cobra.ExactArgs(1),
		Long: strings.TrimSpace(
			fmt.Sprintf(`Make a new request via an existing oracle script with the configuration flags.
Example:
$ reserved tx oracle request-band-rates 2 --from mykey
`),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			requestID, err := strconv.Atoi(args[0])
			if err != nil {
				return errors.New("requestID should be a positive number")
			} else if requestID <= 0 {
				return errors.New("requestID should be a positive number")
			}

			msg := types.NewMsgRequestBandIBCRates(
				clientCtx.GetFromAddress(),
				uint64(requestID),
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
