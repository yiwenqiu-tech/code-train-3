package list

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head = &ListNode{} // 保护节点
	var cur = head
	var list1Cur = list1
	var list2Cur = list2
	for list1Cur != nil && list2Cur != nil {
		if list1Cur.Val.(int) < list2Cur.Val.(int) {
			cur.Next = list1Cur
			list1Cur = list1Cur.Next
			cur = cur.Next
		} else {
			cur.Next = list2Cur
			list2Cur = list2Cur.Next
			cur = cur.Next
		}
	}
	if list1Cur == nil {
		for list2Cur != nil {
			cur.Next = list2Cur
			list2Cur = list2Cur.Next
			cur = cur.Next
		}
	}

	if list2Cur == nil {
		for list1Cur != nil {
			cur.Next = list1Cur
			list1Cur = list1Cur.Next
			cur = cur.Next
		}
	}
	return head.Next
}
