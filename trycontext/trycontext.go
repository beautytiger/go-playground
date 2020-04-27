package main

import "fmt"
import "context"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for n := range generator(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	fmt.Println("main done")
}

func generator(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("done generating")
				return
			case dst <- n:
				n++
			}
		}
	}()
	fmt.Println("done return chan")
	return dst
}
