package train

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head ListNode
	var p = &head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}
		p = p.Next
	}
	// 当非空时，逐个迭代
	/*for cur != nil {
		p.Next = cur
		cur = cur.Next
		p = p.Next
	}
	for cur2 != nil {
		p.Next = cur2
		cur2 = cur2.Next
		p = p.Next
	}*/

	// 直接放到最后即可
	if list1 != nil {
		p.Next = list1
	}
	if list2 != nil {
		p.Next = list2
	}
	return head.Next
}

// 递归做法
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val < list2.Val {
		list1.Next = mergeTwoLists2(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists2(list1, list2.Next)
		return list2
	}
}
