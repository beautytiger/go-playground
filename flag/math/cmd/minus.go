package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func NewMinusCmd() *cobra.Command {
	minusCmd := &cobra.Command{
		Use: "minus",
		Short: "minus subcommand minus all parsed args.",
		Run: func(cmd *cobra.Command, args []string) {
			values := ConvertArgsToFloat64Slice(args, ErrorHandling(parseHandling))
			result := calc(values, MINUS)
			fmt.Printf("%s = %.2f\n", strings.Join(args, "-"), result)
		},
	}
	return minusCmd
}