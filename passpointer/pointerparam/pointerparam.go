package main

import "fmt"

//pointer as param

type ST struct{
	s string
}

//pointer是值传递，没有办法把一个pointer变量间接设置为nil
func changePointer(p *int) {
	fmt.Println("inside func")
	fmt.Println(p)
	p = nil
	fmt.Println(p)
	fmt.Println("finish inside func")
}

func main() {
	var b *int
	fmt.Println(b)
	a := 999
	b = &a
	fmt.Println(b)
	changePointer(b)
	fmt.Println(b)

	l := ST{}
	m := ST{}
	fmt.Println(l==m)
}
