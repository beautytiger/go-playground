package main

import "fmt"
import "reflect"

type Author struct {
	name string
	branch string
	language string
	Particle int
}

type People struct {
	name string
	weight int
	gender string
	age int
}

func main() {
	a1 := Author{
		name : "Bob",
		branch: "EN",
		language: "ZH",
		Particle: 40,
	}
	a2 := Author{
		name : "Bob",
		branch: "EN",
		language: "ZH",
		Particle: 40,
	}
	a3 := Author{"Tom", "Head", "EN", 33}

	p1 := People{"Tom", 90, "M", 30}
	p2 := People{
		name: "tom",
		weight: 95,
		gender: "M",
		age: 44,
	}

	if a1 == a2 {
		fmt.Println("a1 is equal to a2")
	} else {
		fmt.Println("a1 is not equal to a2")
	}

	if a1 == a3 {
		fmt.Println("a1 is equal to a3")
	} else {
		fmt.Println("a1 is not equal to a3")
	}

	fmt.Println("a1 is equal to a2:", reflect.DeepEqual(a1, a2))
	fmt.Println("a2 is equal to a3:", reflect.DeepEqual(a2, a3))
	fmt.Println("p1 is equal to p2:", reflect.DeepEqual(p1, p2))
	fmt.Println("a1 is equal to p2:", reflect.DeepEqual(a1, p2))
	// below is not correct since different types can not compared
	//fmt.Println("a1 is equal to p1:", a1 == p1)
}