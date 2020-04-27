package main

import "fmt"

func main() {
	var x int = 99999
	var y *int
	fmt.Println("starting value of pointer y=", y)
	y = &x
	var p = &x
	z := &x
	fmt.Println("value stored in x=", x)
	fmt.Println("address of x=", &x)
	fmt.Println("value stored in variable y=", y)
	fmt.Println("value stored in variable p=", p)
	fmt.Println("value stored in variable z=", z)
	fmt.Printf("variable y type %T\n", y)
	fmt.Printf("variable p type %T\n", p)
	fmt.Printf("variable z type %T\n", z)
	fmt.Println("value stored in *z", *z)
}
