package train

func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head
	var merge *ListNode
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			merge = slow
			break
		}
	}
	if merge == nil {
		return nil
	}
	fast = head
	slow = merge
	for slow != fast {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
