package train

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func maxDepth(root *TreeNode) int {
//	return dfsForMaxDepth(root, 0)
//}

//func dfsForMaxDepth(root *TreeNode, depth int) int {
//	if root == nil {
//		return depth
//	}
//	leftMax := dfsForMaxDepth(root.Left, depth+1)
//	rightMax := dfsForMaxDepth(root.Right, depth+1)
//	return max(leftMax, rightMax)
//}
//
//func dfsForMaxDepth2(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	leftMax := dfsForMaxDepth2(root.Left)
//	rightMax := dfsForMaxDepth2(root.Right)
//	return max(leftMax, rightMax) + 1
//}
