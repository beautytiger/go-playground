package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	path := "/file/not/exist"
	openFile(path)
	fmt.Println("returned in main func")
}

func openFile(path string) {
	f, err := os.OpenFile(path, syscall.O_RDONLY|syscall.O_DIRECT, 0)
	fmt.Println("returned file", f)
	rerr := f.Close()
	fmt.Println(rerr)
	if err != nil {
		fmt.Println("false")
		return
	}
	fmt.Println("true")
}
