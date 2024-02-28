package divide_conquer

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	return recurForMergeKLists2(lists)
}

func recurForMergeKLists2(lists []*ListNode) *ListNode { // 将数组划分为两份，分治计算每一份后再合并起来
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	listLen := len(lists)
	headHalfRes := recurForMergeKLists2(lists[0 : listLen/2])
	lastHalfRes := recurForMergeKLists2(lists[listLen/2:])
	var res = &ListNode{}
	var resCur = res
	for headHalfRes != nil && lastHalfRes != nil {
		if headHalfRes.Val > lastHalfRes.Val {
			resCur.Next = lastHalfRes
			resCur = resCur.Next
			lastHalfRes = lastHalfRes.Next
		} else {
			resCur.Next = headHalfRes
			resCur = resCur.Next
			headHalfRes = headHalfRes.Next
		}
	}
	for headHalfRes != nil {
		resCur.Next = headHalfRes
		resCur = resCur.Next
		headHalfRes = headHalfRes.Next
	}
	for lastHalfRes != nil {
		resCur.Next = lastHalfRes
		resCur = resCur.Next
		lastHalfRes = lastHalfRes.Next
	}
	return res.Next
}

func recurForMergeKLists(lists []*ListNode) *ListNode { // 每次只合并前两个元素，然后将结果重新放入数组里，参与下一轮合并，直到数组元素只有一个
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	listA := lists[0]
	listB := lists[1]
	var listC = &ListNode{}
	var headC = listC
	for listA != nil && listB != nil {
		if listA.Val > listB.Val {
			listC.Next = &ListNode{
				Val:  listB.Val,
				Next: nil,
			}
			listC = listC.Next
			listB = listB.Next
		} else {
			listC.Next = &ListNode{
				Val:  listA.Val,
				Next: nil,
			}
			listC = listC.Next
			listA = listA.Next
		}
	}
	for listA != nil {
		listC.Next = listA
		listC = listC.Next
		listA = listA.Next
	}
	for listB != nil {
		listC.Next = listB
		listC = listC.Next
		listB = listB.Next
	}
	lists = lists[2:]
	lists = append(lists, headC.Next)
	return recurForMergeKLists(lists)
}
