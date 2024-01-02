package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
)

// Proposal flags
const (
	FlagTitle        = "title"
	FlagDescription  = "description"
	FlagProposalType = "type"
	FlagDeposit      = "deposit"
	flagVoter        = "voter"
	flagDepositor    = "depositor"
	flagStatus       = "status"
	FlagProposal     = "proposal"
)

type proposal struct {
	Title       string
	Description string
	Type        string
	Deposit     string
}

// ProposalFlags defines the core required fields of a proposal. It is used to
// verify that these values are not provided in conjunction with a JSON proposal
// file.
var ProposalFlags = []string{
	FlagTitle,
	FlagDescription,
	FlagProposalType,
	FlagDeposit,
}

// NewTxCmd returns the transaction commands for this module
// governance ModuleClient is slightly different from other ModuleClients in that
// it contains a slice of "proposal" child commands. These commands are respective
// to proposal type handlers that are implemented in other modules but are mounted
// under the governance CLI (eg. parameter change proposals).
func NewTxCmd(propCmds []*cobra.Command) *cobra.Command {
	govTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Governance transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmdSubmitProp := NewCmdSubmitProposal()
	for _, propCmd := range propCmds {
		flags.AddTxFlagsToCmd(propCmd)
		cmdSubmitProp.AddCommand(propCmd)
	}

	govTxCmd.AddCommand(
		cmdSubmitProp,
	)

	return govTxCmd
}

// NewCmdSubmitProposal implements submitting a proposal transaction command.
func NewCmdSubmitProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-proposal",
		Short: "Submit a proposal along with an initial deposit",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit a proposal along with an initial deposit.
Proposal title, description, type and deposit can be given directly or through a proposal JSON file.

Example:
$ %s tx gov submit-proposal --proposal="path/to/proposal.json" --from mykey

Where proposal.json contains:

{
  "title": "Test Proposal",
  "description": "My awesome proposal",
  "type": "Text",
  "deposit": "10test"
}

Which is equivalent to:

$ %s tx gov submit-proposal --title="Test Proposal" --description="My awesome proposal" --type="Text" --deposit="10test" --from mykey
`,
				version.AppName, version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			proposal, err := parseSubmitProposalFlags(cmd.Flags())
			if err != nil {
				return fmt.Errorf("failed to parse proposal: %w", err)
			}

			amount, err := sdk.ParseCoinsNormalized(proposal.Deposit)
			if err != nil {
				return err
			}

			content := types.ContentFromProposalType(proposal.Title, proposal.Description, proposal.Type)

			msg, err := types.NewMsgSubmitProposal(content, amount, clientCtx.GetFromAddress())
			if err != nil {
				return fmt.Errorf("invalid message: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagTitle, "", "The proposal title")
	cmd.Flags().String(FlagDescription, "", "The proposal description")
	cmd.Flags().String(FlagProposalType, "", "The proposal Type")
	cmd.Flags().String(FlagDeposit, "", "The proposal deposit")
	cmd.Flags().String(FlagProposal, "", "Proposal file path (if this path is given, other proposal flags are ignored)")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
