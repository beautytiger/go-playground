package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	a := 9799
	b := int32(a)
	c := int64(b)
	d := int8(c)

	fmt.Printf("value: %v, type: %T\n", a, a)
	fmt.Printf("value: %v, type: %T\n", b, b)
	fmt.Printf("value: %v, type: %T\n", c, c)
	fmt.Printf("value: %v, type: %T\n", d, d)

	l := strconv.Itoa(a)
	m := string(c)
	o := strconv.FormatInt(c,17)
	fmt.Printf("value: %v, type: %T\n", l, l)
	fmt.Printf("value: %v, type: %T\n", m, m)
	fmt.Printf("value: %v, type: %T\n", o, o)

	s := "84734"

	if n, err := strconv.Atoi(s); err == nil {
		fmt.Printf("value: %v, type: %T\n", n, n)
	} else {
		fmt.Printf("s is not an integer")
	}

	if p, err := strconv.ParseInt(s, 10, 64); err == nil {
		fmt.Printf("value: %v, type: %T\n", p, p)
	} else {
		fmt.Printf("s is not an integer")
	}

	var x uint64 = 1 << 63
	var y float64 = math.Pow(2, 63)
	fmt.Printf("value: %v, type: %T\n", x, x)
	fmt.Printf("value: %v, type: %T\n", y, y)

}
