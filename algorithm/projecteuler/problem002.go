package main

import "fmt"

func main() {
	var limit int
	var sum int
	// go 无内置指数运算，下面这种做法行不通
	// limit = 4*10 ^ 6
	limit = 4000000
	var i, j, tmp = 1, 2, 0
	for i <= limit {
		if i%2 == 0 {
			sum += i
		}
		tmp = i + j
		i = j
		j = tmp
	}

	fmt.Println("result: ", sum)
}
