package main

import (
	"fmt"
	"github.com/beautytiger/go-playground/pkgvar/pkga"
	"github.com/beautytiger/go-playground/pkgvar/pkgb"
	"github.com/beautytiger/go-playground/pkgvar/pkgc"
)

func main() {
	fmt.Println(pkga.V1)
	pkga.V1 = "in main"
	//fmt.Println(pkgb.V2)
	s := pkga.NewStruct()
	fmt.Println(s.Name)

	//c :=pkgc.NewC()
	//fmt.Println(c.Name)

	pkgc.Run()
	pkgb.Run()
}
