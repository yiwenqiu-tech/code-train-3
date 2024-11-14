package train

// TODO REVIEW
func sortList(head *ListNode) *ListNode {
	return sortL(head, nil)
}

func sortL(head *ListNode, tail *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == tail {
		head.Next = nil //PS
		return head
	}
	slow := head
	fast := head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	list1 := sortL(head, mid)
	list2 := sortL(mid, tail)
	return merge(list1, list2)
}

func merge(head1, head2 *ListNode) *ListNode {
	var tmpHead = &ListNode{}
	temp := tmpHead
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			temp.Next = head1
			head1 = head1.Next
		} else {
			temp.Next = head2
			head2 = head2.Next
		}
		temp = temp.Next
	}
	if head1 != nil {
		temp.Next = head1
	}
	if head2 != nil {
		temp.Next = head2
	}
	return tmpHead.Next
}
