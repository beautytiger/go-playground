package main

import "fmt"

func main() {
	//nilpointer()
	nilslice()
}

// will cause runtime panic
//staticcheck cannot check this
func nilpointer() {
	var a *[]int
	// potential nil pointer deref
	*a = append(*a, 1)
	fmt.Println(a)
}

//append to nil slice is ok
func nilslice() {
	var a []int
	a = append(a, 1)
	fmt.Println(a)
}