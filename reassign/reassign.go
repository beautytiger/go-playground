package main

import (
	"fmt"
	"math"
)

func main() {
	abs := math.Abs(float64(3.14))
	a := byte('f')
	// Note: Must use float32 comparisons for underlying float32 value to get precise cutoffs right.
	if abs != 0 {
		if abs < 1e-6 || abs >= 1e21 {
			a = 'e'
			a = 'e'
		}
	}
	fmt.Println(string(a))
}
