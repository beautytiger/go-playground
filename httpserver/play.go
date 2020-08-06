package main

import (
	"fmt"
)

type query func(string) string

func exec(name string, vs ...query) string {
	ch := make(chan string)
	fn := func(i int) {
		ch <- vs[i](name)
	}
	var a string
	for i, _ := range vs {
		k := i
		go fn(k)
		// 异步变同步
		s := <- ch
		a += s
	}
	// 这里只会从chan中接收一个值，后面的会一直阻塞
	//此时的返回值是随机的
	//return <-ch
	return a
}

func main() {
	ret := exec("111", func(n string) string {
		return n + "func1"
	}, func(n string) string {
		return n + "func2"
	}, func(n string) string {
		return n + "func3"
	}, func(n string) string {
		return n + "func4"
	})
	fmt.Println(ret)
}
