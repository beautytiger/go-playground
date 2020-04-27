package main
import "fmt"
func main() {
	var a int = 1
	var b int = 1
	var c int = 2
	//var d string = "3"

	x := &a
	y := &b
	z := &c
	//o := &d
	p := &a

	m := x == y
	fmt.Println("pointer x == y", m)
	vm := *x == *y
	fmt.Println("pointer *x == *y", vm)
	pm := x == p
	fmt.Println("pointer x == p", pm)
	pp := &x == &p
	fmt.Println("pointer &x == &p", pp)
	n := x == z
	fmt.Println("pointer x == z", n)
	q := y == z
	fmt.Println("pointer y == z", q)
	//different types of pointer cannot compare
	//p := x == o
	//fmt.Println("pointer x == o", p)
}
