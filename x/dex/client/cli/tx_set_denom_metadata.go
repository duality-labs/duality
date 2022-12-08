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
		Use:   "set-denom-metadata [description] [DenomUnits (unit1:exponent1,unit2:exponent2...)] [Display] [Name] [Symbol] [Base]",
		Short: "Broadcast message setDenomMetadata",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDescription := args[0]
			argDenomUnits := strings.Split(args[1], ",")
			argDisplay := args[2]
			argName := args[3]
			argSymbol := args[4]
			argBase := args[5]

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
				Name: argName,
				Symbol: argSymbol,
				DenomUnits: denomUnits,
				Base: argBase,
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
