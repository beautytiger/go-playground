package main

import (
	"fmt"
)

func main() {
	//var a interface{} = nilobj
	m := map[string]interface{}{"m": nil}
	switch at := m["m"].(type) {
	case interface{}:
		fmt.Println("interface")
	case nil:
		fmt.Println("nilobj")
	default:
		fmt.Printf("invalid type encountered %T", at)
	}
}
