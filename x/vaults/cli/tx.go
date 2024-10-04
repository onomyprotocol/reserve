package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/onomyprotocol/reserve/x/vaults/types"
)

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2, // nolint:gomnd
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(NewCreateVaultCmd())
	cmd.AddCommand(NewDepositCmd())
	cmd.AddCommand(NewWithdrawCmd())
	cmd.AddCommand(NewCloseCmd())
	cmd.AddCommand(NewRepayCmd())
	cmd.AddCommand(NewMintCmd())

	return cmd
}

func NewCreateVaultCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-vault [collateral] [minted]",
		Args:  cobra.ExactArgs(2),
		Short: "create vaults ",
		Long: `create vaults.

			Example:
			$ onomyd tx vautls create-vault 1000atom 8000nomUSD --from mykey
	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			addr := clientCtx.GetFromAddress()

			collateral, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}
			minted, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}
			msg := types.NewMsgCreateVault(addr.String(), collateral, minted)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewDepositCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit [vault-id] [amount]",
		Args:  cobra.ExactArgs(2),
		Short: "deposit ",
		Long: `deposit.

			Example:
			$ onomyd tx vaults 1 1000atom --from mykey
	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			addr := clientCtx.GetFromAddress()

			vaultID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgDeposit(vaultID, addr.String(), amount)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewWithdrawCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw [vault-id] [amount]",
		Args:  cobra.ExactArgs(2),
		Short: "withdraw ",
		Long: `withdraw.

			Example:
			$ onomyd tx withdraw 1 1000atom --from mykey
	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			addr := clientCtx.GetFromAddress()

			vaultID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgWithdraw(vaultID, addr.String(), amount)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewMintCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [vault-id] [amount]",
		Args:  cobra.ExactArgs(2),
		Short: "mint ",
		Long: `mint.

			Example:
			$ onomyd tx mint 1 1000atom --from mykey
	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			addr := clientCtx.GetFromAddress()

			vaultID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgMint(vaultID, addr.String(), amount)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewRepayCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repay [vault-id] [amount]",
		Args:  cobra.ExactArgs(2),
		Short: "repay ",
		Long: `repay.

			Example:
			$ onomyd tx repay 1 1000atom --from mykey
	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			addr := clientCtx.GetFromAddress()

			vaultID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgRepay(vaultID, addr.String(), amount)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCloseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "close [vault-id]",
		Args:  cobra.ExactArgs(1),
		Short: "close ",
		Long: `close.

			Example:
			$ onomyd tx close 1 --from mykey
	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			addr := clientCtx.GetFromAddress()

			vaultID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgClose(vaultID, addr.String())
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
