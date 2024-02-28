package tree

func minDepth(root *TreeNode) int {
	return recurForMinDepth(root)
}

func recurForMinDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	if node.Right == nil && node.Left == nil {
		return 1
	}
	if node.Right == nil {
		return recurForMinDepth(node.Left) + 1
	}
	if node.Left == nil {
		return recurForMinDepth(node.Right) + 1
	}
	left := recurForMinDepth(node.Left)
	right := recurForMinDepth(node.Right)
	if left > right {
		return right + 1
	} else {
		return left + 1
	}
}
