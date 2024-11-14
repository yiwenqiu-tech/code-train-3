package train

var sum113 int
var targetSum113 int
var res [][]int

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	res = nil
	targetSum113 = targetSum
	dfsForHasPathSum113(root, nil)
	return res
}

func dfsForHasPathSum113(root *TreeNode, curRes []int) {
	sum113 += root.Val
	curRes = append(curRes, root.Val)
	if root.Left == nil && root.Right == nil { // 到叶子节点
		if sum113 == targetSum113 { // 和跟target一致，添加到结果里
			var tmp []int
			tmp = append(tmp, curRes...)
			res = append(res, tmp)
		}
		// 还原现场
		curRes = curRes[:len(curRes)-1]
		sum113 -= root.Val
		return
	}
	if root.Left != nil {
		dfsForHasPathSum113(root.Left, curRes)
	}

	if root.Right != nil {
		dfsForHasPathSum113(root.Right, curRes)
	}
	// 还原现场
	curRes = curRes[:len(curRes)-1]
	sum113 -= root.Val
	return
}
