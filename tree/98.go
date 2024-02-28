package tree

func isValidBST(root *TreeNode) bool {
	_, _, valid := recurForIsValidBSTRight(root)
	return valid
}

type integer struct {
	val int
}

func recurForIsValidBSTRight(node *TreeNode) (*integer, *integer, bool) {
	if node == nil {
		return nil, nil, true
	}
	minForLeft, maxForLeft, isValidForLeft := recurForIsValidBSTRight(node.Left)
	minForRight, maxForRight, isValidForRight := recurForIsValidBSTRight(node.Right)
	
	if !isValidForLeft || !isValidForRight { // 如果左右子树存在一个不符合要求，则整颗树也无法满足
		return nil, nil, false
	}
	if maxForLeft != nil && maxForLeft.val >= node.Val { // 如果左子树的最大值大于等于当前节点，则不满足定义，直接返回
		return nil, nil, false
	}
	if minForRight != nil && minForRight.val <= node.Val { // 如果右子树的最小值小于等于当前节点，则不满足定义，直接返回
		return nil, nil, false
	}
	var max integer
	if maxForRight == nil { // 右子树最大值不存在时，则max为node的值
		max.val = node.Val
	} else { // 否则为右子树的最大值
		max = *maxForRight
	}

	var min integer
	if minForLeft == nil { // 左子树最小值不存在时，则min为node的值
		min.val = node.Val
	} else { // 否则为左子树的最小值
		min = *minForLeft
	}
	return &min, &max, true
}

func recurForIsValidBST(node *TreeNode) bool { // 这种做法只能保证当前层的有效，却不能保证跨层, 错误例子：{5,4,6,null,null,3,7}
	if node == nil {
		return true
	}
	if (node.Left != nil && node.Left.Val >= node.Val) ||
		(node.Right != nil && node.Right.Val <= node.Val) {
		return false
	}
	return recurForIsValidBST(node.Left) && recurForIsValidBST(node.Right)
}
