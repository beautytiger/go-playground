package main

import "fmt"

func sum(nums ...int) int {
	fmt.Printf("get data %v type %T\n", nums, nums)
	total := 0
	for _, n := range nums {
		total += n
	}
	nums[0] = 1000
	return total
}

var data []int = []int{1, 2, 3, 4, 5}
//数组不能使用...传递
var array = [...]int{1, 2, 3, 4, 5}
func main() {
	fmt.Println(sum(1,2,3,4,5))
	fmt.Println(sum(data...))
	fmt.Println(data)
	//fmt.Println(sum(array...))
}
