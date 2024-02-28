package list

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	var slow = head
	var fast = head
	for slow.Next != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}
