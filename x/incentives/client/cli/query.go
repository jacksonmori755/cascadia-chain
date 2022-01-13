package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/tharsis/evmos/x/incentives/types"
)

// GetQueryCmd returns the parent command for all incentives CLI query commands.
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the incentives module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetIncentivesCmd(),
		GetIncentiveCmd(),
		GetGasMetersCmd(),
		GetGasMeterCmd(),
		GetParamsCmd(),
	)
	return cmd
}

// GetIncentivesCmd queries the list of incentives
func GetIncentivesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "incentives",
		Short: "Gets incentives",
		Long:  "Gets incentives",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			req := &types.QueryIncentivesRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.Incentives(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetIncentiveCmd queries a given contract incentive
func GetIncentiveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "incentive [contract-address]",
		Short: "Gets incentive",
		Long:  "Gets incentive",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if !common.IsHexAddress(args[0]) {
				return fmt.Errorf("invalid contract address: %s", args[0])
			}

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryIncentiveRequest{
				Contract: args[0],
			}

			res, err := queryClient.Incentive(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetGasMetersCmd queries the list of incentives
func GetGasMetersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gas-meters [contract-address]",
		Short: "Gets meters for a given incentive contract",
		Long:  "Gets meters for a given incentive contract",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if !common.IsHexAddress(args[0]) {
				return fmt.Errorf("invalid contract address: %s", args[0])
			}

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			req := &types.QueryGasMetersRequest{
				Contract:   args[0],
				Pagination: pageReq,
			}

			res, err := queryClient.GasMeters(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetGasMeterCmd queries the list of incentives
func GetGasMeterCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gas-meter [contract-address] [participant-address]",
		Short: "Gets meter for a given incentive contract and user address",
		Long:  "Gets meter for a given incentive contract and user address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			if !common.IsHexAddress(args[0]) {
				return fmt.Errorf("invalid contract address: %s", args[0])
			}

			if !common.IsHexAddress(args[1]) {
				return fmt.Errorf("invalid user address: %s", args[0])
			}

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryGasMeterRequest{
				Contract:    args[0],
				Participant: args[1],
			}

			res, err := queryClient.GasMeter(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetParamsCmd queries the module parameters
func GetParamsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Gets incentives params",
		Long:  "Gets incentives params",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryParamsRequest{}

			res, err := queryClient.Params(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
