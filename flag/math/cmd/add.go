package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func NewAddCmd() *cobra.Command {
	addCmd := &cobra.Command{
		Use: "add",
		Short: "add subcommand add all parsed args.",
		Run: func(cmd *cobra.Command, args []string) {
			values := ConvertArgsToFloat64Slice(args, ErrorHandling(parseHandling))
			result := calc(values, ADD)
			fmt.Printf("%s = %.2f\n", strings.Join(args, "+"), result)
		},
	}
	return addCmd
}