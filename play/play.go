package main

import "fmt"

func main() {
	a := make([]int, 5)
	for i := 1; i < 20; i++ {
		a = append(a, i)
		fmt.Printf("%d: slice len %d, cap %d\n", i, len(a), cap(a))
	}
}
