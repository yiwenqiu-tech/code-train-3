package train

import (
	"container/heap"
	"fmt"
	"math/rand"
)

type pointHeap [][]int

func (ih *pointHeap) Len() int {
	return len(*ih)
}

func (ih *pointHeap) Less(i, j int) bool {
	pi := (*ih)[i]
	pj := (*ih)[j]
	return pi[0]*pi[0]+pi[1]*pi[1] < pj[0]*pj[0]+pj[1]*pj[1]
}

func (ih *pointHeap) Swap(i, j int) {
	(*ih)[i], (*ih)[j] = (*ih)[j], (*ih)[i]
}

func (ih *pointHeap) Push(x any) {
	*ih = append(*ih, x.([]int))
}

func (ih *pointHeap) Pop() any {
	item := (*ih)[len(*ih)-1]
	*ih = (*ih)[:len(*ih)-1]
	return item
}

func kClosest(points [][]int, k int) [][]int {
	var pHeap pointHeap = points
	heap.Init(&pHeap) // O(N)
	fmt.Println(pHeap)
	var res [][]int
	for k > 0 {
		res = append(res, heap.Pop(&pHeap).([]int))
		k--
	}
	return res
}

type PointInfo struct {
	point    []int
	distinct int
}

func kClosest2(points [][]int, k int) [][]int {
	if len(points) <= k {
		return points
	}
	ps := make([]*PointInfo, len(points))
	for i := range points {
		ps[i] = &PointInfo{
			point:    points[i],
			distinct: points[i][0]*points[i][0] + points[i][1]*points[i][1],
		}
	}

	fastChoose973(ps, 0, len(points)-1, k)

	var res [][]int
	for i := range ps[0:k] {
		res = append(res, ps[i].point)
	}
	return res
}

func fastChoose973(ps []*PointInfo, left, right, index int) {
	pivot := partition973(ps, left, right)
	if pivot == index {
		return
	} else if pivot < index {
		fastChoose973(ps, pivot+1, right, index)
	} else {
		fastChoose973(ps, left, pivot-1, index)
	}
}

func partition973(ps []*PointInfo, left, right int) int {
	pivot := left + rand.Intn(right-left+1) // 随机数
	pivotValue := ps[pivot]
	ps[left], ps[pivot] = ps[pivot], ps[left] // 交换
	for left < right {
		for left < right && ps[right].distinct >= pivotValue.distinct {
			right--
		}
		if left >= right {
			break
		}
		ps[left] = ps[right]
		for left < right && ps[left].distinct <= pivotValue.distinct {
			left++
		}
		if left >= right {
			break
		}
		ps[right] = ps[left]
	}
	ps[left] = pivotValue
	return left
}
