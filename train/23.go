package train

import "container/heap"

type listNodeHeap []*ListNode

func (l *listNodeHeap) Len() int {
	return len(*l)
}

func (l *listNodeHeap) Less(i, j int) bool {
	return (*l)[i].Val < (*l)[j].Val
}

func (l *listNodeHeap) Swap(i, j int) {
	(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
}

func (l *listNodeHeap) Push(x any) {
	*l = append(*l, x.(*ListNode))
}

func (l *listNodeHeap) Pop() any {
	item := (*l)[len(*l)-1]
	*l = (*l)[0 : len(*l)-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
	var head = &ListNode{}
	nodeHeap := listNodeHeap{}
	heap.Init(&nodeHeap)
	for _, l := range lists {
		if l == nil {
			continue
		}
		heap.Push(&nodeHeap, l)
	}
	cur := head
	for nodeHeap.Len() > 0 {
		node := heap.Pop(&nodeHeap).(*ListNode)
		if node.Next != nil {
			heap.Push(&nodeHeap, node.Next)
		}
		cur.Next = node
		cur = node
	}
	return head.Next
}

func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	half1 := mergeKLists(lists[0 : len(lists)/2])
	half2 := mergeKLists(lists[len(lists)/2:])
	return merge2List(half1, half2)
}

func merge2List(list1 *ListNode, list2 *ListNode) *ListNode {
	var head = &ListNode{}
	cur := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
			cur = cur.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
			cur = cur.Next
		}
	}
	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}
	return head.Next
}
