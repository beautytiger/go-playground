package main

import "runtime"
import "fmt"

func main() {
	thought2()
}

func thought1() {
	fmt.Println(runtime.GOMAXPROCS(0))
	fmt.Println(runtime.NumCPU())
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}

func thought2() {
	go func() {
		for {
			fmt.Print(0)
		}
	}()
	for {
		fmt.Print(1)
	}
}