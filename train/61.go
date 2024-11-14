package train

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if k == 0 {
		return head
	}
	listLen := 1
	cur := head
	for cur.Next != nil {
		listLen++
		cur = cur.Next
	}
	if k%listLen == 0 { // k与list模为0，则不需要移动
		return head
	}

	cur.Next = head // 首尾相连成环

	// 找到最后一个node，断开与下一个node的连接
	index := 0
	cur = head
	lastNodeIndex := listLen - 1 // 原来最后一个节点
	lastNodeIndex = lastNodeIndex - k%listLen
	for index < lastNodeIndex {
		cur = cur.Next
		index++
	}
	ans := cur.Next
	cur.Next = nil

	return ans
}
