package tree

//https://www.geeksforgeeks.org/print-postorder-from-given-inorder-and-preorder-traversals/
func PostFromInAndPre(in, pre []int) {
	if len(in) < 1 || len(pre) < 1 || len(in) != len(pre) {
		return
	}
	root := pre[0]
	rootIndex := getIndex(in, root)
	if rootIndex != 0 {
		PostFromInAndPre(in[:rootIndex], pre[1:rootIndex+1])
	}
	if rootIndex != len(in)-1 {
		PostFromInAndPre(in[rootIndex+1:], pre[rootIndex+1:])
	}
	NodePrint(NewNode(root))
}

func getIndex(array []int, n int) int {
	for index, value := range array {
		if value == n {
			return index
		}
	}
	return -1
}

//another version solution of above problem
func PostFromInAndPre2(in, pre []int) {
	if len(in) < 1 || len(pre) < 1 || len(in) != len(pre) {
		return
	}
	count := len(in)
	index := 0
	root := NewEmptyBTree(count)
	var assignValueFunc NodeActionFunc = func(n *Node) {
		n.Data = in[index]
		index += 1
	}
	InOrder(root, assignValueFunc)
	PostOrder(root, NodePrint)
}
