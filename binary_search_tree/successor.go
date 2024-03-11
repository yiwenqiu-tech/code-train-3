package binary_search_tree

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	/*	// 当右子树存在，则从右子树的左子树去找
		if p.Right != nil {
			res := p.Right
			for res.Left != nil {
				res = res.Left
			}
			return res
		}

		// 右子树不存在，则表示再root->p的路径上，查找一个比p大的值中最小的值，即为后继者
		var res *TreeNode
		var cur = root
		for cur != nil {
			if cur.Val == p.Val {
				break
			}
			if cur.Val > p.Val {
				if res == nil || cur.Val < res.Val {
					res = cur
				}
				cur = cur.Left
			} else {
				cur = cur.Right
			}
		}
		return res*/
	result = nil
	return inorderSuccessorForVal(root, p.Val)
}

var result *TreeNode

func inorderSuccessorForVal(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		if root.Right != nil {
			result = root.Right
			for result.Left != nil {
				result = result.Left
			}
			return result
		} else {
			return result // 直接返回result
		}
	}
	if root.Val < val {
		return inorderSuccessorForVal(root.Right, val)
	} else {
		if result == nil {
			result = root
		} else {
			if root.Val < result.Val {
				result = root
			}
		}
		return inorderSuccessorForVal(root.Left, val)
	}
}
