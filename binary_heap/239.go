package binary_heap

import "container/heap"

type valWithIndex struct {
	val   int
	index int
}

type valSlice []*valWithIndex

func (s *valSlice) Len() int {
	return len(*s)
}

func (s *valSlice) Less(i, j int) bool {
	return (*s)[i].val > (*s)[j].val
}

func (s *valSlice) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *valSlice) Push(x any) {
	*s = append(*s, x.(*valWithIndex))
}

func (s *valSlice) Pop() any {
	item := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return item
}

func maxSlidingWindow(nums []int, k int) []int {
	var valHeap valSlice
	var result []int
	for i := 0; i < len(nums); i++ {
		heap.Push(&valHeap, &valWithIndex{
			val:   nums[i],
			index: i,
		})
		if i >= k-1 { // 如果堆里数量达到k，则pop出来一个答案
			res := valHeap[0]      //取堆顶
			for res.index <= i-k { // 当index<= i-k时，将堆顶pop出来，然后再取堆顶数据来对比看是否有效
				heap.Pop(&valHeap)
				res = valHeap[0]
			}
			result = append(result, res.val)
		}
	}
	return result
}
