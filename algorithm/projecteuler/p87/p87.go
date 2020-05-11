package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(P87(50000000))
}

func P87(n int) int {
	primeUplimit := int(math.Sqrt(float64(n)))
	primes := PrimeGriddle(primeUplimit)
	var power2, power3, power4 []int
	for _, p := range primes {
		temp := p * p
		power2 = append(power2, temp)
		temp *= p
		if temp < n {
			power3 = append(power3, temp)
		}
		temp *= p
		if temp < n {
			power4 = append(power4, temp)
		}
	}
	var result = make(map[int]bool)
	for _, i := range power2 {
		for _, j := range power3 {
			if i + j >= n {
				break
			}
			for _, k := range power4 {
				if i + j + k >= n {
					break
				}
				result[i+j+k] = true
			}
		}
	}
	//fmt.Println(result)
	return len(result)
}


// return prime number array below number n
func PrimeGriddle(n int) []int {
	var temp = make([]int, n+1)
	for i:= 2; i <= n; i++ {
		temp[i] = i
	}
	for i :=2; i <= n; i++ {
		for j := 2; i*j <= n; j++ {
			temp[i*j] = 0
		}
	}
	var primes []int
	for i := 2; i <= n; i++ {
		if temp[i] != 0 {
			primes = append(primes, temp[i])
		}
	}
	return primes
}
