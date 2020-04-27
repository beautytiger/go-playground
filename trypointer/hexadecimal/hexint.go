package main

import "fmt"

func main() {
	x := 0xABCDEF
	y := 0xabcdef

	fmt.Printf("type of variable x: %T\n", x)
	fmt.Printf("value of x in dexadecimal: %X\n", x)
	fmt.Printf("value of x in decimal: %v\n", x)

	fmt.Printf("type of variable y: %T\n", y)
	fmt.Printf("value of y in dexadecimal: %X\n", y)
	fmt.Printf("value of y in decimal: %v\n", y)
}
