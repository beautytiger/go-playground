package main

import "fmt"

func main() {
	var m map[string]interface{}
	var n map[string]string
	var o map[int]string
	var p map[string]int
	var q map[string]interface{}

	n = map[string]string{
		"n1": "s1",
		"n2": "s2",
	}
	o = map[int]string{
		1: "o1",
		2: "o2",
	}
	p = map[string]int{
		"p1": 1,
		"p2": 2,
	}
	q = map[string]interface{}{
		"q1": "i1",
		"q2": "i2",
	}
	m = map[string]interface{}{
		"m1": "v1",
		"m2": n,
		"m3": o,
		"m4": p,
		"m5": q,
	}
	iterMap("", m)
}

func iterMap(prefix string, m map[string]interface{}) {
	for k, v := range m {
		switch rv := v.(type) {
		case map[string]interface{}:
			iterMap(k, rv)
		default:
			fmt.Printf("%v.%v: %v\n", prefix, k, v)
		}
	}
}
