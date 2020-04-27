package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use: "clone url [dest]",
	Short: "clone a repository into a new directory",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%v -> %T\n", cmd, cmd)
		fmt.Println("----")
		fmt.Printf("%v -> %T\n", args, args)
		output, err := ExecuteCommand("git", "clone", args...)
		fmt.Println(output)
		if err != nil {
			Error(cmd, args, err)
		}
	},
}