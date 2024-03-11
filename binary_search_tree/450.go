package binary_search_tree

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Right == nil && root.Left == nil { // 左右子树都为空，直接删除即可
			return nil
		} else if root.Right == nil { // 右子树为空时，返回其左子树
			return root.Left
		} else if root.Left == nil { // 左子树为空时，返回其右子树
			return root.Right
		} else {
			res := root.Right
			parent := root
			for res.Left != nil {
				parent = res
				res = res.Left
			}
			if parent == root { // 右节点没有往左的节点，则删除右节点。注意，删除右节点，还要分情况，若右节点存在右节点，则将其右节点往上提，否则直接置空
				if parent.Right.Right != nil {
					parent.Right = parent.Right.Right
				} else {
					parent.Right = nil // 删除后继
				}
			} else { // 右节点存在往左的节点，则删除其左节点。注意，删除左节点，还要分情况，若左节点存在右节点，则将其右节点往上提，否则直接置空
				if parent.Left.Right != nil {
					parent.Left = parent.Left.Right
				} else {
					parent.Left = nil // 删除后继
				}
			}
			root.Val = res.Val // 更新root Val
			return root
		}
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		root.Left = deleteNode(root.Left, key)
	}
	return root
}
