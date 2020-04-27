package cmd

import (
	"errors"
	"github.com/spf13/cobra"
)

type ErrorHandling int

const (
	ContinueOnParseError ErrorHandling = 1
	ExitOnParseError ErrorHandling = 2
	PanicOnParseError ErrorHandling = 3
	ReturnOnDevidedByZero ErrorHandling = 4
	PanicOnDividedByZero ErrorHandling = 5
)

type OpType int

const (
	ADD OpType = 1
	MINUS OpType = 2
	MULTIPLY OpType = 3
	DIVIDE OpType = 4
)

var (
	parseHandling int
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "math",
		Short: "math calc the accumulative result",
		Run: func(cmd *cobra.Command, args []string){
			Error(cmd, args, errors.New("unrecognized subcommand"))
		},
	}
	rootCmd.PersistentFlags().IntVarP(&parseHandling, "parse_error", "p", int(ContinueOnParseError), "do what when parse args error")
	// add sub commands
	rootCmd.AddCommand(NewAddCmd())
	rootCmd.AddCommand(NewMinusCmd())
	rootCmd.AddCommand(NewDividedCmd())
	rootCmd.AddCommand(NewMultiplyCmd())
	return rootCmd
}