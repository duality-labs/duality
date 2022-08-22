package cli

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListNodes() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-nodes",
		Short: "list all nodes",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllNodesRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.NodesAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowNodes() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-nodes [node] [outgoing-edges]",
		Short: "shows a nodes",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argNode := args[0]
			argOutgoingEdges := args[1]

			params := &types.QueryGetNodesRequest{
				Node:          argNode,
				OutgoingEdges: argOutgoingEdges,
			}

			res, err := queryClient.Nodes(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
