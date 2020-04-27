package main

import (
	"fmt"
)

func main(){
	var a map[string]string
	var b []string
	m := new(map[string]string)
	s := new([]string)
	fmt.Println(m)
	fmt.Println(s)
	fmt.Println(a)
	fmt.Println(b)

	for k, v := range *m {
		fmt.Println(k, v)
	}
	for i, t := range *s {
		fmt.Println(i, t)
	}
	for k, v := range a {
		fmt.Println(k, v)
	}
	for i, t := range b {
		fmt.Println(i, t)
	}
}
