package train

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left >= right {
		return head
	}
	var leftLast *ListNode  // 记录左的前一个元素
	var leftNode = head     // 记录左元素
	var rightNode *ListNode // 记录右元素
	var rightNext *ListNode // 记录右元素的下一个元素
	var last *ListNode
	var cur = head
	index := 1
	for cur != nil {
		if index < left { // 当index小于left时，找到left位置
			leftLast = cur
			cur = cur.Next
			leftNode = cur
			index++
			continue
		}
		if index <= right { // 当index小于right时
			if index == right {
				rightNext = cur.Next
				rightNode = cur
			}
			// 开始反转
			next := cur.Next
			cur.Next = last
			last = cur
			cur = next
			index++
		} else {
			break
		}
	}
	// 最后，连接被反转中间的元素
	leftNode.Next = rightNext
	if leftLast != nil { // 注意判断为空
		leftLast.Next = rightNode
	}

	if left == 1 {
		return rightNode
	} else {
		return head
	}
}
