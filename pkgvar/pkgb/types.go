package pkgb

import (
	"fmt"
	"github.com/beautytiger/go-playground/pkgvar/pkga"
)

var V2 string = "in pkg 2"


func init() {
	pkga.V1 = "changed in pkg b"
}

func Run() {
	fmt.Println("in package b", pkga.V1)
}
