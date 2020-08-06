package main

import (
	"fmt"
	"regexp"
)

func main() {
	var r1 = regexp.MustCompile(`(((\d+)-(\d+))-((\d+)-(\d+)))`)
	var r2 = regexp.MustCompile(`(?:\d+)-(\d+)-(\d+)-(\d+)`)
	var r3 = regexp.MustCompile(`(?P<first>\d+)-(\d+)-(\d+)-(\d+)`)
	var s string = "abc1-2-3-4-5-6-7"
	fmt.Println(r1.FindStringSubmatch(s))
	fmt.Println(r2.FindStringSubmatch(s))
	fmt.Println(r3.FindStringSubmatch(s))
	fmt.Println(r3.SubexpNames()[5])
}