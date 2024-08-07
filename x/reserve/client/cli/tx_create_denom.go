package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	"reserve/x/reserve/types"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// CmdCreateDenomProposal implements the command to submit a create-denom proposal.
func CmdCreateDenomProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-denom rate metadata-path",
		Args:  cobra.ExactArgs(6),
		Short: "Submit a create denom proposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit a create denom proposal.
Example:
$ %s tx gov submit-proposal create-denom peg-pair debt-interest-rate bond-interest-rate denom-metadata bond-metadata --title="Test Proposal" --description="My awesome proposal" --deposit="10000000000000000000aonex"`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			pegPair := args[0]

			debtInterestRate, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			if debtInterestRate > 100000 {
				return types.ErrInterestGtLimit
			}

			bondInterestRate, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			if bondInterestRate > 100000 {
				return types.ErrInterestGtLimit
			}

			path := args[3]

			metadataFile, err := os.Open(path)
			if err != nil {
				return err
			}

			byteMetadata, err := io.ReadAll(metadataFile)
			if err != nil {
				return err
			}

			var denomMetadata banktypes.Metadata

			err = json.Unmarshal(byteMetadata, &denomMetadata)
			if err != nil {
				return err
			}

			path = args[4]

			metadataFile, err = os.Open(path)
			if err != nil {
				return err
			}

			byteMetadata, err = io.ReadAll(metadataFile)
			if err != nil {
				return err
			}

			var bondMetadata banktypes.Metadata

			err = json.Unmarshal(byteMetadata, &bondMetadata)
			if err != nil {
				return err
			}

			proposalFlags, err := parseProposalFlags(cmd.Flags())
			if err != nil {
				return err
			}

			deposit, err := sdk.ParseCoinsNormalized(proposalFlags.Deposit)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			content := types.NewCreateDenomProposal(
				from,
				proposalFlags.Title,
				proposalFlags.Description,
				denomMetadata,
				bondMetadata,
				pegPair,
				debtInterestRate,
				bondInterestRate,
			)

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	addProposalFlags(cmd)

	return cmd
}
