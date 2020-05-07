package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := map[string]interface{}{}
	json.Unmarshal([]byte(`{"foo": null}`), &m)

	// json/yaml null mapped to nilobj
	fmt.Println(m["foo"] == nil)

	// switch with nilobj
	switch m["foo"].(type) {
	case interface{}:
		fmt.Println("1: interface{}")
	case nil:
		fmt.Println("1: nilobj")
	default:
		fmt.Println("1: default")
	}

	// switch with nilobj removed (as in PR)
	switch m["foo"].(type) {
	case interface{}:
		fmt.Println("2: interface{}")
	default:
		fmt.Println("2: default")
	}
}
