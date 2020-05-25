package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(100)
	fmt.Printf("%v at %p", i, &i)
	for {
	}
}
