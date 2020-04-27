package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a
	b = []int{2, 3, 4}
	fmt.Println("vim-go, %v", b)
}
