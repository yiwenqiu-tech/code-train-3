package main

import (
	"container/heap"
	"fmt"
)

func main() {
	var input [9]string
	fmt.Scanf("%s %s %s %s %s %s %s %s %s", &input[0], &input[1], &input[2],
		&input[3], &input[4], &input[5], &input[6], &input[7], &input[8])
	var board [3][3]int
	for i, c := range input {
		x := i / 3
		y := i % 3
		if c == "x" {
			board[x][y] = 0
		} else {
			board[x][y] = int([]byte(c)[0] - '0')
		}
	}
	fmt.Println(slidingPuzzle(board))
}

var target = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}}

type boardItem struct {
	estimateCost int
	board        [3][3]int
	change       string
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

func (b *boardItems) Push(x interface{}) {
	*b = append(*b, x.(*boardItem))
}

func (b *boardItems) Pop() interface{} {
	item := (*b)[len(*b)-1]

	*b = (*b)[0 : len(*b)-1]

	return item
}

func slidingPuzzle(board [3][3]int) string {
	if board == target {
		return ""
	}
	var depthMap = make(map[[3][3]int]int)
	depthMap[board] = 0

	var h boardItems
	heap.Push(&h, &boardItem{
		board:        board,
		estimateCost: 0 + estimateCost(&board),
	})

	for h.Len() > 0 {
		front := heap.Pop(&h)
		frontItem := front.(*boardItem)
		frontItemLevel := depthMap[frontItem.board]
		for _, s := range transformNewState(frontItem.board) {
			if _, exist := depthMap[s.newState]; exist {
				continue
			}
			if s.newState == target {
				return frontItem.change + s.change
			}
			depthMap[s.newState] = frontItemLevel + 1
			heap.Push(&h, &boardItem{
				board:        s.newState,
				estimateCost: frontItemLevel + 1 + estimateCost(&s.newState), // 注意是，当前代价+未来估价
				change:       frontItem.change + s.change,
			})
		}
	}
	return "unsolvable"
}

func estimateCost(items *[3][3]int) int {
	target := [9][2]int{{1, 1}, {0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}, {2, 1}, {2, 0}, {1, 0}}
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

func abs(value int) int {
	if value < 0 {
		return -value
	} else {
		return value
	}
}

type state struct {
	newState [3][3]int
	change   string
}

func transformNewState(items [3][3]int) []state {
	var res []state
	for i := 0; i < len(items); i++ {
		for j := 0; j < len(items[i]); j++ {
			if items[i][j] == 0 {
				if i == 0 {
					tmp := items
					tmp[i][j], tmp[i+1][j] = tmp[i+1][j], tmp[i][j]
					res = append(res, state{
						newState: tmp,
						change:   "d",
					})
				} else if i == 2 {
					tmp := items
					tmp[i][j], tmp[i-1][j] = tmp[i-1][j], tmp[i][j]
					res = append(res, state{
						newState: tmp,
						change:   "u",
					})
				} else {
					tmp := items
					tmp[i][j], tmp[i+1][j] = tmp[i+1][j], tmp[i][j]
					res = append(res, state{
						newState: tmp,
						change:   "d",
					})
					tmp = items
					tmp[i][j], tmp[i-1][j] = tmp[i-1][j], tmp[i][j]
					res = append(res, state{
						newState: tmp,
						change:   "u",
					})
				}
				if j == 0 {
					tmp := items
					tmp[i][j], tmp[i][j+1] = tmp[i][j+1], tmp[i][j]
					res = append(res, state{
						newState: tmp,
						change:   "r",
					})
				} else if j == 2 {
					tmp := items
					tmp[i][j], tmp[i][j-1] = tmp[i][j-1], tmp[i][j]
					res = append(res, state{
						newState: tmp,
						change:   "l",
					})
				} else {
					tmp := items
					tmp[i][j], tmp[i][j+1] = tmp[i][j+1], tmp[i][j]
					res = append(res, state{
						newState: tmp,
						change:   "r",
					})
					tmp = items
					tmp[i][j], tmp[i][j-1] = tmp[i][j-1], tmp[i][j]
					res = append(res, state{
						newState: tmp,
						change:   "l",
					})
				}
			}
		}
	}
	return res
}
