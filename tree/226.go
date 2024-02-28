package tree

func invertTree(root *TreeNode) *TreeNode {
	recurForInvertTree(root)
	return root
}

func recurForInvertTree(node *TreeNode) {
	if node == nil {
		return
	}
	tmp := node.Left
	node.Left = node.Right
	node.Right = tmp               // 交换当前层的
	recurForInvertTree(node.Left)  // 递归交换左节点
	recurForInvertTree(node.Right) // 递归交换右节点
}
