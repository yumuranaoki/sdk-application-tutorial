package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

// GetCmdGetValue queries information about a name
func GetCmdGetValue(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-value [key]",
		Short: "get value",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			key := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/get-value/%s", queryRoute, key), nil)
			if err != nil {
				fmt.Printf("could not resolve name - %s \n", string(key))
				return nil
			}

			fmt.Println(string(res))

			return nil
		},
	}
}
