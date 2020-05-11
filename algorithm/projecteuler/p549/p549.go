package main

import "fmt"

func main() {
	//fmt.Println(FirstFactorialDivisibleBy(25))
	//var factorial, prime int = 4, 2
	//fmt.Println(PrimeCount(factorial, prime))
	//fmt.Println(PrimeCountRecursion(factorial, prime))
	//fmt.Println(HighestPowerDividesFactorial(factorial, prime))
	fmt.Println(p549(100000))
}

//还没有得到答案
func p549(n int) int {
	if n < 2 {
		return 0
	}
	var sum int
	for i := 2; i <= n; i++ {
		temp := FirstFactorialDivisibleBy(i)
		sum += temp
	}
	return sum
}

func GreatestCommonDivisor(x, y int) int {
	if x % y == 0 {
		return y
	}
	return GreatestCommonDivisor(y, x % y)
}

func FactorialLimit(n int) int {
	var limit int
	divs := Divisor(n)
	for divisor, exp := range divs {
		temp := ExpFactorialLimit(divisor, exp)
		if limit < ExpFactorialLimit(divisor, exp) {
			limit = temp
		}
	}
	return limit
}

//PrimeCount返回n!中包含prime因数的个数
//比如 4! 中包含3个2，则 PrimeCount(4, 2) -> 3
//自己通过规律推算的，求相关数学理论:)
//see HighestPowerDividesFactorial
func PrimeCount(n, prime int) int {
	var count int
	for n >= prime {
		count += n / prime
		n = n / prime
	}
	return count
}

//PrimeCountRecursion 是 PrimeCount的递归版本
func PrimeCountRecursion(n, prime int) int {
	if n < prime {
		return 0
	}
	if n == prime {
		return 1
	}
	return n / prime + PrimeCount(n/prime, prime)
}

//PrimeCount的简化版
//https://undergroundmathematics.org/divisibility-and-induction/factorial-fun/solution
func HighestPowerDividesFactorial(n, p int) int {
	var power, div int
	div = p
	for div <= n {
		power = power + n/div
		div = div*p
	}
	return power
}

//ExpFactorialLimit 返回能够整除x**y的阶乘的最后一个乘数
//比如 5**2 最小能够被10的阶乘整除 FactorialLimit(5, 2) -> 10
//使用的逆推法，十分的不高效
func ExpFactorialLimit(x, y int) int {
	var start int
	start = x*y-y/x*x
	for {
		if PrimeCount(start, x) >= y {
			return start
		}
		start += x
	}
}

//Divisor返回一个x的除数map，key为除数，value为除数的个数
func Divisor(x int) map[int]int {
	var divisor = make(map[int]int)
	for {
		if x % 2 == 0 {
			x = x / 2
			divisor[2]++
		} else {
			break
		}
	}
	for i := 3; i <= x; i=i+2 {
		for {
			if x % i == 0 {
				x = x / i
				divisor[i]++
			} else {
				break
			}
		}
	}
	return divisor
}

//这个算法也不是很高效
//https://www.geeksforgeeks.org/find-first-natural-number-whose-factorial-divisible-x/
func FirstFactorialDivisibleBy(n int) int {
	var f int = 1
	var temp int = n
	for f = 1; f <= n; f++ {
		temp = temp / GreatestCommonDivisor(f, temp)
		if temp == 1 {
			break
		}
	}
	return f
}

//有个人解了300多道题，可以作为参考
//https://en.wikipedia.org/wiki/Kempner_function
//https://euler.stephan-brumme.com/549/