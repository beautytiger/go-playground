package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var (
	dividedByZeroHandling int
)

func NewDividedCmd() *cobra.Command {
	divideCmd := &cobra.Command{
		Use: "divide",
		Short: "divide subcommand divede all parsed args.",
		Run: func(cmd *cobra.Command, args []string) {
			values := ConvertArgsToFloat64Slice(args, ErrorHandling(parseHandling))
			result := calc(values, DIVIDE)
			fmt.Printf("%s = %.2f\n", strings.Join(args, "/"), result)
		},
	}
	divideCmd.Flags().IntVarP(&dividedByZeroHandling, "divide_by_zero", "d", int(PanicOnDividedByZero), "do what when divided by zero")
	return divideCmd
}