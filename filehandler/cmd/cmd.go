package main

import (
	"fmt"
	"github.com/beautytiger/go-playground/filehandler"
	"github.com/mitchellh/go-homedir"
	flag "github.com/spf13/pflag"
)

var fileDir string

func main() {
	flag.Parse()
	//if flag.NArg() != 1 {
	//	fmt.Fprintf(os.Stderr, "missing required %s argument\n", "path")
	//	return
	//}
	//dir := flag.Arg(0)
	fileDir, err := homedir.Expand(fileDir)
	if err != nil {
		fmt.Println("path error")
		panic(err)
	}
	filehandler.ScanRun(fileDir)
}

func init() {
	flag.StringVarP(&fileDir, "dir", "d", "~", "the dir to examine")
}
