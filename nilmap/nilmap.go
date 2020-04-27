package main

import "fmt"

func main() {
	//var set map[string]int
	var set = make(map[string]int)
	set["a"] = 1
	set["b"] = 2
	fmt.Println(set)
}
