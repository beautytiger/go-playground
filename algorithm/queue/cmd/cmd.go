package main

import (
	"fmt"
	"github.com/beautytiger/go-playground/algorithm/queue"
)

func main(){
	fmt.Println("start")
	q := queue.NewQueue()
	q.Push(0)
	q.Print()
	q.Push(1)
	q.Print()
	q.Push(2)
	q.Print()
	q.Push(3)
	q.Print()
	q.Push(4)
	q.Print()
	q.Pop()
	q.Push(5)
	q.Print()
	q.Pop()
	q.Pop()
	q.Pop()
	q.Pop()
	q.Pop()
	q.Pop()
	q.Print()
}
