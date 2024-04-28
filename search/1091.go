package search

import (
	"container/heap"
	"container/list"
)

// 普通BFS
func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	depthMap := make(map[int]int)
	queue := list.New()
	begin := transfer(0, 0, n)
	depthMap[begin] = 1
	queue.PushBack(begin)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	target := transfer(n-1, n-1, n)
	if begin == target {
		return 1
	}

	for queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)
		value := front.Value.(int)
		frontLevel := depthMap[value]
		for _, n := range move(grid, value, n) {
			if n == target {
				return frontLevel + 1
			}
			if _, exist := depthMap[n]; exist {
				continue
			}
			queue.PushBack(n)
			depthMap[n] = frontLevel + 1
		}
	}
	return -1
}

// 双向BFS
func shortestPathBinaryMatrix2(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}

	begin := transfer(0, 0, n)
	target := transfer(n-1, n-1, n)
	if begin == target {
		return 1
	}

	beginDepthMap := make(map[int]int)
	beginQueue := list.New()
	beginDepthMap[begin] = 1
	beginQueue.PushBack(begin)

	endDepthMap := make(map[int]int)
	endQueue := list.New()
	endDepthMap[target] = 1
	endQueue.PushBack(target)

	for beginQueue.Len() > 0 && endQueue.Len() > 0 {

		beginLen := beginQueue.Len()
		for beginLen > 0 {
			front := beginQueue.Front()
			beginQueue.Remove(front)
			beginLen--
			value := front.Value.(int)
			frontLevel := beginDepthMap[value]
			for _, n := range move(grid, value, n) {
				if value, exist := endDepthMap[n]; exist {
					return frontLevel + value
				}
				if _, exist := beginDepthMap[n]; exist {
					continue
				}
				beginQueue.PushBack(n)
				beginDepthMap[n] = frontLevel + 1
			}
		}

		endLen := endQueue.Len()
		for endLen > 0 {
			front := endQueue.Front()
			endQueue.Remove(front)
			endLen--
			value := front.Value.(int)
			frontLevel := endDepthMap[value]
			for _, n := range move(grid, value, n) {
				if value, exist := beginDepthMap[n]; exist {
					return frontLevel + value
				}
				if _, exist := endDepthMap[n]; exist {
					continue
				}
				endQueue.PushBack(n)
				endDepthMap[n] = frontLevel + 1
			}
		}

	}
	return -1
}

func transfer(i, j int, n int) int {
	return i*n + j
}

// 可以优化为方向数组！
func move(grid [][]int, i int, n int) []int {
	srcx := i / n
	srcy := i % n
	var res []int
	if srcx > 0 { // 向上移动
		nx := srcx - 1
		ny := srcy
		if grid[nx][ny] == 0 {
			res = append(res, nx*n+ny)
		}
	}
	if srcx < n-1 { // 向下移动
		nx := srcx + 1
		ny := srcy
		if grid[nx][ny] == 0 {
			res = append(res, nx*n+ny)
		}
	}
	if srcy > 0 { // 向左移动
		nx := srcx
		ny := srcy - 1
		if grid[nx][ny] == 0 {
			res = append(res, nx*n+ny)
		}
	}

	if srcy < n-1 { // 向右移动
		nx := srcx
		ny := srcy + 1
		if grid[nx][ny] == 0 {
			res = append(res, nx*n+ny)
		}
	}
	if srcx > 0 && srcy > 0 { // 向左上移动
		nx := srcx - 1
		ny := srcy - 1
		if grid[nx][ny] == 0 {
			res = append(res, nx*n+ny)
		}
	}
	if srcx > 0 && srcy < n-1 { // 向右上移动
		nx := srcx - 1
		ny := srcy + 1
		if grid[nx][ny] == 0 {
			res = append(res, nx*n+ny)
		}
	}
	if srcx < n-1 && srcy > 0 { // 向左下移动
		nx := srcx + 1
		ny := srcy - 1
		if grid[nx][ny] == 0 {
			res = append(res, nx*n+ny)
		}
	}
	if srcx < n-1 && srcy < n-1 { // 向右下移动
		nx := srcx + 1
		ny := srcy + 1
		if grid[nx][ny] == 0 {
			res = append(res, nx*n+ny)
		}
	}
	return res
}

type item struct {
	estimateCost int
	num          int
	heapIndex    int
}

type items []*item

func (b *items) Len() int {
	return len(*b)
}

func (b *items) Less(i, j int) bool {
	return (*b)[i].estimateCost < (*b)[j].estimateCost
}

func (b *items) Swap(i, j int) {
	(*b)[i], (*b)[j] = (*b)[j], (*b)[i]
	(*b)[i].heapIndex = i
	(*b)[j].heapIndex = j
}

func (b *items) Push(x interface{}) {
	pushItem := x.(*item)
	pushItem.heapIndex = len(*b)
	*b = append(*b, pushItem)

}

func (b *items) Pop() interface{} {
	it := (*b)[len(*b)-1]
	it.heapIndex = -1
	*b = (*b)[0 : len(*b)-1]
	return it
}

func shortestPathBinaryMatrix3(grid [][]int) int {
	n := len(grid)
	depthMap := make(map[int]int)
	itemMap := make(map[int]*item)
	begin := transfer(0, 0, n)
	beginItem := &item{
		estimateCost: estiCost(begin, n) + 0,
		num:          begin,
	}
	var itemHeap items
	depthMap[begin] = 1
	heap.Push(&itemHeap, beginItem)
	itemMap[begin] = beginItem
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	target := transfer(n-1, n-1, n)
	if begin == target {
		return 1
	}

	for itemHeap.Len() > 0 {
		front := heap.Pop(&itemHeap)
		value := front.(*item)
		frontLevel := depthMap[value.num]
		if value.num == target {
			return depthMap[value.num]
		}
		for _, ns := range move(grid, value.num, n) {
			_, exist := depthMap[ns]
			if !exist {
				newItem := &item{
					estimateCost: estiCost(ns, n) + frontLevel + 1,
					num:          ns,
				}
				heap.Push(&itemHeap, newItem)
				depthMap[ns] = frontLevel + 1
				itemMap[ns] = newItem
				continue
			}
			if depthMap[ns] > frontLevel+1 { // 与普通BFS不一样，优先队列的BFS需要校验 新的代价是否小于旧的代价，是的话更新！
				nsItem := itemMap[ns]
				nsItem.estimateCost = estiCost(ns, n) + frontLevel + 1
				depthMap[ns] = frontLevel + 1
				heap.Fix(&itemHeap, nsItem.heapIndex)
			}

		}
	}
	return -1
}

func estiCost(i, n int) int {
	x := i / n
	y := i % n
	return max(n-1-x, n-1-y)
}
