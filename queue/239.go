package queue

import "container/list"

type Item struct {
	Index int
	Value int
}

func maxSlidingWindow(nums []int, k int) []int {
	deque := list.New()
	var res []int
	for index, n := range nums {
		for deque.Len() > 0 && index-deque.Front().Value.(*Item).Index >= k { // 移除过期元素
			deque.Remove(deque.Front())
		}
		for deque.Len() > 0 && n >= deque.Back().Value.(*Item).Value {
			deque.Remove(deque.Back())
		}
		deque.PushBack(&Item{
			Index: index,
			Value: n,
		})
		if index >= k-1 {
			res = append(res, deque.Front().Value.(*Item).Value)
		}
	}
	return res
}
