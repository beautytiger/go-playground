package main

import (
	"fmt"
	"github.com/beautytiger/go-playground/algorithm/linkedlist"
)

func main() {
	data := []int{1,2,3,4,5,6,7,8,9}
	//data := []int{1, 2}
	singly := linkedlist.NewSinglyFromArray(data)
	linkedlist.SinglyPrint(singly)
	singly = linkedlist.Reverse(singly)
	linkedlist.SinglyPrint(singly)
	singly = linkedlist.ReverseRecursive(singly)
	linkedlist.SinglyPrint(singly)
	linkedlist.ReverseTailRecursive(&singly)
	linkedlist.SinglyPrint(singly)
	//fmt.Println("group reverse")
	//singly = linkedlist.GroupReverse(singly, 1)
	//linkedlist.SinglyPrint(singly)
	fmt.Println("group reverse by stack")
	singly = linkedlist.GroupReverseByStack(singly, 7)
	linkedlist.SinglyPrint(singly)
}
