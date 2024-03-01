package tree

var inorderTraversalRes []int

func inorderTraversal(root *TreeNode) []int {
	inorderTraversalRes = nil
	recurForInOrderTraversal(root)
	return inorderTraversalRes
}

func recurForInOrderTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	recurForInOrderTraversal(node.Left)
	inorderTraversalRes = append(inorderTraversalRes, node.Val)
	recurForInOrderTraversal(node.Right)
}
