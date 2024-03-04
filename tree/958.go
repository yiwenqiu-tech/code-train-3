package tree

import "container/list"

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return false
	}
	queue := list.List{}
	queue.PushFront(root)

	var shouldFinish bool
	for queue.Len() != 0 {
		first := queue.Back()
		queue.Remove(first)
		node := first.Value.(*TreeNode)
		if shouldFinish && (node.Left != nil || node.Right != nil) {
			return false
		}
		if node.Left == nil && node.Right != nil {
			return false
		}
		if node.Left != nil && node.Right != nil {
			queue.PushFront(node.Left)
			queue.PushFront(node.Right)
		}
		if node.Left != nil && node.Right == nil {
			shouldFinish = true
			queue.PushFront(node.Left)
		}
		if node.Left == nil && node.Right == nil {
			shouldFinish = true
		}
	}
	return true
}
