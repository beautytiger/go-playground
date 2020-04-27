package main

import (
	"os"
	"github.com/beautytiger/go-playground/flag/startcobra/cmd"
)

func main(){
	cmdRoot := cmd.NewRootCommand()
	if err := cmdRoot.Execute(); err != nil {
		os.Exit(1)
	}
}
