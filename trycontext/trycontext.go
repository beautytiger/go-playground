package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	handler(ctx)
	time.Sleep(6*time.Second)
}

func handler(ctx context.Context) {
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()
	go dispatcher(cancelCtx)
	select {
	case <- time.NewTimer(2*time.Second).C:
		log.Println("handler will cancel all the work")
	}
}

func dispatcher(ctx context.Context) {
	c := make(chan string, 1)
	timeoutCtx, timeoutCancel := context.WithTimeout(ctx, 3*time.Second)
	defer timeoutCancel()
	go worker(timeoutCtx, c)
	select {
	case <- ctx.Done():
		log.Println("parent context cancelled")
	case <- timeoutCtx.Done():
		log.Println("current context timeout")
	case r := <- c:
		// 这里有问题，和上面的ctx.Done()会发生竞态
		// https://blog.golang.org/race-detector
		log.Println("get result from worker:", r)
		return
	}
	r := <- c
	log.Println("final get result from worker:", r)
}

func worker(ctx context.Context, c chan<- string) {
	t := time.NewTimer(4*time.Second)
	select {
	case <- t.C:
		log.Println("worker work done")
		c <- "done"
		return
	case <- ctx.Done():
		log.Println("parent cancel worker's work")
		c <- "stopped"
		return
	}
}