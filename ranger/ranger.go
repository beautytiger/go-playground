package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	//rangeString()
	//rangeArray()
	//rangeSlice()
	//appendSlice()
	//rangeMap()
	//mapAssign()
	//chanUse()
	chanRange()
	//nilInterface()
}

// int cannot be ranged
//func rangeInt() {
//	var a int
//	for i := range a {
//		fmt.Println(i)
//	}
//}

//can not range over bool
//func rangeBool() {
//	fmt.Println("start range a nil string")
//	var bo bool
//	for s := range bo {
//		fmt.Println("inside a for range loop")
//		fmt.Println(s)
//	}
//	fmt.Println("end range a nil string")
//}

func rangeString() {
	//ways to define a string variable
	//var s string
	//var s = ""
	//s := new(string)
	//var s *string

	fmt.Println("start range a nil string")
	var str string
	fmt.Println("declare a string:", str)
	fmt.Println(`str == ""`, str == "")
	//string cannot compare to nil
	//fmt.Println("str == nil", str == nil)
	//cannot convert nil to type string
	//fmt.Println(`"" == string(nil)`, "" == string(nil))
	for s := range str {
		fmt.Println("inside a for range loop")
		fmt.Println(s)
	}
	fmt.Println("end range a nil string")

	var content string = "ab"
	for i, j := range content {
		fmt.Println("inside b for range loop")
		fmt.Println(i, j)
	}

	var pon = new(string)
	fmt.Println(`*pon == ""`, *pon == "")
	fmt.Println(`pon == nil`, pon == nil)
	fmt.Println("new a string pointer: ", pon)
	fmt.Println("new a string: ", *pon)
	for i, j := range *pon {
		fmt.Println("inside c for range loop")
		fmt.Println(i, j)
	}

	var p *string
	//invalid memory address or nil pointer dereference
	//fmt.Println(`*p == ""`, *p == "")
	fmt.Println(`p == nil`, p == nil)
	//always check nil before pointer dereference
	if p != nil {
		fmt.Println("string pointer is not nil, can be ranged")
		// must dereference here, cannot range over a string pointer
		for i, j := range *p {
			fmt.Println("inside d for range loop")
			fmt.Println(i, j)
		}
	} else {
		fmt.Println("string pointer is nil")
	}

	t := "ab"
	p = &t
	fmt.Println(`p == nil`, p == nil)
	//always check nil before pointer dereference
	if p != nil {
		fmt.Println("string pointer is not nil, can be ranged")
		for i, j := range *p {
			fmt.Println("inside d for range loop")
			fmt.Println(i, j)
		}
	} else {
		fmt.Println("string pointer is nil")
	}
}

func rangeArray() {
	//ways to define array
	var a [2]int
	var b *[2]int
	c := new([2]int)
	var d = [2]int{}

	fmt.Println(`var b *[2]int == c := new([2]int)`, b == c)
	fmt.Println("var a [2]int:")
	for i, j := range a {
		fmt.Println(i, j)
	}

	fmt.Println("var b *[2]int == nil", b == nil)
	//nil check is must, or it will cause runtime panic
	if b != nil {
		for i, j := range b {
			fmt.Println(i, j)
		}
	}
	b = &a
	fmt.Println("b = &a")
	if b != nil {
		//you can range over a array pointer, it will dereferenced automatically
		for i, j := range b {
			fmt.Println(i, j)
		}
	}
	fmt.Println("c := new([2]int) == nil", c == nil)
	fmt.Println("range over array pointer c")
	for i, j := range c {
		fmt.Println(i, j)
	}
	fmt.Println("range over array pointer *c")
	for i, j := range *c {
		fmt.Println(i, j)
	}
	fmt.Println("range over var d = [2]int{}")
	for i, j := range d {
		fmt.Println(i, j)
	}
}

func rangeSlice() {
	//ways to define array
	var a []int
	var b *[]int
	c := new([]int)
	var d = []int{}
	e := make([]int, 0)

	fmt.Println(`var b *[]int == c := new([]int)`, b == c)
	fmt.Println("var a []int == nil", a == nil)
	for i, j := range a {
		fmt.Println(i, j)
	}

	fmt.Println("var b *[]int == nil", b == nil)
	//nil check is must, or it will cause runtime panic
	//for i, j := range *b {
	//	fmt.Println(i, j)
	//}
	if b != nil {
		//cannot range over slice pointer, must deref first
		for i, j := range *b {
			fmt.Println(i, j)
		}
	}
	b = &a
	fmt.Println("b = &a")
	if b != nil {
		for i, j := range *b {
			fmt.Println(i, j)
		}
	}
	fmt.Println("c := new([]int) == nil", c == nil)
	//range over slice pointer is not a good idea
	//fmt.Println("range over array pointer c")
	//for i, j := range c {
	//	fmt.Println(i, j)
	//}
	fmt.Println("range over array pointer *c")
	for i, j := range *c {
		fmt.Println(i, j)
	}
	fmt.Println("range over var d = [2]int{}")
	fmt.Println("d = [2]int{} == nil", d == nil)
	fmt.Println("d =", d)
	//empty slice is not nil slice
	//https://www.pixelstech.net/article/1539870875-Empty-slice-vs-nil-slice-in-GoLang
	//fmt.Println("d[0] = ", d[0])
	for i, j := range d {
		fmt.Println(i, j)
	}
	//slice can only be compared to nil
	//fmt.Println(`var d = []int{} == e := make([]int, 0)`, d == e)
	fmt.Println("range over e := make([]int, 0)")
	fmt.Println("e := make([]int, 0) == nil", e == nil)
	for i, j := range e {
		fmt.Println(i, j)
	}
}

func appendSlice() {
	//ways to define array
	var a []int
	var b *[]int
	c := new([]int)
	var d = []int{}
	e := make([]int, 0)

	a = append(a, 1)
	fmt.Println("a=", a)
	//nil slice pointer can not be append
	//*b = append(*b, 1)
	fmt.Println("b == nil", b == nil)
	fmt.Println("b=", b)
	//nil slice can be append
	*c = append(*c, 1)
	fmt.Println("c=", c)
	fmt.Println("*c=", *c)
	d = append(d, 1)
	fmt.Println("d=", d)
	e = append(e, 1)
	fmt.Println("e=", e)
}

func rangeMap() {
	var a map[int]int
	var b *map[int]int
	c := new(map[int]int)
	var d = map[int]int{}
	var e = make(map[int]int)
	fmt.Printf("a=%v\nb=%v\nc=%v\nd=%v\ne=%v\n", a, b, c, d, e)
	//map can only be compared to nil
	//fmt.Println(`var d = map[int]int{} == var e = make(map[int]int)`, d == e)
	fmt.Println("a==nil", a==nil)
	fmt.Println("b==nil", b==nil)
	//nil pointer deref
	//fmt.Println("*b==nil", *b==nil)
	fmt.Println("c==nil", c==nil)
	fmt.Println("*c==nil", *c==nil)
	fmt.Println("d==nil", d==nil)
	fmt.Println("e==nil", e==nil)
	// nil map can be ranged
	for k, v := range a {
		fmt.Println(k, v)
	}
	//range map pointer, compile time error
	//for k, v := range b {
	//	fmt.Println(k, v)
	//}
	//range map pointer deref, runtime error
	//for k, v := range *b {
	//	fmt.Println(k, v)
	//}
	if b != nil {
		for k, v := range *b {
			fmt.Println(k, v)
		}
	}
	for k, v := range *c {
		fmt.Println(k, v)
	}
	if c != nil {
		for k, v := range *c {
			fmt.Println(k, v)
		}
	}
	for k, v := range d {
		fmt.Println(k, v)
	}
	for k, v := range e {
		fmt.Println(k, v)
	}
}

func mapAssign() {
	var a map[int]int
	var b *map[int]int
	c := new(map[int]int)
	var d = map[int]int{}
	var e = make(map[int]int)
	fmt.Printf("a=%v\nb=%v\nc=%v\nd=%v\ne=%v\n", a, b, c, d, e)
	//runtime error
	//assignment to entry in nil map
	//a[1] = 1
	//runtime error
	//invalid memory address or nil pointer dereference
	//(*b)[1] = 1
	//runtime error
	//assignment to entry in nil map
	//(*c)[1] = 1
	d[1] = 1
	e[1] = 1
	fmt.Printf("a=%v\nb=%v\nc=%v\nd=%v\ne=%v\n", a, b, c, d, e)
}

func chanUse() {
	var a chan int
	b := new(chan int)
	c := make(chan int)
	fmt.Printf("a=%v\nb=%v\nc=%v\n", a, b, c)
	fmt.Printf("a=%T\nb=%T\nc=%T\n", a, b, c)
	num := 999
	fmt.Println("num", num)
	var res int
	fmt.Println("init res", res)
	//send to nil chan
	//runtime error
	//go echo(num, a)
	//receive from nil chan
	//runtime error
	//res = <- a
	fmt.Println("get res", res)
	//runtime error
	//new create a pointer pointing to nil chan
	//go echo(num, *b)
	//res = <- *b
	fmt.Println("get res", res)
	go echo(num, c)
	res = <- c
	fmt.Println("get res", res)

	//runtime error
	//send on closed channel
	//fmt.Println("close c")
	//close(c)
	//go echo(num, c)
	//time.Sleep(time.Second*1)
	//receive from close chan will get zero value of type
	//res = <- c
	//fmt.Println("get res", res)

	fmt.Println("read from closed chan")
	go func(c chan int){
		fmt.Println("send value to chan")
		c <- 657
		fmt.Println("close chan")
		close(c)
		fmt.Println("chan closed")
	}(c)
	fmt.Println("receive value from chan")
	//https://stackoverflow.com/questions/16105325/how-to-check-a-channel-is-closed-or-not-without-reading-it
	//the answer is no
	v, ok := <- c
	fmt.Println("<-c", v, ok)
	//closed chan will not block
	//and ok will be false
	v, ok = <- c
	fmt.Println("<-c", v, ok)
	v, ok = <- c
	fmt.Println("<-c", v, ok)
}

func echo(i int, c chan int) {
	fmt.Println("get int", i)
	c <- i
}

func chanRange() {
	a := make(chan int, 20)
	b := make(chan int)
	//Channel values can only be compared for equality.
	//Two channel values are considered equal if they originated
	//from the same make call (meaning they refer to the same channel value in memory).
	fmt.Println(`make(chan int, 0) == make(chan int)`, a==b)
	go func(ch chan int) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}(a)
	//range can only get one value from chan
	for k := range a {
		fmt.Println("get value from ch:", k)
	}
	k, ok := <- a
	fmt.Println("final result from ch", k, ok)
	//send to a closed chan will cause runtime panic
	//fmt.Println("a <- 333")
	//a <- 333

	var ch chan int
	go func(ch chan int) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}(ch)
	for k := range a {
		fmt.Println("get value from nil ch:", k)
	}
	chp := new(chan int)
	go func(ch *chan int) {
		for i := 0; i < 10; i++ {
			*ch <- i
		}
		close(*ch)
	}(chp)
	for k := range *chp {
		fmt.Println("get value from *chp:", k)
	}
}

func nilInterface() {
	var a interface{}
	fmt.Println("var a interface{} == nil", a == nil)
	a = 1
	fmt.Println("a = 1 == nil", a == nil)
	a = '1'
	fmt.Println("a = '1' == nil", a == nil)
	a = "1"
	fmt.Println("a = \"1\" == nil", a == nil)
	a = nil
	fmt.Println("a = nil == nil", a == nil)
	a = 1.1
	fmt.Println("a = 1.1 == nil", a == nil)
	a = false
	fmt.Println("a = false == nil", a == nil)
	a = struct{}{}
	fmt.Println("a = struct{}{} == nil", a == nil)


}