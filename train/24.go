package train

func swapPairs(head *ListNode) *ListNode {
	tmpHead := &ListNode{0, head}
	cur := tmpHead.Next
	pre := tmpHead // 记录上一组的节点，用来关联每一组
	for cur != nil && cur.Next != nil {
		tmp := cur.Next.Next
		cur.Next.Next = cur

		pre.Next = cur.Next
		cur.Next = tmp

		pre = cur
		cur = tmp
	}
	return tmpHead.Next
}
