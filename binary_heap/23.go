package binary_heap

import (
	"code-train-3/selfcontainer"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type listNodes []*ListNode

func (l *listNodes) Len() int {
	return len(*l)
}

func (l *listNodes) Less(i, j int) bool {
	return (*l)[i].Val < (*l)[j].Val
}
func (l *listNodes) Swap(i, j int) {
	(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
}

func (l *listNodes) Push(x any) {
	*l = append(*l, x.(*ListNode))
}
func (l *listNodes) Pop() any {
	last := (*l)[l.Len()-1]
	*l = (*l)[0 : (*l).Len()-1]
	return last
}

func mergeKLists(lists []*ListNode) *ListNode {
	var ignoreNilList []*ListNode
	for _, l := range lists {
		if l == nil {
			continue
		}
		ignoreNilList = append(ignoreNilList, l)
	} // 需要过滤空list后，再建堆

	//var lsHeap listNodes
	//for _, l := range lists {
	//	if l != nil {
	//		heap.Push(&lsHeap, l)
	//	}
	//}

	//lsHeap := listNodes(ignoreNilList)
	//heap.Init(&lsHeap)
	//
	//var head *ListNode
	//var cur *ListNode
	//for lsHeap.Len() > 0 {
	//	ls := heap.Pop(&lsHeap).(*ListNode)
	//	if head == nil {
	//		head = ls
	//		cur = head
	//	} else {
	//		cur.Next = ls
	//		cur = cur.Next
	//	}
	//	if ls.Next != nil {
	//		heap.Push(&lsHeap, ls.Next)
	//	}
	//}
	//return head

	ln := listNodes(ignoreNilList)
	selfHeap := selfcontainer.Heap{Interface: &ln}
	selfHeap.Init()

	var head *ListNode
	var cur *ListNode
	for selfHeap.Len() > 0 {
		ls := selfHeap.PopTop().(*ListNode)
		if head == nil {
			head = ls
			cur = head
		} else {
			cur.Next = ls
			cur = cur.Next
		}
		if ls.Next != nil {
			selfHeap.Insert(ls.Next)
		}
	}
	return head
}
