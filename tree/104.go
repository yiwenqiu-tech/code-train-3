package tree

func maxDepth(root *TreeNode) int {
	return recurForMaxDepth(root)
}

func recurForMaxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftMax := recurForMaxDepth(node.Left)
	rightMax := recurForMaxDepth(node.Right)
	if leftMax > rightMax {
		return leftMax + 1
	} else {
		return rightMax + 1
	}
}
