package main

import (
	"fmt"
	"flag"
	"os"
)

var (
	intFlag int
	strFlag string
	boolFlag bool
	longF string
	longFusage = "long input from user"
	longDefault = "default"
)

func init() {
	flag.IntVar(&intFlag, "intf", 0,"please give me an int")
	flag.BoolVar(&boolFlag, "boolf", false, "please input an bool value")
	flag.StringVar(&strFlag, "strf", "default", `please input an bool (default is "default")`)
	flag.StringVar(&longF, "long", longDefault, longFusage)
	flag.StringVar(&longF, "l", longDefault, longFusage + "(for short)")
}

func main() {
	flag.Parse()

	fmt.Printf("intf is %v, type is %T\n", intFlag, intFlag)
	fmt.Printf("boolf is %v, type is %T\n", strFlag, strFlag)
	fmt.Printf("strf is %v, type is %T\n", boolFlag, boolFlag)
	fmt.Printf("longF is %v, type is %T\n", longF, longF)

	fmt.Printf("input args: %v, type: %T, length: %d\n", os.Args, os.Args, len(os.Args))
}
