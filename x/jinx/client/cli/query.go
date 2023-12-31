package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/incubus-network/nemo/x/jinx/types"
)

// flags for cli queries
const (
	flagName  = "name"
	flagDenom = "denom"
	flagOwner = "owner"
)

// GetQueryCmd returns the cli query commands for the  module
func GetQueryCmd() *cobra.Command {
	jinxQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the jinx module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmds := []*cobra.Command{
		queryParamsCmd(),
		queryAccountsCmd(),
		queryDepositsCmd(),
		queryUnsyncedDepositsCmd(),
		queryTotalDepositedCmd(),
		queryBorrowsCmd(),
		queryUnsyncedBorrowsCmd(),
		queryTotalBorrowedCmd(),
		queryInterestRateCmd(),
		queryReserves(),
		queryInterestFactorsCmd(),
	}

	for _, cmd := range cmds {
		flags.AddQueryFlagsToCmd(cmd)
	}

	jinxQueryCmd.AddCommand(cmds...)

	return jinxQueryCmd
}

func queryParamsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Short: "get the jinx module parameters",
		Long:  "Get the current global jinx module parameters.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(&res.Params)
		},
	}
}

func queryAccountsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "accounts",
		Short: "query jinx module accounts",
		Long:  "Query for all jinx module accounts",
		Example: fmt.Sprintf(`%[1]s q %[2]s accounts
%[1]s q %[2]s accounts`, version.AppName, types.ModuleName),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			req := &types.QueryAccountsRequest{}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Accounts(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}

func queryUnsyncedDepositsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unsynced-deposits",
		Short: "query jinx module unsynced deposits with optional filters",
		Long:  "query for all jinx module unsynced deposits or a specific unsynced deposit using flags",
		Example: fmt.Sprintf(`%[1]s q %[2]s unsynced-deposits
%[1]s q %[2]s unsynced-deposits --owner fury1l0xsq2z7gqd7yly0g40y5836g0appuma0grvkv --denom bnb
%[1]s q %[2]s unsynced-deposits --denom ufury
%[1]s q %[2]s unsynced-deposits --denom btcb`, version.AppName, types.ModuleName),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			ownerBech, err := cmd.Flags().GetString(flagOwner)
			if err != nil {
				return err
			}
			denom, err := cmd.Flags().GetString(flagDenom)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			req := &types.QueryUnsyncedDepositsRequest{
				Denom:      denom,
				Pagination: pageReq,
			}

			if len(ownerBech) != 0 {
				depositOwner, err := sdk.AccAddressFromBech32(ownerBech)
				if err != nil {
					return err
				}
				req.Owner = depositOwner.String()
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.UnsyncedDeposits(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "unsynced-deposits")

	cmd.Flags().String(flagOwner, "", "(optional) filter for unsynced deposits by owner address")
	cmd.Flags().String(flagDenom, "", "(optional) filter for unsynced deposits by denom")

	return cmd
}

func queryDepositsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposits",
		Short: "query jinx module deposits with optional filters",
		Long:  "query for all jinx module deposits or a specific deposit using flags",
		Example: fmt.Sprintf(`%[1]s q %[2]s deposits
%[1]s q %[2]s deposits --owner fury1l0xsq2z7gqd7yly0g40y5836g0appuma0grvkv --denom bnb
%[1]s q %[2]s deposits --denom ufury
%[1]s q %[2]s deposits --denom btcb`, version.AppName, types.ModuleName),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			ownerBech, err := cmd.Flags().GetString(flagOwner)
			if err != nil {
				return err
			}
			denom, err := cmd.Flags().GetString(flagDenom)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			req := &types.QueryDepositsRequest{
				Denom:      denom,
				Pagination: pageReq,
			}

			if len(ownerBech) != 0 {
				depositOwner, err := sdk.AccAddressFromBech32(ownerBech)
				if err != nil {
					return err
				}
				req.Owner = depositOwner.String()
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Deposits(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "deposits")

	cmd.Flags().String(flagOwner, "", "(optional) filter for deposits by owner address")
	cmd.Flags().String(flagDenom, "", "(optional) filter for deposits by denom")

	return cmd
}

func queryUnsyncedBorrowsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unsynced-borrows",
		Short: "query jinx module unsynced borrows with optional filters",
		Long:  "query for all jinx module unsynced borrows or a specific unsynced borrow using flags",
		Example: fmt.Sprintf(`%[1]s q %[2]s unsynced-borrows
%[1]s q %[2]s unsynced-borrows --owner fury1l0xsq2z7gqd7yly0g40y5836g0appuma0grvkv
%[1]s q %[2]s unsynced-borrows --denom bnb`, version.AppName, types.ModuleName),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			ownerBech, err := cmd.Flags().GetString(flagOwner)
			if err != nil {
				return err
			}
			denom, err := cmd.Flags().GetString(flagDenom)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			req := &types.QueryUnsyncedBorrowsRequest{
				Denom:      denom,
				Pagination: pageReq,
			}

			if len(ownerBech) != 0 {
				borrowOwner, err := sdk.AccAddressFromBech32(ownerBech)
				if err != nil {
					return err
				}
				req.Owner = borrowOwner.String()
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.UnsyncedBorrows(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "unsynced borrows")

	cmd.Flags().String(flagOwner, "", "(optional) filter for unsynced borrows by owner address")
	cmd.Flags().String(flagDenom, "", "(optional) filter for unsynced borrows by denom")

	return cmd
}

func queryBorrowsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "borrows",
		Short: "query jinx module borrows with optional filters",
		Long:  "query for all jinx module borrows or a specific borrow using flags",
		Example: fmt.Sprintf(`%[1]s q %[2]s borrows
%[1]s q %[2]s borrows --owner fury1l0xsq2z7gqd7yly0g40y5836g0appuma0grvkv
%[1]s q %[2]s borrows --denom bnb`, version.AppName, types.ModuleName),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			ownerBech, err := cmd.Flags().GetString(flagOwner)
			if err != nil {
				return err
			}
			denom, err := cmd.Flags().GetString(flagDenom)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			req := &types.QueryBorrowsRequest{
				Denom:      denom,
				Pagination: pageReq,
			}

			if len(ownerBech) != 0 {
				borrowOwner, err := sdk.AccAddressFromBech32(ownerBech)
				if err != nil {
					return err
				}
				req.Owner = borrowOwner.String()
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Borrows(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "borrows")

	cmd.Flags().String(flagOwner, "", "(optional) filter for borrows by owner address")
	cmd.Flags().String(flagDenom, "", "(optional) filter for borrows by denom")

	return cmd
}

func queryTotalBorrowedCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "total-borrowed",
		Short: "get total current borrowed amount",
		Long:  "get the total amount of coins currently borrowed using flags",
		Example: fmt.Sprintf(`%[1]s q %[2]s total-borrowed
%[1]s q %[2]s total-borrowed --denom bnb`, version.AppName, types.ModuleName),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			denom, err := cmd.Flags().GetString(flagDenom)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.TotalBorrowed(context.Background(), &types.QueryTotalBorrowedRequest{
				Denom: denom,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().String(flagDenom, "", "(optional) filter total borrowed coins by denom")

	return cmd
}

func queryTotalDepositedCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "total-deposited",
		Short: "get total current deposited amount",
		Long:  "get the total amount of coins currently deposited using flags",
		Example: fmt.Sprintf(`%[1]s q %[2]s total-deposited
%[1]s q %[2]s total-deposited --denom bnb`, version.AppName, types.ModuleName),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			denom, err := cmd.Flags().GetString(flagDenom)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.TotalDeposited(context.Background(), &types.QueryTotalDepositedRequest{
				Denom: denom,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().String(flagDenom, "", "(optional) filter total deposited coins by denom")

	return cmd
}

func queryInterestRateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "interest-rate",
		Short: "get current money market interest rates",
		Long:  "get current money market interest rates",
		Example: fmt.Sprintf(`%[1]s q %[2]s interest-rate
%[1]s q %[2]s interest-rate --denom bnb`, version.AppName, types.ModuleName),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			denom, err := cmd.Flags().GetString(flagDenom)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.InterestRate(context.Background(), &types.QueryInterestRateRequest{
				Denom: denom,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().String(flagDenom, "", "(optional) filter interest rates by denom")

	return cmd
}

func queryReserves() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reserves",
		Short: "get total current Jinx module reserves",
		Long:  "get the total amount of coins currently held as reserve by the Jinx module",
		Example: fmt.Sprintf(`%[1]s q %[2]s reserves
%[1]s q %[2]s reserves --denom bnb`, version.AppName, types.ModuleName),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			denom, err := cmd.Flags().GetString(flagDenom)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Reserves(context.Background(), &types.QueryReservesRequest{
				Denom: denom,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().String(flagDenom, "", "(optional) filter reserve coins by denom")

	return cmd
}

func queryInterestFactorsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "interest-factors",
		Short: "get current global interest factors",
		Long:  "get current global interest factors",
		Example: fmt.Sprintf(`%[1]s q %[2]s interest-factors
%[1]s q %[2]s interest-factors --denom bnb`, version.AppName, types.ModuleName),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			denom, err := cmd.Flags().GetString(flagDenom)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.InterestFactors(context.Background(), &types.QueryInterestFactorsRequest{
				Denom: denom,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().String(flagDenom, "", "(optional) filter interest factors by denom")

	return cmd
}
