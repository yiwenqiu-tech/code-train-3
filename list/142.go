package list

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var slow = head
	var fast = head
	var meet *ListNode
	for slow.Next != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			meet = slow
			break
		}
	}
	if meet == nil {
		return nil
	}
	fast = head // 将快节点从头节点开始
	for fast != slow {
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}
