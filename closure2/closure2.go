package main

import (
	"errors"
	"fmt"
)

func main() {
	var err error

	f := func() (string, error) {
		return "placeholder", errors.New("I am an error")
	}

	func() {
		str, err := f()
		fmt.Println(str)
		fmt.Printf("got err in func: %v\n", err)
	}()

	fmt.Printf("the outer err is: %v\n", err)
}
