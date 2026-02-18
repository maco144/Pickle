package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/maco144/pickle/x/workqueue/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinEditDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdSubmitWork(),
		CmdValidateWork(),
		CmdRejectWork(),
	)

	return cmd
}

// CmdSubmitWork creates a command to submit work
func CmdSubmitWork() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-work [work-type] [work-data]",
		Short: "Submit a new work unit for validation",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgSubmitWork{
				Submitter: clientCtx.GetFromAddress().String(),
				WorkType:  args[0],
				WorkData:  []byte(args[1]),
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// CmdValidateWork creates a command to validate work
func CmdValidateWork() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate-work [work-id] [valid] [confidence]",
		Short: "Validate a work unit",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			valid := args[1] == "true"
			confidence := uint32(0)
			_, err = fmt.Sscanf(args[2], "%d", &confidence)
			if err != nil {
				return fmt.Errorf("invalid confidence: %w", err)
			}

			msg := &types.MsgValidateWork{
				Validator:  clientCtx.GetFromAddress().String(),
				WorkID:     args[0],
				Valid:      valid,
				Confidence: confidence,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// CmdRejectWork creates a command to reject work
func CmdRejectWork() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reject-work [work-id] [reason]",
		Short: "Reject a work unit",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgRejectWork{
				Validator: clientCtx.GetFromAddress().String(),
				WorkID:    args[0],
				Reason:    args[1],
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
