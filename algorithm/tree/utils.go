package tree

//获得一个树的高度
func Height(root *Node) int {
	if root == nil {
		return 0
	} else {
		var leftHeight int = Height(root.Left)
		var rightHeight int = Height(root.Right)
		if leftHeight >= rightHeight {
			return leftHeight + 1
		}
		return rightHeight + 1
	}
}

//操作树的某一层
func actionByLevel(root *Node, level int, f NodeActionFunc) {
	if level <= 0 || root == nil {
		return
	}
	if level == 1 {
		f(root)
	} else {
		actionByLevel(root.Left, level-1, f)
		actionByLevel(root.Right, level-1, f)
	}
}

func actionByLevelSpiral(root *Node, level int, toRight bool, f NodeActionFunc) {
	if level <= 0 || root == nil {
		return
	}
	if level == 1 {
		f(root)
	} else {
		if toRight {
			actionByLevelSpiral(root.Left, level-1, toRight, f)
			actionByLevelSpiral(root.Right, level-1, toRight, f)
		} else {
			actionByLevelSpiral(root.Right, level-1, toRight, f)
			actionByLevelSpiral(root.Left, level-1, toRight, f)
		}
	}
}