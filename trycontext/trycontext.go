package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	for n := range generator(ctx) {
		fmt.Println(n)
		if n == 5 {
			cancel()
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
				close(dst)
				return
			case dst <- n:
				n++
			}
		}
	}()
	fmt.Println("done return chan")
	return dst
}
