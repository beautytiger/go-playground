package main

import "fmt"

func main() {
	var a string = "nitin"
	//fmt.Println(isPalindromic(a, 0, len(a)-6))
	allPalPartitions(a)
}

func isPalindromic(s string, start, end int) bool {
	for start <= end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func allPalPartitions(s string) {
	var r = []string{}
	for i:=0; i<len(s); i++ {
		for j:=i; j<len(s);j++{
			if isPalindromic(s, i, j) {
				r = append(r, s[i:j+1])
			}
		}
	}
	for _, c := range r {
		fmt.Printf("%s ", c)
	}
}
