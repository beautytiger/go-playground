package main

import (
	"fmt"
)

func main() {
	a := "$hel.lo wor#ld"
	fmt.Println(a)
	r := reverse(a)
	fmt.Println(r)
}

// reverse string except for special letters
func reverse(s string) string {
	rs := []rune(s)
	i := 0
	j := len(rs) -1
	for i < j {
		if (rs[i] < 'a' || rs[i] > 'z') && (rs[i] < 'A' || rs[i] > 'Z') {
			i++
			continue
		}
		if (rs[j] < 'a' || rs[j] > 'z') && (rs[j] < 'A' || rs[j] > 'Z') {
			j--
			continue
		}
		rs[i], rs[j] = rs[j], rs[i]
		i++
		j--
	}
	r := string(rs)
	return r
}
