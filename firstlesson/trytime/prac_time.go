package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Location: ", t.Location(), " Time: ", t)
	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Location: ", location, " Time: ", t.In(location))

	loc, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(loc)
	fmt.Println("Location: ", loc, " Time: ", now)
}
