package main

import "fmt"

func power(a *int) {
	*a = *a * *a
}

func main() {
	n := 99
	fmt.Println("before calling func, n=", n)
	power(&n)
	fmt.Println("after calling func, n=", n)
}
