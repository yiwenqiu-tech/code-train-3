package main

/*
import (
	"container/heap"
	"fmt"
)

func main() {
	var input string
	fmt.Scanf("%s", &input)
	var board [3][3]int
	for i, c := range input {
		x := i / 3
		y := i % 3
		board[x][y] = int(c - '0')
	}
	fmt.Println(slidingPuzzle(board))
}

var target = [3][3]int{{1, 2, 3}, {8, 0, 4}, {7, 6, 5}}

type boardItem struct {
	estimateCost int
	board        [3][3]int
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

func slidingPuzzle(board [3][3]int) int {
	if board == target {
		return 0
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

func transformNewState(items [3][3]int) [][3][3]int {
	var res [][3][3]int
	for i := 0; i < len(items); i++ {
		for j := 0; j < len(items[i]); j++ {
			if items[i][j] == 0 {
				if i == 0 {
					tmp := items
					tmp[i][j], tmp[i+1][j] = tmp[i+1][j], tmp[i][j]
					res = append(res, tmp)
				} else if i == 2 {
					tmp := items
					tmp[i][j], tmp[i-1][j] = tmp[i-1][j], tmp[i][j]
					res = append(res, tmp)
				} else {
					tmp := items
					tmp[i][j], tmp[i+1][j] = tmp[i+1][j], tmp[i][j]
					res = append(res, tmp)
					tmp = items
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
*/
