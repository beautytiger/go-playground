package main

import "fmt"

func main() {
	//assignment to entry in nil map
	//will cause runtime error
	//var set map[string]int
	//set["a"] = 1
	//set["b"] = 2

	//set := new(map[string]int)
	//(*set)["a"] = 1
	//(*set)["b"] = 2

	//must use make to initialize map
	//the right way to use map
	var set = make(map[string]int)
	set["a"] = 1
	set["b"] = 2
	fmt.Println(set)
}
