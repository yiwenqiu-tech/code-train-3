package tree

import "container/list"

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	nodeParentMap := recordNodeParent(root)
	var pPathMap = make(map[*TreeNode]struct{})
	pPathMap[p] = struct{}{} // 记录当前节点
	pCur := p
	// 逐层往上找p的父节点，找不到则将pCur置为空，终止循环。此时，将路过的节点放入pPathMap
	for pCur != nil {
		if parent, exist := nodeParentMap[pCur]; exist {
			pPathMap[parent] = struct{}{}
			pCur = parent
		} else {
			pCur = nil
		}
	}
	// 逐层往上找q的父节点，若节点在pPathMap，表示该节点为最近公共祖先，直接返回结果。
	qCur := q
	for qCur != nil {
		if _, exist := pPathMap[qCur]; exist {
			return qCur
		}
		if parent, exist := nodeParentMap[qCur]; exist {
			qCur = parent
		} else {
			qCur = nil
		}
	}
	return nil
}

func recordNodeParent(root *TreeNode) map[*TreeNode]*TreeNode {
	nodeParentMap := make(map[*TreeNode]*TreeNode)

	queue := list.List{}
	queue.PushBack(root)

	for queue.Len() != 0 {
		head := queue.Front()
		queue.Remove(head)
		headVal := head.Value.(*TreeNode)

		if headVal.Left != nil {
			nodeParentMap[headVal.Left] = headVal
			queue.PushBack(headVal.Left)
		}
		if headVal.Right != nil {
			nodeParentMap[headVal.Right] = headVal
			queue.PushBack(headVal.Right)
		}
	}
	return nodeParentMap
}
