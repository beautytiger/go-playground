package main

import (
	"fmt"
	"regexp"
)

var kubeVersionRegex = regexp.MustCompile("^v([\\d]+)(?:(alpha|beta)([\\d]+))?$")

func main(){
	//pattern := regexp.MustCompile("(?:my)")
	//str := "myname.youname"
	//result := pattern.FindAllString(str, 0)
	//fmt.Printf("%#s", result)
	//looptest()
	v := "v2"
	fmt.Println(kubeVersionRegex.FindStringSubmatch(v))
	fmt.Println(int(""))
}

func looptest() {
	lastPort := 10000
	for offset := 1; offset < 100; offset++ {
		port := lastPort + offset
	//for port := lastPort + 1; port < 100; port++ {
		fmt.Println(port)
	}
}
