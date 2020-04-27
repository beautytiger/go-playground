package main

import (
	"fmt"
	"time"
)

func main() {
	a := time.Now()
	fmt.Println(a)
	var sleep time.Duration = 1 << 9
	time.Sleep(sleep)
	b := time.Until(a)
	fmt.Println(b)
}
