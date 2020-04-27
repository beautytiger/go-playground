package main

import (
	"fmt"
	"time"
)

func main(){
	t1, _ := time.Parse("2006-01-02 15:04:05", "2019-03-21 08:01:01")
	t2, _ := time.Parse("2006-01-02 15:04:05", "2019-04-01 10:11:53")
	fmt.Println(t1)
	fmt.Println(t2)
}
