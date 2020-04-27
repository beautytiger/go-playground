package main

import (
	"regexp"
	"fmt"
)

var kubeReleaseRegex = regexp.MustCompile(`^v?(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)([-0-9a-zA-Z_\.+]*)?$`)

func main() {
	var a string = "1.14.0bades"
	fmt.Println(kubeReleaseRegex.MatchString(a))
	re := regexp.MustCompile(`a(x*)b(y|z)c`)
	fmt.Printf("%q\n", re.FindStringSubmatch("-axxxbyc-"))
	fmt.Printf("%q\n", re.FindStringSubmatch("-abzc-"))
	reg := regexp.MustCompile(`a(x*)b`)
	fmt.Printf("%q\n", reg.FindAllStringSubmatch("-ab-", -1))
	fmt.Printf("%q\n", reg.FindAllStringSubmatch("-axxb-", 9))
	fmt.Printf("%q\n", reg.FindAllStringSubmatch("-ab-axb-", -1))
	fmt.Printf("%q\n", reg.FindAllStringSubmatch("-axxb-ab-", -1))
}
