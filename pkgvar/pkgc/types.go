package pkgc

import (
	"fmt"
	"github.com/beautytiger/go-playground/pkgvar/pkga"
)

var V3 string = "in package c"

type C struct {
	Name string
}

func NewC() C {
	pkga.V1 = "changed in pkg c"
	return C{pkga.V1}
}

func Run() {
	fmt.Println("work in pkg c", pkga.V1)
}
