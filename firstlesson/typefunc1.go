// https://www.jianshu.com/p/fc4902159cf5
package main

import "fmt"

// Greeting function types
type Greeting func(name string) string

func say(g Greeting, n string) {
	fmt.Println(g(n))
}

func english(name string) string {
	return "Hello, " + name
}

func main() {
	say(english, "World")
}
