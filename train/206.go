package train

func reverseList(head *ListNode) *ListNode {
	var last *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = last
		last = cur
		cur = next
	}
	return last
}
