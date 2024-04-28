package search

import (
	"container/heap"
	"container/list"
)

var target = [2][3]int{{1, 2, 3}, {4, 5, 0}}

func slidingPuzzle(board [][]int) int {
	initBoard := [2][3]int{}
	for i := range board {
		for j := range board[i] {
			initBoard[i][j] = board[i][j]
		}
	}
	if initBoard == target {
		return 0
	}
	var depthMap = make(map[[2][3]int]int)
	queue := list.New()
	queue.PushBack(initBoard)
	depthMap[initBoard] = 0

	for queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)
		frontItem := front.Value.([2][3]int)
		frontItemLevel := depthMap[frontItem]
		for _, s := range transformNewState(frontItem) {
			if _, exist := depthMap[s]; exist {
				continue
			}
			if s == target {
				return frontItemLevel + 1
			}
			depthMap[s] = frontItemLevel + 1
			queue.PushBack(s)
		}
	}
	return -1
}

func transformNewState(items [2][3]int) [][2][3]int {
	var res [][2][3]int
	for i := 0; i < len(items); i++ {
		for j := 0; j < len(items[i]); j++ {
			if items[i][j] == 0 {
				if i == 0 {
					tmp := items
					tmp[i][j], tmp[i+1][j] = tmp[i+1][j], tmp[i][j]
					res = append(res, tmp)
				} else {
					tmp := items
					tmp[i][j], tmp[i-1][j] = tmp[i-1][j], tmp[i][j]
					res = append(res, tmp)
				}
				if j == 0 {
					tmp := items
					tmp[i][j], tmp[i][j+1] = tmp[i][j+1], tmp[i][j]
					res = append(res, tmp)
				} else if j == 2 {
					tmp := items
					tmp[i][j], tmp[i][j-1] = tmp[i][j-1], tmp[i][j]
					res = append(res, tmp)
				} else {
					tmp := items
					tmp[i][j], tmp[i][j+1] = tmp[i][j+1], tmp[i][j]
					res = append(res, tmp)
					tmp = items
					tmp[i][j], tmp[i][j-1] = tmp[i][j-1], tmp[i][j]
					res = append(res, tmp)
				}
			}
		}
	}
	return res
}

type boardItem struct {
	estimateCost int
	board        [2][3]int
}

type boardItems []*boardItem

func (b *boardItems) Len() int {
	return len(*b)
}

func (b *boardItems) Less(i, j int) bool {
	return (*b)[i].estimateCost < (*b)[j].estimateCost
}

func (b *boardItems) Swap(i, j int) {
	(*b)[i], (*b)[j] = (*b)[j], (*b)[i]
}

func (b *boardItems) Push(x any) {
	*b = append(*b, x.(*boardItem))
}

func (b *boardItems) Pop() any {
	item := (*b)[len(*b)-1]

	*b = (*b)[0 : len(*b)-1]

	return item
}

func slidingPuzzle2(board [][]int) int {
	initBoard := [2][3]int{}
	for i := range board {
		for j := range board[i] {
			initBoard[i][j] = board[i][j]
		}
	}
	if initBoard == target {
		return 0
	}
	var depthMap = make(map[[2][3]int]int)
	depthMap[initBoard] = 0

	var h boardItems
	heap.Push(&h, &boardItem{
		board:        initBoard,
		estimateCost: 0 + estimateCost(&initBoard),
	})

	for h.Len() > 0 {
		front := heap.Pop(&h)
		frontItem := front.(*boardItem)
		frontItemLevel := depthMap[frontItem.board]
		for _, s := range transformNewState(frontItem.board) {
			if _, exist := depthMap[s]; exist {
				continue
			}
			if s == target {
				return frontItemLevel + 1
			}
			depthMap[s] = frontItemLevel + 1
			heap.Push(&h, &boardItem{
				board:        s,
				estimateCost: frontItemLevel + 1 + estimateCost(&s), // 注意是，当前代价+未来估价
			})
		}
	}
	return -1
}

func estimateCost(items *[2][3]int) int {
	target := [6][2]int{{1, 2}, {0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 1}}
	var res int
	for i := 0; i < len(items); i++ {
		for j := 0; j < len(items[i]); j++ {
			item := items[i][j]
			if item == 0 {
				continue
			}
			targetForItem := target[item]
			res += abs(targetForItem[0]-i) + abs(targetForItem[1]-j)
		}
	}
	return res
}
