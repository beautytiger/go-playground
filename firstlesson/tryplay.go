package main

import "fmt"

func main () {
	const (
		car int = 1 << iota
		bike
		plane
		train
	)
	fmt.Println(car, bike, plane, train)
}