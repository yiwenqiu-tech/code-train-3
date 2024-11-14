package train

var sum112 int
var targetSum112 int

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum112 = targetSum
	return dfsForHasPathSum(root)
}

func dfsForHasPathSum(root *TreeNode) bool {
	sum112 += root.Val
	defer func() {
		sum112 -= root.Val // 还原现场
	}()
	if root.Left == nil && root.Right == nil { // 到叶子节点
		if sum112 == targetSum112 { // 和跟target一致，返回true
			return true
		}
		return false
	}
	if root.Left != nil {
		leftRes := dfsForHasPathSum(root.Left)
		if leftRes {
			return true
		}
	}

	if root.Right != nil {
		rightRes := dfsForHasPathSum(root.Right)
		if rightRes {
			return true
		}
	}
	return false
}
