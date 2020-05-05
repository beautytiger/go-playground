package main

import (
	"fmt"
	"github.com/beautytiger/go-playground/algorithm/tree"
)


//get random tree :)
func giveMeTree() *tree.Node {
	//root := tree.NewNode(1)
	//root.Left = tree.NewNode(2)
	//root.Right = tree.NewNode(3)
	//root.Left.Left = tree.NewNode(4)
	//root.Left.Right = tree.NewNode(5)
	//root.Right.Left = tree.NewNode(6)
	//root.Right.Right = tree.NewNode(7)
	//root.Left.Left.Left = tree.NewNode(8)
	//root.Left.Left.Right = tree.NewNode(9)
	//root.Left.Right.Left = tree.NewNode(10)
	//root.Left.Right.Right = tree.NewNode(11)
	//root.Right.Left.Left = tree.NewNode(12)
	//root.Right.Left.Right = tree.NewNode(13)
	//root.Right.Right.Left = tree.NewNode(14)
	//root.Right.Right.Right = tree.NewNode(15)
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	//root := tree.NewBSTreeFromArray(data)
	root := tree.NewBTreeFromArray(data)
	return root
}

func runTree() {
	root := giveMeTree()
	f := tree.NodePrint
	fmt.Println("level order")
	tree.LevelOrder(root, f)
	fmt.Println()
	tree.LevelOrderRecursion(root, f)
	fmt.Println()
	tree.LevelOrderSpiralRecursion(root, f)
	fmt.Println()
	fmt.Println("level spiral non recursion")
	tree.LevelOrderSpiral(root, f)
	fmt.Println()
	tree.InOrder(root, f)
	fmt.Println()
	tree.PreOrder(root, f)
	fmt.Println()
	fmt.Println("post order:")
	tree.PostOrder(root, f)
	fmt.Println()
	fmt.Println("in order:")
	tree.InOrderNoRecursion(root, f)
	fmt.Println()
	in := []int{8, 4, 9, 2, 10, 5, 11, 1, 12, 6, 13, 3, 14, 7, 15}
	pre := []int{1, 2, 4, 8, 9, 5, 10, 11, 3, 6, 12, 13, 7, 14, 15}
	tree.PostFromInAndPre(in, pre)
	fmt.Println()
	tree.PostFromInAndPre2(in, pre)
	// BST
	fmt.Println()
	fmt.Println("BT")
	tree.LevelOrder(tree.NewEmptyBTree(8), f)
}

func runArrayTree() {
	t := tree.ArrayTree{}
	t.Add(999)
	t.SetLeft(2, 0)
	t.SetRight(3, 0)
	t.SetRight(777, 2)
	fmt.Println(*t.Item)
	err := t.SetLeft(9, 3)
	fmt.Println(err)
	t.Print()
}

func main() {
	fmt.Println("start")
	runTree()
	//runArrayTree()
}
