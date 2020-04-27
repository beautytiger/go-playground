package main

import (
	"fmt"
	"regexp"
)

func main() {
	var a string
	var c string
	b := 1
	switch b {
	case 1:
		c, err := getValue()
		fmt.Println(c, err)
		fmt.Println(a)
	case 2:
		a, err := getValue()
		fmt.Println(a, err)
	}
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(KubernetesVersionToImageTag("H@ello%worLd&"))
	regGoFile := regexp.MustCompile(".*\\.go")
	fmt.Println("name.go", regGoFile.MatchString("name.go"))
}

func getValue() (string, error) {
	a := "new value"
	return a, nil
}

func KubernetesVersionToImageTag(version string) string {
	allowed := regexp.MustCompile(`[^-a-zA-Z0-9_\.]`)
	return allowed.ReplaceAllString(version, "_")
}
