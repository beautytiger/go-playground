package main

import (
	"fmt"
	"sync"
)

// bufPool stores pointers to scratch buffers.
var bufPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 0, 256)
	},
}

func main() {
	b := bufPool.Get().([]byte)
	b = append(b, byte('a'))
	s := string(b)
	b = b[:0]
	bufPool.Put(&b)
	fmt.Println(s)

	c := bufPool.Get().([]byte)
	c = append(c, byte('a'))
	ss := string(c)
	c = c[:0]
	bufPool.Put(&c)
	fmt.Println(ss)
}