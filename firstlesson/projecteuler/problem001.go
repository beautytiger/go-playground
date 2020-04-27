package main

import "fmt"

func main() {
	target, eleA, eleB := 1000, 3, 5
	sumToA := getSum(eleA, target)
	sumToB := getSum(eleB, target)
	com := leastCommonMultiple(eleA, eleB)
	fmt.Println("lcm: ", com)
	sumToCom := getSum(com, target)
	var result = sumToA + sumToB - sumToCom
	fmt.Println("Result: ", result)
}

func getSum(element, limit int) int {
	// take care of below, so need b - 1
	var count = (limit - 1) / element
	var total = (count + 1) * count * element / 2
	return total
}

func leastCommonMultiple(na, nb int) int {
	return na * nb / greatestCommonDivisor(na, nb)
}

func greatestCommonDivisor(na, nb int) int {
	var tmp int
	if na < nb {
		tmp = na
		na = nb
		nb = tmp
	}
	var result = na % nb
	for {
		if result == 0 {
			break
		}
		tmp = na % nb
		na = nb
		nb = tmp
		result = na % nb
	}
	fmt.Println("gcd: ", nb)
	return nb
}
