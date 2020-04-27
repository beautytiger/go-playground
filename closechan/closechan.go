package main

import (
	"fmt"
)

func main() {
	stopChan := make(chan struct{})
	fmt.Println("starting")
	//close(stopChan)
	var data struct{}
	stopChan <- data
	select{
	case value, open :=<-stopChan:
		fmt.Printf("got stopped chan %v %v\n", value, open)
	}
	fmt.Println("run out of code")
}
