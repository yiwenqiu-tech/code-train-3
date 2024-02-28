package list

func reverseList(head *ListNode) *ListNode {
	var last *ListNode // 存上一个节点
	var cur = head
	for cur != nil {
		next := cur.Next // 缓存下一个节点
		cur.Next = last
		last = cur
		cur = next
	}
	return last
}
