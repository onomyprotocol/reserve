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

// CmdRegisterCollateralProposal implements the command to submit a register-collateral proposal.
func CmdRegisterCollateralProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-collateral metadata-path",
		Args:  cobra.ExactArgs(4),
		Short: "Submit a register collateral proposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit a register collateral proposal.
Example:
$ %s tx gov submit-proposal register-collateral metadata-path minimum-collateral-deposit minting-ratio liquidation-ratio --title="Test Proposal" --description="My awesome proposal" --deposit="10000000000000000000aonex"`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			path := args[0]

			metadataFile, err := os.Open(path)
			if err != nil {
				return err
			}

			byteMetadata, err := io.ReadAll(metadataFile)
			if err != nil {
				return err
			}

			var metadata banktypes.Metadata

			err = json.Unmarshal(byteMetadata, &metadata)
			if err != nil {
				return err
			}

			minCollateralDeposit, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			// Need to make sure these work
			// 100 would be 100%
			// 125 would be 125%
			// 999 would be 999%
			// 99999 would 99999%
			// minting ratio numerator / 100 = minting ratio
			LendingRatio, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			if LendingRatio >= 100000 {
				return fmt.Errorf("minting ratio numerator must be less than 100000: %d", LendingRatio)
			}

			liquidationRatio, err := strconv.ParseUint(args[3], 10, 64)
			if err != nil {
				return err
			}

			if liquidationRatio >= 100000 {
				return fmt.Errorf("liquidation ratio numerator must be less than 100000: %d", liquidationRatio)
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
			content := types.NewRegisterCollateralProposal(
				from,
				proposalFlags.Title,
				proposalFlags.Description,
				metadata,
				minCollateralDeposit,
				LendingRatio,
				liquidationRatio,
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
