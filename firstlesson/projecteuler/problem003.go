package main

import "fmt"

func main() {
	var target = 600851475143
	factor := 2
	for target > factor {
		if target%factor == 0 {
			target = target / factor
			continue
		}
		factor += 1
	}
	fmt.Println("result:", target)
}
