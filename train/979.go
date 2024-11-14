package train

var sum int

func distributeCoins(root *TreeNode) int {
	sum = 0
	dfsFor979(root)
	return sum
}

func dfsFor979(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}
	coinL, nodeL := dfsFor979(node.Left)
	coinR, nodeR := dfsFor979(node.Right)
	coinCur := coinL + coinR + node.Val
	nodeCur := nodeL + nodeR + 1
	sum += abs(coinCur - nodeCur)
	return coinCur, nodeCur
}

func abs(value int) int {
	if value < 0 {
		return -value
	} else {
		return value
	}
}
