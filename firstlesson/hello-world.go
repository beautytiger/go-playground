// this line is comment

/* this line also comment */
package main

import "fmt"

var a int = 100
var c = 3.14

func main() {
	b := 'm'
	const width = 10
	const length = 32
	var area int
	var areaptr *int
	area = width * length
	areaptr = &area
	fmt.Println("hello world") 
	fmt.Println(`this seems function`)
	fmt.Println(&a)
	fmt.Println(&b)
	fmt.Println(&c)
	fmt.Println("area is", area)
	fmt.Println(areaptr)
	fmt.Println(*areaptr)
	if a < 20 {
		fmt.Println("a is less than 20")
	} else {
		fmt.Println("a is not less than 20")
	}
	fmt.Println("a is", a)
	var grade string = "b"
	switch grade {
	case "A":
		fmt.Println("优秀")
	case "b":
		fmt.Println("良好")
	default:
		fmt.Println("未知")
	}
	var i = 0
	for true {
		fmt.Println("这是无限循环")
		i++
		if i > 1000 {
			break
		}
		fmt.Println(i)
	}
	for m := 0; m<10; m++ {
		fmt.Println("m", m)
	}
	x, y := 90, 100
	for x<y {
		x++
		fmt.Println("x", x)
	}

	number := [10]int{1,2,3,4,5}
	for idx, v := range number {
		fmt.Println(idx, v)
	}

	fmt.Println("length of 'hello'", len("hello"))
	fmt.Println(max(10, 20))
	stra, strb := swap("stra", "strb")
	fmt.Println(stra, strb)
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func swap(str1, str2 string) (string, string) {
	return str2, str1
}
