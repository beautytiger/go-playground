package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	fmt.Println(p112(99))
}

//rate为百分比值，整数，比如50表示50%
func p112(rate int) int {
	bouncyCount := 0
	positive := 1
	for {
		if IsBouncyNum(positive) {
			bouncyCount++
		}
		if bouncyCount * 100 == positive * rate {
			return positive
		}
		positive++
	}
}

func IsIncrNum(n int) bool {
	currDigit := n % 10
	n = n / 10
	for n != 0 {
		nextDigit := n % 10
		if nextDigit > currDigit {
			return false
		}
		currDigit = nextDigit
		n = n / 10
	}
	return true
}

func IsDecrNum(n int) bool {
	currDigit := n % 10
	n = n / 10
	for n != 0 {
		nextDigit := n % 10
		if nextDigit < currDigit {
			return false
		}
		currDigit = nextDigit
		n = n / 10
	}
	return true
}

func IsBouncyNum(n int) bool {
	return !IsIncrNum(n) && !IsDecrNum(n)
}
