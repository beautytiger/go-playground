package main

import (
	"fmt"
)

type Foo struct {
	s []string
}

func main() {
	fmt.Println("start")
	var a *Foo
	b := new(Foo)
	c := &Foo{}
	d := &Foo{
		s: []string{"a", "b"},
	}
	fmt.Printf("a value %v type %T\n", a, a)
	fmt.Printf("b value %v type %T\n", b, b)
	fmt.Printf("c value %v type %T\n", c, c)
	//struct一般不太好直接比较
	//struct containing []string cannot be compared
	//fmt.Println("*b == *c", *b == *c)
	fmt.Println("start range over d")
	ranger(d)
	fmt.Println("start range over a")
	ranger(a)
	fmt.Println("start range over b")
	ranger(b)
	fmt.Println("start range over c")
	ranger(c)
}

func ranger(f *Foo) {
	fmt.Println("before range loop")
	if f != nil {
		fmt.Println("f is not nil")
	} else {
		fmt.Println("f is nil")
		return
	}
	if (*f).s != nil {
		fmt.Println("f.s is not nil")
	} else {
		fmt.Println("f.s is nil")
	}
	for k, v := range f.s {
		fmt.Println("inside range loop")
		fmt.Println(k, v)
	}
}
