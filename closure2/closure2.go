package main

import (
	"errors"
	"fmt"
)

func main() {
	errClosure()
	fmt.Println(powerSum(3)(4))
	fmt.Println(sumPower(3)(4)(5))
}

func errClosure() {
	var err error // nil error

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

//求两个数的平方和，使用闭包实现
func powerSum(x int) func(int) int {
	return func(y int) int {
		return x*x + y*y
	}
}

//求3个数的平方和
func sumPower(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}
