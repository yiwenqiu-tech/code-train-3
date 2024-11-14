package train

func trainingPlan(head *ListNode, cnt int) *ListNode {
	// 方法1
	/*	length := 0
		cur := head
		for cur != nil {
			length++
			cur = cur.Next
		}
		cur = head
		loop := length - cnt
		for loop > 0 {
			cur = cur.Next
			loop--
		}
		return cur*/

	// 双指针
	slower := head
	faster := head
	for cnt > 0 {
		faster = faster.Next
		cnt--
	}
	for faster != nil {
		slower = slower.Next
		faster = faster.Next
	}
	return slower
}
