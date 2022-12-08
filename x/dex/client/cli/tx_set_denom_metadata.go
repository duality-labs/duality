package cli

import (
	"strconv"
	"strings"
	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

)

var _ = strconv.Itoa(0)

func CmdSetDenomMetadata() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-denom-metadata [Name] [description] [DenomUnits (unit1:exponent1,unit2:exponent2...)] [Display-Denom] [Base-Denom] [Symbol]",
		Short: "Broadcast message setDenomMetadata",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argDescription := args[1]
			argDenomUnits := strings.Split(args[2], ",")
			argDisplay := args[3]
			argBase := args[4]
			argSymbol := args[5]


			var denomUnits []*banktypes.DenomUnit

			for _, u := range argDenomUnits{
				denomAndExponent := strings.Split(u, ":")
				denom, exponentStr := denomAndExponent[0], denomAndExponent[1]
				exponent, err := strconv.ParseInt(exponentStr, 10, 32)
				if err != nil {
					return err
				}

				denomUnits = append(denomUnits, &banktypes.DenomUnit{
					Denom: denom,
					Exponent: uint32(exponent),
				},
				)
			}

			metadata := banktypes.Metadata{
				Description: argDescription,
				Display: argDisplay,
				Base: argBase,
				Name: argName,
				Symbol: argSymbol,
				DenomUnits: denomUnits,
			}


			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetDenomMetadata(
				clientCtx.GetFromAddress().String(),
				metadata,
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
