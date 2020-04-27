package main

import "time"
import "fmt"

func main() {
	go spinner(time.Second)
	var n int = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib (x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1)+fib(x-2)
}