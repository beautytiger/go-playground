package main

import (
	"encoding/json"
	"fmt"
)

// Marshal function is good for producing JSON that could be returned in a network response, like a Restfult API. The function Marshal returns two values: the encoded JSON data as slice byte and an error value. Using Marshal , we can also encode the values of struct type as JSON values that will help us to quickly build JSON-based APIs.
// http://www.golangprograms.com/advance-programs/golang-program-to-demonstrates-how-to-encode-map-data-into-a-json-string.html
func main() {
	empmap := make(map[string]interface{})
	empmap["testkey"] = "testvalue"
	empmap["info"] = map[string]interface{}{
		"address": "1234-5647",
		"money":   123456,
	}
	empmap["mail"] = "konmyn@163.com"

	mapjson, err := json.Marshal(empmap)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(mapjson)
	jsonstr := string(mapjson)
	fmt.Println(jsonstr)
}
