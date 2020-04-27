package main

import (
	"fmt"
	"os"
	"github.com/beautytiger/go-playground/flag/math/cmd"
)

func main(){
	var a float64 = 5
	var b float64 = 4
	fmt.Println(a/b)
	cmdRoot := cmd.NewRootCmd()
	if err := cmdRoot.Execute(); err != nil {
		os.Exit(1)
	}
}
