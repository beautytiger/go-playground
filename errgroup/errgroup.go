package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.Background()
	eg, ctxn := errgroup.WithContext(context.Background())
	fmt.Println("ctx == ctxn", ctx == ctxn)

	var work = []int{7, 1, 2, 3, 4, 5, 6, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	for i := range work {
		v := work[i]
		eg.Go(func() error {
			r := 0
			for {
				select {
				case <-ctx.Done():
					fmt.Println(v, "job cancelled")
					return ctx.Err()
				default:
					r += v
					if r == 49 {
						return errors.New("I don't like 7")
					}
					if r > 300 {
						fmt.Println(r)
						return nil
					}
				}
			}
		})
	}
	if err := eg.Wait(); err != nil {
		fmt.Println("meet error!", err)
	}
	fmt.Println("process finished")
	fmt.Println("final context error", ctx.Err())
}
