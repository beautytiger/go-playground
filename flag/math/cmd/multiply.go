package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func NewMultiplyCmd() *cobra.Command {
	multiplyCmd := &cobra.Command{
		Use: "multi",
		Short: "multi subcommand multiply all parsed args.",
		Run: func(cmd *cobra.Command, args []string) {
			values := ConvertArgsToFloat64Slice(args, ErrorHandling(parseHandling))
			result := calc(values, MULTIPLY)
			fmt.Printf("%s = %.2f\n", strings.Join(args, "*"), result)
		},
	}
	return multiplyCmd
}