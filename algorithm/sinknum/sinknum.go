package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
}

func NumberSink(array []int) []int {
	countZero := 0
	lastNonZero := 0
	for i := 0; i < len(array); i++ {
		if array[i] == 0 {
			countZero++
		} else {
			array[lastNonZero] = array[i]
			lastNonZero++
		}
	}
	for j := lastNonZero+1; j < len(array); j++{
		array[j] = 0
	}
	return array
}
