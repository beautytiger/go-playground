package main

import "fmt"

type animal struct {
	name string
	weight int
	height int
	alive bool
}

type cat struct {
	animal
	legs int
	tail int
	sound string
}

func main() {
	a := new(cat)
	a.legs = 4
	a.tail = 1
	a.sound = "mou"
	a.name = "miao"
	a.weight = 40
	a.height = 20
	a.alive = true
	fmt.Println(a)
	fmt.Println(*a)
	fmt.Println(a.animal)
}

