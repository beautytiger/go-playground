package linkedlist

import (
	"fmt"
	st "github.com/beautytiger/go-playground/algorithm/stack"
)

type Node struct {
	Data int
	Next *Node
}

type SinglyLinkedList interface {
	//返回链表长度
	Len() int
	//链表头部追加元素
	Push(int)
	//链表头部弹出元素
	Pop(int)
	//清空链表
	Clear()
}

type Singly struct {
	Head *Node
	length int
}

func NewNode(d int) *Node {
	return &Node{Data: d}
}

func NewSinglyFromArray(array []int) *Node {
	var head *Node
	if len(array) == 0 {
		return head
	}
	head = NewNode(array[0])
	current := head
	for _, data := range array[1:] {
		current.Next = NewNode(data)
		current = current.Next
	}
	return head
}

func SinglyPrint(head *Node) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

//翻转单向链表
func Reverse(head *Node) *Node {
	if head == nil {
		return head
	}
	current := head
	var prev *Node = nil
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func ReverseRecursive(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	rest := ReverseRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return rest
}

func tailRecursive(curr, prev *Node) *Node {
	if curr.Next == nil {
		curr.Next = prev
		return curr
	}
	next := curr.Next
	curr.Next = prev
	return tailRecursive(next, curr)
}

//尾递归，原地翻转
func ReverseTailRecursive(head **Node) {
	if *head == nil {
		return
	}
	*head = tailRecursive(*head, nil)
	return
}

//按组翻转链表
//https://www.geeksforgeeks.org/reverse-a-list-in-groups-of-given-size/
func GroupReverse(head *Node, k int) *Node {
	if head == nil || k <= 0 {
		return head
	}
	var next *Node
	var current *Node = head
	var prev *Node
	var count = 0
	for current != nil && count < k {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
		count++
	}

	head.Next = GroupReverse(next, k)
	return prev
}

//非递归方式
func GroupReverseByStack(head *Node, k int) *Node {
	if head == nil || k <= 0 {
		return head
	}
	stack := st.NewGeneric()
	var current *Node = head
	var prev *Node
	var count int
	for current != nil {
		count = 0
		for current != nil && count < k {
			stack.Push(current)
			current = current.Next
			count++
		}
		for stack.Len() > 0 {
			if prev == nil {
				v, ok := stack.Pop().(*Node)
				if !ok {
					panic("type error")
				}
				prev = v
				head = prev
			} else {
				v, ok := stack.Pop().(*Node)
				if !ok {
					panic("type error")
				}
				prev.Next = v
				prev = prev.Next
			}
		}
	}
	prev.Next = nil
	return head
}