package cmd

import (
	"errors"
	"github.com/spf13/cobra"
)



func NewRootCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use: "git",
		Short: "git is a distributed version control system.",
		Long: `git is a free and open source distributed version control system
designed to handle everything from small to very large projects
with speed and efficency.`,
		Run: func(cmd *cobra.Command, args []string){
			Error(cmd, args, errors.New("unrecognized command"))
		},
	}
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(cloneCmd)
	return rootCmd
}
