package cli

import (
	"fmt"
	"strconv"

	errors "github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govcli "github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/onomyprotocol/reserve/x/oracle/types"
)

const (
	flagSymbols                  = "symbols"
	flagRequestedValidatorCount  = "requested-validator-count"
	flagSufficientValidatorCount = "sufficient-validator-count"
	flagMinSourceCount           = "min-source-count"
	flagPrepareGas               = "prepare-gas"
	flagExecuteGas               = "execute-gas"
	flagFeeLimit                 = "fee-limit"
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

	cmd.AddCommand(
		NewRequestBandRatesTxCmd(),
		NewAuthorizeBandOracleRequestProposalTxCmd(),
		NewUpdateBandOracleRequestProposalTxCmd(),
		NewDeleteBandOracleRequestProposalTxCmd(),
	)

	return cmd
}

// NewRequestBandRatesTxCmd implements the request command handler.
func NewRequestBandRatesTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-band-rates [request-id]",
		Short: "Make a new data request via an existing oracle script",
		Args:  cobra.ExactArgs(1),
		Long: `Make a new request via an existing oracle script with the configuration flags.
		Example:
		$ %s tx oracle request-band-rates 2 --from mykey`,
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

			msg := types.NewMsgRequestBandRates(
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

func NewAuthorizeBandOracleRequestProposalTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "authorize-band-oracle-request-proposal [flags]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a proposal to authorize a Band Oracle IBC Request.",
		Long: `Submit a proposal to authorize a Band Oracle IBC Request.
			Example:
			$ %s tx oracle authorize-band-oracle-request-proposal 23 --symbols "BTC,ETH,USDT,USDC" --requested-validator-count 4 --sufficient-validator-count 3 --min-source-count 3 --prepare-gas 20000 --fee-limit "1000uband" --execute-gas 400000 --from mykey
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			content, err := authorizeBandOracleRequestProposalArgsToContent(cmd, args)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			depositStr, err := cmd.Flags().GetString(govcli.FlagDeposit)
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().StringSlice(flagSymbols, []string{}, "Symbols used in calling the oracle script")
	cmd.Flags().Uint64(flagPrepareGas, 50000, "Prepare gas used in fee counting for prepare request")
	cmd.Flags().Uint64(flagExecuteGas, 300000, "Execute gas used in fee counting for execute request")
	cmd.Flags().String(flagFeeLimit, "", "the maximum tokens that will be paid to all data source providers")
	cmd.Flags().Uint64(flagRequestedValidatorCount, 4, "Requested Validator Count")
	cmd.Flags().Uint64(flagSufficientValidatorCount, 10, "Sufficient Validator Count")
	cmd.Flags().Uint64(flagMinSourceCount, 3, "Min Source Count")
	cmd.Flags().String(govcli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(govcli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(govcli.FlagDeposit, "", "deposit of proposal")

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewUpdateBandOracleRequestProposalTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-band-oracle-request-proposal 1 37 [flags]",
		Args:  cobra.ExactArgs(2),
		Short: "Submit a proposal to update a Band Oracle IBC Request.",
		Long: `Submit a proposal to update a Band Oracle IBC Request.
			Example:
			$ %s tx oracle update-band-oracle-request-proposal 1 37 --port-id "oracle" --ibc-version "bandchain-1" --symbols "BTC,ETH,USDT,USDC" --requested-validator-count 4 --sufficient-validator-count 3 --min-source-count 3 --expiration 20 --prepare-gas 50 --execute-gas 5000 --from mykey
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			content, err := updateBandOracleRequestProposalArgsToContent(cmd, args)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			depositStr, err := cmd.Flags().GetString(govcli.FlagDeposit)
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().StringSlice(flagSymbols, []string{}, "Symbols used in calling the oracle script")
	cmd.Flags().Uint64(flagPrepareGas, 0, "Prepare gas used in fee counting for prepare request")
	cmd.Flags().Uint64(flagExecuteGas, 0, "Execute gas used in fee counting for execute request")
	cmd.Flags().String(flagFeeLimit, "", "the maximum tokens that will be paid to all data source providers")
	cmd.Flags().Uint64(flagRequestedValidatorCount, 0, "Requested Validator Count")
	cmd.Flags().Uint64(flagSufficientValidatorCount, 0, "Sufficient Validator Count")
	cmd.Flags().Uint64(flagMinSourceCount, 3, "Min Source Count")
	cmd.Flags().String(govcli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(govcli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(govcli.FlagDeposit, "", "deposit of proposal")

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewDeleteBandOracleRequestProposalTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-band-oracle-request-proposal 1 [flags]",
		Args:  cobra.MinimumNArgs(1),
		Short: "Submit a proposal to Delete a Band Oracle IBC Request.",
		Long: `Submit a proposal to Delete a Band Oracle IBC Request.
			Example:
			$ %s tx oracle delete-band-oracle-request-proposal 1 --from mykey
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			content, err := deleteBandOracleRequestProposalArgsToContent(cmd, args)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			depositStr, err := cmd.Flags().GetString(govcli.FlagDeposit)
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(govcli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(govcli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(govcli.FlagDeposit, "", "deposit of proposal")

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func authorizeBandOracleRequestProposalArgsToContent(
	cmd *cobra.Command,
	args []string,
) (govtypes.Content, error) {
	title, err := cmd.Flags().GetString(govcli.FlagTitle)
	if err != nil {
		return nil, err
	}

	description, err := cmd.Flags().GetString(govcli.FlagDescription)
	if err != nil {
		return nil, err
	}

	int64OracleScriptID, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return nil, err
	}

	askCount, err := cmd.Flags().GetUint64(flagRequestedValidatorCount)
	if err != nil {
		return nil, err
	}

	minCount, err := cmd.Flags().GetUint64(flagSufficientValidatorCount)
	if err != nil {
		return nil, err
	}

	minSourceCount, err := cmd.Flags().GetUint64(flagMinSourceCount)
	if err != nil {
		return nil, err
	}

	symbols, err := cmd.Flags().GetStringSlice(flagSymbols)
	if err != nil {
		return nil, err
	}

	prepareGas, err := cmd.Flags().GetUint64(flagPrepareGas)
	if err != nil {
		return nil, err
	}

	executeGas, err := cmd.Flags().GetUint64(flagExecuteGas)
	if err != nil {
		return nil, err
	}

	coinStr, err := cmd.Flags().GetString(flagFeeLimit)
	if err != nil {
		return nil, err
	}

	feeLimit, err := sdk.ParseCoinsNormalized(coinStr)
	if err != nil {
		return nil, err
	}

	content := &types.AuthorizeBandOracleRequestProposal{
		Title:       title,
		Description: description,
		Request: types.BandOracleRequest{
			OracleScriptId: int64OracleScriptID,
			Symbols:        symbols,
			AskCount:       askCount,
			MinCount:       minCount,
			FeeLimit:       feeLimit,
			PrepareGas:     prepareGas,
			ExecuteGas:     executeGas,
			MinSourceCount: minSourceCount,
		},
	}
	if err := content.ValidateBasic(); err != nil {
		return nil, err
	}
	return content, nil
}

func updateBandOracleRequestProposalArgsToContent(
	cmd *cobra.Command,
	args []string,
) (govtypes.Content, error) {
	title, err := cmd.Flags().GetString(govcli.FlagTitle)
	if err != nil {
		return nil, err
	}

	description, err := cmd.Flags().GetString(govcli.FlagDescription)
	if err != nil {
		return nil, err
	}

	requestID, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return nil, err
	}

	int64OracleScriptID, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return nil, err
	}

	askCount, err := cmd.Flags().GetUint64(flagRequestedValidatorCount)
	if err != nil {
		return nil, err
	}

	minCount, err := cmd.Flags().GetUint64(flagSufficientValidatorCount)
	if err != nil {
		return nil, err
	}
	minSourceCount, err := cmd.Flags().GetUint64(flagMinSourceCount)
	if err != nil {
		return nil, err
	}

	symbols, err := cmd.Flags().GetStringSlice(flagSymbols)
	if err != nil {
		return nil, err
	}

	prepareGas, err := cmd.Flags().GetUint64(flagPrepareGas)
	if err != nil {
		return nil, err
	}

	executeGas, err := cmd.Flags().GetUint64(flagExecuteGas)
	if err != nil {
		return nil, err
	}

	coinStr, err := cmd.Flags().GetString(flagFeeLimit)
	if err != nil {
		return nil, err
	}

	feeLimit, err := sdk.ParseCoinsNormalized(coinStr)
	if err != nil {
		return nil, err
	}

	content := &types.UpdateBandOracleRequestProposal{
		Title:       title,
		Description: description,
		UpdateOracleRequest: &types.BandOracleRequest{
			RequestId:      uint64(requestID),
			OracleScriptId: int64OracleScriptID,
			Symbols:        symbols,
			AskCount:       askCount,
			MinCount:       minCount,
			FeeLimit:       feeLimit,
			PrepareGas:     prepareGas,
			ExecuteGas:     executeGas,
			MinSourceCount: minSourceCount,
		},
	}
	if err := content.ValidateBasic(); err != nil {
		return nil, err
	}

	return content, nil
}

func deleteBandOracleRequestProposalArgsToContent(
	cmd *cobra.Command,
	args []string,
) (govtypes.Content, error) {
	title, err := cmd.Flags().GetString(govcli.FlagTitle)
	if err != nil {
		return nil, err
	}

	description, err := cmd.Flags().GetString(govcli.FlagDescription)
	if err != nil {
		return nil, err
	}

	requestIDs := make([]uint64, 0, len(args))
	for _, arg := range args {
		id, err := strconv.ParseInt(arg, 10, 64)
		if err != nil {
			return nil, err
		}

		requestIDs = append(requestIDs, uint64(id))
	}

	content := &types.DeleteBandOracleRequestProposal{
		Title:            title,
		Description:      description,
		DeleteRequestIds: requestIDs,
	}

	if err := content.ValidateBasic(); err != nil {
		return nil, err
	}

	return content, nil
}
