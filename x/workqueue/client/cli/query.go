package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/maco144/pickle/x/workqueue/types"
)

// GetQueryCmd returns the query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinEditDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdQueryWork(),
		CmdQueryPendingWork(),
		CmdQueryValidatorStats(),
		CmdQueryTotalStats(),
	)

	return cmd
}

// CmdQueryWork creates a command to query a specific work unit
func CmdQueryWork() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "work [work-id]",
		Short: "Query a specific work unit",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			req := &types.QueryWorkRequest{WorkID: args[0]}

			res, err := queryClient.Work(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryPendingWork creates a command to query pending work
func CmdQueryPendingWork() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pending",
		Short: "Query pending work units",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			req := &types.QueryPendingWorkRequest{}

			res, err := queryClient.PendingWork(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryValidatorStats creates a command to query validator statistics
func CmdQueryValidatorStats() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validator-stats [validator-address]",
		Short: "Query validator statistics",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			req := &types.QueryValidatorStatsRequest{ValidatorAddress: args[0]}

			res, err := queryClient.ValidatorStats(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryTotalStats creates a command to query total statistics
func CmdQueryTotalStats() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "total-stats",
		Short: "Query total statistics",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			req := &types.QueryTotalStatsRequest{}

			res, err := queryClient.TotalStats(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
