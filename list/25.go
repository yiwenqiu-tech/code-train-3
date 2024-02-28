package list

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 0 || k == 1 { // k为0或1，直接返回原链表
		return head
	}
	curLoopCount := 0              // 当前遍历的数量
	var curHead, curTail *ListNode // 记录当前组的头和尾
	var lastTail *ListNode         // 记录上一组的尾节点
	var resHead = head             // 结果的头节点
	cur := head                    // 当前遍历的节点
	for cur != nil {               // 当前节点非空
		if curLoopCount == 0 { // 如果新的一轮，则将cur标记为当前组的头节点
			curHead = cur
		}
		curLoopCount++         // 累加loopCount
		if curLoopCount == k { // loopCount到k了
			curTail = cur  // 将当前节点记录为当前组的尾节点
			cur = cur.Next // 先将cur.Next赋值给cur，避免交换后Next丢失了
			newHead, newTail := reverse(curHead, curTail)
			if lastTail != nil {
				lastTail.Next = newHead
			} else {
				resHead = newHead
			}
			lastTail = newTail
			curLoopCount = 0
			curHead = nil
			curTail = nil
		} else { // loopCount没到k，则直接赋值cur
			cur = cur.Next
		}
	}
	if curHead != nil { // 剩下的不够，则将尾节点连向curHead
		lastTail.Next = curHead
	}
	return resHead
}

func reverse(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	var cur = head
	var tailNext = tail.Next // 注意：需要保留tailNext，用于判断是否到了尾节点了
	var last *ListNode
	for cur != tailNext {
		temp := cur.Next
		cur.Next = last
		last = cur
		cur = temp
	}
	return tail, head
}
