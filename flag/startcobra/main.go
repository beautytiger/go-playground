package main

import (
	"os"
	"playground/flag/startcobra/cmd"
)

func main(){
	cmdRoot := cmd.NewRootCommand()
	if err := cmdRoot.Execute(); err != nil {
		os.Exit(1)
	}
}
