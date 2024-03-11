package binary_search_tree

var val int

func convertBST(root *TreeNode) *TreeNode {
	val = 0
	return recurForConvertBST(root)
}

func recurForConvertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	recurForConvertBST(root.Right)

	val = val + root.Val
	root.Val = val

	recurForConvertBST(root.Left)

	return root
}
