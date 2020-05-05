package tree

import (
	"fmt"
	"github.com/beautytiger/go-playground/algorithm/queue"
)

//tree node
type Node struct {
	Data int
	Left *Node
	Right *Node
}

//二叉搜索树添加子节点
func Add(n *Node, d int) *Node {
	if n == nil {
		n = NewNode(d)
		return n
	}
	if d < n.Data {
		n.Left = Add(n.Left, d)
	} else {
		n.Right = Add(n.Right ,d)
	}
	return n
}

//从队列构建一个二叉查找树，这棵树很可能是很不平衡的
func NewBSTreeFromArray(array []int) *Node {
	var root *Node
	for i := range array {
		root = Add(root, array[i])
	}
	return root
}

//act function for traversal result
type NodeActionFunc func(*Node)

func NodePrint(n *Node) {
	if n == nil {
		return
	}
	fmt.Printf("%d, ", n.Data)
}

//构建一个新的树节点
func NewNode(n int) *Node {
	return &Node{Data: n}
}

//https://www.geeksforgeeks.org/find-postorder-traversal-of-bst-from-preorder-traversal/
// 构建一个空的二叉树(BT)，节点数为n
func NewEmptyBTree(n int) *Node {
	if n <= 0 {
		return nil
	}
	root := NewNode(0)
	q := queue.NewGeneric()
	q.Push(root)
	count := 1
	for count < n {
		v, ok := q.Pop().(*Node)
		if ok == false {
			panic("type error")
		}
		v.Left = NewNode(0)
		q.Push(v.Left)
		count = count + 1
		if count < n {
			v.Right = NewNode(0)
			q.Push(v.Right)
		}
		count = count + 1
	}
	return root
}

//从队列构建一颗二叉树，逐行添加子节点
func NewBTreeFromArray(array []int) *Node {
	var root *Node
	if len(array) < 1 {
		return root
	}
	root = NewNode(array[0])
	qi := queue.NewQueue()
	for _, value := range array[1:] {
		qi.Push(value)
	}
	q := queue.NewGeneric()
	q.Push(root)
	for {
		v, ok := q.Pop().(*Node)
		if !ok {
			panic("type error")
		}
		if data := qi.Pop(); data >0 {
			v.Left = NewNode(data)
			q.Push(v.Left)
		} else {
			break // array used out
		}
		if data := qi.Pop(); data >0 {
			v.Right = NewNode(data)
			q.Push(v.Right)
		} else {
			break // array used out
		}
	}
	return root
}


//Insertion in a Binary Tree in level order
func Insert(root *Node, data int) {
	q := make([]*Node, 0)
	q = append(q, root)
	for len(q)>0 {
		op := q[0]
		q = q[1:]
		if op.Left == nil {
			op.Left = NewNode(data)
			break
		} else {
			q = append(q, op.Left)
		}
		if op.Right == nil {
			op.Right = NewNode(data)
			break
		} else {
			q = append(q, op.Right)
		}
	}
}

//inorder traversal
//树的中序遍历打印
func TreePrint(root *Node) {
	if root == nil {
		return
	}
	TreePrint(root.Left)
	fmt.Println(root.Data)
	TreePrint(root.Right)
}

//删除并修理树
func TreeDelete(root *Node, data int) *Node {
	if root == nil {
		 return nil
	}
	if root.Left == nil && root.Right == nil {
		if root.Data == data {
			return nil
		}
		return root
	}
	var lastNode *Node
	var deleteNode *Node
	q := make([]*Node, 0)
	q = append(q, root)
	for len(q) > 0 {
		current := q[0]
		lastNode = current
		q = q[1:]
		if current.Data == data {
			deleteNode = current
		}
		if current.Left != nil {
			q = append(q, current.Left)
		}
		if current.Right != nil {
			q = append(q, current.Right)
		}
	}
	if deleteNode == nil {
		return root
	}
	deleteNode.Data = lastNode.Data
	root = treeDeleteLastNode(root, lastNode)
	return root
}

//delete last from tree root
func treeDeleteLastNode(root, last *Node) *Node {
	q := make([]*Node, 0)
	q = append(q, root)
	for len(q) > 0 {
		temp := q[0]
		q = q[1:]
		if temp.Right != nil {
			if temp.Right == last {
				temp.Right = nil
				return root
			}
			q = append(q, temp.Right)
		}
		if temp.Left != nil {
			if temp.Left == last {
				temp.Left = nil
				return root
			}
			q = append(q, temp.Left)
		}
	}
	return root
}
