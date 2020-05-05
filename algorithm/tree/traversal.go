package tree

import (
	"github.com/beautytiger/go-playground/algorithm/queue"
	stack "github.com/beautytiger/go-playground/algorithm/stack"
)

//深度优先
//中序遍历
func InOrder(root *Node, f NodeActionFunc) {
	if root == nil {
		return
	}
	InOrder(root.Left, f)
	f(root)
	InOrder(root.Right, f)
}

//非递归的中序遍历
func InOrderNoRecursion(root *Node, f NodeActionFunc) {
	if root == nil {
		return
	}
	st := stack.NewGeneric()
	current := root
	for {
		if current != nil {
			st.Push(current)
			current = current.Left
		} else if st.Len() > 0 {
			v, ok := st.Pop().(*Node)
			if ok == false {
				panic("err type")
			}
			current = v
			f(current)
			current = current.Right
		} else {
			break
		}
	}
}

//前序遍历
func PreOrder(root *Node, f NodeActionFunc) {
	if root == nil {
		return
	}
	f(root)
	PreOrder(root.Left, f)
	PreOrder(root.Right, f)
}

//后序遍历
func PostOrder(root *Node, f NodeActionFunc) {
	if root == nil {
		return
	}
	PostOrder(root.Left, f)
	PostOrder(root.Right, f)
	f(root)
}

//广度优先
//层次遍历
func LevelOrder(root *Node, f NodeActionFunc) {
	if root == nil {
		return
	}
	q := queue.NewGeneric()
	q.Push(root)
	for q.Len() > 0 {
		temp := q.Pop()
		//需要首先进行类型转换
		r, ok := temp.(*Node)
		if ok == false {
			panic("err type")
		}
		if r.Left != nil {
			q.Push(r.Left)
		}
		if r.Right != nil {
			q.Push(r.Right)
		}
		f(r)
	}
}

//层次遍历
//递归版本
func LevelOrderRecursion(root *Node, f NodeActionFunc) {
	h := Height(root)
	for i:=0; i<=h; i++ {
		actionByLevel(root, i+1, f)
	}
}

//z字型层次遍历，递归版
func LevelOrderSpiralRecursion(root *Node, f NodeActionFunc) {
	h := Height(root)
	toRight := true
	for i:=0; i<h; i++ {
		actionByLevelSpiral(root, i+1, toRight, f)
		toRight = !toRight
	}
}

//z字型层次遍历，非递归版
func LevelOrderSpiral(root *Node, f NodeActionFunc) {
	if root == nil {
		return
	}
	//left to right stack
	s1 := stack.NewGeneric()
	//right to left stack
	s2 := stack.NewGeneric()
	s1.Push(root)
	for s1.Len() != 0 || s2.Len() != 0 {
		for s1.Len() != 0 {
			k, v := s1.Pop().(*Node)
			if !v {
				panic("type error")
			}
			f(k)
			if k.Left != nil {
				s2.Push(k.Left)
			}
			if k.Right != nil {
				s2.Push(k.Right)
			}
		}
		for s2.Len() != 0 {
			k, v := s2.Pop().(*Node)
			if !v {
				panic("type error")
			}
			f(k)
			if k.Right != nil {
				s1.Push(k.Right)
			}
			if k.Left != nil {
				s1.Push(k.Left)
			}
		}
	}
}
